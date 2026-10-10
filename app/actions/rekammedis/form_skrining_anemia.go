package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningAnemia skrining anemia (RMSkriningAnemia).
var FormSkriningAnemia = &Form[model.SkriningAnemia, request.SkriningAnemiaData]{
	Slug:  "skrining-anemia",
	Label: "skrining anemia",
	Spec: repo.Spec{
		Table:  "skrining_anemia",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningAnemia, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningAnemia) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningAnemia) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningAnemia) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningAnemia, d request.SkriningAnemiaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("kadar_hb", d.KadarHb, ">= 12 g/dl", "11,9 - 11 g/dl", "10.9 - 8 g/dl", "< 8 g/dl"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.MudahLelah = support.Nullable(d.MudahLelah)
		m.BuahSayur = support.Nullable(d.BuahSayur)
		m.ProteinHewani = support.Nullable(d.ProteinHewani)
		m.MasalahPubertas = support.Nullable(d.MasalahPubertas)
		m.RisikoIms = support.Nullable(d.RisikoIms)
		m.KekerasanSeksual = support.Nullable(d.KekerasanSeksual)
		m.SudahMenstruasi = support.Nullable(d.SudahMenstruasi)
		m.GangguanMenstruasi = support.Nullable(d.GangguanMenstruasi)
		m.TambahDarah = support.Nullable(d.TambahDarah)
		m.KelainanDarah = support.Nullable(d.KelainanDarah)
		m.KeluargaThalasemia = support.Nullable(d.KeluargaThalasemia)
		m.Rambut = support.Nullable(d.Rambut)
		m.Kulit = support.Nullable(d.Kulit)
		m.BekasSutikan = support.Nullable(d.BekasSutikan)
		m.Kuku = support.Nullable(d.Kuku)
		m.TandaKlinis = strings.TrimSpace(d.TandaKlinis)
		m.PemeriksaanHb = support.Nullable(d.PemeriksaanHb)
		m.KadarHb = support.Nullable(d.KadarHb)
		m.JenisAnemia = support.Nullable(d.JenisAnemia)
		m.HasilSkrining = strings.TrimSpace(d.HasilSkrining)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningAnemia]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningAnemia) any { return Str(m.Nip) }},
	},
}
