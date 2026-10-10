package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKesehatanGigiMulutDewasa skrining kesehatan gigi mulut dewasa (RMSkriningKesehatanGigiMulutDewasa).
var FormSkriningKesehatanGigiMulutDewasa = &Form[model.SkriningKesehatanGigiMulutDewasa, request.SkriningKesehatanGigiMulutDewasaData]{
	Slug:  "skrining-kesehatan-gigi-mulut-dewasa",
	Label: "skrining kesehatan gigi mulut dewasa",
	Spec: repo.Spec{
		Table:  "skrining_kesehatan_gigi_mulut_dewasa",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKesehatanGigiMulutDewasa, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKesehatanGigiMulutDewasa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKesehatanGigiMulutDewasa) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKesehatanGigiMulutDewasa) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKesehatanGigiMulutDewasa, d request.SkriningKesehatanGigiMulutDewasaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KontrolGigi = support.Nullable(d.KontrolGigi)
		m.GigiBungsuTumbuh = support.Nullable(d.GigiBungsuTumbuh)
		m.GigiHilang = support.Nullable(d.GigiHilang)
		m.GigiBerlubang = support.Nullable(d.GigiBerlubang)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningKesehatanGigiMulutDewasa]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKesehatanGigiMulutDewasa) any { return Str(m.Nip) }},
	},
}
