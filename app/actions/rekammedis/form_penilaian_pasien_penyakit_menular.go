package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPasienPenyakitMenular penilaian pasien penyakit menular (RMPenilaianPasienPenyakitMenular).
var FormPenilaianPasienPenyakitMenular = &Form[model.PenilaianPasienPenyakitMenular, request.PenilaianPasienPenyakitMenularData]{
	Slug:  "penilaian-pasien-penyakit-menular",
	Label: "penilaian pasien penyakit menular",
	Spec: repo.Spec{
		Table:  "penilaian_pasien_penyakit_menular",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianPasienPenyakitMenular, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianPasienPenyakitMenular) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianPasienPenyakitMenular) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPasienPenyakitMenular) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianPasienPenyakitMenular, d request.PenilaianPasienPenyakitMenularData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.Hubungan = strings.TrimSpace(d.Hubungan)
		m.PasienMengetahuiKondisiPenyakitnya = support.Nullable(d.PasienMengetahuiKondisiPenyakitnya)
		m.PenyakitSamaSerumah = support.Nullable(d.PenyakitSamaSerumah)
		m.RiwayatKontak = support.Nullable(d.RiwayatKontak)
		m.KeteranganRiwayatKontak = support.Nullable(d.KeteranganRiwayatKontak)
		m.TransmisiPenularanPenyakit = support.Nullable(d.TransmisiPenularanPenyakit)
		m.KeteranganTransmisiPenularanPenyakit = support.Nullable(d.KeteranganTransmisiPenularanPenyakit)
		m.KebutuhanRuangRawat = support.Nullable(d.KebutuhanRuangRawat)
		m.KeluhanYangDirasakanSaatIni = support.Nullable(d.KeluhanYangDirasakanSaatIni)
		m.RiwayatPenyakitKeluarga = support.Nullable(d.RiwayatPenyakitKeluarga)
		m.RiwayatAlergi = support.Nullable(d.RiwayatAlergi)
		m.RiwayatVaksinasi = support.Nullable(d.RiwayatVaksinasi)
		m.RiwayatPengobatan = support.Nullable(d.RiwayatPengobatan)
		m.DiagnosaUtama = support.Nullable(d.DiagnosaUtama)
		m.DiagnosaTambahan = support.Nullable(d.DiagnosaTambahan)
		return nil
	},
	Refs: []Ref[model.PenilaianPasienPenyakitMenular]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianPasienPenyakitMenular) any { return Str(m.KdDokter) }},
	},
}
