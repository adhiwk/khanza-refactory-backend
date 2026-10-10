package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanUsgAbdomen hasil pemeriksaan USG abdomen (RMHasilPemeriksaanUSGAbdomen).
var FormHasilPemeriksaanUsgAbdomen = &Form[model.HasilPemeriksaanUsgAbdomen, request.HasilPemeriksaanUsgAbdomenData]{
	Slug:  "hasil-pemeriksaan-usg-abdomen",
	Label: "hasil pemeriksaan USG abdomen",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_usg_abdomen",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanUsgAbdomen, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanUsgAbdomen) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanUsgAbdomen) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanUsgAbdomen) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanUsgAbdomen, d request.HasilPemeriksaanUsgAbdomenData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = strings.TrimSpace(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.Esofagus = support.Nullable(d.Esofagus)
		m.Colon = support.Nullable(d.Colon)
		m.Gaster = support.Nullable(d.Gaster)
		m.Hepar = support.Nullable(d.Hepar)
		m.GallBlader = support.Nullable(d.GallBlader)
		m.Lien = support.Nullable(d.Lien)
		m.Pancreas = support.Nullable(d.Pancreas)
		m.GinjalDextra = support.Nullable(d.GinjalDextra)
		m.GinjalSinistra = support.Nullable(d.GinjalSinistra)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanUsgAbdomen]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanUsgAbdomen) any { return Str(m.KdDokter) }},
	},
}
