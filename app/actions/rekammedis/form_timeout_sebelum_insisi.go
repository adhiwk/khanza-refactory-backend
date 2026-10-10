package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormTimeoutSebelumInsisi time out sebelum insisi (RMTimeOutSebelumInsisi).
var FormTimeoutSebelumInsisi = &Form[model.TimeoutSebelumInsisi, request.TimeoutSebelumInsisiData]{
	Slug:  "timeout-sebelum-insisi",
	Label: "time out sebelum insisi",
	Spec: repo.Spec{
		Table:  "timeout_sebelum_insisi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter_bedah", "kd_dokter_anestesi", "nip_perawat_ok"},
	},
	SetKey: func(m *model.TimeoutSebelumInsisi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.TimeoutSebelumInsisi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.TimeoutSebelumInsisi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.TimeoutSebelumInsisi) []string {
		return []string{m.KdDokterBedah, m.KdDokterAnestesi, derefStr(m.NipPerawatOk)}
	},
	Fill: func(m *model.TimeoutSebelumInsisi, d request.TimeoutSebelumInsisiData) error {
		vTanggalSteril, err := support.ParseDate(d.TanggalSteril)
		if err != nil {
			return err
		}
		m.Sncn = strings.TrimSpace(d.Sncn)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.KdDokterBedah = strings.TrimSpace(d.KdDokterBedah)
		m.KdDokterAnestesi = strings.TrimSpace(d.KdDokterAnestesi)
		m.VerbalIdentitas = support.Nullable(d.VerbalIdentitas)
		m.VerbalTindakan = support.Nullable(d.VerbalTindakan)
		m.VerbalAreaInsisi = support.Nullable(d.VerbalAreaInsisi)
		m.PenandaanAreaOperasi = support.Nullable(d.PenandaanAreaOperasi)
		m.LamaOperasi = strings.TrimSpace(d.LamaOperasi)
		m.PenayanganRadiologi = support.Nullable(d.PenayanganRadiologi)
		m.PenayanganCtscan = support.Nullable(d.PenayanganCtscan)
		m.PenayanganMri = support.Nullable(d.PenayanganMri)
		m.AntibiotikProfilaks = support.Nullable(d.AntibiotikProfilaks)
		m.NamaAntibiotik = strings.TrimSpace(d.NamaAntibiotik)
		m.JamPemberian = strings.TrimSpace(d.JamPemberian)
		m.AntisipasiKehilanganDarah = strings.TrimSpace(d.AntisipasiKehilanganDarah)
		m.HalKhusus = support.Nullable(d.HalKhusus)
		m.HalKhususDiperhatikan = strings.TrimSpace(d.HalKhususDiperhatikan)
		m.TanggalSteril = vTanggalSteril
		m.PetujukSterilisasi = support.Nullable(d.PetujukSterilisasi)
		m.VerifikasiPreoperatif = support.Nullable(d.VerifikasiPreoperatif)
		m.NipPerawatOk = support.Nullable(d.NipPerawatOk)
		return nil
	},
	Refs: []Ref[model.TimeoutSebelumInsisi]{
		{Column: "kd_dokter_bedah", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.TimeoutSebelumInsisi) any { return Str(m.KdDokterBedah) }},
		{Column: "kd_dokter_anestesi", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.TimeoutSebelumInsisi) any { return Str(m.KdDokterAnestesi) }},
		{Column: "nip_perawat_ok", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.TimeoutSebelumInsisi) any { return StrPtr(m.NipPerawatOk) }},
	},
}
