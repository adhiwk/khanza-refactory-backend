package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKesehatanGigiMulutBalita skrining kesehatan gigi mulut balita (RMSkriningKesehatanGigiMulutBalita).
var FormSkriningKesehatanGigiMulutBalita = &Form[model.SkriningKesehatanGigiMulutBalita, request.SkriningKesehatanGigiMulutBalitaData]{
	Slug:  "skrining-kesehatan-gigi-mulut-balita",
	Label: "skrining kesehatan gigi mulut balita",
	Spec: repo.Spec{
		Table:  "skrining_kesehatan_gigi_mulut_balita",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKesehatanGigiMulutBalita, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKesehatanGigiMulutBalita) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKesehatanGigiMulutBalita) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKesehatanGigiMulutBalita) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKesehatanGigiMulutBalita, d request.SkriningKesehatanGigiMulutBalitaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.PernahPemeriksaanGigimulut = support.Nullable(d.PernahPemeriksaanGigimulut)
		m.SudahTumbuhGigi = support.Nullable(d.SudahTumbuhGigi)
		m.JumlahGigiTumbuh = support.Nullable(d.JumlahGigiTumbuh)
		m.KondisiKebersihanGigimulut = support.Nullable(d.KondisiKebersihanGigimulut)
		m.KebiasaanSusuBotol = support.Nullable(d.KebiasaanSusuBotol)
		m.MengemilManis = support.Nullable(d.MengemilManis)
		m.MenyikatGigiSebelumTidur = support.Nullable(d.MenyikatGigiSebelumTidur)
		m.MengemutMakanan = support.Nullable(d.MengemutMakanan)
		m.LidahKotor = support.Nullable(d.LidahKotor)
		m.CelahBibir = support.Nullable(d.CelahBibir)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningKesehatanGigiMulutBalita]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKesehatanGigiMulutBalita) any { return Str(m.Nip) }},
	},
}
