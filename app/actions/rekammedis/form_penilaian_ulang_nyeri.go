package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
)

// FormPenilaianUlangNyeri penilaian ulang nyeri (RMPenilaianUlangNyeri).
var FormPenilaianUlangNyeri = &Form[model.PenilaianUlangNyeri, request.PenilaianUlangNyeriData]{
	Slug:  "penilaian-ulang-nyeri",
	Label: "penilaian ulang nyeri",
	Spec: repo.Spec{
		Table:  "penilaian_ulang_nyeri",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianUlangNyeri, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianUlangNyeri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianUlangNyeri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianUlangNyeri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianUlangNyeri, d request.PenilaianUlangNyeriData) error {
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Provokes = strings.TrimSpace(d.Provokes)
		m.KetProvokes = strings.TrimSpace(d.KetProvokes)
		m.Quality = strings.TrimSpace(d.Quality)
		m.KetQuality = strings.TrimSpace(d.KetQuality)
		m.Lokasi = strings.TrimSpace(d.Lokasi)
		m.Menyebar = strings.TrimSpace(d.Menyebar)
		m.SkalaNyeri = strings.TrimSpace(d.SkalaNyeri)
		m.Durasi = strings.TrimSpace(d.Durasi)
		m.NyeriHilang = strings.TrimSpace(d.NyeriHilang)
		m.KetNyeri = strings.TrimSpace(d.KetNyeri)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianUlangNyeri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianUlangNyeri) any { return Str(m.Nip) }},
	},
}
