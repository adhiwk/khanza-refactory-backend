package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPasienImunitasRendah penilaian pasien imunitas rendah (RMPenilaianPasienImunitasRendah).
var FormPenilaianPasienImunitasRendah = &Form[model.PenilaianPasienImunitasRendah, request.PenilaianPasienImunitasRendahData]{
	Slug:  "penilaian-pasien-imunitas-rendah",
	Label: "penilaian pasien imunitas rendah",
	Spec: repo.Spec{
		Table:  "penilaian_pasien_imunitas_rendah",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianPasienImunitasRendah, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianPasienImunitasRendah) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianPasienImunitasRendah) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPasienImunitasRendah) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianPasienImunitasRendah, d request.PenilaianPasienImunitasRendahData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Anamnesis = support.Nullable(d.Anamnesis)
		m.Hubungan = strings.TrimSpace(d.Hubungan)
		m.PasienMengetahuiKondisiPenyakitnya = support.Nullable(d.PasienMengetahuiKondisiPenyakitnya)
		m.KebutuhanRuangPerawatan = support.Nullable(d.KebutuhanRuangPerawatan)
		m.RiwayatPenyakitKeluhan = support.Nullable(d.RiwayatPenyakitKeluhan)
		m.RiwayatPenyakitKeluarga = support.Nullable(d.RiwayatPenyakitKeluarga)
		m.RiwayatAlergi = support.Nullable(d.RiwayatAlergi)
		m.RiwayatVaksinasi = support.Nullable(d.RiwayatVaksinasi)
		m.RiwayatPengobatan = support.Nullable(d.RiwayatPengobatan)
		m.DiagnosaUtama = support.Nullable(d.DiagnosaUtama)
		m.DiagnosaTambahan = support.Nullable(d.DiagnosaTambahan)
		return nil
	},
	Refs: []Ref[model.PenilaianPasienImunitasRendah]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianPasienImunitasRendah) any { return Str(m.KdDokter) }},
	},
}
