package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianDehidrasi penilaian derajat dehidrasi (RMPenilaianDerajatDehidrasi).
var FormPenilaianDehidrasi = &Form[model.PenilaianDehidrasi, request.PenilaianDehidrasiData]{
	Slug:  "penilaian-dehidrasi",
	Label: "penilaian derajat dehidrasi",
	Spec: repo.Spec{
		Table:  "penilaian_dehidrasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianDehidrasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianDehidrasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianDehidrasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianDehidrasi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianDehidrasi, d request.PenilaianDehidrasiData) error {
		if err := Enum("penilaian1", d.Penilaian1, "Baik", "Lesu/Haus", "Gelisah, Haus, Mengantuk, Hingga Syok"); err != nil {
			return err
		}
		m.Penilaian1 = support.Nullable(d.Penilaian1)
		m.PenilaianNilai1 = d.PenilaianNilai1
		m.Penilaian2 = support.Nullable(d.Penilaian2)
		m.PenilaianNilai2 = d.PenilaianNilai2
		m.Penilaian3 = support.Nullable(d.Penilaian3)
		m.PenilaianNilai3 = d.PenilaianNilai3
		m.Penilaian4 = support.Nullable(d.Penilaian4)
		m.PenilaianNilai4 = d.PenilaianNilai4
		m.Penilaian5 = support.Nullable(d.Penilaian5)
		m.PenilaianNilai5 = d.PenilaianNilai5
		m.Penilaian6 = support.Nullable(d.Penilaian6)
		m.PenilaianNilai6 = d.PenilaianNilai6
		m.PenilaianTotalnilai = d.PenilaianTotalnilai
		m.HasilPenilaian = support.Nullable(d.HasilPenilaian)
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.PenilaianDehidrasi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianDehidrasi) any { return Str(m.KdDokter) }},
	},
}
