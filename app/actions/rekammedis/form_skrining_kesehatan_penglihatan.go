package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKesehatanPenglihatan skrining kesehatan penglihatan (RMSkriningKesehatanPenglihatan).
var FormSkriningKesehatanPenglihatan = &Form[model.SkriningKesehatanPenglihatan, request.SkriningKesehatanPenglihatanData]{
	Slug:  "skrining-kesehatan-penglihatan",
	Label: "skrining kesehatan penglihatan",
	Spec: repo.Spec{
		Table:  "skrining_kesehatan_penglihatan",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKesehatanPenglihatan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKesehatanPenglihatan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKesehatanPenglihatan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKesehatanPenglihatan) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKesehatanPenglihatan, d request.SkriningKesehatanPenglihatanData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.MataLuar = support.Nullable(d.MataLuar)
		m.TajamKiri = support.Nullable(d.TajamKiri)
		m.TajamKanan = support.Nullable(d.TajamKanan)
		m.ButaWarnaKiri = support.Nullable(d.ButaWarnaKiri)
		m.ButaWarnaKanan = support.Nullable(d.ButaWarnaKanan)
		m.Kacamata = support.Nullable(d.Kacamata)
		m.VisusKiri = support.Nullable(d.VisusKiri)
		m.VisusKanan = support.Nullable(d.VisusKanan)
		m.RefraksiKiri = support.Nullable(d.RefraksiKiri)
		m.RefraksiKanan = support.Nullable(d.RefraksiKanan)
		m.RujukRefraksi = support.Nullable(d.RujukRefraksi)
		m.KatarakKiri = support.Nullable(d.KatarakKiri)
		m.KatarakKanan = support.Nullable(d.KatarakKanan)
		m.RujukKatarak = support.Nullable(d.RujukKatarak)
		m.HasilSkrining = strings.TrimSpace(d.HasilSkrining)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningKesehatanPenglihatan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKesehatanPenglihatan) any { return Str(m.Nip) }},
	},
}
