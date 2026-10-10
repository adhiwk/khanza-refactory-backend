package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanMata penilaian awal keperawatan mata (RMPenilaianAwalKeperawatanMata).
var FormPenilaianAwalKeperawatanMata = &Form[model.PenilaianAwalKeperawatanMata, request.PenilaianAwalKeperawatanMataData]{
	Slug:  "penilaian-awal-keperawatan-mata",
	Label: "penilaian awal keperawatan mata",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_mata",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_mata_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan", RefColumn: "kode_masalah"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanMata, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanMata) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanMata) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanMata) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanMata, d request.PenilaianAwalKeperawatanMataData) error {
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
		m.Rps = strings.TrimSpace(d.Rps)
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
		m.Sg3 = strings.TrimSpace(d.Sg3)
		m.Nilai3 = strings.TrimSpace(d.Nilai3)
		m.Sg4 = strings.TrimSpace(d.Sg4)
		m.Nilai4 = strings.TrimSpace(d.Nilai4)
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
		m.Visuskanan = strings.TrimSpace(d.Visuskanan)
		m.Visuskiri = strings.TrimSpace(d.Visuskiri)
		m.Refraksikanan = strings.TrimSpace(d.Refraksikanan)
		m.Refraksikiri = strings.TrimSpace(d.Refraksikiri)
		m.Tiokanan = strings.TrimSpace(d.Tiokanan)
		m.Tiokiri = strings.TrimSpace(d.Tiokiri)
		m.Palberakanan = strings.TrimSpace(d.Palberakanan)
		m.Palberakiri = strings.TrimSpace(d.Palberakiri)
		m.Konjungtivakanan = strings.TrimSpace(d.Konjungtivakanan)
		m.Konjungtivakiri = strings.TrimSpace(d.Konjungtivakiri)
		m.Sklerakanan = strings.TrimSpace(d.Sklerakanan)
		m.Sklerakiri = strings.TrimSpace(d.Sklerakiri)
		m.Korneakanan = strings.TrimSpace(d.Korneakanan)
		m.Korneakiri = strings.TrimSpace(d.Korneakiri)
		m.Bmdkanan = strings.TrimSpace(d.Bmdkanan)
		m.Bmdkiri = strings.TrimSpace(d.Bmdkiri)
		m.Iriskanan = strings.TrimSpace(d.Iriskanan)
		m.Iriskiri = strings.TrimSpace(d.Iriskiri)
		m.Pupilkanan = strings.TrimSpace(d.Pupilkanan)
		m.Pupilkiri = strings.TrimSpace(d.Pupilkiri)
		m.Lensakanan = strings.TrimSpace(d.Lensakanan)
		m.Lensakiri = strings.TrimSpace(d.Lensakiri)
		m.Oftalmoskopikanan = strings.TrimSpace(d.Oftalmoskopikanan)
		m.Oftalmoskopikiri = strings.TrimSpace(d.Oftalmoskopikiri)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanMata]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanMata) any { return Str(m.Nip) }},
	},
}
