package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilEndoskopiHidung hasil endoskopi hidung (RMHasilEndoskopiHidung).
var FormHasilEndoskopiHidung = &Form[model.HasilEndoskopiHidung, request.HasilEndoskopiHidungData]{
	Slug:  "hasil-endoskopi-hidung",
	Label: "hasil endoskopi hidung",
	Spec: repo.Spec{
		Table:  "hasil_endoskopi_hidung",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilEndoskopiHidung, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilEndoskopiHidung) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilEndoskopiHidung) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilEndoskopiHidung) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilEndoskopiHidung, d request.HasilEndoskopiHidungData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = strings.TrimSpace(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.KondisiHidungKanan = support.Nullable(d.KondisiHidungKanan)
		m.KondisiHidungKiri = support.Nullable(d.KondisiHidungKiri)
		m.KavumNasiKanan = support.Nullable(d.KavumNasiKanan)
		m.KavumNasiKiri = support.Nullable(d.KavumNasiKiri)
		m.KonkaInferiorKanan = support.Nullable(d.KonkaInferiorKanan)
		m.KonkaInferiorKiri = support.Nullable(d.KonkaInferiorKiri)
		m.MeatusMediusKanan = support.Nullable(d.MeatusMediusKanan)
		m.MeatusMediusKiri = support.Nullable(d.MeatusMediusKiri)
		m.SeptumKanan = support.Nullable(d.SeptumKanan)
		m.SeptumKiri = support.Nullable(d.SeptumKiri)
		m.NasofaringKanan = support.Nullable(d.NasofaringKanan)
		m.NasofaringKiri = support.Nullable(d.NasofaringKiri)
		m.LainlainKanan = support.Nullable(d.LainlainKanan)
		m.LainlainKiri = support.Nullable(d.LainlainKiri)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilEndoskopiHidung]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilEndoskopiHidung) any { return Str(m.KdDokter) }},
	},
}
