package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningPuma skrining PUMA (RMSkriningPUMA).
var FormSkriningPuma = &Form[model.SkriningPuma, request.SkriningPumaData]{
	Slug:  "skrining-puma",
	Label: "skrining PUMA",
	Spec: repo.Spec{
		Table:  "skrining_puma",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningPuma, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningPuma) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningPuma) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningPuma) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningPuma, d request.SkriningPumaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Jk = support.Nullable(d.Jk)
		m.NilaiJk = d.NilaiJk
		m.Usia = support.Nullable(d.Usia)
		m.NilaiUsia = d.NilaiUsia
		m.PernahMerokok = support.Nullable(d.PernahMerokok)
		m.NilaiPernahMerokok = d.NilaiPernahMerokok
		m.JumlahRokokPerhari = support.Nullable(d.JumlahRokokPerhari)
		m.LamaMerokok = support.Nullable(d.LamaMerokok)
		m.NapasPendek = support.Nullable(d.NapasPendek)
		m.NilaiNapasPendek = d.NilaiNapasPendek
		m.PunyaDahak = support.Nullable(d.PunyaDahak)
		m.NilaiPunyaDahak = d.NilaiPunyaDahak
		m.BiasaBatuk = support.Nullable(d.BiasaBatuk)
		m.NilaiBiasaBatuk = d.NilaiBiasaBatuk
		m.Spirometri = support.Nullable(d.Spirometri)
		m.NilaiSpirometri = d.NilaiSpirometri
		m.NilaiTotal = d.NilaiTotal
		m.KeteranganHasilSkrining = support.Nullable(d.KeteranganHasilSkrining)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningPuma]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningPuma) any { return Str(m.Nip) }},
	},
}
