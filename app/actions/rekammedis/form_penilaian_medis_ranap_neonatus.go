package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRanapNeonatus penilaian awal medis ranap neonatus (RMPenilaianAwalMedisRanapNeonatus).
var FormPenilaianMedisRanapNeonatus = &Form[model.PenilaianMedisRanapNeonatus, request.PenilaianMedisRanapNeonatusData]{
	Slug:  "penilaian-medis-ranap-neonatus",
	Label: "penilaian awal medis ranap neonatus",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ranap_neonatus",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRanapNeonatus, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRanapNeonatus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRanapNeonatus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRanapNeonatus) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRanapNeonatus, d request.PenilaianMedisRanapNeonatusData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		vTanggalPersalinan, err := support.ParseDateTime(d.TanggalPersalinan)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.NoRkmMedisIbu = support.Nullable(d.NoRkmMedisIbu)
		m.G = support.Nullable(d.G)
		m.P = support.Nullable(d.P)
		m.A = support.Nullable(d.A)
		m.Hidup = support.Nullable(d.Hidup)
		m.Usiahamil = support.Nullable(d.Usiahamil)
		m.Hbsag = support.Nullable(d.Hbsag)
		m.Hiv = support.Nullable(d.Hiv)
		m.Syphilis = support.Nullable(d.Syphilis)
		m.RiwayatObstetriIbu = support.Nullable(d.RiwayatObstetriIbu)
		m.KeteranganRiwayatObstetriIbu = support.Nullable(d.KeteranganRiwayatObstetriIbu)
		m.FaktorRisikoNeonatal = support.Nullable(d.FaktorRisikoNeonatal)
		m.KeteranganFaktorRisikoNeonatal = support.Nullable(d.KeteranganFaktorRisikoNeonatal)
		m.TanggalPersalinan = vTanggalPersalinan
		m.BersalinDi = support.Nullable(d.BersalinDi)
		m.InisiasiMenyusui = support.Nullable(d.InisiasiMenyusui)
		m.JenisPersalinan = support.Nullable(d.JenisPersalinan)
		m.Indikasi = support.Nullable(d.Indikasi)
		m.Aterm = support.Nullable(d.Aterm)
		m.Bernafas = support.Nullable(d.Bernafas)
		m.TanusOtot = support.Nullable(d.TanusOtot)
		m.CairanAmnion = support.Nullable(d.CairanAmnion)
		m.F1 = strings.TrimSpace(d.F1)
		m.U1 = strings.TrimSpace(d.U1)
		m.T1 = strings.TrimSpace(d.T1)
		m.R1 = strings.TrimSpace(d.R1)
		m.W1 = strings.TrimSpace(d.W1)
		m.N1 = strings.TrimSpace(d.N1)
		m.F5 = strings.TrimSpace(d.F5)
		m.U5 = strings.TrimSpace(d.U5)
		m.T5 = strings.TrimSpace(d.T5)
		m.R5 = strings.TrimSpace(d.R5)
		m.W5 = strings.TrimSpace(d.W5)
		m.N5 = strings.TrimSpace(d.N5)
		m.F10 = strings.TrimSpace(d.F10)
		m.U10 = strings.TrimSpace(d.U10)
		m.T10 = strings.TrimSpace(d.T10)
		m.R10 = strings.TrimSpace(d.R10)
		m.W10 = strings.TrimSpace(d.W10)
		m.N10 = strings.TrimSpace(d.N10)
		m.FrekuensiNapas = support.Nullable(d.FrekuensiNapas)
		m.NilaiFrekuensiNapas = d.NilaiFrekuensiNapas
		m.Retraksi = support.Nullable(d.Retraksi)
		m.NilaiRetraksi = d.NilaiRetraksi
		m.Sianosis = support.Nullable(d.Sianosis)
		m.NilaiSianosis = d.NilaiSianosis
		m.JalanMasukUdara = support.Nullable(d.JalanMasukUdara)
		m.NilaiJalanMasukUdara = d.NilaiJalanMasukUdara
		m.Grunting = support.Nullable(d.Grunting)
		m.NilaiGrunting = d.NilaiGrunting
		m.TotalDownScore = d.TotalDownScore
		m.KeteranganDownScore = support.Nullable(d.KeteranganDownScore)
		m.Nadi = support.Nullable(d.Nadi)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Saturasi = support.Nullable(d.Saturasi)
		m.Bb = support.Nullable(d.Bb)
		m.Pb = support.Nullable(d.Pb)
		m.Lk = support.Nullable(d.Lk)
		m.Ld = support.Nullable(d.Ld)
		m.KeadaanUmum = strings.TrimSpace(d.KeadaanUmum)
		m.KeteranganKeadaanUmum = support.Nullable(d.KeteranganKeadaanUmum)
		m.Kulit = strings.TrimSpace(d.Kulit)
		m.KeteranganKulit = support.Nullable(d.KeteranganKulit)
		m.Kepala = strings.TrimSpace(d.Kepala)
		m.KeteranganKepala = support.Nullable(d.KeteranganKepala)
		m.Mata = strings.TrimSpace(d.Mata)
		m.KeteranganMata = support.Nullable(d.KeteranganMata)
		m.Telinga = strings.TrimSpace(d.Telinga)
		m.KeteranganTelinga = support.Nullable(d.KeteranganTelinga)
		m.Hidung = strings.TrimSpace(d.Hidung)
		m.KeteranganHidung = support.Nullable(d.KeteranganHidung)
		m.Mulut = strings.TrimSpace(d.Mulut)
		m.KeteranganMulut = support.Nullable(d.KeteranganMulut)
		m.Tenggorokan = strings.TrimSpace(d.Tenggorokan)
		m.KeteranganTenggorokan = support.Nullable(d.KeteranganTenggorokan)
		m.Leher = strings.TrimSpace(d.Leher)
		m.KeteranganLeher = support.Nullable(d.KeteranganLeher)
		m.Thorax = strings.TrimSpace(d.Thorax)
		m.KeteranganThorax = support.Nullable(d.KeteranganThorax)
		m.Abdomen = strings.TrimSpace(d.Abdomen)
		m.KeteranganAbdomen = support.Nullable(d.KeteranganAbdomen)
		m.Genitalia = strings.TrimSpace(d.Genitalia)
		m.KeteranganGenitalia = support.Nullable(d.KeteranganGenitalia)
		m.Anus = strings.TrimSpace(d.Anus)
		m.KeteranganAnus = support.Nullable(d.KeteranganAnus)
		m.Muskulos = strings.TrimSpace(d.Muskulos)
		m.KeteranganMuskulos = support.Nullable(d.KeteranganMuskulos)
		m.Ekstrimitas = strings.TrimSpace(d.Ekstrimitas)
		m.KeteranganEkstrimitas = support.Nullable(d.KeteranganEkstrimitas)
		m.Paru = strings.TrimSpace(d.Paru)
		m.KeteranganParu = support.Nullable(d.KeteranganParu)
		m.Refleks = strings.TrimSpace(d.Refleks)
		m.KeteranganRefleks = support.Nullable(d.KeteranganRefleks)
		m.KelainanLainnya = support.Nullable(d.KelainanLainnya)
		m.PemeriksaanRegional = strings.TrimSpace(d.PemeriksaanRegional)
		m.Lab = strings.TrimSpace(d.Lab)
		m.Radiologi = strings.TrimSpace(d.Radiologi)
		m.Penunjanglainnya = strings.TrimSpace(d.Penunjanglainnya)
		m.Diagnosis = strings.TrimSpace(d.Diagnosis)
		m.Tata = strings.TrimSpace(d.Tata)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRanapNeonatus]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRanapNeonatus) any { return Str(m.KdDokter) }},
		{Column: "no_rkm_medis_ibu", Table: "pasien", RefCol: "no_rkm_medis", Label: "pasien", Value: func(m *model.PenilaianMedisRanapNeonatus) any { return StrPtr(m.NoRkmMedisIbu) }},
	},
}
