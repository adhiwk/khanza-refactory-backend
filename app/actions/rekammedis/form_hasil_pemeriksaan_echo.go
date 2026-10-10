package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanEcho hasil pemeriksaan echo (RMHasilPemeriksaanEcho).
var FormHasilPemeriksaanEcho = &Form[model.HasilPemeriksaanEcho, request.HasilPemeriksaanEchoData]{
	Slug:  "hasil-pemeriksaan-echo",
	Label: "hasil pemeriksaan echo",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_echo",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanEcho, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanEcho) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanEcho) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanEcho) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanEcho, d request.HasilPemeriksaanEchoData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Sistolik = support.Nullable(d.Sistolik)
		m.Diastolic = support.Nullable(d.Diastolic)
		m.Kontraktilitas = support.Nullable(d.Kontraktilitas)
		m.DimensiRuang = support.Nullable(d.DimensiRuang)
		m.Katup = support.Nullable(d.Katup)
		m.AnalisaSegmental = support.Nullable(d.AnalisaSegmental)
		m.Erap = support.Nullable(d.Erap)
		m.LainLain = support.Nullable(d.LainLain)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanEcho]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanEcho) any { return Str(m.KdDokter) }},
	},
}
