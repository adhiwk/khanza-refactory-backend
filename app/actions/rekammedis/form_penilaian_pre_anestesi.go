package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPreAnestesi penilaian pre anastesi (RMPenilaianPreAnastesi).
var FormPenilaianPreAnestesi = &Form[model.PenilaianPreAnestesi, request.PenilaianPreAnestesiData]{
	Slug:  "penilaian-pre-anestesi",
	Label: "penilaian pre anastesi",
	Spec: repo.Spec{
		Table:  "penilaian_pre_anestesi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianPreAnestesi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianPreAnestesi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianPreAnestesi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPreAnestesi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianPreAnestesi, d request.PenilaianPreAnestesiData) error {
		vTanggalOperasi, err := support.ParseDateTime(d.TanggalOperasi)
		if err != nil {
			return err
		}
		vPuasa, err := support.ParseDateTime(d.Puasa)
		if err != nil {
			return err
		}
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.TanggalOperasi = vTanggalOperasi
		m.Diagnosa = support.Nullable(d.Diagnosa)
		m.RencanaTindakan = support.Nullable(d.RencanaTindakan)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Td = strings.TrimSpace(d.Td)
		m.Io2 = strings.TrimSpace(d.Io2)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Pernapasan = strings.TrimSpace(d.Pernapasan)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.FisikCardiovasculer = support.Nullable(d.FisikCardiovasculer)
		m.FisikParu = support.Nullable(d.FisikParu)
		m.FisikAbdomen = support.Nullable(d.FisikAbdomen)
		m.FisikExtrimitas = support.Nullable(d.FisikExtrimitas)
		m.FisikEndokrin = support.Nullable(d.FisikEndokrin)
		m.FisikGinjal = support.Nullable(d.FisikGinjal)
		m.FisikObatobatan = support.Nullable(d.FisikObatobatan)
		m.FisikLaborat = support.Nullable(d.FisikLaborat)
		m.FisikPenunjang = support.Nullable(d.FisikPenunjang)
		m.RiwayatPenyakitAlergiobat = support.Nullable(d.RiwayatPenyakitAlergiobat)
		m.RiwayatPenyakitAlergilainnya = support.Nullable(d.RiwayatPenyakitAlergilainnya)
		m.RiwayatPenyakitTerapi = support.Nullable(d.RiwayatPenyakitTerapi)
		m.RiwayatKebiasaanMerokok = strings.TrimSpace(d.RiwayatKebiasaanMerokok)
		m.RiwayatKebiasaanKetMerokok = strings.TrimSpace(d.RiwayatKebiasaanKetMerokok)
		m.RiwayatKebiasaanAlkohol = strings.TrimSpace(d.RiwayatKebiasaanAlkohol)
		m.RiwayatKebiasaanKetAlkohol = strings.TrimSpace(d.RiwayatKebiasaanKetAlkohol)
		m.RiwayatKebiasaanObat = strings.TrimSpace(d.RiwayatKebiasaanObat)
		m.RiwayatKebiasaanKetObat = strings.TrimSpace(d.RiwayatKebiasaanKetObat)
		m.RiwayatMedisCardiovasculer = support.Nullable(d.RiwayatMedisCardiovasculer)
		m.RiwayatMedisRespiratory = support.Nullable(d.RiwayatMedisRespiratory)
		m.RiwayatMedisEndocrine = support.Nullable(d.RiwayatMedisEndocrine)
		m.RiwayatMedisLainnya = support.Nullable(d.RiwayatMedisLainnya)
		m.Asa = support.Nullable(d.Asa)
		m.Puasa = vPuasa
		m.RencanaAnestesi = support.Nullable(d.RencanaAnestesi)
		m.RencanaPerawatan = support.Nullable(d.RencanaPerawatan)
		m.CatatanKhusus = support.Nullable(d.CatatanKhusus)
		return nil
	},
	Refs: []Ref[model.PenilaianPreAnestesi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianPreAnestesi) any { return Str(m.KdDokter) }},
	},
}
