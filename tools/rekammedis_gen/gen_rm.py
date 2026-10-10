"""Generator form asesmen rekam medis: model, request, descriptor Form, routes.

python3 gen_rm.py            -> tulis ke backend (FORCE=1 untuk menimpa)
"""
import json
import os
import re
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import ddl  # noqa: E402

ROOT = os.environ.get('GEN_ROOT', os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', '..'))
HERE = os.path.dirname(os.path.abspath(__file__))
T = ddl.tables()

PETUGAS_TABLES = {('pegawai', 'nik'), ('dokter', 'kd_dokter'), ('petugas', 'nip')}
EXCLUDE_TABLES = {'data_triase_igd', 'rekonsiliasi_obat', 'skrining_rawat_jalan', 'master_kesimpulan_anjuran_mcu'}
STRINGS = {'varchar', 'char', 'text', 'mediumtext', 'longtext', 'tinytext', 'enum'}


def write(path, content):
    full = os.path.join(ROOT, path)
    os.makedirs(os.path.dirname(full), exist_ok=True)
    if os.path.exists(full) and not os.environ.get('FORCE'):
        return
    with open(full, 'w') as f:
        f.write(content)


def kind(c):
    t = c['type']
    if t in ddl.INTS:
        return 'int'
    if t in ddl.FLOATS:
        return 'float'
    if t == 'date':
        return 'date'
    if t in ('datetime', 'timestamp'):
        return 'datetime'
    if t == 'time':
        return 'time'
    if t in STRINGS:
        return 'string'
    return 'skip'


def model_type(c):
    k = kind(c)
    if k in ('date', 'datetime'):
        return '*time.Time'
    base = {'int': 'int', 'float': 'float64', 'time': 'string', 'string': 'string'}[k]
    return base if c['notnull'] else '*' + base


def req_type(c):
    k = kind(c)
    if k == 'int':
        return 'int' if c['notnull'] else '*int'
    if k == 'float':
        return 'float64' if c['notnull'] else '*float64'
    return 'string'


def label_of(cls):
    name = re.sub(r'^RM', '', cls)
    name = re.sub(r'^Data(?=[A-Z])', '', name)
    words = re.findall(r'[A-Z]+(?=[A-Z][a-z]|$)|[A-Z]?[a-z]+|\d+', name)
    return ' '.join(w.lower() if not w.isupper() or len(w) == 1 else w for w in words)


def select_forms():
    scan = json.load(open(os.path.join(HERE, 'rm_scan.json')))
    forms, seen = [], set()
    for r in scan:
        main = r.get('main')
        if not main or r['file'].startswith('Master') or main.startswith('temporary') or main in EXCLUDE_TABLES:
            continue
        t = T[main]
        if 'no_rawat' not in t['pk'] or main in seen:
            continue
        if any(kind(c) == 'skip' for c in t['cols']):
            print('skip (tipe kolom)', main)
            continue
        seen.add(main)
        forms.append((r['file'], main, r.get('children', [])))
    return forms


def detail_specs(main, children):
    out = []
    for ch in children:
        t = T.get(ch)
        if not t:
            continue
        cols = [c['name'] for c in t['cols']]
        if len(cols) != 2 or cols[0] != 'no_rawat' or not cols[1].startswith('kode_'):
            continue
        if t['fks'].get('no_rawat', (None,))[0] != main or cols[1] not in t['fks']:
            continue
        rt, rc = t['fks'][cols[1]]
        out.append({'name': cols[1], 'table': ch, 'column': cols[1], 'ref_table': rt, 'ref_col': rc})
    return out


def gen(cls, table, children):
    t = T[table]
    st = ddl.camel(table)
    label = label_of(cls)
    slug = table.replace('_', '-')
    cols = t['cols']
    keys = t['pk']
    colmap = {c['name']: c for c in cols}
    inputs = [c for c in cols if c['name'] not in keys]
    details = detail_specs(table, children)
    petugas = [c for c in cols if c['name'] in t['fks'] and t['fks'][c['name']] in PETUGAS_TABLES]
    refs = [c for c in cols if c['name'] in t['fks'] and c['name'] != 'no_rawat' and t['fks'][c['name']][0] not in (table, 'reg_periksa')]
    needs_time = any(model_type(c) == '*time.Time' for c in cols)
    fname = {c['name']: ddl.camel(c['name']) for c in cols}
    if len(set(fname.values())) != len(fname):
        raise SystemExit('nama field bentrok: ' + table)

    # waktu asesmen
    if 'tanggal' in colmap and kind(colmap['tanggal']) == 'datetime':
        waktu_go, waktu_sql = 'action.Tm(m.Tanggal)', 'tanggal'
    elif 'tgl_perawatan' in colmap and 'jam_rawat' in colmap:
        waktu_go, waktu_sql = 'action.TglJam(m.TglPerawatan, ' + deref(colmap['jam_rawat'], 'm.JamRawat') + ')', "concat(tgl_perawatan,' ',jam_rawat)"
    elif 'tanggal' in colmap and kind(colmap['tanggal']) == 'date' and 'jam' in colmap:
        waktu_go, waktu_sql = 'action.TglJam(m.Tanggal, ' + deref(colmap['jam'], 'm.Jam') + ')', "concat(tanggal,' ',jam)"
    else:
        dt = [k for k in keys if kind(colmap[k]) == 'datetime']
        if dt:
            waktu_go, waktu_sql = f'action.Tm(m.{fname[dt[0]]})', dt[0]
        else:
            waktu_go, waktu_sql = 'time.Time{}', ''

    # ---------- model
    w = max(len(fname[c['name']]) for c in cols)
    tw = max(len(model_type(c)) for c in cols)
    lines = []
    for c in cols:
        tag = f"column:{c['name']}"
        if c['name'] in keys:
            tag += ';primaryKey;autoIncrement:false'
        lines.append(f"\t{fname[c['name']].ljust(w)} {model_type(c).ljust(tw)} `gorm:\"{tag}\" json:\"{c['name']}\"`")
    write(f'app/models/rekammedis/{table}.go', f"""package rekammedis
{chr(10) + 'import "time"' + chr(10) if needs_time else ''}
// {st} tabel `{table}` ({label}, {cls}).
type {st} struct {{
{chr(10).join(lines)}
}}

func ({st}) TableName() string {{
	return "{table}"
}}
""")

    # ---------- request
    rw = max([len(fname[c['name']]) for c in inputs] + [1])
    rtw = max([len(req_type(c)) for c in inputs] + [1])
    fields = '\n'.join(f"\t{fname[c['name']].ljust(rw)} {req_type(c).ljust(rtw)} `form:\"{c['name']}\" json:\"{c['name']}\"`" for c in inputs)
    pet_names = {c['name'] for c in petugas}

    def rule(c, is_key=False):
        k = kind(c)
        parts = []
        if is_key:
            if c['name'] == 'no_rawat' or k == 'string':
                parts.append('required')
        elif c['notnull'] and (c['type'] == 'enum' or k in ('date', 'datetime', 'time') or c['name'] in pet_names):
            parts.append('required')
        if c['type'] == 'enum' and enum_in_action(c):
            if '' in ddl.enum_values(c) and 'required' in parts:
                parts.remove('required')
            parts.append('string')
        elif c['type'] == 'enum':
            parts.append('in:' + ','.join(ddl.enum_values(c)))
        elif k == 'int':
            parts.append('int')
        elif k == 'float':
            parts.append('numeric')
        elif k in ('date', 'datetime'):
            parts.append('date')
        elif k == 'time':
            parts.append('string|len:8')
        elif c['type'] in ('varchar', 'char') and c['arg']:
            parts += ['string', 'max_len:' + c['arg']]
        else:
            parts.append('string')
        return '|'.join(parts)

    kw = max(len(c['name']) for c in cols) + 3
    rules = '\n'.join(f"\t\t{('\"' + c['name'] + '\":').ljust(kw)} \"{rule(c)}\"," for c in inputs)
    key_fields = '\n'.join(f"\t{fname[k]} string `form:\"{k}\" json:\"{k}\"`" for k in keys)
    key_rules = '\n'.join(f"\trules[\"{k}\"] = \"{rule(colmap[k], True)}\"" for k in keys)
    key_values = ', '.join(f'"{k}": r.{fname[k]}' for k in keys)
    det_fields = ''.join(f"\t{ddl.camel(d['name'])} []string `form:\"{d['name']}\" json:\"{d['name']}\"`\n" for d in details)
    det_rules = ''.join(f"\trules[\"{d['name']}\"] = \"slice\"\n\trules[\"{d['name']}.*\"] = \"string|max_len:20\"\n" for d in details)
    det_values = ', '.join(f'"{d["name"]}": r.{ddl.camel(d["name"])}' for d in details)
    det_struct = f"\n// {st}Detail daftar kode detail (diganti seluruhnya saat simpan/ubah).\ntype {st}Detail struct {{\n{det_fields}}}\n" if details else ''
    det_embed = f'\t{st}Detail\n' if details else ''
    det_return = f'map[string][]string{{{det_values}}}' if details else 'nil'
    write(f'app/http/requests/rekammedis/{table}.go', f"""package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// {st}Data isian {label}.
type {st}Data struct {{
{fields}
}}
{det_struct}
func {lc(st)}Rules() map[string]any {{
	rules := map[string]any{{
{rules}
	}}
{det_rules}	return rules
}}

// {st}Store simpan {label}; kolom waktu kunci kosong = sekarang.
type {st}Store struct {{
{key_fields}
	{st}Data
{det_embed}}}

func (r *{st}Store) Authorize(ctx http.Context) error {{ return nil }}

func (r *{st}Store) Rules(ctx http.Context) map[string]any {{
	rules := {lc(st)}Rules()
{key_rules}
	return rules
}}

func (r *{st}Store) KeyValues() repo.Key {{ return repo.Key{{{key_values}}} }}

func (r *{st}Store) Payload() {st}Data {{ return r.{st}Data }}

func (r *{st}Store) DetailValues() map[string][]string {{ return {det_return} }}

// {st}Update ubah {label} (PUT); kunci lewat query string.
type {st}Update struct {{
	{st}Data
{det_embed}}}

func (r *{st}Update) Authorize(ctx http.Context) error {{ return nil }}

func (r *{st}Update) Rules(ctx http.Context) map[string]any {{ return {lc(st)}Rules() }}

func (r *{st}Update) Payload() {st}Data {{ return r.{st}Data }}

func (r *{st}Update) DetailValues() map[string][]string {{ return {det_return} }}
""")

    # ---------- form descriptor (action)
    setkey, keyof = [], []
    for k in keys:
        c = colmap[k]
        f = fname[k]
        kk = kind(c)
        if kk == 'datetime':
            setkey.append(f'\tif m.{f}, err = action.KeyDateTime(key["{k}"]); err != nil {{\n\t\treturn err\n\t}}')
            keyof.append(f'"{k}": action.FmtDateTime(m.{f})')
        elif kk == 'date':
            setkey.append(f'\tif m.{f}, err = action.KeyDate(key["{k}"]); err != nil {{\n\t\treturn err\n\t}}')
            keyof.append(f'"{k}": action.FmtDate(m.{f})')
        elif kk == 'time':
            setkey.append(f'\tif m.{f}, err = action.KeyTime(key["{k}"]); err != nil {{\n\t\treturn err\n\t}}')
            keyof.append(f'"{k}": m.{f}')
        else:
            setkey.append(f'\tm.{f} = strings.TrimSpace(key["{k}"])')
            keyof.append(f'"{k}": m.{f}')
    fill, pre = [], []
    for c in inputs:
        f = fname[c['name']]
        k = kind(c)
        v = 'v' + f
        if k == 'date':
            pre.append(f'\t{v}, err := support.ParseDate(d.{f})\n\tif err != nil {{\n\t\treturn err\n\t}}')
            fill.append(f'\tm.{f} = {v}')
        elif k == 'datetime':
            pre.append(f'\t{v}, err := support.ParseDateTime(d.{f})\n\tif err != nil {{\n\t\treturn err\n\t}}')
            fill.append(f'\tm.{f} = {v}')
        elif k == 'time':
            pre.append(f'\t{v}, err := action.OptTime(d.{f})\n\tif err != nil {{\n\t\treturn err\n\t}}')
            fill.append(f'\tm.{f} = {v}' if c['notnull'] else f'\tm.{f} = support.Nullable({v})')
        elif k == 'string':
            if c['type'] == 'enum' and enum_in_action(c):
                vals = ', '.join(json.dumps(v) for v in ddl.enum_values(c))
                pre.append(f'\tif err := Enum("{c["name"]}", d.{f}, {vals}); err != nil {{\n\t\treturn err\n\t}}')
            fill.append(f'\tm.{f} = strings.TrimSpace(d.{f})' if c['notnull'] else f'\tm.{f} = support.Nullable(d.{f})')
        else:
            fill.append(f'\tm.{f} = d.{f}')
    pet = ', '.join(deref(c, 'm.' + fname[c['name']]) for c in petugas)
    ref_lines = []
    for c in refs:
        rt, rc = t['fks'][c['name']]
        f = fname[c['name']]
        k = kind(c)
        if k == 'string':
            val = f'action.Str(m.{f})' if c['notnull'] else f'action.StrPtr(m.{f})'
        elif c['notnull']:
            val = f'm.{f}'
        else:
            val = f'derefAny(m.{f})'
        ref_lines.append(f'\t\t{{Column: "{c["name"]}", Table: "{rt}", RefCol: "{rc}", Label: "{rt.replace("_", " ")}", '
                         f'Value: func(m *model.{st}) any {{ return {val} }}}},')
    det_spec = ''.join(f'\t\t\t{{Name: "{d["name"]}", Table: "{d["table"]}", Column: "{d["column"]}", RefTable: "{d["ref_table"]}", RefColumn: "{d["ref_col"]}"}},\n'
                       for d in details)
    search = ', '.join(f'"{s}"' for s in ['no_rawat'] + [c['name'] for c in petugas])
    body = '\n'.join(pre + fill)
    uses = body + '\n'.join(setkey) + '\n'.join(ref_lines) + waktu_go
    imports = []
    if 'strings.' in uses:
        imports.append('"strings"')
    imports.append('"time"')
    imp = ('\t' + '\n\t'.join(imports) + '\n\n') if imports else ''
    sup = '\t"goravel/app/support"\n' if 'support.' in uses else ''
    err_decl = '\tvar err error\n' if any('err =' in s for s in setkey) else ''
    waktu_go = waktu_go.replace('action.', '')
    setkey = [x.replace('action.', '') for x in setkey]
    keyof = [x.replace('action.', '') for x in keyof]
    pre = [x.replace('action.', '') for x in pre]
    ref_lines = [x.replace('action.', '') for x in ref_lines]
    body = '\n'.join(pre + fill)
    write(f'app/actions/rekammedis/form_{table}.go', f"""package rekammedis

import (
{imp}	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
{sup})

// Form{st} {label} ({cls}).
var Form{st} = &Form[model.{st}, request.{st}Data]{{
	Slug:  "{slug}",
	Label: "{label}",
	Spec: repo.Spec{{
		Table:  "{table}",
		Keys:   []string{{{', '.join(f'"{k}"' for k in keys)}}},
		Waktu:  "{waktu_sql}",
		Search: []string{{{search}}},
{('		Details: []repo.Detail{' + chr(10) + det_spec + '		},' + chr(10)) if details else ''}	}},
	SetKey: func(m *model.{st}, key repo.Key) error {{
{err_decl}{chr(10).join(setkey)}
		return {'err' if err_decl else 'nil'}
	}},
	KeyOf: func(m *model.{st}) repo.Key {{
		return repo.Key{{{', '.join(keyof)}}}
	}},
	Waktu:   func(m *model.{st}) time.Time {{ return {waktu_go} }},
	Petugas: func(m *model.{st}) []string {{ return []string{{{pet}}} }},
	Fill: func(m *model.{st}, d request.{st}Data) error {{
{body}
		return nil
	}},
	Refs: []Ref[model.{st}]{{
{chr(10).join(ref_lines)}
	}},
}}
""")
    return table, st, label


def enum_in_action(c):
    return any(',' in v or v == '' or '|' in v for v in ddl.enum_values(c))


def deref(c, expr):
    return expr if c['notnull'] else f'derefStr({expr})'


def lc(s):
    return s[:1].lower() + s[1:]


def routes(items):
    body = '\n'.join(
        f'\treg("{t.replace("_", "-")}", rmctrl.NewController(rmaction.Form{st}, '
        f'func() rmreq.StoreRequest[rmreq.{st}Data] {{ return &rmreq.{st}Store{{}} }}, '
        f'func() rmreq.UpdateRequest[rmreq.{st}Data] {{ return &rmreq.{st}Update{{}} }}))'
        for t, st, _ in items)
    reg = '\n'.join(f'\t{{Slug: "{t.replace("_", "-")}", Label: "{lb}", Table: "{t}"}},' for t, st, lb in items)
    write('app/actions/rekammedis/daftar_form.go', f"""// Kode dibangkitkan dari sik.sql + source/src/rekammedis.
package rekammedis

// FormInfo identitas form asesmen untuk ringkasan riwayat.
type FormInfo struct {{
	Slug  string `json:"slug"`
	Label string `json:"label"`
	Table string `json:"-"`
}}

// DaftarForm seluruh form asesmen rekam medis.
var DaftarForm = []FormInfo{{
{reg}
}}
""")
    write('routes/api_rekam_medis_asesmen.go', f"""// Kode dibangkitkan dari sik.sql + source/src/rekammedis; satu baris per form asesmen.
package routes

import (
	"github.com/goravel/framework/contracts/route"

	rmaction "goravel/app/actions/rekammedis"
	rmctrl "goravel/app/http/controllers/rekammedis"
	rmreq "goravel/app/http/requests/rekammedis"
	"goravel/app/modules/rbac"
)

// registerRekamMedisAsesmenRoutes /rekam-medis/<form>: GET (list), GET /detail, POST, PUT, DELETE; kunci lewat query string.
func registerRekamMedisAsesmenRoutes(router route.Router) {{
	perm := rbac.RequirePermission
	reg := func(slug string, c rmctrl.Handler) {{
		p := "/rekam-medis/" + slug
		router.Middleware(perm("rekam_medis.view")).Get(p, c.Index)
		router.Middleware(perm("rekam_medis.view")).Get(p+"/detail", c.Show)
		router.Middleware(perm("rekam_medis.create")).Post(p, c.Store)
		router.Middleware(perm("rekam_medis.update")).Put(p, c.Update)
		router.Middleware(perm("rekam_medis.delete")).Delete(p, c.Destroy)
	}}
{body}
}}
""")


if __name__ == '__main__':
    items = [gen(*f) for f in select_forms()]
    routes(items)
    print(len(items), 'form')
