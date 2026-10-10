package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanKebidanan penilaian awal keperawatan kebidanan (RMPenilaianAwalKeperawatanKebidanan).
var FormPenilaianAwalKeperawatanKebidanan = &Form[model.PenilaianAwalKeperawatanKebidanan, request.PenilaianAwalKeperawatanKebidananData]{
	Slug:  "penilaian-awal-keperawatan-kebidanan",
	Label: "penilaian awal keperawatan kebidanan",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_kebidanan",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanKebidanan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanKebidanan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanKebidanan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanKebidanan) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanKebidanan, d request.PenilaianAwalKeperawatanKebidananData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		vHpht, err := support.ParseDate(d.Hpht)
		if err != nil {
			return err
		}
		vTp, err := support.ParseDate(d.Tp)
		if err != nil {
			return err
		}
		if err := Enum("sg1", d.Sg1, "Ya", "Tidak", "Tidak Yakin", "Ya, 1-5 Kg", "Ya, 6-10 Kg", "Ya, 11-15 Kg", "Ya, >15 Kg"); err != nil {
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
		m.Lila = strings.TrimSpace(d.Lila)
		m.Bmi = strings.TrimSpace(d.Bmi)
		m.Tfu = strings.TrimSpace(d.Tfu)
		m.Tbj = strings.TrimSpace(d.Tbj)
		m.Letak = strings.TrimSpace(d.Letak)
		m.Presentasi = strings.TrimSpace(d.Presentasi)
		m.Penurunan = strings.TrimSpace(d.Penurunan)
		m.His = strings.TrimSpace(d.His)
		m.Kekuatan = strings.TrimSpace(d.Kekuatan)
		m.Lamanya = strings.TrimSpace(d.Lamanya)
		m.Bjj = strings.TrimSpace(d.Bjj)
		m.KetBjj = strings.TrimSpace(d.KetBjj)
		m.Portio = strings.TrimSpace(d.Portio)
		m.Serviks = strings.TrimSpace(d.Serviks)
		m.Ketuban = strings.TrimSpace(d.Ketuban)
		m.Hodge = strings.TrimSpace(d.Hodge)
		m.Inspekulo = strings.TrimSpace(d.Inspekulo)
		m.KetInspekulo = strings.TrimSpace(d.KetInspekulo)
		m.Ctg = strings.TrimSpace(d.Ctg)
		m.KetCtg = strings.TrimSpace(d.KetCtg)
		m.Usg = strings.TrimSpace(d.Usg)
		m.KetUsg = strings.TrimSpace(d.KetUsg)
		m.Lab = strings.TrimSpace(d.Lab)
		m.KetLab = strings.TrimSpace(d.KetLab)
		m.Lakmus = strings.TrimSpace(d.Lakmus)
		m.KetLakmus = strings.TrimSpace(d.KetLakmus)
		m.Panggul = strings.TrimSpace(d.Panggul)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Umur = strings.TrimSpace(d.Umur)
		m.Lama = strings.TrimSpace(d.Lama)
		m.Banyaknya = strings.TrimSpace(d.Banyaknya)
		m.Haid = strings.TrimSpace(d.Haid)
		m.Siklus = strings.TrimSpace(d.Siklus)
		m.KetSiklus = strings.TrimSpace(d.KetSiklus)
		m.KetSiklus1 = strings.TrimSpace(d.KetSiklus1)
		m.Status = strings.TrimSpace(d.Status)
		m.Kali = strings.TrimSpace(d.Kali)
		m.Usia1 = strings.TrimSpace(d.Usia1)
		m.Ket1 = strings.TrimSpace(d.Ket1)
		m.Usia2 = support.Nullable(d.Usia2)
		m.Ket2 = support.Nullable(d.Ket2)
		m.Usia3 = support.Nullable(d.Usia3)
		m.Ket3 = support.Nullable(d.Ket3)
		m.Hpht = vHpht
		m.UsiaKehamilan = strings.TrimSpace(d.UsiaKehamilan)
		m.Tp = vTp
		m.Imunisasi = strings.TrimSpace(d.Imunisasi)
		m.KetImunisasi = strings.TrimSpace(d.KetImunisasi)
		m.G = strings.TrimSpace(d.G)
		m.P = strings.TrimSpace(d.P)
		m.A = strings.TrimSpace(d.A)
		m.Hidup = strings.TrimSpace(d.Hidup)
		m.Ginekologi = strings.TrimSpace(d.Ginekologi)
		m.Kebiasaan = strings.TrimSpace(d.Kebiasaan)
		m.KetKebiasaan = strings.TrimSpace(d.KetKebiasaan)
		m.Kebiasaan1 = strings.TrimSpace(d.Kebiasaan1)
		m.KetKebiasaan1 = strings.TrimSpace(d.KetKebiasaan1)
		m.Kebiasaan2 = strings.TrimSpace(d.Kebiasaan2)
		m.KetKebiasaan2 = strings.TrimSpace(d.KetKebiasaan2)
		m.Kebiasaan3 = strings.TrimSpace(d.Kebiasaan3)
		m.Kb = strings.TrimSpace(d.Kb)
		m.KetKb = strings.TrimSpace(d.KetKb)
		m.Komplikasi = strings.TrimSpace(d.Komplikasi)
		m.KetKomplikasi = strings.TrimSpace(d.KetKomplikasi)
		m.Berhenti = strings.TrimSpace(d.Berhenti)
		m.Alasan = strings.TrimSpace(d.Alasan)
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
		m.TotalHasil = strings.TrimSpace(d.TotalHasil)
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
		m.Masalah = strings.TrimSpace(d.Masalah)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanKebidanan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanKebidanan) any { return Str(m.Nip) }},
	},
}
