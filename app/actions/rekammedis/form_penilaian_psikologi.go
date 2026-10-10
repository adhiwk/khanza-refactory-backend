package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPsikologi penilaian psikologi (RMPenilaianPsikologi).
var FormPenilaianPsikologi = &Form[model.PenilaianPsikologi, request.PenilaianPsikologiData]{
	Slug:  "penilaian-psikologi",
	Label: "penilaian psikologi",
	Spec: repo.Spec{
		Table:  "penilaian_psikologi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianPsikologi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianPsikologi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianPsikologi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPsikologi) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianPsikologi, d request.PenilaianPsikologiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.DikirimDari = strings.TrimSpace(d.DikirimDari)
		m.TujuanPemeriksaan = strings.TrimSpace(d.TujuanPemeriksaan)
		m.KetAnamnesis = strings.TrimSpace(d.KetAnamnesis)
		m.Rupa = strings.TrimSpace(d.Rupa)
		m.BentukTubuh = strings.TrimSpace(d.BentukTubuh)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Pakaian = strings.TrimSpace(d.Pakaian)
		m.Ekspresi = strings.TrimSpace(d.Ekspresi)
		m.Berbicara = strings.TrimSpace(d.Berbicara)
		m.PenggunaanKata = strings.TrimSpace(d.PenggunaanKata)
		m.CiriMenyolok = strings.TrimSpace(d.CiriMenyolok)
		m.HasilPsikotes = strings.TrimSpace(d.HasilPsikotes)
		m.Kepribadian = strings.TrimSpace(d.Kepribadian)
		m.Psikodinamika = strings.TrimSpace(d.Psikodinamika)
		m.KesimpulanPsikolog = strings.TrimSpace(d.KesimpulanPsikolog)
		return nil
	},
	Refs: []Ref[model.PenilaianPsikologi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianPsikologi) any { return Str(m.Nip) }},
	},
}
