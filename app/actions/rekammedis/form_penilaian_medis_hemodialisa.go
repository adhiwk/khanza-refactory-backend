package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisHemodialisa penilaian awal medis hemodialisa (RMPenilaianAwalMedisHemodialisa).
var FormPenilaianMedisHemodialisa = &Form[model.PenilaianMedisHemodialisa, request.PenilaianMedisHemodialisaData]{
	Slug:  "penilaian-medis-hemodialisa",
	Label: "penilaian awal medis hemodialisa",
	Spec: repo.Spec{
		Table:  "penilaian_medis_hemodialisa",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisHemodialisa, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisHemodialisa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisHemodialisa) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisHemodialisa) []string { return []string{derefStr(m.KdDokter)} },
	Fill: func(m *model.PenilaianMedisHemodialisa, d request.PenilaianMedisHemodialisaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		vDialisisPertama, err := support.ParseDate(d.DialisisPertama)
		if err != nil {
			return err
		}
		vTanggalCpad, err := support.ParseDate(d.TanggalCpad)
		if err != nil {
			return err
		}
		vTanggalTransplantasi, err := support.ParseDate(d.TanggalTransplantasi)
		if err != nil {
			return err
		}
		vTanggalThorax, err := support.ParseDate(d.TanggalThorax)
		if err != nil {
			return err
		}
		vTanggalEkg, err := support.ParseDate(d.TanggalEkg)
		if err != nil {
			return err
		}
		vTanggalBno, err := support.ParseDate(d.TanggalBno)
		if err != nil {
			return err
		}
		vTanggalUsg, err := support.ParseDate(d.TanggalUsg)
		if err != nil {
			return err
		}
		vTanggalRenogram, err := support.ParseDate(d.TanggalRenogram)
		if err != nil {
			return err
		}
		vTanggalBiopsi, err := support.ParseDate(d.TanggalBiopsi)
		if err != nil {
			return err
		}
		vTanggalCtscan, err := support.ParseDate(d.TanggalCtscan)
		if err != nil {
			return err
		}
		vTanggalArteriografi, err := support.ParseDate(d.TanggalArteriografi)
		if err != nil {
			return err
		}
		vTanggalKulturUrin, err := support.ParseDate(d.TanggalKulturUrin)
		if err != nil {
			return err
		}
		vTanggalLaborat, err := support.ParseDate(d.TanggalLaborat)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = support.Nullable(d.KdDokter)
		m.Anamnesis = support.Nullable(d.Anamnesis)
		m.Hubungan = support.Nullable(d.Hubungan)
		m.Ruangan = support.Nullable(d.Ruangan)
		m.Alergi = support.Nullable(d.Alergi)
		m.Nyeri = support.Nullable(d.Nyeri)
		m.StatusNutrisi = support.Nullable(d.StatusNutrisi)
		m.Hipertensi = support.Nullable(d.Hipertensi)
		m.KeteranganHipertensi = support.Nullable(d.KeteranganHipertensi)
		m.Diabetes = support.Nullable(d.Diabetes)
		m.KeteranganDiabetes = support.Nullable(d.KeteranganDiabetes)
		m.BatuSaluranKemih = support.Nullable(d.BatuSaluranKemih)
		m.KeteranganBatuSaluranKemih = support.Nullable(d.KeteranganBatuSaluranKemih)
		m.OperasiSaluranKemih = support.Nullable(d.OperasiSaluranKemih)
		m.KeteranganOperasiSaluranKemih = support.Nullable(d.KeteranganOperasiSaluranKemih)
		m.InfeksiSaluranKemih = support.Nullable(d.InfeksiSaluranKemih)
		m.KeteranganInfeksiSaluranKemih = support.Nullable(d.KeteranganInfeksiSaluranKemih)
		m.BengkakSeluruhTubuh = support.Nullable(d.BengkakSeluruhTubuh)
		m.KeteranganBengkakSeluruhTubuh = support.Nullable(d.KeteranganBengkakSeluruhTubuh)
		m.UrinBerdarah = support.Nullable(d.UrinBerdarah)
		m.KeteranganUrinBerdarah = support.Nullable(d.KeteranganUrinBerdarah)
		m.PenyakitGinjalLaom = support.Nullable(d.PenyakitGinjalLaom)
		m.KeteranganPenyakitGinjalLaom = support.Nullable(d.KeteranganPenyakitGinjalLaom)
		m.PenyakitLain = support.Nullable(d.PenyakitLain)
		m.KeteranganPenyakitLain = support.Nullable(d.KeteranganPenyakitLain)
		m.KonsumsiObatNefro = support.Nullable(d.KonsumsiObatNefro)
		m.KeteranganKonsumsiObatNefro = support.Nullable(d.KeteranganKonsumsiObatNefro)
		m.DialisisPertama = vDialisisPertama
		m.PernahCpad = strings.TrimSpace(d.PernahCpad)
		m.TanggalCpad = vTanggalCpad
		m.PernahTransplantasi = strings.TrimSpace(d.PernahTransplantasi)
		m.TanggalTransplantasi = vTanggalTransplantasi
		m.KeadaanUmum = strings.TrimSpace(d.KeadaanUmum)
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Td = strings.TrimSpace(d.Td)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Napas = strings.TrimSpace(d.Napas)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Hepatomegali = strings.TrimSpace(d.Hepatomegali)
		m.Splenomegali = strings.TrimSpace(d.Splenomegali)
		m.Ascites = strings.TrimSpace(d.Ascites)
		m.Edema = strings.TrimSpace(d.Edema)
		m.Whezzing = strings.TrimSpace(d.Whezzing)
		m.Ronchi = strings.TrimSpace(d.Ronchi)
		m.Ikterik = strings.TrimSpace(d.Ikterik)
		m.TekananVena = strings.TrimSpace(d.TekananVena)
		m.Anemia = strings.TrimSpace(d.Anemia)
		m.Kardiomegali = strings.TrimSpace(d.Kardiomegali)
		m.Bising = strings.TrimSpace(d.Bising)
		m.Thorax = strings.TrimSpace(d.Thorax)
		m.TanggalThorax = vTanggalThorax
		m.Ekg = strings.TrimSpace(d.Ekg)
		m.TanggalEkg = vTanggalEkg
		m.Bno = strings.TrimSpace(d.Bno)
		m.TanggalBno = vTanggalBno
		m.Usg = strings.TrimSpace(d.Usg)
		m.TanggalUsg = vTanggalUsg
		m.Renogram = strings.TrimSpace(d.Renogram)
		m.TanggalRenogram = vTanggalRenogram
		m.Biopsi = strings.TrimSpace(d.Biopsi)
		m.TanggalBiopsi = vTanggalBiopsi
		m.Ctscan = strings.TrimSpace(d.Ctscan)
		m.TanggalCtscan = vTanggalCtscan
		m.Arteriografi = strings.TrimSpace(d.Arteriografi)
		m.TanggalArteriografi = vTanggalArteriografi
		m.KulturUrin = strings.TrimSpace(d.KulturUrin)
		m.TanggalKulturUrin = vTanggalKulturUrin
		m.Laborat = strings.TrimSpace(d.Laborat)
		m.TanggalLaborat = vTanggalLaborat
		m.Hematokrit = strings.TrimSpace(d.Hematokrit)
		m.Hemoglobin = strings.TrimSpace(d.Hemoglobin)
		m.Leukosit = strings.TrimSpace(d.Leukosit)
		m.Trombosit = strings.TrimSpace(d.Trombosit)
		m.HitungJenis = strings.TrimSpace(d.HitungJenis)
		m.Ureum = strings.TrimSpace(d.Ureum)
		m.UrinLengkap = strings.TrimSpace(d.UrinLengkap)
		m.Kreatinin = strings.TrimSpace(d.Kreatinin)
		m.Cct = strings.TrimSpace(d.Cct)
		m.Sgot = strings.TrimSpace(d.Sgot)
		m.Sgpt = strings.TrimSpace(d.Sgpt)
		m.Ct = strings.TrimSpace(d.Ct)
		m.AsamUrat = strings.TrimSpace(d.AsamUrat)
		m.Hbsag = strings.TrimSpace(d.Hbsag)
		m.AntiHcv = strings.TrimSpace(d.AntiHcv)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisHemodialisa]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisHemodialisa) any { return StrPtr(m.KdDokter) }},
	},
}
