package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningRisikoKankerParu skrining risiko kanker paru (RMSkriningRisikoKankerParu).
var FormSkriningRisikoKankerParu = &Form[model.SkriningRisikoKankerParu, request.SkriningRisikoKankerParuData]{
	Slug:  "skrining-risiko-kanker-paru",
	Label: "skrining risiko kanker paru",
	Spec: repo.Spec{
		Table:  "skrining_risiko_kanker_paru",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningRisikoKankerParu, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningRisikoKankerParu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningRisikoKankerParu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningRisikoKankerParu) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningRisikoKankerParu, d request.SkriningRisikoKankerParuData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("pernah_kanker", d.PernahKanker, "Ya, Pernah > 5 Tahun Yang Lalu", "Ya, Pernah < 5 Tahun Yang Lalu", "Tidak Pernah"); err != nil {
			return err
		}
		if err := Enum("ada_keluarga_kanker", d.AdaKeluargaKanker, "Ya, Kanker Paru", "Ya, Kanker Jenis Lain", "Tidak Ada"); err != nil {
			return err
		}
		if err := Enum("riwayat_rokok", d.RiwayatRokok, "Perokok Aktif, Masih Merokok 1 Tahun Ini", "Bekas Perokok, Berhenti < 15 Tahun", "Perokok Pasif", "Tidak Merokok"); err != nil {
			return err
		}
		if err := Enum("pernah_paru_kronik", d.PernahParuKronik, "Ya, Pernah Tuberkolosis (TBC)", "Ya, Pernah Penyakit Paru Kronik(PPOK)", "Tidak"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.JenisKelamin = support.Nullable(d.JenisKelamin)
		m.NilaiJenisKelamin = support.Nullable(d.NilaiJenisKelamin)
		m.Umur = support.Nullable(d.Umur)
		m.NilaiUmur = support.Nullable(d.NilaiUmur)
		m.PernahKanker = support.Nullable(d.PernahKanker)
		m.NilaiPernahKanker = support.Nullable(d.NilaiPernahKanker)
		m.AdaKeluargaKanker = support.Nullable(d.AdaKeluargaKanker)
		m.NilaiAdaKeluargaKanker = support.Nullable(d.NilaiAdaKeluargaKanker)
		m.RiwayatRokok = support.Nullable(d.RiwayatRokok)
		m.NilaiRiwayatRokok = support.Nullable(d.NilaiRiwayatRokok)
		m.RiwayatBekerjaMengandungKarsinogen = support.Nullable(d.RiwayatBekerjaMengandungKarsinogen)
		m.NilaiRiwayatBekerjaMengandungKarsinogen = support.Nullable(d.NilaiRiwayatBekerjaMengandungKarsinogen)
		m.LingkunganTinggalPolusiTinggi = support.Nullable(d.LingkunganTinggalPolusiTinggi)
		m.NilaiLingkunganTinggalPolusiTinggi = support.Nullable(d.NilaiLingkunganTinggalPolusiTinggi)
		m.LingkunganRumahTidakSehat = support.Nullable(d.LingkunganRumahTidakSehat)
		m.NilaiLingkunganRumahTidakSehat = support.Nullable(d.NilaiLingkunganRumahTidakSehat)
		m.PernahParuKronik = support.Nullable(d.PernahParuKronik)
		m.NilaiPernahParuKronik = support.Nullable(d.NilaiPernahParuKronik)
		m.TotalSkor = support.Nullable(d.TotalSkor)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningRisikoKankerParu]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningRisikoKankerParu) any { return Str(m.Nip) }},
	},
}
