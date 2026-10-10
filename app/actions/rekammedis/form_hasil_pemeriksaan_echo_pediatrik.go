package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilPemeriksaanEchoPediatrik hasil pemeriksaan echo pediatrik (RMHasilPemeriksaanEchoPediatrik).
var FormHasilPemeriksaanEchoPediatrik = &Form[model.HasilPemeriksaanEchoPediatrik, request.HasilPemeriksaanEchoPediatrikData]{
	Slug:  "hasil-pemeriksaan-echo-pediatrik",
	Label: "hasil pemeriksaan echo pediatrik",
	Spec: repo.Spec{
		Table:  "hasil_pemeriksaan_echo_pediatrik",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilPemeriksaanEchoPediatrik, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilPemeriksaanEchoPediatrik) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilPemeriksaanEchoPediatrik) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilPemeriksaanEchoPediatrik) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilPemeriksaanEchoPediatrik, d request.HasilPemeriksaanEchoPediatrikData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = support.Nullable(d.DiagnosaKlinis)
		m.KirimanDari = support.Nullable(d.KirimanDari)
		m.Situs = support.Nullable(d.Situs)
		m.AvVa = support.Nullable(d.AvVa)
		m.DrainaseVenaPulmonalis = support.Nullable(d.DrainaseVenaPulmonalis)
		m.KatupMitral = support.Nullable(d.KatupMitral)
		m.KatupAorta = support.Nullable(d.KatupAorta)
		m.KatupTricuspid = support.Nullable(d.KatupTricuspid)
		m.KatupPulmonal = support.Nullable(d.KatupPulmonal)
		m.KatupSeptumAtrium = support.Nullable(d.KatupSeptumAtrium)
		m.KatupSeptumVentrikal = support.Nullable(d.KatupSeptumVentrikal)
		m.KatupArkusAorta = support.Nullable(d.KatupArkusAorta)
		m.KatupKeteranganLainnya = support.Nullable(d.KatupKeteranganLainnya)
		m.RuangJantung = support.Nullable(d.RuangJantung)
		m.ModeIvds = support.Nullable(d.ModeIvds)
		m.ModeIvss = support.Nullable(d.ModeIvss)
		m.ModeLvidDextra = support.Nullable(d.ModeLvidDextra)
		m.ModeLvidSinistra = support.Nullable(d.ModeLvidSinistra)
		m.ModeLvpwDextra = support.Nullable(d.ModeLvpwDextra)
		m.ModeLvpwSinistra = support.Nullable(d.ModeLvpwSinistra)
		m.ModeEjectionFraction = support.Nullable(d.ModeEjectionFraction)
		m.ModeFractionShotening = support.Nullable(d.ModeFractionShotening)
		m.Doppler = support.Nullable(d.Doppler)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.Saran = support.Nullable(d.Saran)
		return nil
	},
	Refs: []Ref[model.HasilPemeriksaanEchoPediatrik]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilPemeriksaanEchoPediatrik) any { return Str(m.KdDokter) }},
	},
}
