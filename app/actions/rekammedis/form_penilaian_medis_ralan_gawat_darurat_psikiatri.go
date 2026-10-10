package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanGawatDaruratPsikiatri penilaian awal medis IGD psikiatri (RMPenilaianAwalMedisIGDPsikiatri).
var FormPenilaianMedisRalanGawatDaruratPsikiatri = &Form[model.PenilaianMedisRalanGawatDaruratPsikiatri, request.PenilaianMedisRalanGawatDaruratPsikiatriData]{
	Slug:  "penilaian-medis-ralan-gawat-darurat-psikiatri",
	Label: "penilaian awal medis IGD psikiatri",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_gawat_darurat_psikiatri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanGawatDaruratPsikiatri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanGawatDaruratPsikiatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanGawatDaruratPsikiatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanGawatDaruratPsikiatri) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMedisRalanGawatDaruratPsikiatri, d request.PenilaianMedisRalanGawatDaruratPsikiatriData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.Hubungan = strings.TrimSpace(d.Hubungan)
		m.KeluhanUtama = support.Nullable(d.KeluhanUtama)
		m.GejalaMenyertai = support.Nullable(d.GejalaMenyertai)
		m.FaktorPencetus = support.Nullable(d.FaktorPencetus)
		m.RiwayatPenyakitDahulu = support.Nullable(d.RiwayatPenyakitDahulu)
		m.KeteranganRiwayatPenyakitDahulu = support.Nullable(d.KeteranganRiwayatPenyakitDahulu)
		m.RiwayatKehamilan = support.Nullable(d.RiwayatKehamilan)
		m.RiwayatSosial = support.Nullable(d.RiwayatSosial)
		m.KeteranganRiwayatSosial = support.Nullable(d.KeteranganRiwayatSosial)
		m.RiwayatPekerjaan = support.Nullable(d.RiwayatPekerjaan)
		m.KeteranganRiwayatPekerjaan = support.Nullable(d.KeteranganRiwayatPekerjaan)
		m.RiwayatObatDiminum = support.Nullable(d.RiwayatObatDiminum)
		m.FaktorKepribadianPremorbid = support.Nullable(d.FaktorKepribadianPremorbid)
		m.FaktorKeturunan = support.Nullable(d.FaktorKeturunan)
		m.KeteranganFaktorKeturunan = support.Nullable(d.KeteranganFaktorKeturunan)
		m.FaktorOrganik = support.Nullable(d.FaktorOrganik)
		m.KeteranganFaktorOrganik = support.Nullable(d.KeteranganFaktorOrganik)
		m.RiwayatAlergi = support.Nullable(d.RiwayatAlergi)
		m.FisikKesadaran = support.Nullable(d.FisikKesadaran)
		m.FisikTd = support.Nullable(d.FisikTd)
		m.FisikRr = support.Nullable(d.FisikRr)
		m.FisikSuhu = support.Nullable(d.FisikSuhu)
		m.FisikNyeri = support.Nullable(d.FisikNyeri)
		m.FisikNadi = support.Nullable(d.FisikNadi)
		m.FisikBb = support.Nullable(d.FisikBb)
		m.FisikTb = support.Nullable(d.FisikTb)
		m.FisikStatusNutrisi = support.Nullable(d.FisikStatusNutrisi)
		m.FisikGcs = support.Nullable(d.FisikGcs)
		m.StatusKelainanKepala = support.Nullable(d.StatusKelainanKepala)
		m.KeteranganStatusKelainanKepala = support.Nullable(d.KeteranganStatusKelainanKepala)
		m.StatusKelainanLeher = support.Nullable(d.StatusKelainanLeher)
		m.KeteranganStatusKelainanLeher = support.Nullable(d.KeteranganStatusKelainanLeher)
		m.StatusKelainanDada = support.Nullable(d.StatusKelainanDada)
		m.KeteranganStatusKelainanDada = support.Nullable(d.KeteranganStatusKelainanDada)
		m.StatusKelainanPerut = support.Nullable(d.StatusKelainanPerut)
		m.KeteranganStatusKelainanPerut = support.Nullable(d.KeteranganStatusKelainanPerut)
		m.StatusKelainanAnggotaGerak = support.Nullable(d.StatusKelainanAnggotaGerak)
		m.KeteranganStatusKelainanAnggotaGerak = support.Nullable(d.KeteranganStatusKelainanAnggotaGerak)
		m.StatusLokalisata = support.Nullable(d.StatusLokalisata)
		m.PsikiatrikKesanUmum = support.Nullable(d.PsikiatrikKesanUmum)
		m.PsikiatrikSikapPrilaku = support.Nullable(d.PsikiatrikSikapPrilaku)
		m.PsikiatrikKesadaran = support.Nullable(d.PsikiatrikKesadaran)
		m.PsikiatrikOrientasi = support.Nullable(d.PsikiatrikOrientasi)
		m.PsikiatrikDayaIngat = support.Nullable(d.PsikiatrikDayaIngat)
		m.PsikiatrikPersepsi = support.Nullable(d.PsikiatrikPersepsi)
		m.PsikiatrikPikiran = support.Nullable(d.PsikiatrikPikiran)
		m.PsikiatrikInsight = support.Nullable(d.PsikiatrikInsight)
		m.Laborat = support.Nullable(d.Laborat)
		m.Radiologi = support.Nullable(d.Radiologi)
		m.Ekg = support.Nullable(d.Ekg)
		m.Diagnosis = support.Nullable(d.Diagnosis)
		m.Permasalahan = support.Nullable(d.Permasalahan)
		m.InstruksiMedis = support.Nullable(d.InstruksiMedis)
		m.RencanaTarget = support.Nullable(d.RencanaTarget)
		m.PulangDipulangkan = support.Nullable(d.PulangDipulangkan)
		m.KeteranganPulangDipulangkan = support.Nullable(d.KeteranganPulangDipulangkan)
		m.PulangDirawatDiruang = support.Nullable(d.PulangDirawatDiruang)
		m.PulangIndikasiRanap = support.Nullable(d.PulangIndikasiRanap)
		m.PulangDirujukKe = support.Nullable(d.PulangDirujukKe)
		m.PulangAlasanDirujuk = support.Nullable(d.PulangAlasanDirujuk)
		m.PulangPaksa = support.Nullable(d.PulangPaksa)
		m.KeteranganPulangPaksa = support.Nullable(d.KeteranganPulangPaksa)
		m.PulangMeninggalIgd = support.Nullable(d.PulangMeninggalIgd)
		m.PulangPenyebabKematian = support.Nullable(d.PulangPenyebabKematian)
		m.FisikPulangKesadaran = support.Nullable(d.FisikPulangKesadaran)
		m.FisikPulangTd = support.Nullable(d.FisikPulangTd)
		m.FisikPulangNadi = support.Nullable(d.FisikPulangNadi)
		m.FisikPulangGcs = support.Nullable(d.FisikPulangGcs)
		m.FisikPulangSuhu = support.Nullable(d.FisikPulangSuhu)
		m.FisikPulangRr = support.Nullable(d.FisikPulangRr)
		m.Edukasi = support.Nullable(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanGawatDaruratPsikiatri]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanGawatDaruratPsikiatri) any { return Str(m.KdDokter) }},
	},
}
