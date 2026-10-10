package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormUjiFungsiKfr uji fungsi KFR (RMUjiFungsiKFR).
var FormUjiFungsiKfr = &Form[model.UjiFungsiKfr, request.UjiFungsiKfrData]{
	Slug:  "uji-fungsi-kfr",
	Label: "uji fungsi KFR",
	Spec: repo.Spec{
		Table:  "uji_fungsi_kfr",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.UjiFungsiKfr, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.UjiFungsiKfr) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.UjiFungsiKfr) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.UjiFungsiKfr) []string { return []string{derefStr(m.KdDokter)} },
	Fill: func(m *model.UjiFungsiKfr, d request.UjiFungsiKfrData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.DiagnosisFungsional = support.Nullable(d.DiagnosisFungsional)
		m.DiagnosisMedis = support.Nullable(d.DiagnosisMedis)
		m.HasilDidapat = support.Nullable(d.HasilDidapat)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.Rekomedasi = support.Nullable(d.Rekomedasi)
		m.KdDokter = support.Nullable(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.UjiFungsiKfr]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.UjiFungsiKfr) any { return StrPtr(m.KdDokter) }},
	},
}
