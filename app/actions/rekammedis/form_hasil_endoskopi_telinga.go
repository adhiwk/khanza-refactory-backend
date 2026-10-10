package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilEndoskopiTelinga hasil endoskopi telinga (RMHasilEndoskopiTelinga).
var FormHasilEndoskopiTelinga = &Form[model.HasilEndoskopiTelinga, request.HasilEndoskopiTelingaData]{
	Slug:  "hasil-endoskopi-telinga",
	Label: "hasil endoskopi telinga",
	Spec: repo.Spec{
		Table:  "hasil_endoskopi_telinga",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilEndoskopiTelinga, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilEndoskopiTelinga) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilEndoskopiTelinga) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilEndoskopiTelinga) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilEndoskopiTelinga, d request.HasilEndoskopiTelingaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = strings.TrimSpace(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.BentukLiangTelingaKanan = support.Nullable(d.BentukLiangTelingaKanan)
		m.BentukLiangTelingaKiri = support.Nullable(d.BentukLiangTelingaKiri)
		m.KondisiLiangTelingaKanan = support.Nullable(d.KondisiLiangTelingaKanan)
		m.KeteranganKondisiLiangTelingaKanan = support.Nullable(d.KeteranganKondisiLiangTelingaKanan)
		m.KondisiLiangTelingaKiri = support.Nullable(d.KondisiLiangTelingaKiri)
		m.KeteranganKondisiLiangTelingaKiri = support.Nullable(d.KeteranganKondisiLiangTelingaKiri)
		m.MembranTimpaniIntakKanan = support.Nullable(d.MembranTimpaniIntakKanan)
		m.MembranTimpaniIntakKiri = support.Nullable(d.MembranTimpaniIntakKiri)
		m.MembranTimpaniPerforasiKanan = support.Nullable(d.MembranTimpaniPerforasiKanan)
		m.KeteranganMembranTimpaniPerforasiKanan = support.Nullable(d.KeteranganMembranTimpaniPerforasiKanan)
		m.MembranTimpaniPerforasiKiri = support.Nullable(d.MembranTimpaniPerforasiKiri)
		m.KeteranganMembranTimpaniPerforasiKiri = support.Nullable(d.KeteranganMembranTimpaniPerforasiKiri)
		m.KavumTimpaniMukosaKanan = support.Nullable(d.KavumTimpaniMukosaKanan)
		m.KavumTimpaniMukosaKiri = support.Nullable(d.KavumTimpaniMukosaKiri)
		m.KavumTimpaniOsikelKanan = support.Nullable(d.KavumTimpaniOsikelKanan)
		m.KavumTimpaniOsikelKiri = support.Nullable(d.KavumTimpaniOsikelKiri)
		m.KavumTimpaniIsthmusKanan = support.Nullable(d.KavumTimpaniIsthmusKanan)
		m.KavumTimpaniIsthmusKiri = support.Nullable(d.KavumTimpaniIsthmusKiri)
		m.KavumTimpaniAnteriorKanan = support.Nullable(d.KavumTimpaniAnteriorKanan)
		m.KavumTimpaniAnteriorKiri = support.Nullable(d.KavumTimpaniAnteriorKiri)
		m.KavumTimpaniPosteriorKanan = support.Nullable(d.KavumTimpaniPosteriorKanan)
		m.KavumTimpaniPosteriorKiri = support.Nullable(d.KavumTimpaniPosteriorKiri)
		m.LainlainKanan = support.Nullable(d.LainlainKanan)
		m.LainlainKiri = support.Nullable(d.LainlainKiri)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.Anjuran = support.Nullable(d.Anjuran)
		return nil
	},
	Refs: []Ref[model.HasilEndoskopiTelinga]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilEndoskopiTelinga) any { return Str(m.KdDokter) }},
	},
}
