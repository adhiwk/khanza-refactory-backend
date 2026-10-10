package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningTbc skrining TBC (RMSkriningTBC).
var FormSkriningTbc = &Form[model.SkriningTbc, request.SkriningTbcData]{
	Slug:  "skrining-tbc",
	Label: "skrining TBC",
	Spec: repo.Spec{
		Table:  "skrining_tbc",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningTbc, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningTbc) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningTbc) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningTbc) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningTbc, d request.SkriningTbcData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.BeratBadan = support.Nullable(d.BeratBadan)
		m.TinggiBadan = support.Nullable(d.TinggiBadan)
		m.Imt = support.Nullable(d.Imt)
		m.KasifikasiImt = support.Nullable(d.KasifikasiImt)
		m.LingkarPinggang = support.Nullable(d.LingkarPinggang)
		m.RisikoLingkarPinggang = support.Nullable(d.RisikoLingkarPinggang)
		m.RiwayatKontakTbc = support.Nullable(d.RiwayatKontakTbc)
		m.JenisKontakTbc = support.Nullable(d.JenisKontakTbc)
		m.FaktorResikoPernahTerdiagnosaTbc = support.Nullable(d.FaktorResikoPernahTerdiagnosaTbc)
		m.KeteranganPernahTerdiagnosa = support.Nullable(d.KeteranganPernahTerdiagnosa)
		m.FaktorResikoPernahBerobatTbc = support.Nullable(d.FaktorResikoPernahBerobatTbc)
		m.FaktorResikoMalnutrisi = support.Nullable(d.FaktorResikoMalnutrisi)
		m.FaktorResikoMerokok = support.Nullable(d.FaktorResikoMerokok)
		m.FaktorResikoRiwayatDm = support.Nullable(d.FaktorResikoRiwayatDm)
		m.FaktorResikoOdhiv = support.Nullable(d.FaktorResikoOdhiv)
		m.FaktorResikoLansia = support.Nullable(d.FaktorResikoLansia)
		m.FaktorResikoIbuHamil = support.Nullable(d.FaktorResikoIbuHamil)
		m.FaktorResikoWbp = support.Nullable(d.FaktorResikoWbp)
		m.FaktorResikoTinggalDiwilayahPadatKumuh = support.Nullable(d.FaktorResikoTinggalDiwilayahPadatKumuh)
		m.AbnormalitasTbc = support.Nullable(d.AbnormalitasTbc)
		m.GejalaTbcBatuk = support.Nullable(d.GejalaTbcBatuk)
		m.GejalaTbcBbTurun = support.Nullable(d.GejalaTbcBbTurun)
		m.GejalaTbcDemam = support.Nullable(d.GejalaTbcDemam)
		m.GejalaTbcBerkeringatMalamHari = support.Nullable(d.GejalaTbcBerkeringatMalamHari)
		m.KeteranganGejalaPenyakitLain = support.Nullable(d.KeteranganGejalaPenyakitLain)
		m.KesimpulanSkrining = support.Nullable(d.KesimpulanSkrining)
		m.KeteranganHasilSkrining = support.Nullable(d.KeteranganHasilSkrining)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningTbc]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningTbc) any { return Str(m.Nip) }},
	},
}
