package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistPreOperasi checklist pre operasi (RMChecklistPreOperasi).
var FormChecklistPreOperasi = &Form[model.ChecklistPreOperasi, request.ChecklistPreOperasiData]{
	Slug:  "checklist-pre-operasi",
	Label: "checklist pre operasi",
	Spec: repo.Spec{
		Table:  "checklist_pre_operasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter_bedah", "kd_dokter_anestesi", "nip_petugas_ruangan", "nip_perawat_ok"},
	},
	SetKey: func(m *model.ChecklistPreOperasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistPreOperasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.ChecklistPreOperasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistPreOperasi) []string {
		return []string{m.KdDokterBedah, m.KdDokterAnestesi, derefStr(m.NipPetugasRuangan), derefStr(m.NipPerawatOk)}
	},
	Fill: func(m *model.ChecklistPreOperasi, d request.ChecklistPreOperasiData) error {
		m.Sncn = strings.TrimSpace(d.Sncn)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.KdDokterBedah = strings.TrimSpace(d.KdDokterBedah)
		m.KdDokterAnestesi = strings.TrimSpace(d.KdDokterAnestesi)
		m.Identitas = support.Nullable(d.Identitas)
		m.SuratIjinBedah = support.Nullable(d.SuratIjinBedah)
		m.SuratIjinAnestesi = support.Nullable(d.SuratIjinAnestesi)
		m.SuratIjinTransfusi = support.Nullable(d.SuratIjinTransfusi)
		m.PenandaanAreaOperasi = support.Nullable(d.PenandaanAreaOperasi)
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
		m.PersiapanDarah = support.Nullable(d.PersiapanDarah)
		m.KeteranganPersiapanDarah = support.Nullable(d.KeteranganPersiapanDarah)
		m.PerlengkapanKhusus = support.Nullable(d.PerlengkapanKhusus)
		m.NipPetugasRuangan = support.Nullable(d.NipPetugasRuangan)
		m.NipPerawatOk = support.Nullable(d.NipPerawatOk)
		return nil
	},
	Refs: []Ref[model.ChecklistPreOperasi]{
		{Column: "kd_dokter_bedah", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ChecklistPreOperasi) any { return Str(m.KdDokterBedah) }},
		{Column: "kd_dokter_anestesi", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ChecklistPreOperasi) any { return Str(m.KdDokterAnestesi) }},
		{Column: "nip_petugas_ruangan", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.ChecklistPreOperasi) any { return StrPtr(m.NipPetugasRuangan) }},
		{Column: "nip_perawat_ok", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.ChecklistPreOperasi) any { return StrPtr(m.NipPerawatOk) }},
	},
}
