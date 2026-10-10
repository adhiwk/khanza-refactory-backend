package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanBedahMulut penilaian awal medis ralan bedah mulut (RMPenilaianAwalMedisRalanBedahMulut).
var FormPenilaianMedisRalanBedahMulut = &Form[model.PenilaianMedisRalanBedahMulut, request.PenilaianMedisRalanBedahMulutData]{
	Slug:  "penilaian-medis-ralan-bedah-mulut",
	Label: "penilaian awal medis ralan bedah mulut",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_bedah_mulut",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanBedahMulut, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanBedahMulut) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanBedahMulut) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanBedahMulut) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanBedahMulut, d request.PenilaianMedisRalanBedahMulutData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.Hubungan = strings.TrimSpace(d.Hubungan)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Rps = strings.TrimSpace(d.Rps)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Keadaan = strings.TrimSpace(d.Keadaan)
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.StatusNutrisi = strings.TrimSpace(d.StatusNutrisi)
		m.Kulit = strings.TrimSpace(d.Kulit)
		m.KeteranganKulit = strings.TrimSpace(d.KeteranganKulit)
		m.Kepala = strings.TrimSpace(d.Kepala)
		m.KeteranganKepala = strings.TrimSpace(d.KeteranganKepala)
		m.Mata = strings.TrimSpace(d.Mata)
		m.KeteranganMata = strings.TrimSpace(d.KeteranganMata)
		m.Leher = strings.TrimSpace(d.Leher)
		m.KeteranganLeher = strings.TrimSpace(d.KeteranganLeher)
		m.Kelenjar = strings.TrimSpace(d.Kelenjar)
		m.KeteranganKelenjar = strings.TrimSpace(d.KeteranganKelenjar)
		m.Dada = strings.TrimSpace(d.Dada)
		m.KeteranganDada = strings.TrimSpace(d.KeteranganDada)
		m.Perut = strings.TrimSpace(d.Perut)
		m.KeteranganPerut = strings.TrimSpace(d.KeteranganPerut)
		m.Ekstremitas = strings.TrimSpace(d.Ekstremitas)
		m.KeteranganEkstremitas = strings.TrimSpace(d.KeteranganEkstremitas)
		m.Wajah = strings.TrimSpace(d.Wajah)
		m.Intra = strings.TrimSpace(d.Intra)
		m.Gigigeligi = strings.TrimSpace(d.Gigigeligi)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Rad = strings.TrimSpace(d.Rad)
		m.Penunjang = strings.TrimSpace(d.Penunjang)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosis2 = strings.TrimSpace(d.Diagnosis2)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanBedahMulut]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanBedahMulut) any { return Str(m.KdDokter) }},
	},
}
