package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanUsgNeonatus hasil pemeriksaan USG neonatus (RMHasilPemeriksaanUSGNeonatus).
var FormHasilPemeriksaanUsgNeonatus = &Form[model.HasilPemeriksaanUsgNeonatus, request.HasilPemeriksaanUsgNeonatusData]{
	Slug:  "hasil-pemeriksaan-usg-neonatus",
	Label: "hasil pemeriksaan USG neonatus",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_usg_neonatus",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanUsgNeonatus, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanUsgNeonatus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanUsgNeonatus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanUsgNeonatus) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanUsgNeonatus, d request.HasilPemeriksaanUsgNeonatusData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = support.Nullable(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.VentrikalSinistra = support.Nullable(d.VentrikalSinistra)
		m.VentrikalDextra = support.Nullable(d.VentrikalDextra)
		m.Kesan = support.Nullable(d.Kesan)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.Saran = support.Nullable(d.Saran)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanUsgNeonatus]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanUsgNeonatus) any { return Str(m.KdDokter) }},
	},
}
