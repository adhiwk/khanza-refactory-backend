package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanRalanGeriatri penilaian awal keperawatan ralan geriatri (RMPenilaianAwalKeperawatanRalanGeriatri).
var FormPenilaianAwalKeperawatanRalanGeriatri = &Form[model.PenilaianAwalKeperawatanRalanGeriatri, request.PenilaianAwalKeperawatanRalanGeriatriData]{
	Slug:  "penilaian-awal-keperawatan-ralan-geriatri",
	Label: "penilaian awal keperawatan ralan geriatri",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_ralan_geriatri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_ralan_masalah_geriatri", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_geriatri", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ralan_rencana_geriatri", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_geriatri", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanRalanGeriatri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanRalanGeriatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanRalanGeriatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanRalanGeriatri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanRalanGeriatri, d request.PenilaianAwalKeperawatanRalanGeriatriData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("sg1", d.Sg1, "Tidak", "Tidak Yakin", "Ya, 1-5 Kg", "Ya, 6-10 Kg", "Ya, 11-15 Kg", "Ya, >15 Kg"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Gcs = strings.TrimSpace(d.Gcs)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Bmi = strings.TrimSpace(d.Bmi)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.AlatBantu = strings.TrimSpace(d.AlatBantu)
		m.KetBantu = strings.TrimSpace(d.KetBantu)
		m.Prothesa = strings.TrimSpace(d.Prothesa)
		m.KetPro = strings.TrimSpace(d.KetPro)
		m.Adl = strings.TrimSpace(d.Adl)
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
		m.Sg1 = strings.TrimSpace(d.Sg1)
		m.Nilai1 = strings.TrimSpace(d.Nilai1)
		m.Sg2 = strings.TrimSpace(d.Sg2)
		m.Nilai2 = strings.TrimSpace(d.Nilai2)
		m.TotalHasil = d.TotalHasil
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Provokes = strings.TrimSpace(d.Provokes)
		m.KetProvokes = strings.TrimSpace(d.KetProvokes)
		m.Quality = strings.TrimSpace(d.Quality)
		m.KetQuality = strings.TrimSpace(d.KetQuality)
		m.Lokasi = strings.TrimSpace(d.Lokasi)
		m.Menyebar = strings.TrimSpace(d.Menyebar)
		m.SkalaNyeri = strings.TrimSpace(d.SkalaNyeri)
		m.Durasi = strings.TrimSpace(d.Durasi)
		m.NyeriHilang = strings.TrimSpace(d.NyeriHilang)
		m.KetNyeri = strings.TrimSpace(d.KetNyeri)
		m.PadaDokter = strings.TrimSpace(d.PadaDokter)
		m.KetDokter = strings.TrimSpace(d.KetDokter)
		m.EdukasiKemampuanBacatulis = support.Nullable(d.EdukasiKemampuanBacatulis)
		m.EdukasiKebutuhanPenerjemah = support.Nullable(d.EdukasiKebutuhanPenerjemah)
		m.EdukasiKeteranganKebutuhanPenerjemah = support.Nullable(d.EdukasiKeteranganKebutuhanPenerjemah)
		m.EdukasiHambatan = support.Nullable(d.EdukasiHambatan)
		m.EdukasiHambatanKategori = strings.TrimSpace(d.EdukasiHambatanKategori)
		m.EdukasiKeteranganHambatan = support.Nullable(d.EdukasiKeteranganHambatan)
		m.EdukasiCaraBicara = support.Nullable(d.EdukasiCaraBicara)
		m.EdukasiBahasaIsyarat = support.Nullable(d.EdukasiBahasaIsyarat)
		m.EdukasiMenerimaInformasi = support.Nullable(d.EdukasiMenerimaInformasi)
		m.EdukasiKeteranganMenerimaInformasi = support.Nullable(d.EdukasiKeteranganMenerimaInformasi)
		m.EdukasiMetodeBelajar = support.Nullable(d.EdukasiMetodeBelajar)
		m.FrailyPhenotypeBeratBadan = support.Nullable(d.FrailyPhenotypeBeratBadan)
		m.FrailyPhenotypeBeratBadanNilai = d.FrailyPhenotypeBeratBadanNilai
		m.FrailyPhenotypeAktifitasFisik = support.Nullable(d.FrailyPhenotypeAktifitasFisik)
		m.FrailyPhenotypeAktifitasFisikNilai = d.FrailyPhenotypeAktifitasFisikNilai
		m.FrailyPhenotypeKelelahan = support.Nullable(d.FrailyPhenotypeKelelahan)
		m.FrailyPhenotypeKelelahanNilai = d.FrailyPhenotypeKelelahanNilai
		m.FrailyPhenotypeKekuatan = support.Nullable(d.FrailyPhenotypeKekuatan)
		m.FrailyPhenotypeKekuatanNilai = d.FrailyPhenotypeKekuatanNilai
		m.FrailyPhenotypeWaktuBerjalan = support.Nullable(d.FrailyPhenotypeWaktuBerjalan)
		m.FrailyPhenotypeWaktuBerjalanNilai = d.FrailyPhenotypeWaktuBerjalanNilai
		m.FrailyPhenotypeNilaiTotal = d.FrailyPhenotypeNilaiTotal
		m.FrailyPhenotypeStatus = support.Nullable(d.FrailyPhenotypeStatus)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanRalanGeriatri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRalanGeriatri) any { return Str(m.Nip) }},
	},
}
