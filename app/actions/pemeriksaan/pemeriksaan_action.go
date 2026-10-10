// Package pemeriksaan use case SOAP & tanda vital rawat jalan / rawat inap (tab Pemeriksaan DlgRawatJalan / DlgRawatInap).
package pemeriksaan

import (
	"strings"
	"time"

	request "goravel/app/http/requests/pemeriksaan"
	model "goravel/app/models/perawatan"
	repo "goravel/app/repository/pemeriksaan"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrNotFound           = support.NotFound("data pemeriksaan tidak ditemukan")
	ErrPegawaiNotFound    = support.NotFound("data dokter/paramedis pemeriksa tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("pemeriksaan pada tanggal & jam tersebut sudah ada")
	ErrKosong             = support.Invalid("isi minimal satu data pemeriksaan")
	ErrSebelumRegistrasi  = support.Invalid("waktu pemeriksaan tidak boleh sebelum waktu registrasi")
	ErrKadaluarsa         = support.Forbidden("perubahan data / penghapusan data tidak boleh lebih dari 2 x 24 jam")
)

const batasUbah = 48 * time.Hour

type Action struct {
	rawat model.Rawat
	repo  repo.Repository
	now   func() time.Time
}

func NewAction(rawat model.Rawat, repo repo.Repository) *Action {
	return &Action{rawat: rawat, repo: repo, now: time.Now}
}

func (a *Action) konteks(noRawat string) (*model.Konteks, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(noRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	return k, nil
}

func (a *Action) List(noRawat string) ([]model.Pemeriksaan, error) {
	k, err := a.konteks(noRawat)
	if err != nil {
		return nil, err
	}
	return a.repo.List(a.rawat, k.NoRawat)
}

func (a *Action) Create(req request.StoreRequest) (*model.Pemeriksaan, error) {
	k, err := a.konteks(req.NoRawat)
	if err != nil {
		return nil, err
	}
	tgl, err := support.DateOrToday(req.TglPerawatan)
	if err != nil {
		return nil, err
	}
	jam, err := support.TimeOrNow(req.JamRawat)
	if err != nil {
		return nil, err
	}
	waktu, _ := time.ParseInLocation(support.DateTimeLayout, tgl.Format(support.DateLayout)+" "+jam, time.Local)
	if k.SebelumRegistrasi(waktu) {
		return nil, ErrSebelumRegistrasi
	}

	p := &model.Pemeriksaan{NoRawat: k.NoRawat, TglPerawatan: tgl.Format(support.DateLayout), JamRawat: jam}
	existing, err := a.repo.Find(a.rawat, p.NoRawat, p.TglPerawatan, p.JamRawat)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrAlreadyExists
	}
	if err := a.fill(p, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(a.rawat, p); err != nil {
		return nil, err
	}
	return a.repo.Find(a.rawat, p.NoRawat, p.TglPerawatan, p.JamRawat)
}

func (a *Action) Update(noRawat, tgl, jam string, d request.Data, superAdmin bool) (*model.Pemeriksaan, error) {
	p, err := a.editable(noRawat, tgl, jam, superAdmin)
	if err != nil {
		return nil, err
	}
	if err := a.fill(p, d); err != nil {
		return nil, err
	}
	if err := a.repo.Update(a.rawat, p); err != nil {
		return nil, err
	}
	return a.repo.Find(a.rawat, p.NoRawat, p.TglPerawatan, p.JamRawat)
}

func (a *Action) Delete(noRawat, tgl, jam string, superAdmin bool) error {
	p, err := a.editable(noRawat, tgl, jam, superAdmin)
	if err != nil {
		return err
	}
	return a.repo.Delete(a.rawat, p.NoRawat, p.TglPerawatan, p.JamRawat)
}

func (a *Action) editable(noRawat, tgl, jam string, superAdmin bool) (*model.Pemeriksaan, error) {
	p, err := a.repo.Find(a.rawat, strings.TrimSpace(noRawat), strings.TrimSpace(tgl), strings.TrimSpace(jam))
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, ErrNotFound
	}
	if !superAdmin {
		waktu, err := time.ParseInLocation(support.DateTimeLayout, p.TglPerawatan+" "+p.JamRawat, time.Local)
		if err == nil && a.now().Sub(waktu) > batasUbah {
			return nil, ErrKadaluarsa
		}
	}
	return p, nil
}

func (a *Action) fill(p *model.Pemeriksaan, d request.Data) error {
	nip := strings.TrimSpace(d.Nip)
	ok, err := a.repo.PegawaiExists(nip)
	if err != nil {
		return err
	}
	if !ok {
		return ErrPegawaiNotFound
	}

	p.SuhuTubuh = support.Nullable(d.SuhuTubuh)
	p.Tensi = strings.TrimSpace(d.Tensi)
	p.Nadi = support.Nullable(d.Nadi)
	p.Respirasi = support.Nullable(d.Respirasi)
	p.Tinggi = support.Nullable(d.Tinggi)
	p.Berat = support.Nullable(d.Berat)
	p.Spo2 = strings.TrimSpace(d.Spo2)
	p.Gcs = support.Nullable(d.Gcs)
	p.Kesadaran = support.OrDefault(d.Kesadaran, "Compos Mentis")
	p.Keluhan = support.Nullable(d.Keluhan)
	p.Pemeriksaan = support.Nullable(d.Pemeriksaan)
	p.Alergi = support.Nullable(d.Alergi)
	p.LingkarPerut = nil
	if a.rawat == model.Ralan {
		p.LingkarPerut = support.Nullable(d.LingkarPerut)
	}
	p.Rtl = strings.TrimSpace(d.Rtl)
	p.Penilaian = strings.TrimSpace(d.Penilaian)
	p.Instruksi = strings.TrimSpace(d.Instruksi)
	p.Evaluasi = strings.TrimSpace(d.Evaluasi)
	p.Nip = nip
	if p.Kosong() {
		return ErrKosong
	}
	return nil
}
