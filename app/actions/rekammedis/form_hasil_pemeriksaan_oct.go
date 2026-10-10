package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanOct hasil pemeriksaan OCT (RMHasilPemeriksaanOCT).
var FormHasilPemeriksaanOct = &Form[model.HasilPemeriksaanOct, request.HasilPemeriksaanOctData]{
	Slug:  "hasil-pemeriksaan-oct",
	Label: "hasil pemeriksaan OCT",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_oct",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanOct, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanOct) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanOct) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanOct) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanOct, d request.HasilPemeriksaanOctData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = support.Nullable(d.DiagnosaKlinis)
		m.KirimanDari = support.Nullable(d.KirimanDari)
		m.HasilPemeriksaan = support.Nullable(d.HasilPemeriksaan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanOct]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanOct) any { return Str(m.KdDokter) }},
	},
}
