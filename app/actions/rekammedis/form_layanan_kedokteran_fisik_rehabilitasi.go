package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormLayananKedokteranFisikRehabilitasi layanan kedokteran fisik rehabilitasi (RMLayananKedokteranFisikRehabilitasi).
var FormLayananKedokteranFisikRehabilitasi = &Form[model.LayananKedokteranFisikRehabilitasi, request.LayananKedokteranFisikRehabilitasiData]{
	Slug:  "layanan-kedokteran-fisik-rehabilitasi",
	Label: "layanan kedokteran fisik rehabilitasi",
	Spec: repo.Spec{
		Table:  "layanan_kedokteran_fisik_rehabilitasi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.LayananKedokteranFisikRehabilitasi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.LayananKedokteranFisikRehabilitasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.LayananKedokteranFisikRehabilitasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.LayananKedokteranFisikRehabilitasi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.LayananKedokteranFisikRehabilitasi, d request.LayananKedokteranFisikRehabilitasiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Pendamping = support.Nullable(d.Pendamping)
		m.KeteranganPendamping = support.Nullable(d.KeteranganPendamping)
		m.Anamnesa = support.Nullable(d.Anamnesa)
		m.PemeriksaanFisik = support.Nullable(d.PemeriksaanFisik)
		m.DiagnosaMedis = support.Nullable(d.DiagnosaMedis)
		m.DiagnosaFungsi = support.Nullable(d.DiagnosaFungsi)
		m.Tatalaksana = support.Nullable(d.Tatalaksana)
		m.Anjuran = strings.TrimSpace(d.Anjuran)
		m.Evaluasi = strings.TrimSpace(d.Evaluasi)
		m.SuspekPenyakitKerja = support.Nullable(d.SuspekPenyakitKerja)
		m.KeteranganSuspekPenyakitKerja = support.Nullable(d.KeteranganSuspekPenyakitKerja)
		m.StatusProgram = strings.TrimSpace(d.StatusProgram)
		return nil
	},
	Refs: []Ref[model.LayananKedokteranFisikRehabilitasi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.LayananKedokteranFisikRehabilitasi) any { return Str(m.KdDokter) }},
	},
}
