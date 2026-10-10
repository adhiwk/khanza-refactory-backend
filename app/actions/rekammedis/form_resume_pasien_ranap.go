package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormResumePasienRanap resume pasien ranap (RMDataResumePasienRanap).
var FormResumePasienRanap = &Form[model.ResumePasienRanap, request.ResumePasienRanapData]{
	Slug:  "resume-pasien-ranap",
	Label: "resume pasien ranap",
	Spec: repo.Spec{
		Table:  "resume_pasien_ranap",
		Keys:   []string{"no_rawat"},
		Waktu:  "",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.ResumePasienRanap, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.ResumePasienRanap) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.ResumePasienRanap) time.Time { return time.Time{} },
	Petugas: func(m *model.ResumePasienRanap) []string { return []string{m.KdDokter} },
	Fill: func(m *model.ResumePasienRanap, d request.ResumePasienRanapData) error {
		vKontrol, err := support.ParseDateTime(d.Kontrol)
		if err != nil {
			return err
		}
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaAwal = strings.TrimSpace(d.DiagnosaAwal)
		m.Alasan = strings.TrimSpace(d.Alasan)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.PemeriksaanFisik = strings.TrimSpace(d.PemeriksaanFisik)
		m.JalannyaPenyakit = strings.TrimSpace(d.JalannyaPenyakit)
		m.PemeriksaanPenunjang = strings.TrimSpace(d.PemeriksaanPenunjang)
		m.HasilLaborat = strings.TrimSpace(d.HasilLaborat)
		m.TindakanDanOperasi = strings.TrimSpace(d.TindakanDanOperasi)
		m.ObatDiRs = strings.TrimSpace(d.ObatDiRs)
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
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Diet = strings.TrimSpace(d.Diet)
		m.LabBelum = strings.TrimSpace(d.LabBelum)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		m.CaraKeluar = strings.TrimSpace(d.CaraKeluar)
		m.KetKeluar = support.Nullable(d.KetKeluar)
		m.Keadaan = strings.TrimSpace(d.Keadaan)
		m.KetKeadaan = support.Nullable(d.KetKeadaan)
		m.Dilanjutkan = strings.TrimSpace(d.Dilanjutkan)
		m.KetDilanjutkan = support.Nullable(d.KetDilanjutkan)
		m.Kontrol = vKontrol
		m.ObatPulang = strings.TrimSpace(d.ObatPulang)
		return nil
	},
	Refs: []Ref[model.ResumePasienRanap]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ResumePasienRanap) any { return Str(m.KdDokter) }},
	},
}
