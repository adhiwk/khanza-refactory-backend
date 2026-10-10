package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
)

// FormResumePasien resume pasien (RMDataResumePasien).
var FormResumePasien = &Form[model.ResumePasien, request.ResumePasienData]{
	Slug:  "resume-pasien",
	Label: "resume pasien",
	Spec: repo.Spec{
		Table:  "resume_pasien",
		Keys:   []string{"no_rawat"},
		Waktu:  "",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.ResumePasien, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.ResumePasien) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.ResumePasien) time.Time { return time.Time{} },
	Petugas: func(m *model.ResumePasien) []string { return []string{m.KdDokter} },
	Fill: func(m *model.ResumePasien, d request.ResumePasienData) error {
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.JalannyaPenyakit = strings.TrimSpace(d.JalannyaPenyakit)
		m.PemeriksaanPenunjang = strings.TrimSpace(d.PemeriksaanPenunjang)
		m.HasilLaborat = strings.TrimSpace(d.HasilLaborat)
		m.DiagnosaUtama = strings.TrimSpace(d.DiagnosaUtama)
		m.KdDiagnosaUtama = strings.TrimSpace(d.KdDiagnosaUtama)
		m.DiagnosaSekunder = strings.TrimSpace(d.DiagnosaSekunder)
		m.KdDiagnosaSekunder = strings.TrimSpace(d.KdDiagnosaSekunder)
		m.DiagnosaSekunder2 = strings.TrimSpace(d.DiagnosaSekunder2)
		m.KdDiagnosaSekunder2 = strings.TrimSpace(d.KdDiagnosaSekunder2)
		m.DiagnosaSekunder3 = strings.TrimSpace(d.DiagnosaSekunder3)
		m.KdDiagnosaSekunder3 = strings.TrimSpace(d.KdDiagnosaSekunder3)
		m.DiagnosaSekunder4 = strings.TrimSpace(d.DiagnosaSekunder4)
		m.KdDiagnosaSekunder4 = strings.TrimSpace(d.KdDiagnosaSekunder4)
		m.ProsedurUtama = strings.TrimSpace(d.ProsedurUtama)
		m.KdProsedurUtama = strings.TrimSpace(d.KdProsedurUtama)
		m.ProsedurSekunder = strings.TrimSpace(d.ProsedurSekunder)
		m.KdProsedurSekunder = strings.TrimSpace(d.KdProsedurSekunder)
		m.ProsedurSekunder2 = strings.TrimSpace(d.ProsedurSekunder2)
		m.KdProsedurSekunder2 = strings.TrimSpace(d.KdProsedurSekunder2)
		m.ProsedurSekunder3 = strings.TrimSpace(d.ProsedurSekunder3)
		m.KdProsedurSekunder3 = strings.TrimSpace(d.KdProsedurSekunder3)
		m.KondisiPulang = strings.TrimSpace(d.KondisiPulang)
		m.ObatPulang = strings.TrimSpace(d.ObatPulang)
		return nil
	},
	Refs: []Ref[model.ResumePasien]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ResumePasien) any { return Str(m.KdDokter) }},
	},
}
