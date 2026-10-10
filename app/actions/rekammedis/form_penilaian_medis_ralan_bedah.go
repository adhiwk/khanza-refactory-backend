package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanBedah penilaian awal medis ralan bedah (RMPenilaianAwalMedisRalanBedah).
var FormPenilaianMedisRalanBedah = &Form[model.PenilaianMedisRalanBedah, request.PenilaianMedisRalanBedahData]{
	Slug:  "penilaian-medis-ralan-bedah",
	Label: "penilaian awal medis ralan bedah",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_bedah",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanBedah, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanBedah) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanBedah) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanBedah) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanBedah, d request.PenilaianMedisRalanBedahData) error {
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
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.Status = strings.TrimSpace(d.Status)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Gcs = strings.TrimSpace(d.Gcs)
		m.Kepala = strings.TrimSpace(d.Kepala)
		m.Thoraks = strings.TrimSpace(d.Thoraks)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.Ekstremitas = strings.TrimSpace(d.Ekstremitas)
		m.Genetalia = strings.TrimSpace(d.Genetalia)
		m.Columna = strings.TrimSpace(d.Columna)
		m.Muskulos = strings.TrimSpace(d.Muskulos)
		m.Lainnya = strings.TrimSpace(d.Lainnya)
		m.KetLokalis = strings.TrimSpace(d.KetLokalis)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Rad = strings.TrimSpace(d.Rad)
		m.Pemeriksaan = strings.TrimSpace(d.Pemeriksaan)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosis2 = strings.TrimSpace(d.Diagnosis2)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanBedah]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanBedah) any { return Str(m.KdDokter) }},
	},
}
