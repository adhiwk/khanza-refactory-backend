// Package templatedokter use case template pemeriksaan dokter (MasterTemplatePemeriksaanDokter).
// Template milik dokter: hanya dokter pembuat yang boleh mengubah/menghapus, kecuali Admin Utama.
package templatedokter

import (
	"fmt"
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/templatedokter"
	repo "goravel/app/repository/templatedokter"
	"goravel/app/support"
)

var (
	ErrNotFound       = support.NotFound("data template pemeriksaan dokter tidak ditemukan")
	ErrDokterNotFound = support.NotFound("data dokter tidak ditemukan")
	ErrTanpaPegawai   = support.Forbidden("akun belum dihubungkan ke pegawai (users.kd_pegawai)")
	ErrBukanDokter    = support.Forbidden("hanya bisa diisi/diubah/dihapus oleh dokter yang bersangkutan")
	ErrRencana        = support.Invalid("rencana & instruksi wajib diisi")
)

type Actor struct {
	KodePegawai string
	SuperAdmin  bool
}

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search, kdDokter string, page, limit int) ([]model.Template, int64, error) {
	return a.repo.Paginate(strings.TrimSpace(search), strings.TrimSpace(kdDokter), page, limit)
}

func (a *Action) Detail(no string) (*model.Template, error) {
	t, err := a.repo.Find(strings.TrimSpace(no))
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNotFound
	}
	return t, nil
}

func (a *Action) Create(in model.Template, actor Actor) (*model.Template, error) {
	if err := a.validasi(&in, actor); err != nil {
		return nil, err
	}
	err := support.Transaction(func(tx contractsorm.Query) error {
		no, err := a.repo.NextKode(tx)
		if err != nil {
			return err
		}
		in.NoTemplate = no
		return a.repo.Insert(tx, in)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(in.NoTemplate)
}

func (a *Action) Update(no string, in model.Template, actor Actor) (*model.Template, error) {
	old, err := a.editable(no, actor)
	if err != nil {
		return nil, err
	}
	in.NoTemplate = old.NoTemplate
	if err := a.validasi(&in, actor); err != nil {
		return nil, err
	}
	if err := support.Transaction(func(tx contractsorm.Query) error { return a.repo.Update(tx, in) }); err != nil {
		return nil, err
	}
	return a.Detail(in.NoTemplate)
}

func (a *Action) Delete(no string, actor Actor) error {
	old, err := a.editable(no, actor)
	if err != nil {
		return err
	}
	return support.Transaction(func(tx contractsorm.Query) error { return a.repo.Delete(tx, old.NoTemplate) })
}

func (a *Action) editable(no string, actor Actor) (*model.Template, error) {
	t, err := a.Detail(no)
	if err != nil {
		return nil, err
	}
	if actor.SuperAdmin {
		return t, nil
	}
	return t, pemilik(t.KdDokter, actor)
}

func pemilik(kdDokter string, actor Actor) error {
	if actor.KodePegawai == "" {
		return ErrTanpaPegawai
	}
	if kdDokter != actor.KodePegawai {
		return ErrBukanDokter
	}
	return nil
}

func (a *Action) cek(table, column string, value any, label string) error {
	ok, err := a.repo.RefExists(table, column, value)
	if err != nil {
		return err
	}
	if !ok {
		return support.NotFound(fmt.Sprintf("%s %v tidak ditemukan", label, value))
	}
	return nil
}

// validasi membersihkan & memeriksa seluruh isi; kode duplikat dibuang, urut kosong diisi berurutan.
func (a *Action) validasi(t *model.Template, actor Actor) error {
	t.KdDokter = strings.TrimSpace(t.KdDokter)
	if !actor.SuperAdmin {
		if err := pemilik(t.KdDokter, actor); err != nil {
			return err
		}
	}
	if ok, err := a.repo.RefExists("dokter", "kd_dokter", t.KdDokter); err != nil {
		return err
	} else if !ok {
		return ErrDokterNotFound
	}
	if strings.TrimSpace(t.Rencana) == "" || strings.TrimSpace(t.Instruksi) == "" {
		return ErrRencana
	}

	var err error
	if t.Diagnosa, err = a.kode(t.Diagnosa, "penyakit", "kd_penyakit", "diagnosa", true); err != nil {
		return err
	}
	if t.Prosedur, err = a.kode(t.Prosedur, "icd9", "kode", "prosedur", true); err != nil {
		return err
	}
	for i := range t.Prosedur {
		t.Prosedur[i].Jumlah = support.OrDefault(t.Prosedur[i].Jumlah, "1")
	}
	if t.Radiologi, err = a.kode(t.Radiologi, "jns_perawatan_radiologi", "kd_jenis_prw", "pemeriksaan radiologi", false); err != nil {
		return err
	}
	if t.Tindakan, err = a.kode(t.Tindakan, "jns_perawatan", "kd_jenis_prw", "tindakan", false); err != nil {
		return err
	}

	labs := []model.Lab{}
	seenLab := map[string]bool{}
	for _, l := range t.Lab {
		if l.KdJenisPrw = strings.TrimSpace(l.KdJenisPrw); l.KdJenisPrw == "" || seenLab[l.KdJenisPrw] {
			continue
		}
		seenLab[l.KdJenisPrw] = true
		if err := a.cek("jns_perawatan_lab", "kd_jenis_prw", l.KdJenisPrw, "pemeriksaan laboratorium"); err != nil {
			return err
		}
		ids := []int{}
		seenID := map[int]bool{}
		for _, id := range l.IDTemplate {
			if seenID[id] {
				continue
			}
			seenID[id] = true
			ok, err := a.repo.LabTemplateExists(l.KdJenisPrw, id)
			if err != nil {
				return err
			}
			if !ok {
				return support.NotFound(fmt.Sprintf("item laboratorium %d bukan bagian pemeriksaan %s", id, l.KdJenisPrw))
			}
			ids = append(ids, id)
		}
		l.IDTemplate = ids
		labs = append(labs, l)
	}
	t.Lab = labs

	resep := []model.Resep{}
	seenObat := map[string]bool{}
	for _, o := range t.Resep {
		if o.KodeBrng = strings.TrimSpace(o.KodeBrng); o.KodeBrng == "" || seenObat[o.KodeBrng] {
			continue
		}
		seenObat[o.KodeBrng] = true
		if err := a.cek("databarang", "kode_brng", o.KodeBrng, "obat"); err != nil {
			return err
		}
		resep = append(resep, o)
	}
	t.Resep = resep

	racikan := []model.Racikan{}
	for i, rc := range t.Racikan {
		if len(rc.Detail) == 0 {
			continue
		}
		rc.NoRacik = support.OrDefault(rc.NoRacik, fmt.Sprint(i+1))
		if rc.KdRacik = strings.TrimSpace(rc.KdRacik); rc.KdRacik != "" {
			if err := a.cek("metode_racik", "kd_racik", rc.KdRacik, "metode racik"); err != nil {
				return err
			}
		}
		detail := []model.RacikanDetail{}
		seen := map[string]bool{}
		for _, d := range rc.Detail {
			if d.KodeBrng = strings.TrimSpace(d.KodeBrng); d.KodeBrng == "" || seen[d.KodeBrng] {
				continue
			}
			seen[d.KodeBrng] = true
			if err := a.cek("databarang", "kode_brng", d.KodeBrng, "obat racikan"); err != nil {
				return err
			}
			detail = append(detail, d)
		}
		rc.Detail = detail
		racikan = append(racikan, rc)
	}
	t.Racikan = racikan
	return nil
}

func (a *Action) kode(in []model.Kode, table, column, label string, urut bool) ([]model.Kode, error) {
	out := []model.Kode{}
	seen := map[string]bool{}
	for _, k := range in {
		if k.Kode = strings.TrimSpace(k.Kode); k.Kode == "" || seen[k.Kode] {
			continue
		}
		seen[k.Kode] = true
		if err := a.cek(table, column, k.Kode, label); err != nil {
			return nil, err
		}
		if urut && k.Urut <= 0 {
			k.Urut = len(out) + 1
		}
		out = append(out, k)
	}
	return out, nil
}
