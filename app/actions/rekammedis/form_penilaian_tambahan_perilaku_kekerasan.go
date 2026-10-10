package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianTambahanPerilakuKekerasan penilaian tambahan perilaku kekerasan (RMPenilaianTambahanPerilakuKekerasan).
var FormPenilaianTambahanPerilakuKekerasan = &Form[model.PenilaianTambahanPerilakuKekerasan, request.PenilaianTambahanPerilakuKekerasanData]{
	Slug:  "penilaian-tambahan-perilaku-kekerasan",
	Label: "penilaian tambahan perilaku kekerasan",
	Spec: repo.Spec{
		Table:  "penilaian_tambahan_perilaku_kekerasan",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianTambahanPerilakuKekerasan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianTambahanPerilakuKekerasan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianTambahanPerilakuKekerasan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianTambahanPerilakuKekerasan) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianTambahanPerilakuKekerasan, d request.PenilaianTambahanPerilakuKekerasanData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.StatikInsidenKekerasanBaruIni = support.Nullable(d.StatikInsidenKekerasanBaruIni)
		m.StatikSkorinsidenKekerasanBaruIni = d.StatikSkorinsidenKekerasanBaruIni
		m.StatikRiwayatPenggunaanSenjata = support.Nullable(d.StatikRiwayatPenggunaanSenjata)
		m.StatikSkorriwayatPenggunaanSenjata = d.StatikSkorriwayatPenggunaanSenjata
		m.StatikLakiLaki = support.Nullable(d.StatikLakiLaki)
		m.StatikSkorlakiLaki = d.StatikSkorlakiLaki
		m.StatikUsiaDibawah35 = support.Nullable(d.StatikUsiaDibawah35)
		m.StatikSkorusiaDibawah35 = d.StatikSkorusiaDibawah35
		m.StatikRiwayatKriminal = support.Nullable(d.StatikRiwayatKriminal)
		m.StatikSkorriwayatKriminal = d.StatikSkorriwayatKriminal
		m.StatikIdeKekerasan = support.Nullable(d.StatikIdeKekerasan)
		m.StatikSkorideKekerasan = d.StatikSkorideKekerasan
		m.StatikKekerasanAnakAnak = support.Nullable(d.StatikKekerasanAnakAnak)
		m.StatikSkorkekerasanAnakAnak = d.StatikSkorkekerasanAnakAnak
		m.StatikPeranDalamHidup = support.Nullable(d.StatikPeranDalamHidup)
		m.StatikSkorperanDalamHidup = d.StatikSkorperanDalamHidup
		m.StatikPenggunaanNapza = support.Nullable(d.StatikPenggunaanNapza)
		m.StatikSkorpenggunaanNapza = d.StatikSkorpenggunaanNapza
		m.StatikSkortotal = d.StatikSkortotal
		m.DinamisIdeMelukaiOrangLain = support.Nullable(d.DinamisIdeMelukaiOrangLain)
		m.DinamisSkorideMelukaiOrangLain = d.DinamisSkorideMelukaiOrangLain
		m.DinamisAksesKekerasan = support.Nullable(d.DinamisAksesKekerasan)
		m.DinamisSkoraksesKekerasan = d.DinamisSkoraksesKekerasan
		m.DinamisIdeParanoid = support.Nullable(d.DinamisIdeParanoid)
		m.DinamisSkorideParanoid = d.DinamisSkorideParanoid
		m.DinamisPerintahHalusinasi = support.Nullable(d.DinamisPerintahHalusinasi)
		m.DinamisSkorperintahHalusinasi = d.DinamisSkorperintahHalusinasi
		m.DinamisFrustasiAgitasi = support.Nullable(d.DinamisFrustasiAgitasi)
		m.DinamisSkorfrustasiAgitasi = d.DinamisSkorfrustasiAgitasi
		m.DinamisKesenanganKekerasan = support.Nullable(d.DinamisKesenanganKekerasan)
		m.DinamisSkorkesenanganKekerasan = d.DinamisSkorkesenanganKekerasan
		m.DinamisSeksualTidakWajar = support.Nullable(d.DinamisSeksualTidakWajar)
		m.DinamisSkorseksualTidakWajar = d.DinamisSkorseksualTidakWajar
		m.DinamisHilangnyaKontrolDiri = support.Nullable(d.DinamisHilangnyaKontrolDiri)
		m.DinamisSkorhilangnyaKontrolDiri = d.DinamisSkorhilangnyaKontrolDiri
		m.DinamisPengguaanNapza = support.Nullable(d.DinamisPengguaanNapza)
		m.DinamisSkorpengguaanNapza = d.DinamisSkorpengguaanNapza
		m.DinamisSkortotal = d.DinamisSkortotal
		m.FaktorFaktorPencegahan = support.Nullable(d.FaktorFaktorPencegahan)
		m.TotalSkor = d.TotalSkor
		m.LevelSkor = support.Nullable(d.LevelSkor)
		return nil
	},
	Refs: []Ref[model.PenilaianTambahanPerilakuKekerasan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianTambahanPerilakuKekerasan) any { return Str(m.Nip) }},
	},
}
