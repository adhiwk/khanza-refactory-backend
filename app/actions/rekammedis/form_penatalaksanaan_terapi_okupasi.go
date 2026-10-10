package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenatalaksanaanTerapiOkupasi penatalaksanaan terapi okupasi (RMPenatalaksanaanTerapiOkupasi).
var FormPenatalaksanaanTerapiOkupasi = &Form[model.PenatalaksanaanTerapiOkupasi, request.PenatalaksanaanTerapiOkupasiData]{
	Slug:  "penatalaksanaan-terapi-okupasi",
	Label: "penatalaksanaan terapi okupasi",
	Spec: repo.Spec{
		Table:  "penatalaksanaan_terapi_okupasi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenatalaksanaanTerapiOkupasi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenatalaksanaanTerapiOkupasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenatalaksanaanTerapiOkupasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenatalaksanaanTerapiOkupasi) []string { return []string{m.Nip} },
	Fill: func(m *model.PenatalaksanaanTerapiOkupasi, d request.PenatalaksanaanTerapiOkupasiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rps = strings.TrimSpace(d.Rps)
		m.AnamnesaGeneral = strings.TrimSpace(d.AnamnesaGeneral)
		m.TandaVital = strings.TrimSpace(d.TandaVital)
		m.PemeriksaanPenunjang = strings.TrimSpace(d.PemeriksaanPenunjang)
		m.Spesialisasi = strings.TrimSpace(d.Spesialisasi)
		m.KeteranganSpesialisasi = strings.TrimSpace(d.KeteranganSpesialisasi)
		m.PemeriksaanOkupasiTerapi = strings.TrimSpace(d.PemeriksaanOkupasiTerapi)
		m.Aset = strings.TrimSpace(d.Aset)
		m.Limitasi = strings.TrimSpace(d.Limitasi)
		m.DiagnosaTerapiOkupasi = strings.TrimSpace(d.DiagnosaTerapiOkupasi)
		m.RencanaIntervensi = strings.TrimSpace(d.RencanaIntervensi)
		return nil
	},
	Refs: []Ref[model.PenatalaksanaanTerapiOkupasi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenatalaksanaanTerapiOkupasi) any { return Str(m.Nip) }},
	},
}
