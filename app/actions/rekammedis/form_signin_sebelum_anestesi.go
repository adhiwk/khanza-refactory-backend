package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSigninSebelumAnestesi sign in sebelum anastesi (RMSignInSebelumAnastesi).
var FormSigninSebelumAnestesi = &Form[model.SigninSebelumAnestesi, request.SigninSebelumAnestesiData]{
	Slug:  "signin-sebelum-anestesi",
	Label: "sign in sebelum anastesi",
	Spec: repo.Spec{
		Table:  "signin_sebelum_anestesi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter_bedah", "kd_dokter_anestesi", "nip_perawat_ok"},
	},
	SetKey: func(m *model.SigninSebelumAnestesi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.SigninSebelumAnestesi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.SigninSebelumAnestesi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SigninSebelumAnestesi) []string {
		return []string{m.KdDokterBedah, m.KdDokterAnestesi, derefStr(m.NipPerawatOk)}
	},
	Fill: func(m *model.SigninSebelumAnestesi, d request.SigninSebelumAnestesiData) error {
		m.Sncn = strings.TrimSpace(d.Sncn)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.KdDokterBedah = strings.TrimSpace(d.KdDokterBedah)
		m.KdDokterAnestesi = strings.TrimSpace(d.KdDokterAnestesi)
		m.Identitas = support.Nullable(d.Identitas)
		m.PenandaanAreaOperasi = support.Nullable(d.PenandaanAreaOperasi)
		m.Alergi = support.Nullable(d.Alergi)
		m.ResikoAspirasi = support.Nullable(d.ResikoAspirasi)
		m.ResikoAspirasiRencanaAntisipasi = support.Nullable(d.ResikoAspirasiRencanaAntisipasi)
		m.ResikoKehilanganDarah = support.Nullable(d.ResikoKehilanganDarah)
		m.ResikoKehilanganDarahLine = support.Nullable(d.ResikoKehilanganDarahLine)
		m.ResikoKehilanganDarahRencanaAntisipasi = support.Nullable(d.ResikoKehilanganDarahRencanaAntisipasi)
		m.KesiapanAlatObatAnestesi = support.Nullable(d.KesiapanAlatObatAnestesi)
		m.KesiapanAlatObatAnestesiRencanaAntisipasi = support.Nullable(d.KesiapanAlatObatAnestesiRencanaAntisipasi)
		m.NipPerawatOk = support.Nullable(d.NipPerawatOk)
		return nil
	},
	Refs: []Ref[model.SigninSebelumAnestesi]{
		{Column: "kd_dokter_bedah", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SigninSebelumAnestesi) any { return Str(m.KdDokterBedah) }},
		{Column: "kd_dokter_anestesi", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SigninSebelumAnestesi) any { return Str(m.KdDokterAnestesi) }},
		{Column: "nip_perawat_ok", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SigninSebelumAnestesi) any { return StrPtr(m.NipPerawatOk) }},
	},
}
