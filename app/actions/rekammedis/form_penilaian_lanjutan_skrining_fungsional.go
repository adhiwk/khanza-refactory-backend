package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianLanjutanSkriningFungsional penilaian lanjutan skrining fungsional (RMPenilaianLanjutanSkriningFungsional).
var FormPenilaianLanjutanSkriningFungsional = &Form[model.PenilaianLanjutanSkriningFungsional, request.PenilaianLanjutanSkriningFungsionalData]{
	Slug:  "penilaian-lanjutan-skrining-fungsional",
	Label: "penilaian lanjutan skrining fungsional",
	Spec: repo.Spec{
		Table:  "penilaian_lanjutan_skrining_fungsional",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianLanjutanSkriningFungsional, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianLanjutanSkriningFungsional) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianLanjutanSkriningFungsional) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianLanjutanSkriningFungsional) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianLanjutanSkriningFungsional, d request.PenilaianLanjutanSkriningFungsionalData) error {
		m.PenilaianSkriningSkala1 = support.Nullable(d.PenilaianSkriningSkala1)
		m.PenilaianSkriningNilai1 = d.PenilaianSkriningNilai1
		m.PenilaianSkriningSkala2 = support.Nullable(d.PenilaianSkriningSkala2)
		m.PenilaianSkriningNilai2 = d.PenilaianSkriningNilai2
		m.PenilaianSkriningSkala3 = support.Nullable(d.PenilaianSkriningSkala3)
		m.PenilaianSkriningNilai3 = d.PenilaianSkriningNilai3
		m.PenilaianSkriningSkala4 = support.Nullable(d.PenilaianSkriningSkala4)
		m.PenilaianSkriningNilai4 = d.PenilaianSkriningNilai4
		m.PenilaianSkriningSkala5 = support.Nullable(d.PenilaianSkriningSkala5)
		m.PenilaianSkriningNilai5 = d.PenilaianSkriningNilai5
		m.PenilaianSkriningSkala6 = support.Nullable(d.PenilaianSkriningSkala6)
		m.PenilaianSkriningNilai6 = d.PenilaianSkriningNilai6
		m.PenilaianSkriningSkala7 = support.Nullable(d.PenilaianSkriningSkala7)
		m.PenilaianSkriningNilai7 = d.PenilaianSkriningNilai7
		m.PenilaianSkriningSkala8 = support.Nullable(d.PenilaianSkriningSkala8)
		m.PenilaianSkriningNilai8 = d.PenilaianSkriningNilai8
		m.PenilaianSkriningSkala9 = support.Nullable(d.PenilaianSkriningSkala9)
		m.PenilaianSkriningNilai9 = d.PenilaianSkriningNilai9
		m.PenilaianSkriningSkala10 = support.Nullable(d.PenilaianSkriningSkala10)
		m.PenilaianSkriningNilai10 = d.PenilaianSkriningNilai10
		m.PenilaianSkriningTotalnilai = d.PenilaianSkriningTotalnilai
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianLanjutanSkriningFungsional]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianLanjutanSkriningFungsional) any { return Str(m.Nip) }},
	},
}
