package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanSlitLamp hasil pemeriksaan slit lamp (RMHasilPemeriksaanSlitLamp).
var FormHasilPemeriksaanSlitLamp = &Form[model.HasilPemeriksaanSlitLamp, request.HasilPemeriksaanSlitLampData]{
	Slug:  "hasil-pemeriksaan-slit-lamp",
	Label: "hasil pemeriksaan slit lamp",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_slit_lamp",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanSlitLamp, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanSlitLamp) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanSlitLamp) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanSlitLamp) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanSlitLamp, d request.HasilPemeriksaanSlitLampData) error {
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
	Refs: []Ref[model.HasilPemeriksaanSlitLamp]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanSlitLamp) any { return Str(m.KdDokter) }},
	},
}
