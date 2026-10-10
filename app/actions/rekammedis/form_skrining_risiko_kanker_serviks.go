package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningRisikoKankerServiks skrining risiko kanker serviks (RMSkriningRisikoKankerServiks).
var FormSkriningRisikoKankerServiks = &Form[model.SkriningRisikoKankerServiks, request.SkriningRisikoKankerServiksData]{
	Slug:  "skrining-risiko-kanker-serviks",
	Label: "skrining risiko kanker serviks",
	Spec: repo.Spec{
		Table:  "skrining_risiko_kanker_serviks",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningRisikoKankerServiks, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningRisikoKankerServiks) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningRisikoKankerServiks) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningRisikoKankerServiks) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningRisikoKankerServiks, d request.SkriningRisikoKankerServiksData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.RiwayatPenyakitKeluarga = support.Nullable(d.RiwayatPenyakitKeluarga)
		m.RiwayatPenyakitSendiri = support.Nullable(d.RiwayatPenyakitSendiri)
		m.RisikoMerokok = support.Nullable(d.RisikoMerokok)
		m.RisikoKurangFisik = support.Nullable(d.RisikoKurangFisik)
		m.RisikoGulaBerlebihan = support.Nullable(d.RisikoGulaBerlebihan)
		m.RisikoGaramBerlebihan = support.Nullable(d.RisikoGaramBerlebihan)
		m.RisikoLemakBerlebihan = support.Nullable(d.RisikoLemakBerlebihan)
		m.RisikoKurangBuahSayur = support.Nullable(d.RisikoKurangBuahSayur)
		m.RisikoAlkohol = support.Nullable(d.RisikoAlkohol)
		m.HasilIva = support.Nullable(d.HasilIva)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningRisikoKankerServiks]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningRisikoKankerServiks) any { return Str(m.Nip) }},
	},
}
