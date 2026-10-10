package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormEdukasiPasienKeluargaRj edukasi pasien keluarga rawat jalan (RMEdukasiPasienKeluargaRawatJalan).
var FormEdukasiPasienKeluargaRj = &Form[model.EdukasiPasienKeluargaRj, request.EdukasiPasienKeluargaRjData]{
	Slug:  "edukasi-pasien-keluarga-rj",
	Label: "edukasi pasien keluarga rawat jalan",
	Spec: repo.Spec{
		Table:  "edukasi_pasien_keluarga_rj",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.EdukasiPasienKeluargaRj, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.EdukasiPasienKeluargaRj) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.EdukasiPasienKeluargaRj) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.EdukasiPasienKeluargaRj) []string { return []string{m.Nip} },
	Fill: func(m *model.EdukasiPasienKeluargaRj, d request.EdukasiPasienKeluargaRjData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Bicara = support.Nullable(d.Bicara)
		m.KeteranganBicara = support.Nullable(d.KeteranganBicara)
		m.BahasaSehari = support.Nullable(d.BahasaSehari)
		m.PerluPenerjemah = support.Nullable(d.PerluPenerjemah)
		m.KeteranganPenerjemah = support.Nullable(d.KeteranganPenerjemah)
		m.BahasaIsyarat = support.Nullable(d.BahasaIsyarat)
		m.CaraBelajar = support.Nullable(d.CaraBelajar)
		m.HambatanBelajar = support.Nullable(d.HambatanBelajar)
		m.KeteranganHambatanBelajar = support.Nullable(d.KeteranganHambatanBelajar)
		m.KemampuanBelajar = support.Nullable(d.KemampuanBelajar)
		m.KeteranganKemampuanBelajar = support.Nullable(d.KeteranganKemampuanBelajar)
		m.PenyakitnyaMerupakan = strings.TrimSpace(d.PenyakitnyaMerupakan)
		m.KeteranganPenyakitnyaMerupakan = strings.TrimSpace(d.KeteranganPenyakitnyaMerupakan)
		m.KeputusanMemilihLayanan = strings.TrimSpace(d.KeputusanMemilihLayanan)
		m.KeteranganKeputusanMemilihLayanan = strings.TrimSpace(d.KeteranganKeputusanMemilihLayanan)
		m.KeyakinanTerhadapTerapi = strings.TrimSpace(d.KeyakinanTerhadapTerapi)
		m.KeteranganKeyakinanTerhadapTerapi = strings.TrimSpace(d.KeteranganKeyakinanTerhadapTerapi)
		m.AspekKeyakinanDipertimbangkan = strings.TrimSpace(d.AspekKeyakinanDipertimbangkan)
		m.KeteranganAspekKeyakinanDipertimbangkan = strings.TrimSpace(d.KeteranganAspekKeyakinanDipertimbangkan)
		m.KesediaanMenerimaInformasi = strings.TrimSpace(d.KesediaanMenerimaInformasi)
		m.TopikEdukasiPenyakit = strings.TrimSpace(d.TopikEdukasiPenyakit)
		m.TopikEdukasiRencanaTindakan = strings.TrimSpace(d.TopikEdukasiRencanaTindakan)
		m.TopikEdukasiPengobatan = strings.TrimSpace(d.TopikEdukasiPengobatan)
		m.TopikEdukasiHasilLayanan = strings.TrimSpace(d.TopikEdukasiHasilLayanan)
		return nil
	},
	Refs: []Ref[model.EdukasiPasienKeluargaRj]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.EdukasiPasienKeluargaRj) any { return Str(m.Nip) }},
	},
}
