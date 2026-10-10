package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRanap penilaian awal medis ranap dewasa (RMPenilaianAwalMedisRanapDewasa).
var FormPenilaianMedisRanap = &Form[model.PenilaianMedisRanap, request.PenilaianMedisRanapData]{
	Slug:  "penilaian-medis-ranap",
	Label: "penilaian awal medis ranap dewasa",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ranap",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRanap, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRanap) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRanap) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRanap) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRanap, d request.PenilaianMedisRanapData) error {
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
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Keadaan = strings.TrimSpace(d.Keadaan)
		m.Gcs = strings.TrimSpace(d.Gcs)
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Spo = strings.TrimSpace(d.Spo)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Kepala = strings.TrimSpace(d.Kepala)
		m.Mata = strings.TrimSpace(d.Mata)
		m.Gigi = strings.TrimSpace(d.Gigi)
		m.Tht = strings.TrimSpace(d.Tht)
		m.Thoraks = strings.TrimSpace(d.Thoraks)
		m.Jantung = strings.TrimSpace(d.Jantung)
		m.Paru = strings.TrimSpace(d.Paru)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.Genital = strings.TrimSpace(d.Genital)
		m.Ekstremitas = strings.TrimSpace(d.Ekstremitas)
		m.Kulit = strings.TrimSpace(d.Kulit)
		m.KetFisik = strings.TrimSpace(d.KetFisik)
		m.KetLokalis = strings.TrimSpace(d.KetLokalis)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Rad = strings.TrimSpace(d.Rad)
		m.Penunjang = strings.TrimSpace(d.Penunjang)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Tata = strings.TrimSpace(d.Tata)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRanap]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRanap) any { return Str(m.KdDokter) }},
	},
}
