package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanPenyakitDalam penilaian awal medis ralan penyakit dalam (RMPenilaianAwalMedisRalanPenyakitDalam).
var FormPenilaianMedisRalanPenyakitDalam = &Form[model.PenilaianMedisRalanPenyakitDalam, request.PenilaianMedisRalanPenyakitDalamData]{
	Slug:  "penilaian-medis-ralan-penyakit-dalam",
	Label: "penilaian awal medis ralan penyakit dalam",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_penyakit_dalam",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanPenyakitDalam, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanPenyakitDalam) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanPenyakitDalam) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanPenyakitDalam) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanPenyakitDalam, d request.PenilaianMedisRalanPenyakitDalamData) error {
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
		m.Kondisi = strings.TrimSpace(d.Kondisi)
		m.Status = strings.TrimSpace(d.Status)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Gcs = strings.TrimSpace(d.Gcs)
		m.Kepala = strings.TrimSpace(d.Kepala)
		m.KeteranganKepala = strings.TrimSpace(d.KeteranganKepala)
		m.Thoraks = strings.TrimSpace(d.Thoraks)
		m.KeteranganThorak = strings.TrimSpace(d.KeteranganThorak)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.KeteranganAbdomen = strings.TrimSpace(d.KeteranganAbdomen)
		m.Ekstremitas = strings.TrimSpace(d.Ekstremitas)
		m.KeteranganEkstremitas = strings.TrimSpace(d.KeteranganEkstremitas)
		m.Lainnya = strings.TrimSpace(d.Lainnya)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Rad = strings.TrimSpace(d.Rad)
		m.Penunjanglain = strings.TrimSpace(d.Penunjanglain)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosis2 = strings.TrimSpace(d.Diagnosis2)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanPenyakitDalam]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanPenyakitDalam) any { return Str(m.KdDokter) }},
	},
}
