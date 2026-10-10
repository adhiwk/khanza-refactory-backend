package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanGeriatri penilaian awal medis ralan geriatri (RMPenilaianAwalMedisRalanGeriatri).
var FormPenilaianMedisRalanGeriatri = &Form[model.PenilaianMedisRalanGeriatri, request.PenilaianMedisRalanGeriatriData]{
	Slug:  "penilaian-medis-ralan-geriatri",
	Label: "penilaian awal medis ralan geriatri",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_geriatri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanGeriatri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanGeriatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanGeriatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanGeriatri) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanGeriatri, d request.PenilaianMedisRalanGeriatriData) error {
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
		m.TulangBelakang = strings.TrimSpace(d.TulangBelakang)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Rr = strings.TrimSpace(d.Rr)
		m.KondisiUmum = strings.TrimSpace(d.KondisiUmum)
		m.StatusPsikologisGds = strings.TrimSpace(d.StatusPsikologisGds)
		m.KondisiSosial = strings.TrimSpace(d.KondisiSosial)
		m.StatusKognitifMmse = strings.TrimSpace(d.StatusKognitifMmse)
		m.Kepala = strings.TrimSpace(d.Kepala)
		m.KeteranganKepala = strings.TrimSpace(d.KeteranganKepala)
		m.Thoraks = strings.TrimSpace(d.Thoraks)
		m.KeteranganThoraks = strings.TrimSpace(d.KeteranganThoraks)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.KeteranganAbdomen = strings.TrimSpace(d.KeteranganAbdomen)
		m.Ekstremitas = strings.TrimSpace(d.Ekstremitas)
		m.KeteranganEkstremitas = strings.TrimSpace(d.KeteranganEkstremitas)
		m.IntegumentKebersihan = strings.TrimSpace(d.IntegumentKebersihan)
		m.IntegumentWarna = strings.TrimSpace(d.IntegumentWarna)
		m.IntegumentKelembaban = strings.TrimSpace(d.IntegumentKelembaban)
		m.IntegumentGangguanKulit = strings.TrimSpace(d.IntegumentGangguanKulit)
		m.StatusFungsional = strings.TrimSpace(d.StatusFungsional)
		m.SkriningJatuh = strings.TrimSpace(d.SkriningJatuh)
		m.StatusNutrisi = strings.TrimSpace(d.StatusNutrisi)
		m.Lainnya = strings.TrimSpace(d.Lainnya)
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
	Refs: []Ref[model.PenilaianMedisRalanGeriatri]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanGeriatri) any { return Str(m.KdDokter) }},
	},
}
