package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanRalanPsikiatri penilaian awal keperawatan ralan psikiatri (RMPenilaianAwalKeperawatanRalanPsikiatri).
var FormPenilaianAwalKeperawatanRalanPsikiatri = &Form[model.PenilaianAwalKeperawatanRalanPsikiatri, request.PenilaianAwalKeperawatanRalanPsikiatriData]{
	Slug:  "penilaian-awal-keperawatan-ralan-psikiatri",
	Label: "penilaian awal keperawatan ralan psikiatri",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_ralan_psikiatri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_ralan_masalah_psikiatri", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_psikiatri", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ralan_rencana_psikiatri", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_psikiatri", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanRalanPsikiatri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanRalanPsikiatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanRalanPsikiatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanRalanPsikiatri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanRalanPsikiatri, d request.PenilaianAwalKeperawatanRalanPsikiatriData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("rkd_berobat", d.RkdBerobat, "Tidak", "Ya, Alternatif", "Ya, RS", "Ya, Puskesmas"); err != nil {
			return err
		}
		if err := Enum("sg1", d.Sg1, "Tidak", "Tidak Yakin", "Ya, 1-5 Kg", "Ya, 6-10 Kg", "Ya, 11-15 Kg", "Ya, >15 Kg"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.RkdSakitSejak = strings.TrimSpace(d.RkdSakitSejak)
		m.RkdKeluhan = strings.TrimSpace(d.RkdKeluhan)
		m.RkdBerobat = strings.TrimSpace(d.RkdBerobat)
		m.RkdHasilPengobatan = strings.TrimSpace(d.RkdHasilPengobatan)
		m.FpPutusObat = strings.TrimSpace(d.FpPutusObat)
		m.KetPutusObat = strings.TrimSpace(d.KetPutusObat)
		m.FpEkonomi = strings.TrimSpace(d.FpEkonomi)
		m.KetMasalahEkonomi = strings.TrimSpace(d.KetMasalahEkonomi)
		m.FpMasalahFisik = strings.TrimSpace(d.FpMasalahFisik)
		m.KetMasalahFisik = strings.TrimSpace(d.KetMasalahFisik)
		m.FpMasalahPsikososial = strings.TrimSpace(d.FpMasalahPsikososial)
		m.KetMasalahPsikososial = strings.TrimSpace(d.KetMasalahPsikososial)
		m.RhKeluarga = strings.TrimSpace(d.RhKeluarga)
		m.KetRhKeluarga = strings.TrimSpace(d.KetRhKeluarga)
		m.ResikoBunuhDiri = strings.TrimSpace(d.ResikoBunuhDiri)
		m.RbdIde = strings.TrimSpace(d.RbdIde)
		m.KetRbdIde = strings.TrimSpace(d.KetRbdIde)
		m.RbdRencana = strings.TrimSpace(d.RbdRencana)
		m.KetRbdRencana = strings.TrimSpace(d.KetRbdRencana)
		m.RbdAlat = strings.TrimSpace(d.RbdAlat)
		m.KetRbdAlat = strings.TrimSpace(d.KetRbdAlat)
		m.RbdPercobaan = strings.TrimSpace(d.RbdPercobaan)
		m.KetRbdPercobaan = strings.TrimSpace(d.KetRbdPercobaan)
		m.RbdKeinginan = strings.TrimSpace(d.RbdKeinginan)
		m.KetRbdKeinginan = strings.TrimSpace(d.KetRbdKeinginan)
		m.RpoPenggunaan = strings.TrimSpace(d.RpoPenggunaan)
		m.KetRpoPenggunaan = strings.TrimSpace(d.KetRpoPenggunaan)
		m.RpoEfekSamping = strings.TrimSpace(d.RpoEfekSamping)
		m.KetRpoEfekSamping = strings.TrimSpace(d.KetRpoEfekSamping)
		m.RpoNapza = strings.TrimSpace(d.RpoNapza)
		m.KetRpoNapza = strings.TrimSpace(d.KetRpoNapza)
		m.KetLamaPemakaian = strings.TrimSpace(d.KetLamaPemakaian)
		m.KetCaraPemakaian = strings.TrimSpace(d.KetCaraPemakaian)
		m.KetLatarBelakangPemakaian = strings.TrimSpace(d.KetLatarBelakangPemakaian)
		m.RpoPenggunaanObatLainnya = strings.TrimSpace(d.RpoPenggunaanObatLainnya)
		m.KetPenggunaanObatLainnya = strings.TrimSpace(d.KetPenggunaanObatLainnya)
		m.KetAlasanPenggunaan = strings.TrimSpace(d.KetAlasanPenggunaan)
		m.RpoAlergiObat = strings.TrimSpace(d.RpoAlergiObat)
		m.KetAlergiObat = strings.TrimSpace(d.KetAlergiObat)
		m.RpoMerokok = strings.TrimSpace(d.RpoMerokok)
		m.KetMerokok = strings.TrimSpace(d.KetMerokok)
		m.RpoMinumKopi = strings.TrimSpace(d.RpoMinumKopi)
		m.KetMinumKopi = strings.TrimSpace(d.KetMinumKopi)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Gcs = strings.TrimSpace(d.Gcs)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.PfKeluhanFisik = strings.TrimSpace(d.PfKeluhanFisik)
		m.KetKeluhanFisik = strings.TrimSpace(d.KetKeluhanFisik)
		m.SkalaNyeri = strings.TrimSpace(d.SkalaNyeri)
		m.Durasi = strings.TrimSpace(d.Durasi)
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Provokes = strings.TrimSpace(d.Provokes)
		m.KetProvokes = strings.TrimSpace(d.KetProvokes)
		m.Quality = strings.TrimSpace(d.Quality)
		m.KetQuality = strings.TrimSpace(d.KetQuality)
		m.Lokasi = strings.TrimSpace(d.Lokasi)
		m.Menyebar = strings.TrimSpace(d.Menyebar)
		m.PadaDokter = strings.TrimSpace(d.PadaDokter)
		m.KetDokter = strings.TrimSpace(d.KetDokter)
		m.NyeriHilang = strings.TrimSpace(d.NyeriHilang)
		m.KetNyeri = strings.TrimSpace(d.KetNyeri)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Bmi = strings.TrimSpace(d.Bmi)
		m.LaporStatusNutrisi = strings.TrimSpace(d.LaporStatusNutrisi)
		m.KetLaporStatusNutrisi = strings.TrimSpace(d.KetLaporStatusNutrisi)
		m.Sg1 = strings.TrimSpace(d.Sg1)
		m.Nilai1 = strings.TrimSpace(d.Nilai1)
		m.Sg2 = strings.TrimSpace(d.Sg2)
		m.Nilai2 = strings.TrimSpace(d.Nilai2)
		m.TotalHasil = d.TotalHasil
		m.Resikojatuh = strings.TrimSpace(d.Resikojatuh)
		m.Bjm = strings.TrimSpace(d.Bjm)
		m.Msa = strings.TrimSpace(d.Msa)
		m.Hasil = strings.TrimSpace(d.Hasil)
		m.Lapor = strings.TrimSpace(d.Lapor)
		m.KetLapor = strings.TrimSpace(d.KetLapor)
		m.AdlMandi = strings.TrimSpace(d.AdlMandi)
		m.AdlBerpakaian = strings.TrimSpace(d.AdlBerpakaian)
		m.AdlMakan = strings.TrimSpace(d.AdlMakan)
		m.AdlBak = strings.TrimSpace(d.AdlBak)
		m.AdlBab = strings.TrimSpace(d.AdlBab)
		m.AdlHobi = strings.TrimSpace(d.AdlHobi)
		m.KetAdlHobi = strings.TrimSpace(d.KetAdlHobi)
		m.AdlSosialisasi = strings.TrimSpace(d.AdlSosialisasi)
		m.KetAdlSosialisasi = strings.TrimSpace(d.KetAdlSosialisasi)
		m.AdlKegiatan = strings.TrimSpace(d.AdlKegiatan)
		m.KetAdlKegiatan = strings.TrimSpace(d.KetAdlKegiatan)
		m.SkPenampilan = strings.TrimSpace(d.SkPenampilan)
		m.SkAlamPerasaan = strings.TrimSpace(d.SkAlamPerasaan)
		m.SkPembicaraan = strings.TrimSpace(d.SkPembicaraan)
		m.SkAfek = strings.TrimSpace(d.SkAfek)
		m.SkAktifitasMotorik = strings.TrimSpace(d.SkAktifitasMotorik)
		m.SkGangguanRingan = strings.TrimSpace(d.SkGangguanRingan)
		m.SkProsesPikir = strings.TrimSpace(d.SkProsesPikir)
		m.SkOrientasi = strings.TrimSpace(d.SkOrientasi)
		m.SkTingkatKesadaranOrientasi = strings.TrimSpace(d.SkTingkatKesadaranOrientasi)
		m.SkMemori = strings.TrimSpace(d.SkMemori)
		m.SkInteraksi = strings.TrimSpace(d.SkInteraksi)
		m.SkKonsentrasi = strings.TrimSpace(d.SkKonsentrasi)
		m.SkPersepsi = strings.TrimSpace(d.SkPersepsi)
		m.KetSkPersepsi = strings.TrimSpace(d.KetSkPersepsi)
		m.SkIsiPikir = strings.TrimSpace(d.SkIsiPikir)
		m.SkWaham = strings.TrimSpace(d.SkWaham)
		m.KetSkWaham = strings.TrimSpace(d.KetSkWaham)
		m.SkDayaTilikDiri = strings.TrimSpace(d.SkDayaTilikDiri)
		m.KetSkDayaTilikDiri = strings.TrimSpace(d.KetSkDayaTilikDiri)
		m.KkPembelajaran = strings.TrimSpace(d.KkPembelajaran)
		m.KetKkPembelajaran = strings.TrimSpace(d.KetKkPembelajaran)
		m.KetKkPembelajaranLainnya = strings.TrimSpace(d.KetKkPembelajaranLainnya)
		m.KkPenerjamah = strings.TrimSpace(d.KkPenerjamah)
		m.KetKkPenerjamahLainnya = strings.TrimSpace(d.KetKkPenerjamahLainnya)
		m.KkBahasaIsyarat = strings.TrimSpace(d.KkBahasaIsyarat)
		m.KkKebutuhanEdukasi = strings.TrimSpace(d.KkKebutuhanEdukasi)
		m.KetKkKebutuhanEdukasi = strings.TrimSpace(d.KetKkKebutuhanEdukasi)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanRalanPsikiatri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRalanPsikiatri) any { return Str(m.Nip) }},
	},
}
