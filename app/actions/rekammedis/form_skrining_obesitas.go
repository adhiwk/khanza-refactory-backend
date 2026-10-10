package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningObesitas skrining obesitas (RMSkriningObesitas).
var FormSkriningObesitas = &Form[model.SkriningObesitas, request.SkriningObesitasData]{
	Slug:  "skrining-obesitas",
	Label: "skrining obesitas",
	Spec: repo.Spec{
		Table:  "skrining_obesitas",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningObesitas, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningObesitas) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningObesitas) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningObesitas) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningObesitas, d request.SkriningObesitasData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KebiasaanMakanManis = support.Nullable(d.KebiasaanMakanManis)
		m.AktifitasFisikSetiapHari = support.Nullable(d.AktifitasFisikSetiapHari)
		m.IstirahatCukup = support.Nullable(d.IstirahatCukup)
		m.RisikoMerokok = support.Nullable(d.RisikoMerokok)
		m.RiwayatMinumAlkoholMerokokKeluarga = support.Nullable(d.RiwayatMinumAlkoholMerokokKeluarga)
		m.RiwayatPenggunaanObatSteroid = support.Nullable(d.RiwayatPenggunaanObatSteroid)
		m.BeratBadan = support.Nullable(d.BeratBadan)
		m.TinggiBadan = support.Nullable(d.TinggiBadan)
		m.Imt = support.Nullable(d.Imt)
		m.KasifikasiImt = support.Nullable(d.KasifikasiImt)
		m.LingkarPinggang = support.Nullable(d.LingkarPinggang)
		m.RisikoLingkarPinggang = support.Nullable(d.RisikoLingkarPinggang)
		m.StatusObesitas = support.Nullable(d.StatusObesitas)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningObesitas]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningObesitas) any { return Str(m.Nip) }},
	},
}
