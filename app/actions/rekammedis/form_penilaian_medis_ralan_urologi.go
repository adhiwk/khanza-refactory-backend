package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanUrologi penilaian awal medis ralan urologi (RMPenilaianAwalMedisRalanUrologi).
var FormPenilaianMedisRalanUrologi = &Form[model.PenilaianMedisRalanUrologi, request.PenilaianMedisRalanUrologiData]{
	Slug:  "penilaian-medis-ralan-urologi",
	Label: "penilaian awal medis ralan urologi",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_urologi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanUrologi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanUrologi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanUrologi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanUrologi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanUrologi, d request.PenilaianMedisRalanUrologiData) error {
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
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.RiwayatKebiasaan = strings.TrimSpace(d.RiwayatKebiasaan)
		m.RiwayatOperasiUrologi = strings.TrimSpace(d.RiwayatOperasiUrologi)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Td = strings.TrimSpace(d.Td)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.KeadaanUmum = support.Nullable(d.KeadaanUmum)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.StatusNutrisi = strings.TrimSpace(d.StatusNutrisi)
		m.Thoraks = strings.TrimSpace(d.Thoraks)
		m.KeteranganThoraks = support.Nullable(d.KeteranganThoraks)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.KeteranganAbdomen = support.Nullable(d.KeteranganAbdomen)
		m.Ekstrimitas = strings.TrimSpace(d.Ekstrimitas)
		m.KeteranganEkstrimitas = support.Nullable(d.KeteranganEkstrimitas)
		m.NyeriKetokCva = support.Nullable(d.NyeriKetokCva)
		m.GenitaliaEksternal = support.Nullable(d.GenitaliaEksternal)
		m.ColokDubur = support.Nullable(d.ColokDubur)
		m.Lainnya = strings.TrimSpace(d.Lainnya)
		m.Urinalisis = strings.TrimSpace(d.Urinalisis)
		m.Darah = strings.TrimSpace(d.Darah)
		m.UsgUrologi = strings.TrimSpace(d.UsgUrologi)
		m.Radiologi = strings.TrimSpace(d.Radiologi)
		m.PenunjangLain = strings.TrimSpace(d.PenunjangLain)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosis2 = strings.TrimSpace(d.Diagnosis2)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanUrologi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanUrologi) any { return Str(m.KdDokter) }},
	},
}
