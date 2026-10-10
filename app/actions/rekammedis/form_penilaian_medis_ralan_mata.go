package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanMata penilaian awal medis ralan mata (RMPenilaianAwalMedisRalanMata).
var FormPenilaianMedisRalanMata = &Form[model.PenilaianMedisRalanMata, request.PenilaianMedisRalanMataData]{
	Slug:  "penilaian-medis-ralan-mata",
	Label: "penilaian awal medis ralan mata",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_mata",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanMata, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanMata) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanMata) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanMata) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanMata, d request.PenilaianMedisRalanMataData) error {
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
		m.Status = strings.TrimSpace(d.Status)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Visuskanan = strings.TrimSpace(d.Visuskanan)
		m.Visuskiri = strings.TrimSpace(d.Visuskiri)
		m.Cckanan = strings.TrimSpace(d.Cckanan)
		m.Cckiri = strings.TrimSpace(d.Cckiri)
		m.Palkanan = strings.TrimSpace(d.Palkanan)
		m.Palkiri = strings.TrimSpace(d.Palkiri)
		m.Conkanan = strings.TrimSpace(d.Conkanan)
		m.Conkiri = strings.TrimSpace(d.Conkiri)
		m.Corneakanan = strings.TrimSpace(d.Corneakanan)
		m.Corneakiri = strings.TrimSpace(d.Corneakiri)
		m.Coakanan = strings.TrimSpace(d.Coakanan)
		m.Coakiri = strings.TrimSpace(d.Coakiri)
		m.Pupilkanan = strings.TrimSpace(d.Pupilkanan)
		m.Pupilkiri = strings.TrimSpace(d.Pupilkiri)
		m.Lensakanan = strings.TrimSpace(d.Lensakanan)
		m.Lensakiri = strings.TrimSpace(d.Lensakiri)
		m.Funduskanan = strings.TrimSpace(d.Funduskanan)
		m.Funduskiri = strings.TrimSpace(d.Funduskiri)
		m.Papilkanan = strings.TrimSpace(d.Papilkanan)
		m.Papilkiri = strings.TrimSpace(d.Papilkiri)
		m.Retinakanan = strings.TrimSpace(d.Retinakanan)
		m.Retinakiri = strings.TrimSpace(d.Retinakiri)
		m.Makulakanan = strings.TrimSpace(d.Makulakanan)
		m.Makulakiri = strings.TrimSpace(d.Makulakiri)
		m.Tiokanan = strings.TrimSpace(d.Tiokanan)
		m.Tiokiri = strings.TrimSpace(d.Tiokiri)
		m.Mbokanan = strings.TrimSpace(d.Mbokanan)
		m.Mbokiri = strings.TrimSpace(d.Mbokiri)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Rad = strings.TrimSpace(d.Rad)
		m.Penunjang = strings.TrimSpace(d.Penunjang)
		m.Tes = strings.TrimSpace(d.Tes)
		m.Pemeriksaan = strings.TrimSpace(d.Pemeriksaan)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Diagnosisbdg = strings.TrimSpace(d.Diagnosisbdg)
		m.Permasalahan = strings.TrimSpace(d.Permasalahan)
		m.Terapi = strings.TrimSpace(d.Terapi)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanMata]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanMata) any { return Str(m.KdDokter) }},
	},
}
