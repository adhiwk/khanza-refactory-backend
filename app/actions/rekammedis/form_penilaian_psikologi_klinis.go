package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPsikologiKlinis penilaian psikologi klinis (RMPenilaianPsikologiKlinis).
var FormPenilaianPsikologiKlinis = &Form[model.PenilaianPsikologiKlinis, request.PenilaianPsikologiKlinisData]{
	Slug:  "penilaian-psikologi-klinis",
	Label: "penilaian psikologi klinis",
	Spec: repo.Spec{
		Table:  "penilaian_psikologi_klinis",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianPsikologiKlinis, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianPsikologiKlinis) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianPsikologiKlinis) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPsikologiKlinis) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianPsikologiKlinis, d request.PenilaianPsikologiKlinisData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		vPsikotesTanggalPelaksanaan, err := support.ParseDate(d.PsikotesTanggalPelaksanaan)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.DikirimDari = strings.TrimSpace(d.DikirimDari)
		m.TujuanPemeriksaan = strings.TrimSpace(d.TujuanPemeriksaan)
		m.KetAnamnesis = support.Nullable(d.KetAnamnesis)
		m.KeluhanUtama = support.Nullable(d.KeluhanUtama)
		m.RiwayatPenyakit = support.Nullable(d.RiwayatPenyakit)
		m.RiwayatKeluhan = support.Nullable(d.RiwayatKeluhan)
		m.PermasalahanSaatIni = support.Nullable(d.PermasalahanSaatIni)
		m.PermasalahanAlasan = support.Nullable(d.PermasalahanAlasan)
		m.PermasalahanEkspektasi = support.Nullable(d.PermasalahanEkspektasi)
		m.RiwayatHidupSingkat = support.Nullable(d.RiwayatHidupSingkat)
		m.KondisiPsikologisPenampilan = support.Nullable(d.KondisiPsikologisPenampilan)
		m.KondisiPsikologisEkspresiWajah = support.Nullable(d.KondisiPsikologisEkspresiWajah)
		m.KondisiPsikologisSuasanaHati = support.Nullable(d.KondisiPsikologisSuasanaHati)
		m.KondisiPsikologisTingkahLaku = support.Nullable(d.KondisiPsikologisTingkahLaku)
		m.KondisiPsikologisFungsiUmum = support.Nullable(d.KondisiPsikologisFungsiUmum)
		m.KondisiPsikologisFungsiIntelektual = support.Nullable(d.KondisiPsikologisFungsiIntelektual)
		m.KondisiPsikologisPengalaman = support.Nullable(d.KondisiPsikologisPengalaman)
		m.KondisiPsikologisLainnya = support.Nullable(d.KondisiPsikologisLainnya)
		m.KondisiPatologisDelusi = support.Nullable(d.KondisiPatologisDelusi)
		m.KondisiPatologisProsesPikiran = support.Nullable(d.KondisiPatologisProsesPikiran)
		m.KondisiPatologisHalusinasi = support.Nullable(d.KondisiPatologisHalusinasi)
		m.KondisiPatologisAfek = support.Nullable(d.KondisiPatologisAfek)
		m.KondisiPatologisInsight = support.Nullable(d.KondisiPatologisInsight)
		m.KondisiPatologisKesadaran = support.Nullable(d.KondisiPatologisKesadaran)
		m.KondisiPatologisOrientasi = support.Nullable(d.KondisiPatologisOrientasi)
		m.KondisiPatologisAtensi = support.Nullable(d.KondisiPatologisAtensi)
		m.KondisiPatologisKontrolImpuls = support.Nullable(d.KondisiPatologisKontrolImpuls)
		m.PsikotesTanggalPelaksanaan = vPsikotesTanggalPelaksanaan
		m.PsikotesNamaTes = support.Nullable(d.PsikotesNamaTes)
		m.PsikotesHasil = support.Nullable(d.PsikotesHasil)
		m.DinamikaPsikologis = support.Nullable(d.DinamikaPsikologis)
		m.DiagnosaPsikologis = support.Nullable(d.DiagnosaPsikologis)
		m.ManifestasiFungsiPsikologis = support.Nullable(d.ManifestasiFungsiPsikologis)
		m.RencanaIntervensi = support.Nullable(d.RencanaIntervensi)
		m.TahapanIntervensi1 = support.Nullable(d.TahapanIntervensi1)
		m.TargetTerapi1 = support.Nullable(d.TargetTerapi1)
		m.TahapanIntervensi2 = support.Nullable(d.TahapanIntervensi2)
		m.TargetTerapi2 = support.Nullable(d.TargetTerapi2)
		m.TahapanIntervensi3 = support.Nullable(d.TahapanIntervensi3)
		m.TargetTerapi3 = support.Nullable(d.TargetTerapi3)
		m.TahapanIntervensi4 = support.Nullable(d.TahapanIntervensi4)
		m.TargetTerapi4 = support.Nullable(d.TargetTerapi4)
		m.TahapanIntervensi5 = support.Nullable(d.TahapanIntervensi5)
		m.TargetTerapi5 = support.Nullable(d.TargetTerapi5)
		m.TahapanIntervensi6 = support.Nullable(d.TahapanIntervensi6)
		m.TargetTerapi6 = support.Nullable(d.TargetTerapi6)
		m.TahapanIntervensi7 = support.Nullable(d.TahapanIntervensi7)
		m.TargetTerapi7 = support.Nullable(d.TargetTerapi7)
		m.Evaluasi = support.Nullable(d.Evaluasi)
		return nil
	},
	Refs: []Ref[model.PenilaianPsikologiKlinis]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianPsikologiKlinis) any { return Str(m.Nip) }},
	},
}
