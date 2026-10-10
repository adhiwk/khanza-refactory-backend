import os, re, sys, json, collections
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import ddl
SRC = os.environ.get('RM_SRC', os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', '..', '..', 'source', 'src', 'rekammedis'))
T = ddl.tables()
out = []
for f in sorted(os.listdir(SRC)):
    if not f.endswith('.java'): continue
    s = open(os.path.join(SRC, f), encoding='latin-1').read()
    ins = re.findall(r'Sequel\.menyimpan[a-z0-9]*\("(\w+)"', s)
    upd = re.findall(r'Sequel\.mengedit[a-z0-9]*\("(\w+)"', s)
    dele = re.findall(r'delete from (\w+)', s) + re.findall(r'Sequel\.meghapus[a-z0-9]*\("(\w+)"', s)
    tables = [t for t in ins if t in T]
    cand = [t for t in upd if t in T and t in tables]
    main = cand[0] if cand else (tables[0] if tables else None)
    rec = {'file': f[:-5], 'main': main, 'children': sorted(set(t for t in tables if t != main)),
           'upd': sorted(set(upd)), 'del': sorted(set(dele))}
    if main:
        t = T[main]; cols = [c['name'] for c in t['cols']]
        rec['pk'] = t['pk']
        rec['petugas'] = [c for c in cols if c in ('nip', 'kd_dokter', 'kd_petugas', 'nik', 'kd_dokter_bedah', 'petugas')]
        rec['waktu'] = [c['name'] + ':' + c['type'] for c in t['cols'] if c['type'] in ('date', 'datetime', 'time') and c['name'] in t['pk']]
        rec['ncols'] = len(cols)
        rec['has_norawat'] = 'no_rawat' in cols
        rec['blob'] = [c['name'] for c in t['cols'] if 'blob' in c['type']]
    out.append(rec)
json.dump(out, open(os.path.join(os.path.dirname(__file__), 'rm_scan.json'), 'w'), indent=1)
kinds = collections.Counter()
for r in out:
    if not r['main']: kinds['no-insert'] += 1; continue
    kinds['pk=' + ','.join(r['pk'])] += 1
for k, v in kinds.most_common(): print(v, k)
print('children:', sum(1 for r in out if r.get('children')))
print('no petugas:', [r['file'] for r in out if r['main'] and not r['petugas']][:40])
print('no no_rawat:', [r['file'] for r in out if r['main'] and not r['has_norawat']])
print('blob:', [r['file'] for r in out if r.get('blob')])
