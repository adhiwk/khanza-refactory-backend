package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKesehatanGigiMulutLansia skrining kesehatan gigi mulut lansia (RMSkriningKesehatanGigiMulutLansia).
var FormSkriningKesehatanGigiMulutLansia = &Form[model.SkriningKesehatanGigiMulutLansia, request.SkriningKesehatanGigiMulutLansiaData]{
	Slug:  "skrining-kesehatan-gigi-mulut-lansia",
	Label: "skrining kesehatan gigi mulut lansia",
	Spec: repo.Spec{
		Table:  "skrining_kesehatan_gigi_mulut_lansia",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKesehatanGigiMulutLansia, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKesehatanGigiMulutLansia) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKesehatanGigiMulutLansia) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKesehatanGigiMulutLansia) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKesehatanGigiMulutLansia, d request.SkriningKesehatanGigiMulutLansiaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KontrolGigi = support.Nullable(d.KontrolGigi)
		m.PolaMakan = support.Nullable(d.PolaMakan)
		m.SikatGigi = support.Nullable(d.SikatGigi)
		m.GigiPalsu = support.Nullable(d.GigiPalsu)
		m.GigiBerfungsi = support.Nullable(d.GigiBerfungsi)
		m.MukosaMulut = support.Nullable(d.MukosaMulut)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningKesehatanGigiMulutLansia]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKesehatanGigiMulutLansia) any { return Str(m.Nip) }},
	},
}
