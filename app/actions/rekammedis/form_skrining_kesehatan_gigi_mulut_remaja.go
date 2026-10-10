package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKesehatanGigiMulutRemaja skrining kesehatan gigi mulut remaja (RMSkriningKesehatanGigiMulutRemaja).
var FormSkriningKesehatanGigiMulutRemaja = &Form[model.SkriningKesehatanGigiMulutRemaja, request.SkriningKesehatanGigiMulutRemajaData]{
	Slug:  "skrining-kesehatan-gigi-mulut-remaja",
	Label: "skrining kesehatan gigi mulut remaja",
	Spec: repo.Spec{
		Table:  "skrining_kesehatan_gigi_mulut_remaja",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKesehatanGigiMulutRemaja, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKesehatanGigiMulutRemaja) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKesehatanGigiMulutRemaja) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKesehatanGigiMulutRemaja) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKesehatanGigiMulutRemaja, d request.SkriningKesehatanGigiMulutRemajaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.PernahPemeriksaanGigimulut = support.Nullable(d.PernahPemeriksaanGigimulut)
		m.JumlahGigiTumbuh = support.Nullable(d.JumlahGigiTumbuh)
		m.KondisiKebersihanGigimulut = support.Nullable(d.KondisiKebersihanGigimulut)
		m.PunyaGigiBerlubang = support.Nullable(d.PunyaGigiBerlubang)
		m.PernahGusiBerdarah = support.Nullable(d.PernahGusiBerdarah)
		m.PunyaKarangGigi = support.Nullable(d.PunyaKarangGigi)
		m.GigiDepanTidakTeratur = support.Nullable(d.GigiDepanTidakTeratur)
		m.MenyikatGigiSebelumTidur = support.Nullable(d.MenyikatGigiSebelumTidur)
		m.PunyaSariawan = support.Nullable(d.PunyaSariawan)
		m.PemeriksaanFisik = support.Nullable(d.PemeriksaanFisik)
		m.PemeriksaanPenunjang = support.Nullable(d.PemeriksaanPenunjang)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningKesehatanGigiMulutRemaja]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKesehatanGigiMulutRemaja) any { return Str(m.Nip) }},
	},
}
