package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianBayiBaruLahir penilaian bayi baru lahir (RMPenilaianBayiBaruLahir).
var FormPenilaianBayiBaruLahir = &Form[model.PenilaianBayiBaruLahir, request.PenilaianBayiBaruLahirData]{
	Slug:  "penilaian-bayi-baru-lahir",
	Label: "penilaian bayi baru lahir",
	Spec: repo.Spec{
		Table:  "penilaian_bayi_baru_lahir",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianBayiBaruLahir, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianBayiBaruLahir) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianBayiBaruLahir) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianBayiBaruLahir) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianBayiBaruLahir, d request.PenilaianBayiBaruLahirData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.NoRkmMedisIbu = support.Nullable(d.NoRkmMedisIbu)
		m.PenyakitDideritaIbu = support.Nullable(d.PenyakitDideritaIbu)
		m.KeteranganPenyakitDideritaIbu = support.Nullable(d.KeteranganPenyakitDideritaIbu)
		m.ObatDikonsumsiSelamaKehamilan = support.Nullable(d.ObatDikonsumsiSelamaKehamilan)
		m.PerawatanAntenatal = support.Nullable(d.PerawatanAntenatal)
		m.KeteranganPerawatanAntenatal = support.Nullable(d.KeteranganPerawatanAntenatal)
		m.TerdaftarEkohort = support.Nullable(d.TerdaftarEkohort)
		m.KeteranganTerdaftarEkohort = support.Nullable(d.KeteranganTerdaftarEkohort)
		m.PenyulitKehamilan = support.Nullable(d.PenyulitKehamilan)
		m.KeteranganPenyulitKehamilan = support.Nullable(d.KeteranganPenyulitKehamilan)
		m.Alergi = support.Nullable(d.Alergi)
		m.KeteranganLainnyaRiwayatMaternal = support.Nullable(d.KeteranganLainnyaRiwayatMaternal)
		m.UmurKehamilan = support.Nullable(d.UmurKehamilan)
		m.Kehamilan = support.Nullable(d.Kehamilan)
		m.KeteranganKehamilan = support.Nullable(d.KeteranganKehamilan)
		m.UrutanKehamilan = support.Nullable(d.UrutanKehamilan)
		m.JamKetubanPecah = support.Nullable(d.JamKetubanPecah)
		m.MenitKetubanPecah = support.Nullable(d.MenitKetubanPecah)
		m.JumlahAirKetuban = support.Nullable(d.JumlahAirKetuban)
		m.WarnaAirKetuban = support.Nullable(d.WarnaAirKetuban)
		m.BauAirKetuban = support.Nullable(d.BauAirKetuban)
		m.LetakBayi = support.Nullable(d.LetakBayi)
		m.MacamPersalinan = support.Nullable(d.MacamPersalinan)
		m.KeteranganMacamPersalinan = support.Nullable(d.KeteranganMacamPersalinan)
		m.IndikasiPersalinanOperatif = support.Nullable(d.IndikasiPersalinanOperatif)
		m.KeteranganIndikasiPersalinanOperatif = support.Nullable(d.KeteranganIndikasiPersalinanOperatif)
		m.LamaGawatJanin = support.Nullable(d.LamaGawatJanin)
		m.ObatSelamaPersalinan = support.Nullable(d.ObatSelamaPersalinan)
		m.BeratPlacenta = support.Nullable(d.BeratPlacenta)
		m.KelainanPlacenta = support.Nullable(d.KelainanPlacenta)
		m.KeteranganLainnyaRiwayatPersalinan = support.Nullable(d.KeteranganLainnyaRiwayatPersalinan)
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
		m.Bblahir = support.Nullable(d.Bblahir)
		m.PanjangBadan = support.Nullable(d.PanjangBadan)
		m.LingkarKepala = support.Nullable(d.LingkarKepala)
		m.LingkarDada = support.Nullable(d.LingkarDada)
		m.ResusitasiSaatLahir = support.Nullable(d.ResusitasiSaatLahir)
		m.KeteranganResusitasiSaatLahir = support.Nullable(d.KeteranganResusitasiSaatLahir)
		m.ObatDiberikanSaatLahir = support.Nullable(d.ObatDiberikanSaatLahir)
		m.KeteranganLainnyaKeadaanBayi = support.Nullable(d.KeteranganLainnyaKeadaanBayi)
		m.KondisiUmum = support.Nullable(d.KondisiUmum)
		m.KeteranganKondisiUmum = support.Nullable(d.KeteranganKondisiUmum)
		m.Kulit = support.Nullable(d.Kulit)
		m.KeteranganKulit = support.Nullable(d.KeteranganKulit)
		m.Kepala = support.Nullable(d.Kepala)
		m.KeteranganKepala = support.Nullable(d.KeteranganKepala)
		m.Leher = support.Nullable(d.Leher)
		m.KeteranganLeher = support.Nullable(d.KeteranganLeher)
		m.Mata = support.Nullable(d.Mata)
		m.KeteranganMata = support.Nullable(d.KeteranganMata)
		m.Hidung = support.Nullable(d.Hidung)
		m.KeteranganHidung = support.Nullable(d.KeteranganHidung)
		m.Telinga = support.Nullable(d.Telinga)
		m.KeteranganTelinga = support.Nullable(d.KeteranganTelinga)
		m.Dada = support.Nullable(d.Dada)
		m.KeteranganDada = support.Nullable(d.KeteranganDada)
		m.Paru = support.Nullable(d.Paru)
		m.KeteranganParu = support.Nullable(d.KeteranganParu)
		m.Jantung = support.Nullable(d.Jantung)
		m.KeteranganJantung = support.Nullable(d.KeteranganJantung)
		m.Perut = support.Nullable(d.Perut)
		m.KeteranganPerut = support.Nullable(d.KeteranganPerut)
		m.TaliPusat = support.Nullable(d.TaliPusat)
		m.KeteranganTaliPusat = support.Nullable(d.KeteranganTaliPusat)
		m.AlatKelamin = support.Nullable(d.AlatKelamin)
		m.KeteranganAlatKelamin = support.Nullable(d.KeteranganAlatKelamin)
		m.RuasTulangBelakang = support.Nullable(d.RuasTulangBelakang)
		m.KeteranganRuasTulangBelakang = support.Nullable(d.KeteranganRuasTulangBelakang)
		m.Extrimitas = support.Nullable(d.Extrimitas)
		m.KeteranganExtrimitas = support.Nullable(d.KeteranganExtrimitas)
		m.Anus = support.Nullable(d.Anus)
		m.KeteranganAnus = support.Nullable(d.KeteranganAnus)
		m.Refleks = support.Nullable(d.Refleks)
		m.KeteranganRefleks = support.Nullable(d.KeteranganRefleks)
		m.DenyutFemoral = support.Nullable(d.DenyutFemoral)
		m.KeteranganDenyutFemoral = support.Nullable(d.KeteranganDenyutFemoral)
		m.PemeriksaanFisikLainnya = support.Nullable(d.PemeriksaanFisikLainnya)
		m.PemeriksaanPenunjang = support.Nullable(d.PemeriksaanPenunjang)
		m.Diagnosa = support.Nullable(d.Diagnosa)
		m.Tatalaksana = support.Nullable(d.Tatalaksana)
		return nil
	},
	Refs: []Ref[model.PenilaianBayiBaruLahir]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianBayiBaruLahir) any { return Str(m.KdDokter) }},
		{Column: "no_rkm_medis_ibu", Table: "pasien", RefCol: "no_rkm_medis", Label: "pasien", Value: func(m *model.PenilaianBayiBaruLahir) any { return StrPtr(m.NoRkmMedisIbu) }},
	},
}
