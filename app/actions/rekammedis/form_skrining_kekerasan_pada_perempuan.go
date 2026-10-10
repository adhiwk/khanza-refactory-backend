package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKekerasanPadaPerempuan skrining kekerasan pada perempuan (RMSkriningKekerasanPadaPerempuan).
var FormSkriningKekerasanPadaPerempuan = &Form[model.SkriningKekerasanPadaPerempuan, request.SkriningKekerasanPadaPerempuanData]{
	Slug:  "skrining-kekerasan-pada-perempuan",
	Label: "skrining kekerasan pada perempuan",
	Spec: repo.Spec{
		Table:  "skrining_kekerasan_pada_perempuan",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKekerasanPadaPerempuan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKekerasanPadaPerempuan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKekerasanPadaPerempuan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKekerasanPadaPerempuan) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKekerasanPadaPerempuan, d request.SkriningKekerasanPadaPerempuanData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.MenggambarkanHubungan = support.Nullable(d.MenggambarkanHubungan)
		m.SkorMenggambarkanHubungan = strings.TrimSpace(d.SkorMenggambarkanHubungan)
		m.BerdebatDenganPasangan = support.Nullable(d.BerdebatDenganPasangan)
		m.SkorBerdebatDenganPasangan = strings.TrimSpace(d.SkorBerdebatDenganPasangan)
		m.PertengkaranMembuatSedih = support.Nullable(d.PertengkaranMembuatSedih)
		m.SkorPertengkaranMembuatSedih = strings.TrimSpace(d.SkorPertengkaranMembuatSedih)
		m.PertengkaranMenghasilkanPukulan = support.Nullable(d.PertengkaranMenghasilkanPukulan)
		m.SkorPertengkaranMenghasilkanPukulan = strings.TrimSpace(d.SkorPertengkaranMenghasilkanPukulan)
		m.PernahMerasaTakutDenganPasangan = support.Nullable(d.PernahMerasaTakutDenganPasangan)
		m.SkorPernahMerasaTakutDenganPasangan = strings.TrimSpace(d.SkorPernahMerasaTakutDenganPasangan)
		m.PasanganMelecehkanSecaraFisik = support.Nullable(d.PasanganMelecehkanSecaraFisik)
		m.SkorPasanganMelecehkanSecaraFisik = strings.TrimSpace(d.SkorPasanganMelecehkanSecaraFisik)
		m.PasanganMelecehkanSecaraImosional = support.Nullable(d.PasanganMelecehkanSecaraImosional)
		m.SkorPasanganMelecehkanSecaraImosional = strings.TrimSpace(d.SkorPasanganMelecehkanSecaraImosional)
		m.PasanganMelecehkanSecaraSeksual = support.Nullable(d.PasanganMelecehkanSecaraSeksual)
		m.SkorPasanganMelecehkanSecaraSeksual = strings.TrimSpace(d.SkorPasanganMelecehkanSecaraSeksual)
		m.Totalskor = strings.TrimSpace(d.Totalskor)
		m.HasilSkrining = strings.TrimSpace(d.HasilSkrining)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningKekerasanPadaPerempuan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKekerasanPadaPerempuan) any { return Str(m.Nip) }},
	},
}
