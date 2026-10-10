package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanUsgGynecologi hasil pemeriksaan USG gynecologi (RMHasilPemeriksaanUSGGynecologi).
var FormHasilPemeriksaanUsgGynecologi = &Form[model.HasilPemeriksaanUsgGynecologi, request.HasilPemeriksaanUsgGynecologiData]{
	Slug:  "hasil-pemeriksaan-usg-gynecologi",
	Label: "hasil pemeriksaan USG gynecologi",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_usg_gynecologi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanUsgGynecologi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanUsgGynecologi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanUsgGynecologi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanUsgGynecologi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanUsgGynecologi, d request.HasilPemeriksaanUsgGynecologiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = strings.TrimSpace(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.Uterus = support.Nullable(d.Uterus)
		m.Parametrium = support.Nullable(d.Parametrium)
		m.Ovarium = support.Nullable(d.Ovarium)
		m.Doppler = support.Nullable(d.Doppler)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanUsgGynecologi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanUsgGynecologi) any { return Str(m.KdDokter) }},
	},
}
