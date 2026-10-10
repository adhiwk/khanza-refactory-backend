package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPerencanaanPemulangan perencanaan pemulangan (RMPerencanaanPemulangan).
var FormPerencanaanPemulangan = &Form[model.PerencanaanPemulangan, request.PerencanaanPemulanganData]{
	Slug:  "perencanaan-pemulangan",
	Label: "perencanaan pemulangan",
	Spec: repo.Spec{
		Table:  "perencanaan_pemulangan",
		Keys:   []string{"no_rawat"},
		Waktu:  "",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PerencanaanPemulangan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PerencanaanPemulangan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PerencanaanPemulangan) time.Time { return time.Time{} },
	Petugas: func(m *model.PerencanaanPemulangan) []string { return []string{m.Nip} },
	Fill: func(m *model.PerencanaanPemulangan, d request.PerencanaanPemulanganData) error {
		vRencanaPulang, err := support.ParseDate(d.RencanaPulang)
		if err != nil {
			return err
		}
		m.RencanaPulang = vRencanaPulang
		m.AlasanMasuk = support.Nullable(d.AlasanMasuk)
		m.DiagnosaMedis = support.Nullable(d.DiagnosaMedis)
		m.PengaruhRiPasienDanKeluarga = support.Nullable(d.PengaruhRiPasienDanKeluarga)
		m.KeteranganPengaruhRiPasienDanKeluarga = support.Nullable(d.KeteranganPengaruhRiPasienDanKeluarga)
		m.PengaruhRiPekerjaanSekolah = strings.TrimSpace(d.PengaruhRiPekerjaanSekolah)
		m.KeteranganPengaruhRiPekerjaanSekolah = strings.TrimSpace(d.KeteranganPengaruhRiPekerjaanSekolah)
		m.PengaruhRiKeuangan = strings.TrimSpace(d.PengaruhRiKeuangan)
		m.KeteranganPengaruhRiKeuangan = strings.TrimSpace(d.KeteranganPengaruhRiKeuangan)
		m.AntisipasiMasalahSaatPulang = strings.TrimSpace(d.AntisipasiMasalahSaatPulang)
		m.KeteranganAntisipasiMasalahSaatPulang = strings.TrimSpace(d.KeteranganAntisipasiMasalahSaatPulang)
		m.BantuanDiperlukanDalam = strings.TrimSpace(d.BantuanDiperlukanDalam)
		m.KeteranganBantuanDiperlukanDalam = strings.TrimSpace(d.KeteranganBantuanDiperlukanDalam)
		m.AdakahYangMembantuKeperluan = strings.TrimSpace(d.AdakahYangMembantuKeperluan)
		m.KeteranganAdakahYangMembantuKeperluan = strings.TrimSpace(d.KeteranganAdakahYangMembantuKeperluan)
		m.PasienTinggalSendiri = strings.TrimSpace(d.PasienTinggalSendiri)
		m.KeteranganPasienTinggalSendiri = strings.TrimSpace(d.KeteranganPasienTinggalSendiri)
		m.PasienMenggunakanPeralatanMedis = strings.TrimSpace(d.PasienMenggunakanPeralatanMedis)
		m.KeteranganPasienMenggunakanPeralatanMedis = strings.TrimSpace(d.KeteranganPasienMenggunakanPeralatanMedis)
		m.PasienMemerlukanAlatBantu = strings.TrimSpace(d.PasienMemerlukanAlatBantu)
		m.KeteranganPasienMemerlukanAlatBantu = strings.TrimSpace(d.KeteranganPasienMemerlukanAlatBantu)
		m.MemerlukanPerawatanKhusus = strings.TrimSpace(d.MemerlukanPerawatanKhusus)
		m.KeteranganMemerlukanPerawatanKhusus = strings.TrimSpace(d.KeteranganMemerlukanPerawatanKhusus)
		m.BermasalahMemenuhiKebutuhan = strings.TrimSpace(d.BermasalahMemenuhiKebutuhan)
		m.KeteranganBermasalahMemenuhiKebutuhan = strings.TrimSpace(d.KeteranganBermasalahMemenuhiKebutuhan)
		m.MemilikiNyeriKronis = strings.TrimSpace(d.MemilikiNyeriKronis)
		m.KeteranganMemilikiNyeriKronis = strings.TrimSpace(d.KeteranganMemilikiNyeriKronis)
		m.MemerlukanEdukasiKesehatan = strings.TrimSpace(d.MemerlukanEdukasiKesehatan)
		m.KeteranganMemerlukanEdukasiKesehatan = strings.TrimSpace(d.KeteranganMemerlukanEdukasiKesehatan)
		m.MemerlukanKeterampilkanKhusus = strings.TrimSpace(d.MemerlukanKeterampilkanKhusus)
		m.KeteranganMemerlukanKeterampilkanKhusus = strings.TrimSpace(d.KeteranganMemerlukanKeterampilkanKhusus)
		m.NamaPasienKeluarga = strings.TrimSpace(d.NamaPasienKeluarga)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PerencanaanPemulangan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PerencanaanPemulangan) any { return Str(m.Nip) }},
	},
}
