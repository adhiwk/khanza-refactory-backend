package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanTreadmill hasil pemeriksaan treadmill (RMHasilPemeriksaanTreadmill).
var FormHasilPemeriksaanTreadmill = &Form[model.HasilPemeriksaanTreadmill, request.HasilPemeriksaanTreadmillData]{
	Slug:  "hasil-pemeriksaan-treadmill",
	Label: "hasil pemeriksaan treadmill",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_treadmill",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanTreadmill, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanTreadmill) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanTreadmill) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanTreadmill) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanTreadmill, d request.HasilPemeriksaanTreadmillData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.KirimanDari = support.Nullable(d.KirimanDari)
		m.DiagnosaKlinis = support.Nullable(d.DiagnosaKlinis)
		m.Protokol = support.Nullable(d.Protokol)
		m.KeteranganProtokol = support.Nullable(d.KeteranganProtokol)
		m.TdAwal = support.Nullable(d.TdAwal)
		m.NadiAwal = support.Nullable(d.NadiAwal)
		m.DenyutJantungMaksimal = support.Nullable(d.DenyutJantungMaksimal)
		m.HasilPemeriksaan = support.Nullable(d.HasilPemeriksaan)
		m.TemuanEkg = support.Nullable(d.TemuanEkg)
		m.KapasitasFungsional = support.Nullable(d.KapasitasFungsional)
		m.Interpretasi = support.Nullable(d.Interpretasi)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanTreadmill]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanTreadmill) any { return Str(m.KdDokter) }},
	},
}
