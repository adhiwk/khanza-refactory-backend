package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanIgd penilaian awal keperawatan IGD (RMPenilaianAwalKeperawatanIGD).
var FormPenilaianAwalKeperawatanIgd = &Form[model.PenilaianAwalKeperawatanIgd, request.PenilaianAwalKeperawatanIgdData]{
	Slug:  "penilaian-awal-keperawatan-igd",
	Label: "penilaian awal keperawatan IGD",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_igd",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_igd_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_igd", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ralan_rencana_igd", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_igd", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanIgd, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanIgd) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanIgd) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanIgd) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianAwalKeperawatanIgd, d request.PenilaianAwalKeperawatanIgdData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.StatusKehamilan = strings.TrimSpace(d.StatusKehamilan)
		m.Gravida = support.Nullable(d.Gravida)
		m.Para = support.Nullable(d.Para)
		m.Abortus = support.Nullable(d.Abortus)
		m.Hpht = support.Nullable(d.Hpht)
		m.Tekanan = strings.TrimSpace(d.Tekanan)
		m.Pupil = strings.TrimSpace(d.Pupil)
		m.Neurosensorik = strings.TrimSpace(d.Neurosensorik)
		m.Integumen = strings.TrimSpace(d.Integumen)
		m.Turgor = strings.TrimSpace(d.Turgor)
		m.Edema = strings.TrimSpace(d.Edema)
		m.Mukosa = strings.TrimSpace(d.Mukosa)
		m.Perdarahan = strings.TrimSpace(d.Perdarahan)
		m.JumlahPerdarahan = support.Nullable(d.JumlahPerdarahan)
		m.WarnaPerdarahan = support.Nullable(d.WarnaPerdarahan)
		m.Intoksikasi = strings.TrimSpace(d.Intoksikasi)
		m.Bab = support.Nullable(d.Bab)
		m.Xbab = support.Nullable(d.Xbab)
		m.Kbab = support.Nullable(d.Kbab)
		m.Wbab = support.Nullable(d.Wbab)
		m.Bak = support.Nullable(d.Bak)
		m.Xbak = support.Nullable(d.Xbak)
		m.Wbak = support.Nullable(d.Wbak)
		m.Lbak = support.Nullable(d.Lbak)
		m.Psikologis = strings.TrimSpace(d.Psikologis)
		m.Jiwa = strings.TrimSpace(d.Jiwa)
		m.Perilaku = strings.TrimSpace(d.Perilaku)
		m.Dilaporkan = support.Nullable(d.Dilaporkan)
		m.Sebutkan = support.Nullable(d.Sebutkan)
		m.Hubungan = strings.TrimSpace(d.Hubungan)
		m.TinggalDengan = strings.TrimSpace(d.TinggalDengan)
		m.KetTinggal = support.Nullable(d.KetTinggal)
		m.Budaya = strings.TrimSpace(d.Budaya)
		m.KetBudaya = strings.TrimSpace(d.KetBudaya)
		m.PendidikanPj = strings.TrimSpace(d.PendidikanPj)
		m.KetPendidikanPj = support.Nullable(d.KetPendidikanPj)
		m.Edukasi = strings.TrimSpace(d.Edukasi)
		m.KetEdukasi = strings.TrimSpace(d.KetEdukasi)
		m.Kemampuan = strings.TrimSpace(d.Kemampuan)
		m.Aktifitas = strings.TrimSpace(d.Aktifitas)
		m.AlatBantu = strings.TrimSpace(d.AlatBantu)
		m.KetBantu = support.Nullable(d.KetBantu)
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
		m.KetNyeri = support.Nullable(d.KetNyeri)
		m.PadaDokter = strings.TrimSpace(d.PadaDokter)
		m.KetDokter = support.Nullable(d.KetDokter)
		m.BerjalanA = strings.TrimSpace(d.BerjalanA)
		m.BerjalanB = strings.TrimSpace(d.BerjalanB)
		m.BerjalanC = strings.TrimSpace(d.BerjalanC)
		m.Hasil = strings.TrimSpace(d.Hasil)
		m.Lapor = strings.TrimSpace(d.Lapor)
		m.KetLapor = support.Nullable(d.KetLapor)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanIgd]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanIgd) any { return Str(m.Nip) }},
	},
}
