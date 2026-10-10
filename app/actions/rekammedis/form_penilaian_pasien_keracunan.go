package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPasienKeracunan penilaian pasien keracunan (RMPenilaianPasienKeracunan).
var FormPenilaianPasienKeracunan = &Form[model.PenilaianPasienKeracunan, request.PenilaianPasienKeracunanData]{
	Slug:  "penilaian-pasien-keracunan",
	Label: "penilaian pasien keracunan",
	Spec: repo.Spec{
		Table:  "penilaian_pasien_keracunan",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianPasienKeracunan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianPasienKeracunan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianPasienKeracunan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPasienKeracunan) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianPasienKeracunan, d request.PenilaianPasienKeracunanData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.Hubungan = strings.TrimSpace(d.Hubungan)
		m.TempatKejadian = support.Nullable(d.TempatKejadian)
		m.KeteranganTempatKejadian = support.Nullable(d.KeteranganTempatKejadian)
		m.Keluhan = support.Nullable(d.Keluhan)
		m.RiwayatPenyakitSekarang = support.Nullable(d.RiwayatPenyakitSekarang)
		m.Hamil = support.Nullable(d.Hamil)
		m.Menyusui = support.Nullable(d.Menyusui)
		m.Penyebab = support.Nullable(d.Penyebab)
		m.NamaBahan = support.Nullable(d.NamaBahan)
		m.JumlahBahan = support.Nullable(d.JumlahBahan)
		m.TipePemaparan = support.Nullable(d.TipePemaparan)
		m.KeteranganTipePemaparan = support.Nullable(d.KeteranganTipePemaparan)
		m.TipeKejadian = support.Nullable(d.TipeKejadian)
		m.BauBahan = support.Nullable(d.BauBahan)
		m.KeteranganBauBahan = support.Nullable(d.KeteranganBauBahan)
		m.Pupil = support.Nullable(d.Pupil)
		m.KeteranganPupil = support.Nullable(d.KeteranganPupil)
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.Td = support.Nullable(d.Td)
		m.Nadi = support.Nullable(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Spo = strings.TrimSpace(d.Spo)
		m.Urine = support.Nullable(d.Urine)
		m.PengobatanSebelumIgd = support.Nullable(d.PengobatanSebelumIgd)
		m.Diagnosis = support.Nullable(d.Diagnosis)
		m.PemeriksaanPenunjang = support.Nullable(d.PemeriksaanPenunjang)
		m.PenatalaksanaanDiberikan = support.Nullable(d.PenatalaksanaanDiberikan)
		m.TindakLanjut = support.Nullable(d.TindakLanjut)
		return nil
	},
	Refs: []Ref[model.PenilaianPasienKeracunan]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianPasienKeracunan) any { return Str(m.KdDokter) }},
	},
}
