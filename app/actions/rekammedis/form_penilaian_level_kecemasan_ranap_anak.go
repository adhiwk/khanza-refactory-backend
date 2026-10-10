package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianLevelKecemasanRanapAnak penilaian level kecemasan ranap anak (RMPenilaianLevelKecemasanRanapAnak).
var FormPenilaianLevelKecemasanRanapAnak = &Form[model.PenilaianLevelKecemasanRanapAnak, request.PenilaianLevelKecemasanRanapAnakData]{
	Slug:  "penilaian-level-kecemasan-ranap-anak",
	Label: "penilaian level kecemasan ranap anak",
	Spec: repo.Spec{
		Table:  "penilaian_level_kecemasan_ranap_anak",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianLevelKecemasanRanapAnak, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianLevelKecemasanRanapAnak) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianLevelKecemasanRanapAnak) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianLevelKecemasanRanapAnak) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianLevelKecemasanRanapAnak, d request.PenilaianLevelKecemasanRanapAnakData) error {
		m.Cemas = d.Cemas
		m.FirasatBuruk = d.FirasatBuruk
		m.TakutPikiranSendiri = d.TakutPikiranSendiri
		m.MudahTersinggung = d.MudahTersinggung
		m.MerasaTegang = d.MerasaTegang
		m.Lesu = d.Lesu
		m.TakBisaIstirahatTenang = d.TakBisaIstirahatTenang
		m.MudahTerkejut = d.MudahTerkejut
		m.MudahMenangis = d.MudahMenangis
		m.Gemetar = d.Gemetar
		m.Gelisah = d.Gelisah
		m.TakutPadaGelap = d.TakutPadaGelap
		m.TakutPadaOrangasing = d.TakutPadaOrangasing
		m.TakutPadaKerumunanBanyakOrang = d.TakutPadaKerumunanBanyakOrang
		m.TakutPadaBinatangBesar = d.TakutPadaBinatangBesar
		m.TakutPadaKeramaianLaluLintas = d.TakutPadaKeramaianLaluLintas
		m.TakutDitinggalSendiri = d.TakutDitinggalSendiri
		m.SulitTidur = d.SulitTidur
		m.TerbangunMalamHari = d.TerbangunMalamHari
		m.TidurTidakNyeyak = d.TidurTidakNyeyak
		m.MimpiBuruk = d.MimpiBuruk
		m.BangunDenganLesu = d.BangunDenganLesu
		m.BanyakMengalamiMimpi = d.BanyakMengalamiMimpi
		m.MimpiMenakutkan = d.MimpiMenakutkan
		m.SulitKonsentrasi = d.SulitKonsentrasi
		m.DayaIngatBuruk = d.DayaIngatBuruk
		m.HilangnyaMinat = d.HilangnyaMinat
		m.BerkurangnyaKesenanganPadaHobi = d.BerkurangnyaKesenanganPadaHobi
		m.Sedih = d.Sedih
		m.BangunDiniHari = d.BangunDiniHari
		m.PerasaanBerubah = d.PerasaanBerubah
		m.SakitNyeriDiOtot = d.SakitNyeriDiOtot
		m.Kaku = d.Kaku
		m.KedutanOtot = d.KedutanOtot
		m.GigiGemerutuk = d.GigiGemerutuk
		m.SuaraTidakStabil = d.SuaraTidakStabil
		m.Tinnitus = d.Tinnitus
		m.PenglihatanKabur = d.PenglihatanKabur
		m.MukaMerahGejalaSomatic = d.MukaMerahGejalaSomatic
		m.MerasaLemah = d.MerasaLemah
		m.PerasaanDitusuk = d.PerasaanDitusuk
		m.Takhikardia = d.Takhikardia
		m.Berdebar = d.Berdebar
		m.NyeriDiDada = d.NyeriDiDada
		m.DenyutNadiMengeras = d.DenyutNadiMengeras
		m.PerasaanLesu = d.PerasaanLesu
		m.DetakJantungMenghilang = d.DetakJantungMenghilang
		m.MerasaTertekan = d.MerasaTertekan
		m.PerasaanTercekik = d.PerasaanTercekik
		m.SeringMenarikNapas = d.SeringMenarikNapas
		m.NapasPendek = d.NapasPendek
		m.BuluBerdiri = d.BuluBerdiri
		m.SulitMenelan = d.SulitMenelan
		m.PerutMelilit = d.PerutMelilit
		m.GanguanPencernaan = d.GanguanPencernaan
		m.RasaKembung = d.RasaKembung
		m.NyeriMakan = d.NyeriMakan
		m.TerbakarPerut = d.TerbakarPerut
		m.SukarBab = d.SukarBab
		m.Muntah = d.Muntah
		m.BabLembek = d.BabLembek
		m.KehilanganBb = d.KehilanganBb
		m.Mual = d.Mual
		m.SeringBak = d.SeringBak
		m.TidakBisaMenahanKencing = d.TidakBisaMenahanKencing
		m.MenjadiDingin = d.MenjadiDingin
		m.Manorrhagia = d.Manorrhagia
		m.Amenorrhoea = d.Amenorrhoea
		m.EjakulasiPraecocks = d.EjakulasiPraecocks
		m.EreksiHilang = d.EreksiHilang
		m.Impotensi = d.Impotensi
		m.MulutKering = d.MulutKering
		m.MukaMerahGejalaOtonom = d.MukaMerahGejalaOtonom
		m.MudahBerkeringat = d.MudahBerkeringat
		m.BuluBerdiriGejalaOtonom = d.BuluBerdiriGejalaOtonom
		m.SakitKepala = d.SakitKepala
		m.GelisahWawancara = d.GelisahWawancara
		m.NapasPendekWawancara = d.NapasPendekWawancara
		m.JariGemetar = d.JariGemetar
		m.KerutKening = d.KerutKening
		m.MukaTegang = d.MukaTegang
		m.TonusMeningkat = d.TonusMeningkat
		m.TidakTenang = d.TidakTenang
		m.MukaMerahWawancara = d.MukaMerahWawancara
		m.TotalSkor = d.TotalSkor
		m.KeteranganSkor = support.Nullable(d.KeteranganSkor)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianLevelKecemasanRanapAnak]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianLevelKecemasanRanapAnak) any { return Str(m.Nip) }},
	},
}
