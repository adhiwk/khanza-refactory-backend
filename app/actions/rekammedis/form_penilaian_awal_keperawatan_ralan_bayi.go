package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanRalanBayi penilaian awal keperawatan bayi anak (RMPenilaianAwalKeperawatanBayiAnak).
var FormPenilaianAwalKeperawatanRalanBayi = &Form[model.PenilaianAwalKeperawatanRalanBayi, request.PenilaianAwalKeperawatanRalanBayiData]{
	Slug:  "penilaian-awal-keperawatan-ralan-bayi",
	Label: "penilaian awal keperawatan bayi anak",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_ralan_bayi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_ralan_bayi_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_anak", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ralan_rencana_anak", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_anak", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanRalanBayi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanRalanBayi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanRalanBayi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanRalanBayi) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanRalanBayi, d request.PenilaianAwalKeperawatanRalanBayiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("aktifitas", d.Aktifitas, "Tidur posisi normal, mudah bergerak", "Gerakan menggeliat/berguling, kaku", "Melengkungkan punggung/kaku menghentak"); err != nil {
			return err
		}
		if err := Enum("menangis", d.Menangis, "Tidak menangis (mudah bergerak)", "Mengerang/merengek", "Menangis terus menerus, terisak, menjerit"); err != nil {
			return err
		}
		if err := Enum("bersuara", d.Bersuara, "Bersuara normal/tenang", "Tenang bila dipeluk, digendong/diajak bicara", "Sulit untuk menenangkan"); err != nil {
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
		m.Lp = strings.TrimSpace(d.Lp)
		m.Lk = strings.TrimSpace(d.Lk)
		m.Ld = strings.TrimSpace(d.Ld)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Anakke = strings.TrimSpace(d.Anakke)
		m.Darisaudara = strings.TrimSpace(d.Darisaudara)
		m.Caralahir = strings.TrimSpace(d.Caralahir)
		m.KetCaralahir = strings.TrimSpace(d.KetCaralahir)
		m.Umurkelahiran = strings.TrimSpace(d.Umurkelahiran)
		m.Kelainanbawaan = strings.TrimSpace(d.Kelainanbawaan)
		m.KetKelainanBawaan = strings.TrimSpace(d.KetKelainanBawaan)
		m.Usiatengkurap = strings.TrimSpace(d.Usiatengkurap)
		m.Usiaduduk = strings.TrimSpace(d.Usiaduduk)
		m.Usiaberdiri = strings.TrimSpace(d.Usiaberdiri)
		m.Usiagigipertama = strings.TrimSpace(d.Usiagigipertama)
		m.Usiaberjalan = strings.TrimSpace(d.Usiaberjalan)
		m.Usiabicara = strings.TrimSpace(d.Usiabicara)
		m.Usiamembaca = strings.TrimSpace(d.Usiamembaca)
		m.Usiamenulis = strings.TrimSpace(d.Usiamenulis)
		m.Gangguanemosi = strings.TrimSpace(d.Gangguanemosi)
		m.AlatBantu = strings.TrimSpace(d.AlatBantu)
		m.KetBantu = strings.TrimSpace(d.KetBantu)
		m.Prothesa = strings.TrimSpace(d.Prothesa)
		m.KetPro = strings.TrimSpace(d.KetPro)
		m.Adl = strings.TrimSpace(d.Adl)
		m.StatusPsiko = strings.TrimSpace(d.StatusPsiko)
		m.KetPsiko = strings.TrimSpace(d.KetPsiko)
		m.HubKeluarga = strings.TrimSpace(d.HubKeluarga)
		m.Pengasuh = strings.TrimSpace(d.Pengasuh)
		m.KetPengasuh = strings.TrimSpace(d.KetPengasuh)
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
		m.Wajah = strings.TrimSpace(d.Wajah)
		m.Nilaiwajah = strings.TrimSpace(d.Nilaiwajah)
		m.Kaki = strings.TrimSpace(d.Kaki)
		m.Nilaikaki = strings.TrimSpace(d.Nilaikaki)
		m.Aktifitas = strings.TrimSpace(d.Aktifitas)
		m.Nilaiaktifitas = strings.TrimSpace(d.Nilaiaktifitas)
		m.Menangis = strings.TrimSpace(d.Menangis)
		m.Nilaimenangis = strings.TrimSpace(d.Nilaimenangis)
		m.Bersuara = strings.TrimSpace(d.Bersuara)
		m.Nilaibersuara = strings.TrimSpace(d.Nilaibersuara)
		m.Hasilnyeri = d.Hasilnyeri
		m.Nyeri = strings.TrimSpace(d.Nyeri)
		m.Lokasi = strings.TrimSpace(d.Lokasi)
		m.Durasi = strings.TrimSpace(d.Durasi)
		m.Frekuensi = strings.TrimSpace(d.Frekuensi)
		m.NyeriHilang = strings.TrimSpace(d.NyeriHilang)
		m.KetNyeri = strings.TrimSpace(d.KetNyeri)
		m.PadaDokter = strings.TrimSpace(d.PadaDokter)
		m.KetDokter = strings.TrimSpace(d.KetDokter)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanRalanBayi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRalanBayi) any { return Str(m.Nip) }},
	},
}
