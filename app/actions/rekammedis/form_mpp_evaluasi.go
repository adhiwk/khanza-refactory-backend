package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormMppEvaluasi skrining MPP form a (RMSkriningMPPFormA).
var FormMppEvaluasi = &Form[model.MppEvaluasi, request.MppEvaluasiData]{
	Slug:  "mpp-evaluasi",
	Label: "skrining MPP form a",
	Spec: repo.Spec{
		Table:  "mpp_evaluasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter", "kd_konsulan", "nip"},
	},
	SetKey: func(m *model.MppEvaluasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.MppEvaluasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.MppEvaluasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.MppEvaluasi) []string {
		return []string{derefStr(m.KdDokter), derefStr(m.KdKonsulan), m.Nip}
	},
	Fill: func(m *model.MppEvaluasi, d request.MppEvaluasiData) error {
		m.KdDokter = support.Nullable(d.KdDokter)
		m.KdKonsulan = support.Nullable(d.KdKonsulan)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Kelompok = strings.TrimSpace(d.Kelompok)
		m.Assesmen = strings.TrimSpace(d.Assesmen)
		m.Identifikasi = strings.TrimSpace(d.Identifikasi)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.MppEvaluasi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.MppEvaluasi) any { return StrPtr(m.KdDokter) }},
		{Column: "kd_konsulan", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.MppEvaluasi) any { return StrPtr(m.KdKonsulan) }},
		{Column: "nip", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.MppEvaluasi) any { return Str(m.Nip) }},
	},
}
