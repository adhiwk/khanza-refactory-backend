package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanGigi penilaian awal keperawatan gigi (RMPenilaianAwalKeperawatanGigi).
var FormPenilaianAwalKeperawatanGigi = &Form[model.PenilaianAwalKeperawatanGigi, request.PenilaianAwalKeperawatanGigiData]{
	Slug:  "penilaian-awal-keperawatan-gigi",
	Label: "penilaian awal keperawatan gigi",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_gigi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_gigi_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_gigi", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ralan_rencana_gigi", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_gigi", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanGigi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanGigi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanGigi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanGigi) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanGigi, d request.PenilaianAwalKeperawatanGigiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("riwayat_perawatan_gigi", d.RiwayatPerawatanGigi, "Tidak", "Ya, Kapan"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Bmi = strings.TrimSpace(d.Bmi)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.RiwayatPenyakit = support.Nullable(d.RiwayatPenyakit)
		m.KetRiwayatPenyakit = strings.TrimSpace(d.KetRiwayatPenyakit)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.RiwayatPerawatanGigi = strings.TrimSpace(d.RiwayatPerawatanGigi)
		m.KetRiwayatPerawatanGigi = strings.TrimSpace(d.KetRiwayatPerawatanGigi)
		m.KebiasaanSikatGigi = strings.TrimSpace(d.KebiasaanSikatGigi)
		m.KebiasaanLain = support.Nullable(d.KebiasaanLain)
		m.KetKebiasaanLain = strings.TrimSpace(d.KetKebiasaanLain)
		m.ObatYangDiminumSaatini = support.Nullable(d.ObatYangDiminumSaatini)
		m.AlatBantu = strings.TrimSpace(d.AlatBantu)
		m.KetAlatBantu = strings.TrimSpace(d.KetAlatBantu)
		m.Prothesa = strings.TrimSpace(d.Prothesa)
		m.KetPro = strings.TrimSpace(d.KetPro)
		m.StatusPsiko = strings.TrimSpace(d.StatusPsiko)
		m.KetPsiko = strings.TrimSpace(d.KetPsiko)
		m.HubKeluarga = strings.TrimSpace(d.HubKeluarga)
		m.TinggalDengan = strings.TrimSpace(d.TinggalDengan)
		m.KetTinggal = strings.TrimSpace(d.KetTinggal)
		m.Ekonomi = strings.TrimSpace(d.Ekonomi)
		m.Budaya = strings.TrimSpace(d.Budaya)
		m.KetBudaya = strings.TrimSpace(d.KetBudaya)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		m.KetEdukasi = strings.TrimSpace(d.KetEdukasi)
		m.BerjalanA = strings.TrimSpace(d.BerjalanA)
		m.BerjalanB = strings.TrimSpace(d.BerjalanB)
		m.BerjalanC = strings.TrimSpace(d.BerjalanC)
		m.Hasil = strings.TrimSpace(d.Hasil)
		m.Lapor = strings.TrimSpace(d.Lapor)
		m.KetLapor = strings.TrimSpace(d.KetLapor)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Lokasi = strings.TrimSpace(d.Lokasi)
		m.SkalaNyeri = strings.TrimSpace(d.SkalaNyeri)
		m.Durasi = strings.TrimSpace(d.Durasi)
		m.Frekuensi = strings.TrimSpace(d.Frekuensi)
		m.NyeriHilang = strings.TrimSpace(d.NyeriHilang)
		m.KetNyeri = strings.TrimSpace(d.KetNyeri)
		m.PadaDokter = strings.TrimSpace(d.PadaDokter)
		m.KetDokter = strings.TrimSpace(d.KetDokter)
		m.KebersihanMulut = strings.TrimSpace(d.KebersihanMulut)
		m.MukosaMulut = strings.TrimSpace(d.MukosaMulut)
		m.Karies = strings.TrimSpace(d.Karies)
		m.KarangGigi = strings.TrimSpace(d.KarangGigi)
		m.Gingiva = strings.TrimSpace(d.Gingiva)
		m.Palatum = strings.TrimSpace(d.Palatum)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanGigi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanGigi) any { return Str(m.Nip) }},
	},
}
