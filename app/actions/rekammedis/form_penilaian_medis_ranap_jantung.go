package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRanapJantung penilaian awal medis ranap jantung (RMPenilaianAwalMedisRanapJantung).
var FormPenilaianMedisRanapJantung = &Form[model.PenilaianMedisRanapJantung, request.PenilaianMedisRanapJantungData]{
	Slug:  "penilaian-medis-ranap-jantung",
	Label: "penilaian awal medis ranap jantung",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ranap_jantung",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRanapJantung, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRanapJantung) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRanapJantung) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRanapJantung) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRanapJantung, d request.PenilaianMedisRanapJantungData) error {
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
		m.Jantung = strings.TrimSpace(d.Jantung)
		m.KeteranganJantung = support.Nullable(d.KeteranganJantung)
		m.Paru = strings.TrimSpace(d.Paru)
		m.KeteranganParu = support.Nullable(d.KeteranganParu)
		m.Ekstrimitas = strings.TrimSpace(d.Ekstrimitas)
		m.KeteranganEkstrimitas = support.Nullable(d.KeteranganEkstrimitas)
		m.Lainnya = strings.TrimSpace(d.Lainnya)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Ekg = strings.TrimSpace(d.Ekg)
		m.PenunjangLain = strings.TrimSpace(d.PenunjangLain)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosis2 = strings.TrimSpace(d.Diagnosis2)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRanapJantung]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRanapJantung) any { return Str(m.KdDokter) }},
	},
}
