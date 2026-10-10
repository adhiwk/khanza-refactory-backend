package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanRanapNeonatus penilaian awal keperawatan ranap neonatus (RMPenilaianAwalKeperawatanRanapNeonatus).
var FormPenilaianAwalKeperawatanRanapNeonatus = &Form[model.PenilaianAwalKeperawatanRanapNeonatus, request.PenilaianAwalKeperawatanRanapNeonatusData]{
	Slug:  "penilaian-awal-keperawatan-ranap-neonatus",
	Label: "penilaian awal keperawatan ranap neonatus",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_ranap_neonatus",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip1", "nip2", "kd_dokter"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_ranap_neonatus_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_neonatus", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ranap_neonatus_rencana", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_neonatus", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanRanapNeonatus, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanRanapNeonatus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu: func(m *model.PenilaianAwalKeperawatanRanapNeonatus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanRanapNeonatus) []string {
		return []string{m.Nip1, m.Nip2, m.KdDokter}
	},
	Fill: func(m *model.PenilaianAwalKeperawatanRanapNeonatus, d request.PenilaianAwalKeperawatanRanapNeonatusData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("penilaian_humptydumpty_skala3", d.PenilaianHumptydumptySkala3, "Kelainan Neurologi", "Perubahan Dalam Oksigen(Masalah Saluran Nafas, Dehidrasi, Anemia, Anoreksia / Sakit Kepala, Dll)", "Kelainan Psikis / Perilaku", "Diagnosa Lain"); err != nil {
			return err
		}
		if err := Enum("penilaian_humptydumpty_skala7", d.PenilaianHumptydumptySkala7, "Bermacam-macam Obat Yang Digunakan : Obat Sedative (Kecuali Pasien ICU Yang Menggunakan sedasi dan paralisis), Hipnotik, Barbiturat, Fenoti-Azin, Antidepresan, Laksans/Diuretika,Narkotik", "Salah Satu Dari Pengobatan Di Atas", "Pengobatan Lain"); err != nil {
			return err
		}
		vPerencanaanPulang, err := support.ParseDate(d.PerencanaanPulang)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.AsalPasien = support.Nullable(d.AsalPasien)
		m.CaraMasuk = support.Nullable(d.CaraMasuk)
		m.DiperolehDari = support.Nullable(d.DiperolehDari)
		m.HubunganDenganPasien = support.Nullable(d.HubunganDenganPasien)
		m.KeluhanUtama = support.Nullable(d.KeluhanUtama)
		m.PrenatalG = support.Nullable(d.PrenatalG)
		m.PrenatalP = support.Nullable(d.PrenatalP)
		m.PrenatalA = support.Nullable(d.PrenatalA)
		m.PrenatalUk = support.Nullable(d.PrenatalUk)
		m.PrenatalRiwayatPenyakitIbu = support.Nullable(d.PrenatalRiwayatPenyakitIbu)
		m.PrenatalRiwayatPenyakitIbuKeterangan = support.Nullable(d.PrenatalRiwayatPenyakitIbuKeterangan)
		m.PrenatalRiwayatPengobatanIbuSelamaHamil = support.Nullable(d.PrenatalRiwayatPengobatanIbuSelamaHamil)
		m.PrenatalPernahDirawat = support.Nullable(d.PrenatalPernahDirawat)
		m.PrenatalPernahDirawatKeterangan = support.Nullable(d.PrenatalPernahDirawatKeterangan)
		m.PrenatalStatusGiziIbu = support.Nullable(d.PrenatalStatusGiziIbu)
		m.IntranatalG = support.Nullable(d.IntranatalG)
		m.IntranatalP = support.Nullable(d.IntranatalP)
		m.IntranatalA = support.Nullable(d.IntranatalA)
		m.IntranatalKondisiLahir = support.Nullable(d.IntranatalKondisiLahir)
		m.IntranatalCaraPersalinan = support.Nullable(d.IntranatalCaraPersalinan)
		m.IntranatalCaraPersalinanKeterangan = support.Nullable(d.IntranatalCaraPersalinanKeterangan)
		m.IntranatalApgar = support.Nullable(d.IntranatalApgar)
		m.IntranatalLetak = support.Nullable(d.IntranatalLetak)
		m.IntranatalTaliPusat = support.Nullable(d.IntranatalTaliPusat)
		m.IntranatalKetuban = strings.TrimSpace(d.IntranatalKetuban)
		m.IntranatalBb = strings.TrimSpace(d.IntranatalBb)
		m.IntranatalPb = strings.TrimSpace(d.IntranatalPb)
		m.IntranatalLk = strings.TrimSpace(d.IntranatalLk)
		m.IntranatalLd = strings.TrimSpace(d.IntranatalLd)
		m.IntranatalLp = strings.TrimSpace(d.IntranatalLp)
		m.RisikoInfeksiMayor = strings.TrimSpace(d.RisikoInfeksiMayor)
		m.RisikoInfeksiMayorKeterangan = strings.TrimSpace(d.RisikoInfeksiMayorKeterangan)
		m.RisikoInfeksiMinor = strings.TrimSpace(d.RisikoInfeksiMinor)
		m.RisikoInfeksiMinorKeterangan = strings.TrimSpace(d.RisikoInfeksiMinorKeterangan)
		m.KebutuhanBiologisNutrisi = strings.TrimSpace(d.KebutuhanBiologisNutrisi)
		m.KebutuhanBiologisNutrisiKeterangan = strings.TrimSpace(d.KebutuhanBiologisNutrisiKeterangan)
		m.KebutuhanBiologisNutrisiFrekuensi = strings.TrimSpace(d.KebutuhanBiologisNutrisiFrekuensi)
		m.KebutuhanBiologisNutrisiKali = strings.TrimSpace(d.KebutuhanBiologisNutrisiKali)
		m.KebutuhanBiologisBak = strings.TrimSpace(d.KebutuhanBiologisBak)
		m.KebutuhanBiologisBakKeterangan = strings.TrimSpace(d.KebutuhanBiologisBakKeterangan)
		m.KebutuhanBiologisBab = strings.TrimSpace(d.KebutuhanBiologisBab)
		m.KebutuhanBiologisBabKeterangan = strings.TrimSpace(d.KebutuhanBiologisBabKeterangan)
		m.AlergiObat = strings.TrimSpace(d.AlergiObat)
		m.AlergiObatKeterangan = strings.TrimSpace(d.AlergiObatKeterangan)
		m.AlergiObatReaksi = strings.TrimSpace(d.AlergiObatReaksi)
		m.AlergiMakanan = strings.TrimSpace(d.AlergiMakanan)
		m.AlergiMakananKeterangan = strings.TrimSpace(d.AlergiMakananKeterangan)
		m.AlergiMakananReaksi = strings.TrimSpace(d.AlergiMakananReaksi)
		m.AlergiLainnya = strings.TrimSpace(d.AlergiLainnya)
		m.AlergiLainnyaKeterangan = strings.TrimSpace(d.AlergiLainnyaKeterangan)
		m.AlergiLainnyaReaksi = strings.TrimSpace(d.AlergiLainnyaReaksi)
		m.RiwayatPenyakitKeluarga = strings.TrimSpace(d.RiwayatPenyakitKeluarga)
		m.RiwayatPenyakitKeluargaKeterangan = strings.TrimSpace(d.RiwayatPenyakitKeluargaKeterangan)
		m.RiwayatImunisasi = strings.TrimSpace(d.RiwayatImunisasi)
		m.RiwayatImunisasiKeterangan = strings.TrimSpace(d.RiwayatImunisasiKeterangan)
		m.RiwayatTranfusiDarah = strings.TrimSpace(d.RiwayatTranfusiDarah)
		m.RiwayatTranfusiDarahKeterangan = strings.TrimSpace(d.RiwayatTranfusiDarahKeterangan)
		m.RiwayatTranfusiDarahReaksi = strings.TrimSpace(d.RiwayatTranfusiDarahReaksi)
		m.RiwayatTranfusiDarahReaksiKeterangan = strings.TrimSpace(d.RiwayatTranfusiDarahReaksiKeterangan)
		m.KebiasanIbuObatDiminum = strings.TrimSpace(d.KebiasanIbuObatDiminum)
		m.KebiasanIbuObatDiminumKeterangan = strings.TrimSpace(d.KebiasanIbuObatDiminumKeterangan)
		m.KebiasanIbuNarkoba = strings.TrimSpace(d.KebiasanIbuNarkoba)
		m.KebiasanIbuNarkobaKeterangan = strings.TrimSpace(d.KebiasanIbuNarkobaKeterangan)
		m.KebiasanIbuMerokok = strings.TrimSpace(d.KebiasanIbuMerokok)
		m.KebiasanIbuMerokokKeterangan = strings.TrimSpace(d.KebiasanIbuMerokokKeterangan)
		m.KebiasanIbuAlkohol = strings.TrimSpace(d.KebiasanIbuAlkohol)
		m.KebiasanIbuAlkoholKeterangan = strings.TrimSpace(d.KebiasanIbuAlkoholKeterangan)
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.KeadaanUmum = strings.TrimSpace(d.KeadaanUmum)
		m.Gcs = strings.TrimSpace(d.Gcs)
		m.Td = strings.TrimSpace(d.Td)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Hr = strings.TrimSpace(d.Hr)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Spo2 = strings.TrimSpace(d.Spo2)
		m.DownScore = strings.TrimSpace(d.DownScore)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Lk = strings.TrimSpace(d.Lk)
		m.Ld = strings.TrimSpace(d.Ld)
		m.Lp = strings.TrimSpace(d.Lp)
		m.GdBayi = strings.TrimSpace(d.GdBayi)
		m.GdIbu = strings.TrimSpace(d.GdIbu)
		m.GdAyah = strings.TrimSpace(d.GdAyah)
		m.SarafPusatGerakBayi = strings.TrimSpace(d.SarafPusatGerakBayi)
		m.SarafPusatKepala = strings.TrimSpace(d.SarafPusatKepala)
		m.SarafPusatKepalaKeterangan = strings.TrimSpace(d.SarafPusatKepalaKeterangan)
		m.SarafPusatUbunubun = strings.TrimSpace(d.SarafPusatUbunubun)
		m.SarafPusatUbunubunKeterangan = strings.TrimSpace(d.SarafPusatUbunubunKeterangan)
		m.SarafPusatWajah = strings.TrimSpace(d.SarafPusatWajah)
		m.SarafPusatWajahKeterangan = strings.TrimSpace(d.SarafPusatWajahKeterangan)
		m.SarafPusatKejang = strings.TrimSpace(d.SarafPusatKejang)
		m.SarafPusatKejangKeterangan = strings.TrimSpace(d.SarafPusatKejangKeterangan)
		m.SarafPusatRefleks = strings.TrimSpace(d.SarafPusatRefleks)
		m.SarafPusatRefleksKeterangan = strings.TrimSpace(d.SarafPusatRefleksKeterangan)
		m.SarafPusatTangisbayi = strings.TrimSpace(d.SarafPusatTangisbayi)
		m.SarafPusatTangisbayiKeterangan = strings.TrimSpace(d.SarafPusatTangisbayiKeterangan)
		m.KardiovaskularDenyutnadi = strings.TrimSpace(d.KardiovaskularDenyutnadi)
		m.KardiovaskularSirkulasi = strings.TrimSpace(d.KardiovaskularSirkulasi)
		m.KardiovaskularSirkulasiKeterangan = strings.TrimSpace(d.KardiovaskularSirkulasiKeterangan)
		m.KardiovaskularPulsasi = strings.TrimSpace(d.KardiovaskularPulsasi)
		m.KardiovaskularPulsasiKeterangan = strings.TrimSpace(d.KardiovaskularPulsasiKeterangan)
		m.RespirasiPolanafas = strings.TrimSpace(d.RespirasiPolanafas)
		m.RespirasiJenispernapasan = strings.TrimSpace(d.RespirasiJenispernapasan)
		m.RespirasiJenispernapasanKeterangan = strings.TrimSpace(d.RespirasiJenispernapasanKeterangan)
		m.RespirasiRetraksi = strings.TrimSpace(d.RespirasiRetraksi)
		m.RespirasiAirentry = strings.TrimSpace(d.RespirasiAirentry)
		m.RespirasiMerintih = strings.TrimSpace(d.RespirasiMerintih)
		m.RespirasiSuaraNapas = strings.TrimSpace(d.RespirasiSuaraNapas)
		m.GastrointestinalMulut = strings.TrimSpace(d.GastrointestinalMulut)
		m.GastrointestinalMulutKeterangan = strings.TrimSpace(d.GastrointestinalMulutKeterangan)
		m.GastrointestinalLidah = strings.TrimSpace(d.GastrointestinalLidah)
		m.GastrointestinalLidahKeterangan = strings.TrimSpace(d.GastrointestinalLidahKeterangan)
		m.GastrointestinalTenggorakan = strings.TrimSpace(d.GastrointestinalTenggorakan)
		m.GastrointestinalTenggorakanKeterangan = strings.TrimSpace(d.GastrointestinalTenggorakanKeterangan)
		m.GastrointestinalAbdomen = strings.TrimSpace(d.GastrointestinalAbdomen)
		m.GastrointestinalAbdomenKeterangan = strings.TrimSpace(d.GastrointestinalAbdomenKeterangan)
		m.GastrointestinalBab = strings.TrimSpace(d.GastrointestinalBab)
		m.GastrointestinalBabKeterangan = strings.TrimSpace(d.GastrointestinalBabKeterangan)
		m.GastrointestinalWarnabab = strings.TrimSpace(d.GastrointestinalWarnabab)
		m.GastrointestinalWarnababKeterangan = strings.TrimSpace(d.GastrointestinalWarnababKeterangan)
		m.GastrointestinalBak = strings.TrimSpace(d.GastrointestinalBak)
		m.GastrointestinalBakKeterangan = strings.TrimSpace(d.GastrointestinalBakKeterangan)
		m.GastrointestinalBakwarna = strings.TrimSpace(d.GastrointestinalBakwarna)
		m.GastrointestinalBakwarnaKeterangan = strings.TrimSpace(d.GastrointestinalBakwarnaKeterangan)
		m.NeurologiPosisiMata = strings.TrimSpace(d.NeurologiPosisiMata)
		m.NeurologiKelopakMata = strings.TrimSpace(d.NeurologiKelopakMata)
		m.NeurologiKelopakMataKeterangan = strings.TrimSpace(d.NeurologiKelopakMataKeterangan)
		m.NeurologiBesarPupil = strings.TrimSpace(d.NeurologiBesarPupil)
		m.NeurologiKonjugtiva = strings.TrimSpace(d.NeurologiKonjugtiva)
		m.NeurologiKonjugtivaKeterangan = strings.TrimSpace(d.NeurologiKonjugtivaKeterangan)
		m.NeurologiSklera = strings.TrimSpace(d.NeurologiSklera)
		m.NeurologiSkleraKeterangan = strings.TrimSpace(d.NeurologiSkleraKeterangan)
		m.NeurologiPendengaran = strings.TrimSpace(d.NeurologiPendengaran)
		m.NeurologiPendengaranKeterangan = strings.TrimSpace(d.NeurologiPendengaranKeterangan)
		m.NeurologiPenciuman = strings.TrimSpace(d.NeurologiPenciuman)
		m.NeurologiPenciumanKeterangan = strings.TrimSpace(d.NeurologiPenciumanKeterangan)
		m.IntegumentWarnaKulit = strings.TrimSpace(d.IntegumentWarnaKulit)
		m.IntegumentWarnaKulitKeterangan = strings.TrimSpace(d.IntegumentWarnaKulitKeterangan)
		m.IntegumentVernicKaseosa = strings.TrimSpace(d.IntegumentVernicKaseosa)
		m.IntegumentVernicKaseosaKeterangan = strings.TrimSpace(d.IntegumentVernicKaseosaKeterangan)
		m.IntegumentTurgor = strings.TrimSpace(d.IntegumentTurgor)
		m.IntegumentLanugo = strings.TrimSpace(d.IntegumentLanugo)
		m.IntegumentKulit = strings.TrimSpace(d.IntegumentKulit)
		m.IntegumentRisikoDekubitas = strings.TrimSpace(d.IntegumentRisikoDekubitas)
		m.Reproduksi = strings.TrimSpace(d.Reproduksi)
		m.ReproduksiKeterangan = strings.TrimSpace(d.ReproduksiKeterangan)
		m.MuskuloskeletalRekoilTelinga = strings.TrimSpace(d.MuskuloskeletalRekoilTelinga)
		m.MuskuloskeletalRekoilTelingaKeterangan = strings.TrimSpace(d.MuskuloskeletalRekoilTelingaKeterangan)
		m.MuskuloskeletalLengan = strings.TrimSpace(d.MuskuloskeletalLengan)
		m.MuskuloskeletalLenganKeterangan = strings.TrimSpace(d.MuskuloskeletalLenganKeterangan)
		m.MuskuloskeletalTungkai = strings.TrimSpace(d.MuskuloskeletalTungkai)
		m.MuskuloskeletalTungkaiKeterangan = strings.TrimSpace(d.MuskuloskeletalTungkaiKeterangan)
		m.MuskuloskeletalTelapakKaki = strings.TrimSpace(d.MuskuloskeletalTelapakKaki)
		m.KondisiPsikologis = strings.TrimSpace(d.KondisiPsikologis)
		m.GangguanJiwa = strings.TrimSpace(d.GangguanJiwa)
		m.MenerimaKondisiBayi = strings.TrimSpace(d.MenerimaKondisiBayi)
		m.StatusMenikah = strings.TrimSpace(d.StatusMenikah)
		m.MasalahPernikahan = strings.TrimSpace(d.MasalahPernikahan)
		m.MasalahPernikahanKeterangan = strings.TrimSpace(d.MasalahPernikahanKeterangan)
		m.Pekerjaan = strings.TrimSpace(d.Pekerjaan)
		m.Agama = strings.TrimSpace(d.Agama)
		m.NilaiKepercayaan = strings.TrimSpace(d.NilaiKepercayaan)
		m.NilaiKepercayaanKeterangan = strings.TrimSpace(d.NilaiKepercayaanKeterangan)
		m.Suku = strings.TrimSpace(d.Suku)
		m.Pendidikan = strings.TrimSpace(d.Pendidikan)
		m.Pembayaran = strings.TrimSpace(d.Pembayaran)
		m.TinggalBersama = strings.TrimSpace(d.TinggalBersama)
		m.TinggalBersamaKeterangan = strings.TrimSpace(d.TinggalBersamaKeterangan)
		m.HubunganKeluarga = strings.TrimSpace(d.HubunganKeluarga)
		m.ResponEmosi = strings.TrimSpace(d.ResponEmosi)
		m.BahasaSehariHari = strings.TrimSpace(d.BahasaSehariHari)
		m.KemampuanBacatulis = strings.TrimSpace(d.KemampuanBacatulis)
		m.ButuhPenterjemah = strings.TrimSpace(d.ButuhPenterjemah)
		m.ButuhPenterjemahKeterangan = strings.TrimSpace(d.ButuhPenterjemahKeterangan)
		m.TerdapatHambatanBelajar = strings.TrimSpace(d.TerdapatHambatanBelajar)
		m.HambatanBelajar = strings.TrimSpace(d.HambatanBelajar)
		m.HambatanBelajarKeterangan = strings.TrimSpace(d.HambatanBelajarKeterangan)
		m.HambatanCaraBicara = strings.TrimSpace(d.HambatanCaraBicara)
		m.HambatanBahasaIsyarat = strings.TrimSpace(d.HambatanBahasaIsyarat)
		m.CaraBelajarDisukai = strings.TrimSpace(d.CaraBelajarDisukai)
		m.KesediaanMenerimaInformasi = strings.TrimSpace(d.KesediaanMenerimaInformasi)
		m.KesediaanMenerimaInformasiKeterangan = strings.TrimSpace(d.KesediaanMenerimaInformasiKeterangan)
		m.PemahamanNutrisi = strings.TrimSpace(d.PemahamanNutrisi)
		m.PemahamanPenyakit = strings.TrimSpace(d.PemahamanPenyakit)
		m.PemahamanPengobatan = strings.TrimSpace(d.PemahamanPengobatan)
		m.PemahamanPerawatan = strings.TrimSpace(d.PemahamanPerawatan)
		m.MasalahGizi1 = strings.TrimSpace(d.MasalahGizi1)
		m.NilaiGizi1 = strings.TrimSpace(d.NilaiGizi1)
		m.MasalahGizi2 = strings.TrimSpace(d.MasalahGizi2)
		m.NilaiGizi2 = strings.TrimSpace(d.NilaiGizi2)
		m.MasalahGizi3 = strings.TrimSpace(d.MasalahGizi3)
		m.NilaiGizi3 = strings.TrimSpace(d.NilaiGizi3)
		m.Totalgizi = d.Totalgizi
		m.KeteranganGizi = strings.TrimSpace(d.KeteranganGizi)
		m.PenilaianHumptydumptySkala1 = strings.TrimSpace(d.PenilaianHumptydumptySkala1)
		m.PenilaianHumptydumptyNilai1 = d.PenilaianHumptydumptyNilai1
		m.PenilaianHumptydumptySkala2 = strings.TrimSpace(d.PenilaianHumptydumptySkala2)
		m.PenilaianHumptydumptyNilai2 = d.PenilaianHumptydumptyNilai2
		m.PenilaianHumptydumptySkala3 = strings.TrimSpace(d.PenilaianHumptydumptySkala3)
		m.PenilaianHumptydumptyNilai3 = d.PenilaianHumptydumptyNilai3
		m.PenilaianHumptydumptySkala4 = strings.TrimSpace(d.PenilaianHumptydumptySkala4)
		m.PenilaianHumptydumptyNilai4 = d.PenilaianHumptydumptyNilai4
		m.PenilaianHumptydumptySkala5 = strings.TrimSpace(d.PenilaianHumptydumptySkala5)
		m.PenilaianHumptydumptyNilai5 = d.PenilaianHumptydumptyNilai5
		m.PenilaianHumptydumptySkala6 = strings.TrimSpace(d.PenilaianHumptydumptySkala6)
		m.PenilaianHumptydumptyNilai6 = d.PenilaianHumptydumptyNilai6
		m.PenilaianHumptydumptySkala7 = strings.TrimSpace(d.PenilaianHumptydumptySkala7)
		m.PenilaianHumptydumptyNilai7 = d.PenilaianHumptydumptyNilai7
		m.PenilaianHumptydumptyTotalnilai = d.PenilaianHumptydumptyTotalnilai
		m.PenilaianHumptydumptyHasil = strings.TrimSpace(d.PenilaianHumptydumptyHasil)
		m.SkalaNips1 = strings.TrimSpace(d.SkalaNips1)
		m.SkalaNips1Nilai = strings.TrimSpace(d.SkalaNips1Nilai)
		m.SkalaNips2 = strings.TrimSpace(d.SkalaNips2)
		m.SkalaNips2Nilai = strings.TrimSpace(d.SkalaNips2Nilai)
		m.SkalaNips3 = strings.TrimSpace(d.SkalaNips3)
		m.SkalaNips3Nilai = strings.TrimSpace(d.SkalaNips3Nilai)
		m.SkalaNips4 = strings.TrimSpace(d.SkalaNips4)
		m.SkalaNips4Nilai = strings.TrimSpace(d.SkalaNips4Nilai)
		m.SkalaNips5 = strings.TrimSpace(d.SkalaNips5)
		m.SkalaNips5Nilai = strings.TrimSpace(d.SkalaNips5Nilai)
		m.SkalaNipsTotal = d.SkalaNipsTotal
		m.SkalaNipsKeterangan = strings.TrimSpace(d.SkalaNipsKeterangan)
		m.InformasiPerencanaanPulang = strings.TrimSpace(d.InformasiPerencanaanPulang)
		m.LamaRatarata = strings.TrimSpace(d.LamaRatarata)
		m.PerencanaanPulang = vPerencanaanPulang
		m.KondisiKlinisPulang = strings.TrimSpace(d.KondisiKlinisPulang)
		m.PerawatanLanjutanDirumah = strings.TrimSpace(d.PerawatanLanjutanDirumah)
		m.CaraTransportasiPulang = strings.TrimSpace(d.CaraTransportasiPulang)
		m.TransportasiDigunakan = strings.TrimSpace(d.TransportasiDigunakan)
		m.Rencana = support.Nullable(d.Rencana)
		m.Nip1 = strings.TrimSpace(d.Nip1)
		m.Nip2 = strings.TrimSpace(d.Nip2)
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanRanapNeonatus]{
		{Column: "nip1", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRanapNeonatus) any { return Str(m.Nip1) }},
		{Column: "nip2", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRanapNeonatus) any { return Str(m.Nip2) }},
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianAwalKeperawatanRanapNeonatus) any { return Str(m.KdDokter) }},
	},
}
