package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianLanjutanResikoJatuhGeriatri penilaian lanjutan risiko jatuh geriatri (RMPenilaianLanjutanRisikoJatuhGeriatri).
var FormPenilaianLanjutanResikoJatuhGeriatri = &Form[model.PenilaianLanjutanResikoJatuhGeriatri, request.PenilaianLanjutanResikoJatuhGeriatriData]{
	Slug:  "penilaian-lanjutan-resiko-jatuh-geriatri",
	Label: "penilaian lanjutan risiko jatuh geriatri",
	Spec: repo.Spec{
		Table:  "penilaian_lanjutan_resiko_jatuh_geriatri",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianLanjutanResikoJatuhGeriatri, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianLanjutanResikoJatuhGeriatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianLanjutanResikoJatuhGeriatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianLanjutanResikoJatuhGeriatri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianLanjutanResikoJatuhGeriatri, d request.PenilaianLanjutanResikoJatuhGeriatriData) error {
		m.PenilaianJatuhSkala1 = support.Nullable(d.PenilaianJatuhSkala1)
		m.PenilaianJatuhNilai1 = d.PenilaianJatuhNilai1
		m.PenilaianJatuhSkala2 = support.Nullable(d.PenilaianJatuhSkala2)
		m.PenilaianJatuhNilai2 = d.PenilaianJatuhNilai2
		m.PenilaianJatuhSkala3 = support.Nullable(d.PenilaianJatuhSkala3)
		m.PenilaianJatuhNilai3 = d.PenilaianJatuhNilai3
		m.PenilaianJatuhSkala4 = support.Nullable(d.PenilaianJatuhSkala4)
		m.PenilaianJatuhNilai4 = d.PenilaianJatuhNilai4
		m.PenilaianJatuhSkala5 = support.Nullable(d.PenilaianJatuhSkala5)
		m.PenilaianJatuhNilai5 = d.PenilaianJatuhNilai5
		m.PenilaianJatuhSkala6 = support.Nullable(d.PenilaianJatuhSkala6)
		m.PenilaianJatuhNilai6 = d.PenilaianJatuhNilai6
		m.PenilaianJatuhSkala7 = support.Nullable(d.PenilaianJatuhSkala7)
		m.PenilaianJatuhNilai7 = d.PenilaianJatuhNilai7
		m.PenilaianJatuhSkala8 = support.Nullable(d.PenilaianJatuhSkala8)
		m.PenilaianJatuhNilai8 = d.PenilaianJatuhNilai8
		m.PenilaianJatuhSkala9 = support.Nullable(d.PenilaianJatuhSkala9)
		m.PenilaianJatuhNilai9 = d.PenilaianJatuhNilai9
		m.PenilaianJatuhSkala10 = support.Nullable(d.PenilaianJatuhSkala10)
		m.PenilaianJatuhNilai10 = d.PenilaianJatuhNilai10
		m.PenilaianJatuhSkala11 = support.Nullable(d.PenilaianJatuhSkala11)
		m.PenilaianJatuhNilai11 = d.PenilaianJatuhNilai11
		m.PenilaianJatuhTotalnilai = d.PenilaianJatuhTotalnilai
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Saran = support.Nullable(d.Saran)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianLanjutanResikoJatuhGeriatri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianLanjutanResikoJatuhGeriatri) any { return Str(m.Nip) }},
	},
}
