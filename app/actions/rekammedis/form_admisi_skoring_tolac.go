package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormAdmisiSkoringTolac admisi skoring TOLAC (RMAdmisiSkoringTOLAC).
var FormAdmisiSkoringTolac = &Form[model.AdmisiSkoringTolac, request.AdmisiSkoringTolacData]{
	Slug:  "admisi-skoring-tolac",
	Label: "admisi skoring TOLAC",
	Spec: repo.Spec{
		Table:  "admisi_skoring_tolac",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.AdmisiSkoringTolac, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.AdmisiSkoringTolac) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.AdmisiSkoringTolac) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.AdmisiSkoringTolac) []string { return []string{derefStr(m.KdDokter)} },
	Fill: func(m *model.AdmisiSkoringTolac, d request.AdmisiSkoringTolacData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.HisFrekuensi = support.Nullable(d.HisFrekuensi)
		m.HisDurasiDetik = support.Nullable(d.HisDurasiDetik)
		m.Djj = support.Nullable(d.Djj)
		m.PembukaanCm = support.Nullable(d.PembukaanCm)
		m.PendataranPersen = support.Nullable(d.PendataranPersen)
		m.PenurunanKepala = support.Nullable(d.PenurunanKepala)
		m.PilihanUsia = support.Nullable(d.PilihanUsia)
		m.SkorUsia = support.Nullable(d.SkorUsia)
		m.PilihanRiwayatPervaginam = support.Nullable(d.PilihanRiwayatPervaginam)
		m.SkorRiwayatPervaginam = support.Nullable(d.SkorRiwayatPervaginam)
		m.PilihanIndikasiSc = support.Nullable(d.PilihanIndikasiSc)
		m.SkorIndikasiSc = support.Nullable(d.SkorIndikasiSc)
		m.PilihanPendataran = support.Nullable(d.PilihanPendataran)
		m.SkorPendataran = support.Nullable(d.SkorPendataran)
		m.PilihanPembukaan = support.Nullable(d.PilihanPembukaan)
		m.SkorPembukaan = support.Nullable(d.SkorPembukaan)
		m.TotalSkor = support.Nullable(d.TotalSkor)
		m.PeluangVbac = strings.TrimSpace(d.PeluangVbac)
		m.Keputusan = support.Nullable(d.Keputusan)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.KdDokter = support.Nullable(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.AdmisiSkoringTolac]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.AdmisiSkoringTolac) any { return StrPtr(m.KdDokter) }},
	},
}
