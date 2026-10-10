package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanUsg hasil pemeriksaan USG (RMHasilPemeriksaanUSG).
var FormHasilPemeriksaanUsg = &Form[model.HasilPemeriksaanUsg, request.HasilPemeriksaanUsgData]{
	Slug:  "hasil-pemeriksaan-usg",
	Label: "hasil pemeriksaan USG",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_usg",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanUsg, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanUsg) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanUsg) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanUsg) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanUsg, d request.HasilPemeriksaanUsgData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = support.Nullable(d.DiagnosaKlinis)
		m.KirimanDari = support.Nullable(d.KirimanDari)
		m.Hta = support.Nullable(d.Hta)
		m.KantongGestasi = support.Nullable(d.KantongGestasi)
		m.UkuranBokongkepala = support.Nullable(d.UkuranBokongkepala)
		m.JenisPrestasi = support.Nullable(d.JenisPrestasi)
		m.DiameterBiparietal = support.Nullable(d.DiameterBiparietal)
		m.PanjangFemur = support.Nullable(d.PanjangFemur)
		m.LingkarAbdomen = support.Nullable(d.LingkarAbdomen)
		m.TafsiranBeratJanin = support.Nullable(d.TafsiranBeratJanin)
		m.UsiaKehamilan = support.Nullable(d.UsiaKehamilan)
		m.PlasentaBerimplatansi = support.Nullable(d.PlasentaBerimplatansi)
		m.DerajatMaturitas = support.Nullable(d.DerajatMaturitas)
		m.JumlahAirKetuban = support.Nullable(d.JumlahAirKetuban)
		m.IndekCairanKetuban = support.Nullable(d.IndekCairanKetuban)
		m.KelainanKongenital = support.Nullable(d.KelainanKongenital)
		m.PeluangSex = support.Nullable(d.PeluangSex)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanUsg]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanUsg) any { return Str(m.KdDokter) }},
	},
}
