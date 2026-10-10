package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistPostOperasi checklist post operasi (RMChecklistPostOperasi).
var FormChecklistPostOperasi = &Form[model.ChecklistPostOperasi, request.ChecklistPostOperasiData]{
	Slug:  "checklist-post-operasi",
	Label: "checklist post operasi",
	Spec: repo.Spec{
		Table:  "checklist_post_operasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter_bedah", "kd_dokter_anestesi", "nip_perawat_ok", "nip_perawat_anestesi"},
	},
	SetKey: func(m *model.ChecklistPostOperasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistPostOperasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.ChecklistPostOperasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistPostOperasi) []string {
		return []string{m.KdDokterBedah, m.KdDokterAnestesi, derefStr(m.NipPerawatOk), derefStr(m.NipPerawatAnestesi)}
	},
	Fill: func(m *model.ChecklistPostOperasi, d request.ChecklistPostOperasiData) error {
		vTanggalPemasanganKateter, err := support.ParseDateTime(d.TanggalPemasanganKateter)
		if err != nil {
			return err
		}
		m.Sncn = strings.TrimSpace(d.Sncn)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.KdDokterBedah = strings.TrimSpace(d.KdDokterBedah)
		m.KdDokterAnestesi = strings.TrimSpace(d.KdDokterAnestesi)
		m.KeadaanUmum = support.Nullable(d.KeadaanUmum)
		m.PemeriksaanPenunjangRontgen = support.Nullable(d.PemeriksaanPenunjangRontgen)
		m.KeteranganPemeriksaanPenunjangRontgen = support.Nullable(d.KeteranganPemeriksaanPenunjangRontgen)
		m.PemeriksaanPenunjangEkg = support.Nullable(d.PemeriksaanPenunjangEkg)
		m.KeteranganPemeriksaanPenunjangEkg = support.Nullable(d.KeteranganPemeriksaanPenunjangEkg)
		m.PemeriksaanPenunjangUsg = support.Nullable(d.PemeriksaanPenunjangUsg)
		m.KeteranganPemeriksaanPenunjangUsg = support.Nullable(d.KeteranganPemeriksaanPenunjangUsg)
		m.PemeriksaanPenunjangCtscan = support.Nullable(d.PemeriksaanPenunjangCtscan)
		m.KeteranganPemeriksaanPenunjangCtscan = support.Nullable(d.KeteranganPemeriksaanPenunjangCtscan)
		m.PemeriksaanPenunjangMri = support.Nullable(d.PemeriksaanPenunjangMri)
		m.KeteranganPemeriksaanPenunjangMri = support.Nullable(d.KeteranganPemeriksaanPenunjangMri)
		m.JenisCairanInfus = support.Nullable(d.JenisCairanInfus)
		m.KateterUrine = support.Nullable(d.KateterUrine)
		m.TanggalPemasanganKateter = vTanggalPemasanganKateter
		m.WarnaKateter = support.Nullable(d.WarnaKateter)
		m.JumlahKateter = support.Nullable(d.JumlahKateter)
		m.AreaLukaOperasi = support.Nullable(d.AreaLukaOperasi)
		m.Drain = support.Nullable(d.Drain)
		m.JumlahDrain = support.Nullable(d.JumlahDrain)
		m.LetakDrain = support.Nullable(d.LetakDrain)
		m.WarnaDrain = support.Nullable(d.WarnaDrain)
		m.JaringanPa = support.Nullable(d.JaringanPa)
		m.NipPerawatOk = support.Nullable(d.NipPerawatOk)
		m.NipPerawatAnestesi = support.Nullable(d.NipPerawatAnestesi)
		return nil
	},
	Refs: []Ref[model.ChecklistPostOperasi]{
		{Column: "kd_dokter_bedah", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ChecklistPostOperasi) any { return Str(m.KdDokterBedah) }},
		{Column: "kd_dokter_anestesi", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ChecklistPostOperasi) any { return Str(m.KdDokterAnestesi) }},
		{Column: "nip_perawat_ok", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.ChecklistPostOperasi) any { return StrPtr(m.NipPerawatOk) }},
		{Column: "nip_perawat_anestesi", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.ChecklistPostOperasi) any { return StrPtr(m.NipPerawatAnestesi) }},
	},
}
