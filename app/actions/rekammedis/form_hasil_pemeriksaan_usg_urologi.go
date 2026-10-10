package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanUsgUrologi hasil pemeriksaan USG urologi (RMHasilPemeriksaanUSGUrologi).
var FormHasilPemeriksaanUsgUrologi = &Form[model.HasilPemeriksaanUsgUrologi, request.HasilPemeriksaanUsgUrologiData]{
	Slug:  "hasil-pemeriksaan-usg-urologi",
	Label: "hasil pemeriksaan USG urologi",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_usg_urologi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanUsgUrologi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanUsgUrologi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanUsgUrologi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanUsgUrologi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanUsgUrologi, d request.HasilPemeriksaanUsgUrologiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = strings.TrimSpace(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.GinjalKanan = support.Nullable(d.GinjalKanan)
		m.GinjalKiri = support.Nullable(d.GinjalKiri)
		m.VesicaUrinaria = support.Nullable(d.VesicaUrinaria)
		m.Tambahan = support.Nullable(d.Tambahan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanUsgUrologi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanUsgUrologi) any { return Str(m.KdDokter) }},
	},
}
