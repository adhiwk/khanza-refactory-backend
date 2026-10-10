package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianTerapiWicara penilaian terapi wicara (RMPenilaianTerapiWicara).
var FormPenilaianTerapiWicara = &Form[model.PenilaianTerapiWicara, request.PenilaianTerapiWicaraData]{
	Slug:  "penilaian-terapi-wicara",
	Label: "penilaian terapi wicara",
	Spec: repo.Spec{
		Table:  "penilaian_terapi_wicara",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianTerapiWicara, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianTerapiWicara) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianTerapiWicara) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianTerapiWicara) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianTerapiWicara, d request.PenilaianTerapiWicaraData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.DiagnosaTerapiWicara = support.Nullable(d.DiagnosaTerapiWicara)
		m.DiagnosaMedis = support.Nullable(d.DiagnosaMedis)
		m.Anamnesa = strings.TrimSpace(d.Anamnesa)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Td = strings.TrimSpace(d.Td)
		m.PerilakuAdaptifKontakMata = support.Nullable(d.PerilakuAdaptifKontakMata)
		m.PerilakuAdaptifAtensi = support.Nullable(d.PerilakuAdaptifAtensi)
		m.PerilakuAdaptifPerilaku = support.Nullable(d.PerilakuAdaptifPerilaku)
		m.KemampuanBahasaBicaraSpontan = support.Nullable(d.KemampuanBahasaBicaraSpontan)
		m.KemampuanBahasaPemahamanBahasa = support.Nullable(d.KemampuanBahasaPemahamanBahasa)
		m.KemampuanBahasaPengujaran = support.Nullable(d.KemampuanBahasaPengujaran)
		m.KemampuanBahasaMembaca = support.Nullable(d.KemampuanBahasaMembaca)
		m.KemampuanBahasaPenamaan = support.Nullable(d.KemampuanBahasaPenamaan)
		m.OrganWicaraAnatomisLip = strings.TrimSpace(d.OrganWicaraAnatomisLip)
		m.OrganWicaraAnatomisTongue = strings.TrimSpace(d.OrganWicaraAnatomisTongue)
		m.OrganWicaraAnatomisHardPalate = strings.TrimSpace(d.OrganWicaraAnatomisHardPalate)
		m.OrganWicaraAnatomisSoftPalate = strings.TrimSpace(d.OrganWicaraAnatomisSoftPalate)
		m.OrganWicaraAnatomisUvula = strings.TrimSpace(d.OrganWicaraAnatomisUvula)
		m.OrganWicaraAnatomisMandibula = strings.TrimSpace(d.OrganWicaraAnatomisMandibula)
		m.OrganWicaraAnatomisMaxila = strings.TrimSpace(d.OrganWicaraAnatomisMaxila)
		m.OrganWicaraAnatomisDental = strings.TrimSpace(d.OrganWicaraAnatomisDental)
		m.OrganWicaraAnatomisFaring = strings.TrimSpace(d.OrganWicaraAnatomisFaring)
		m.OrganWicaraFisiologisLip = strings.TrimSpace(d.OrganWicaraFisiologisLip)
		m.OrganWicaraFisiologisTongue = strings.TrimSpace(d.OrganWicaraFisiologisTongue)
		m.OrganWicaraFisiologisHardPalate = strings.TrimSpace(d.OrganWicaraFisiologisHardPalate)
		m.OrganWicaraFisiologisSoftPalate = strings.TrimSpace(d.OrganWicaraFisiologisSoftPalate)
		m.OrganWicaraFisiologisUvula = strings.TrimSpace(d.OrganWicaraFisiologisUvula)
		m.OrganWicaraFisiologisMandibula = strings.TrimSpace(d.OrganWicaraFisiologisMandibula)
		m.OrganWicaraFisiologisMaxilla = strings.TrimSpace(d.OrganWicaraFisiologisMaxilla)
		m.OrganWicaraFisiologisDental = strings.TrimSpace(d.OrganWicaraFisiologisDental)
		m.OrganWicaraFisiologisFaring = strings.TrimSpace(d.OrganWicaraFisiologisFaring)
		m.AktifitasOralMenghisap = strings.TrimSpace(d.AktifitasOralMenghisap)
		m.AktifitasOralMengunyah = strings.TrimSpace(d.AktifitasOralMengunyah)
		m.AktifitasOralMeniup = strings.TrimSpace(d.AktifitasOralMeniup)
		m.KemampuanArtikulasiSubtitusi = strings.TrimSpace(d.KemampuanArtikulasiSubtitusi)
		m.KemampuanArtikulasiOmisi = strings.TrimSpace(d.KemampuanArtikulasiOmisi)
		m.KemampuanArtikulasiDistorsi = strings.TrimSpace(d.KemampuanArtikulasiDistorsi)
		m.KemampuanArtikulasiAdisi = strings.TrimSpace(d.KemampuanArtikulasiAdisi)
		m.Resonasi = strings.TrimSpace(d.Resonasi)
		m.KemampuanSuaraNada = strings.TrimSpace(d.KemampuanSuaraNada)
		m.KemampuanSuaraKualitas = strings.TrimSpace(d.KemampuanSuaraKualitas)
		m.KemampuanSuaraKenyaringan = strings.TrimSpace(d.KemampuanSuaraKenyaringan)
		m.KemampuanIramaKelancaran = strings.TrimSpace(d.KemampuanIramaKelancaran)
		m.KemampuanMenelan = strings.TrimSpace(d.KemampuanMenelan)
		m.Pernafasan = strings.TrimSpace(d.Pernafasan)
		m.TingkatKomunikasiDekodingPendengaran = strings.TrimSpace(d.TingkatKomunikasiDekodingPendengaran)
		m.TingkatKomunikasiDekodingPenglihatan = strings.TrimSpace(d.TingkatKomunikasiDekodingPenglihatan)
		m.TingkatKomunikasiDekodingKinesik = strings.TrimSpace(d.TingkatKomunikasiDekodingKinesik)
		m.TingkatKomunikasiEnkodingBicara = strings.TrimSpace(d.TingkatKomunikasiEnkodingBicara)
		m.TingkatKomunikasiEnkodingTulisan = strings.TrimSpace(d.TingkatKomunikasiEnkodingTulisan)
		m.TingkatKomunikasiEnkodingMimik = strings.TrimSpace(d.TingkatKomunikasiEnkodingMimik)
		m.TingkatKomunikasiEnkodingGesture = strings.TrimSpace(d.TingkatKomunikasiEnkodingGesture)
		m.PenunjangMedis = strings.TrimSpace(d.PenunjangMedis)
		m.PerencanaanTerapiTujuan = strings.TrimSpace(d.PerencanaanTerapiTujuan)
		m.PerencanaanTerapiProgram = strings.TrimSpace(d.PerencanaanTerapiProgram)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		m.TindakLanjut = strings.TrimSpace(d.TindakLanjut)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianTerapiWicara]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianTerapiWicara) any { return Str(m.Nip) }},
	},
}
