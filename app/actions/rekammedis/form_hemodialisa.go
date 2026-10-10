package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHemodialisa hemodialisa (RMHemodialisa).
var FormHemodialisa = &Form[model.Hemodialisa, request.HemodialisaData]{
	Slug:  "hemodialisa",
	Label: "hemodialisa",
	Spec: repo.Spec{
		Table:  "hemodialisa",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.Hemodialisa, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.Hemodialisa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.Hemodialisa) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.Hemodialisa) []string { return []string{derefStr(m.KdDokter)} },
	Fill: func(m *model.Hemodialisa, d request.HemodialisaData) error {
		m.KdDokter = support.Nullable(d.KdDokter)
		m.Lama = support.Nullable(d.Lama)
		m.Akses = support.Nullable(d.Akses)
		m.Dialist = support.Nullable(d.Dialist)
		m.Transfusi = support.Nullable(d.Transfusi)
		m.Penarikan = support.Nullable(d.Penarikan)
		m.Qb = support.Nullable(d.Qb)
		m.Qd = support.Nullable(d.Qd)
		m.Ureum = support.Nullable(d.Ureum)
		m.Hb = support.Nullable(d.Hb)
		m.Hbsag = support.Nullable(d.Hbsag)
		m.Creatinin = support.Nullable(d.Creatinin)
		m.Hiv = support.Nullable(d.Hiv)
		m.Hcv = support.Nullable(d.Hcv)
		m.Lain = support.Nullable(d.Lain)
		m.KdPenyakit = support.Nullable(d.KdPenyakit)
		return nil
	},
	Refs: []Ref[model.Hemodialisa]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.Hemodialisa) any { return StrPtr(m.KdDokter) }},
		{Column: "kd_penyakit", Table: "penyakit", RefCol: "kd_penyakit", Label: "penyakit", Value: func(m *model.Hemodialisa) any { return StrPtr(m.KdPenyakit) }},
	},
}
