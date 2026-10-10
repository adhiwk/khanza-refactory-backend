package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSignoutSebelumMenutupLuka sign out sebelum menutup luka (RMSignOutSebelumMenutupLuka).
var FormSignoutSebelumMenutupLuka = &Form[model.SignoutSebelumMenutupLuka, request.SignoutSebelumMenutupLukaData]{
	Slug:  "signout-sebelum-menutup-luka",
	Label: "sign out sebelum menutup luka",
	Spec: repo.Spec{
		Table:  "signout_sebelum_menutup_luka",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter_bedah", "kd_dokter_anestesi", "nip_perawat_ok"},
	},
	SetKey: func(m *model.SignoutSebelumMenutupLuka, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.SignoutSebelumMenutupLuka) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.SignoutSebelumMenutupLuka) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SignoutSebelumMenutupLuka) []string {
		return []string{m.KdDokterBedah, m.KdDokterAnestesi, derefStr(m.NipPerawatOk)}
	},
	Fill: func(m *model.SignoutSebelumMenutupLuka, d request.SignoutSebelumMenutupLukaData) error {
		m.Sncn = strings.TrimSpace(d.Sncn)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.KdDokterBedah = strings.TrimSpace(d.KdDokterBedah)
		m.KdDokterAnestesi = strings.TrimSpace(d.KdDokterAnestesi)
		m.VerbalTindakan = support.Nullable(d.VerbalTindakan)
		m.VerbalKelengkapanKasa = support.Nullable(d.VerbalKelengkapanKasa)
		m.VerbalInstrumen = support.Nullable(d.VerbalInstrumen)
		m.VerbalAlatTajam = support.Nullable(d.VerbalAlatTajam)
		m.KelengkapanSpecimenLabel = support.Nullable(d.KelengkapanSpecimenLabel)
		m.KelengkapanSpecimenFormulir = support.Nullable(d.KelengkapanSpecimenFormulir)
		m.PeninjauanKegiatanDokterBedah = support.Nullable(d.PeninjauanKegiatanDokterBedah)
		m.PeninjauanKegiatanDokterAnestesi = support.Nullable(d.PeninjauanKegiatanDokterAnestesi)
		m.PeninjauanKegiatanPerawatKamarOk = support.Nullable(d.PeninjauanKegiatanPerawatKamarOk)
		m.PerhatianUtamaFasePemulihan = support.Nullable(d.PerhatianUtamaFasePemulihan)
		m.NipPerawatOk = support.Nullable(d.NipPerawatOk)
		return nil
	},
	Refs: []Ref[model.SignoutSebelumMenutupLuka]{
		{Column: "kd_dokter_bedah", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SignoutSebelumMenutupLuka) any { return Str(m.KdDokterBedah) }},
		{Column: "kd_dokter_anestesi", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SignoutSebelumMenutupLuka) any { return Str(m.KdDokterAnestesi) }},
		{Column: "nip_perawat_ok", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SignoutSebelumMenutupLuka) any { return StrPtr(m.NipPerawatOk) }},
	},
}
