package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanTht penilaian awal medis ralan THT (RMPenilaianAwalMedisRalanTHT).
var FormPenilaianMedisRalanTht = &Form[model.PenilaianMedisRalanTht, request.PenilaianMedisRalanThtData]{
	Slug:  "penilaian-medis-ralan-tht",
	Label: "penilaian awal medis ralan THT",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_tht",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanTht, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanTht) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanTht) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanTht) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanTht, d request.PenilaianMedisRalanThtData) error {
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
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.StatusNutrisi = strings.TrimSpace(d.StatusNutrisi)
		m.Kondisi = strings.TrimSpace(d.Kondisi)
		m.KetLokalis = strings.TrimSpace(d.KetLokalis)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Rad = strings.TrimSpace(d.Rad)
		m.TesPendengaran = strings.TrimSpace(d.TesPendengaran)
		m.Penunjang = strings.TrimSpace(d.Penunjang)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosisbanding = strings.TrimSpace(d.Diagnosisbanding)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Tatalaksana = strings.TrimSpace(d.Tatalaksana)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanTht]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanTht) any { return Str(m.KdDokter) }},
	},
}
