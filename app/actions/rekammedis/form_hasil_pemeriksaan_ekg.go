package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanEkg hasil pemeriksaan EKG (RMHasilPemeriksaanEKG).
var FormHasilPemeriksaanEkg = &Form[model.HasilPemeriksaanEkg, request.HasilPemeriksaanEkgData]{
	Slug:  "hasil-pemeriksaan-ekg",
	Label: "hasil pemeriksaan EKG",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_ekg",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanEkg, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanEkg) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanEkg) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanEkg) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanEkg, d request.HasilPemeriksaanEkgData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = support.Nullable(d.DiagnosaKlinis)
		m.KirimanDari = support.Nullable(d.KirimanDari)
		m.Irama = support.Nullable(d.Irama)
		m.LajuJantung = support.Nullable(d.LajuJantung)
		m.Gelombangp = support.Nullable(d.Gelombangp)
		m.Intervalpr = support.Nullable(d.Intervalpr)
		m.Axis = support.Nullable(d.Axis)
		m.Kompleksqrs = support.Nullable(d.Kompleksqrs)
		m.Segmenst = support.Nullable(d.Segmenst)
		m.Gelombangt = support.Nullable(d.Gelombangt)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanEkg]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanEkg) any { return Str(m.KdDokter) }},
	},
}
