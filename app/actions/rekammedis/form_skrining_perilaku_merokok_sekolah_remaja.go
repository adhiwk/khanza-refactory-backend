package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningPerilakuMerokokSekolahRemaja skrining merokok usia sekolah remaja (RMSkriningMerokokUsiaSekolahRemaja).
var FormSkriningPerilakuMerokokSekolahRemaja = &Form[model.SkriningPerilakuMerokokSekolahRemaja, request.SkriningPerilakuMerokokSekolahRemajaData]{
	Slug:  "skrining-perilaku-merokok-sekolah-remaja",
	Label: "skrining merokok usia sekolah remaja",
	Spec: repo.Spec{
		Table:  "skrining_perilaku_merokok_sekolah_remaja",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningPerilakuMerokokSekolahRemaja, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningPerilakuMerokokSekolahRemaja) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningPerilakuMerokokSekolahRemaja) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningPerilakuMerokokSekolahRemaja) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningPerilakuMerokokSekolahRemaja, d request.SkriningPerilakuMerokokSekolahRemajaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("apakah_anda_merokok", d.ApakahAndaMerokok, "Ya, Setiap Hari", "Ya, Kadang-kadang", "Pernah Mencoba Walau Hanya 1 Hisapan", "Tidak Merokok/Tidak Pernah Mencoba"); err != nil {
			return err
		}
		if err := Enum("jumlah_batang_rokok_hariminggu", d.JumlahBatangRokokHariminggu, "Batang/Hari", "Batang/Minggu", ""); err != nil {
			return err
		}
		if err := Enum("jenis_rokok_yang_digunakan", d.JenisRokokYangDigunakan, "Rokok Konvensional : Rokok Filter/Putih, Kretek, Tingwe, dll", "Rokok Elektronik : Vape, IQOS, dll", "Keduanya", "Lainnya", ""); err != nil {
			return err
		}
		if err := Enum("alasan_mulai_merokok", d.AlasanMulaiMerokok, "Ikut-ikutan Teman", "Pengaruh Keluarga", "Rasa Ingin Tahu", "Terpaksa Oleh Teman/Lingkungan", "Mengisi Waktu Luang", "Menghilangkan Stress", "Lainnya", ""); err != nil {
			return err
		}
		if err := Enum("bagaimana_biasanya_mendapatkan_rokok", d.BagaimanaBiasanyaMendapatkanRokok, "Beli Batangan", "Beli Bungkusan", "Dapat Dari Teman/Saudara/Keluarga", "Lainnya", ""); err != nil {
			return err
		}
		if err := Enum("keinginan_berhenti_merokok", d.KeinginanBerhentiMerokok, "Ya", "Tidak", ""); err != nil {
			return err
		}
		if err := Enum("alasan_utama_berhenti_merokok", d.AlasanUtamaBerhentiMerokok, "Kondisi Kesehatan", "Motivasi Diri Sendiri", "Disarankan Orangtua/Guru/Teman", "Tidak Mampu Beli/Mahal", "Lainnya", ""); err != nil {
			return err
		}
		if err := Enum("dampak_kesehatan_dari_merokok_yang_diketahui", d.DampakKesehatanDariMerokokYangDiketahui, "Kecanduan", "Batuk", "Gigi Kuning & Mulut Berbau", "Sistem Pernapasan Terganggu", ""); err != nil {
			return err
		}
		if err := Enum("orang_yang_paling_sering_merokok_disekolah", d.OrangYangPalingSeringMerokokDisekolah, "Teman/Kakak Kelas", "Guru/Kepala Sekolah", "Satpam/Supir/Penjaga Kantin", "Lainnya", ""); err != nil {
			return err
		}
		if err := Enum("dilakukan_pemeriksaan_kadar_co_pernapasan", d.DilakukanPemeriksaanKadarCoPernapasan, "Ya", "Tidak, Alat & BMHP Tidak Tersedia", "Tidak, BMHP Tidak Tersedia"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdSekolah = support.Nullable(d.KdSekolah)
		m.Kelas = support.Nullable(d.Kelas)
		m.ApakahAndaMerokok = strings.TrimSpace(d.ApakahAndaMerokok)
		m.JumlahBatangRokok = strings.TrimSpace(d.JumlahBatangRokok)
		m.JumlahBatangRokokHariminggu = strings.TrimSpace(d.JumlahBatangRokokHariminggu)
		m.JenisRokokYangDigunakan = strings.TrimSpace(d.JenisRokokYangDigunakan)
		m.JenisRokokYangDigunakanKeterangan = strings.TrimSpace(d.JenisRokokYangDigunakanKeterangan)
		m.UsiaMulaiMerokok = strings.TrimSpace(d.UsiaMulaiMerokok)
		m.AlasanMulaiMerokok = strings.TrimSpace(d.AlasanMulaiMerokok)
		m.AlasanMulaiMerokokKeterangan = strings.TrimSpace(d.AlasanMulaiMerokokKeterangan)
		m.SudahBerapaLamaMerokok = strings.TrimSpace(d.SudahBerapaLamaMerokok)
		m.BagaimanaBiasanyaMendapatkanRokok = strings.TrimSpace(d.BagaimanaBiasanyaMendapatkanRokok)
		m.BagaimanaBiasanyaMendapatkanRokokKeterangan = strings.TrimSpace(d.BagaimanaBiasanyaMendapatkanRokokKeterangan)
		m.KeinginanBerhentiMerokok = strings.TrimSpace(d.KeinginanBerhentiMerokok)
		m.AlasanUtamaBerhentiMerokok = strings.TrimSpace(d.AlasanUtamaBerhentiMerokok)
		m.AlasanUtamaBerhentiMerokokKeterangan = strings.TrimSpace(d.AlasanUtamaBerhentiMerokokKeterangan)
		m.TahuDampakKesehatanMerokok = strings.TrimSpace(d.TahuDampakKesehatanMerokok)
		m.DampakKesehatanDariMerokokYangDiketahui = strings.TrimSpace(d.DampakKesehatanDariMerokokYangDiketahui)
		m.TahuMerokokPintuMasukNarkoba = strings.TrimSpace(d.TahuMerokokPintuMasukNarkoba)
		m.MelihatOrangMerokokDiSekolah = strings.TrimSpace(d.MelihatOrangMerokokDiSekolah)
		m.OrangYangPalingSeringMerokokDisekolah = strings.TrimSpace(d.OrangYangPalingSeringMerokokDisekolah)
		m.OrangYangPalingSeringMerokokDisekolahKeterangan = strings.TrimSpace(d.OrangYangPalingSeringMerokokDisekolahKeterangan)
		m.AdaAnggotaKeluargaDiRumahYangMerokok = strings.TrimSpace(d.AdaAnggotaKeluargaDiRumahYangMerokok)
		m.TemanDekatBanyakyangMerokok = strings.TrimSpace(d.TemanDekatBanyakyangMerokok)
		m.DilakukanPemeriksaanKadarCoPernapasan = strings.TrimSpace(d.DilakukanPemeriksaanKadarCoPernapasan)
		m.HasilPemeriksaanCoPernapasan = strings.TrimSpace(d.HasilPemeriksaanCoPernapasan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningPerilakuMerokokSekolahRemaja]{
		{Column: "kd_sekolah", Table: "master_sekolah", RefCol: "kd_sekolah", Label: "master sekolah", Value: func(m *model.SkriningPerilakuMerokokSekolahRemaja) any { return StrPtr(m.KdSekolah) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningPerilakuMerokokSekolahRemaja) any { return Str(m.Nip) }},
	},
}
