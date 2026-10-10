package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormKonselingFarmasi konseling farmasi (RMKonselingFarmasi).
var FormKonselingFarmasi = &Form[model.KonselingFarmasi, request.KonselingFarmasiData]{
	Slug:  "konseling-farmasi",
	Label: "konseling farmasi",
	Spec: repo.Spec{
		Table:  "konseling_farmasi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.KonselingFarmasi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.KonselingFarmasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.KonselingFarmasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.KonselingFarmasi) []string { return []string{m.Nip} },
	Fill: func(m *model.KonselingFarmasi, d request.KonselingFarmasiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Diagnosa = support.Nullable(d.Diagnosa)
		m.ObatPemakaian = support.Nullable(d.ObatPemakaian)
		m.RiwayatAlergi = support.Nullable(d.RiwayatAlergi)
		m.Keluhan = support.Nullable(d.Keluhan)
		m.PernahDatang = support.Nullable(d.PernahDatang)
		m.TindakLanjut = support.Nullable(d.TindakLanjut)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.KonselingFarmasi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.KonselingFarmasi) any { return Str(m.Nip) }},
	},
}
