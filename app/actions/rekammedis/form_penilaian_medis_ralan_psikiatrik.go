package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanPsikiatrik penilaian awal medis ralan psikiatrik (RMPenilaianAwalMedisRalanPsikiatrik).
var FormPenilaianMedisRalanPsikiatrik = &Form[model.PenilaianMedisRalanPsikiatrik, request.PenilaianMedisRalanPsikiatrikData]{
	Slug:  "penilaian-medis-ralan-psikiatrik",
	Label: "penilaian awal medis ralan psikiatrik",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_psikiatrik",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanPsikiatrik, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanPsikiatrik) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanPsikiatrik) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanPsikiatrik) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanPsikiatrik, d request.PenilaianMedisRalanPsikiatrikData) error {
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
		m.Penampilan = strings.TrimSpace(d.Penampilan)
		m.Pembicaraan = strings.TrimSpace(d.Pembicaraan)
		m.Psikomotor = strings.TrimSpace(d.Psikomotor)
		m.Sikap = strings.TrimSpace(d.Sikap)
		m.Mood = strings.TrimSpace(d.Mood)
		m.FungsiKognitif = strings.TrimSpace(d.FungsiKognitif)
		m.GangguanPersepsi = strings.TrimSpace(d.GangguanPersepsi)
		m.ProsesPikir = strings.TrimSpace(d.ProsesPikir)
		m.PengendalianImpuls = strings.TrimSpace(d.PengendalianImpuls)
		m.Tilikan = strings.TrimSpace(d.Tilikan)
		m.Rta = strings.TrimSpace(d.Rta)
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
		m.Gigi = strings.TrimSpace(d.Gigi)
		m.Tht = strings.TrimSpace(d.Tht)
		m.Thoraks = strings.TrimSpace(d.Thoraks)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.Genital = strings.TrimSpace(d.Genital)
		m.Ekstremitas = strings.TrimSpace(d.Ekstremitas)
		m.Kulit = strings.TrimSpace(d.Kulit)
		m.KetFisik = strings.TrimSpace(d.KetFisik)
		m.Penunjang = strings.TrimSpace(d.Penunjang)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Tata = strings.TrimSpace(d.Tata)
		m.Konsulrujuk = strings.TrimSpace(d.Konsulrujuk)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanPsikiatrik]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanPsikiatrik) any { return Str(m.KdDokter) }},
	},
}
