package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanRanapBayi penilaian awal keperawatan ranap bayi anak (RMPenilaianAwalKeperawatanRanapBayiAnak).
var FormPenilaianAwalKeperawatanRanapBayi = &Form[model.PenilaianAwalKeperawatanRanapBayi, request.PenilaianAwalKeperawatanRanapBayiData]{
	Slug:  "penilaian-awal-keperawatan-ranap-bayi",
	Label: "penilaian awal keperawatan ranap bayi anak",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_ranap_bayi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip1", "nip2", "kd_dokter"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_ranap_bayi_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan_anak", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ranap_bayi_rencana", Column: "kode_rencana", RefTable: "master_rencana_keperawatan_anak", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanRanapBayi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanRanapBayi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanRanapBayi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanRanapBayi) []string { return []string{m.Nip1, m.Nip2, m.KdDokter} },
	Fill: func(m *model.PenilaianAwalKeperawatanRanapBayi, d request.PenilaianAwalKeperawatanRanapBayiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("inte_decubi", d.InteDecubi, "Tidak Ada", "Usia > 65 tahun", "Obesitas", "Imobilisasi", "Paraplegi/Vegetative State", "Dirawat Di HCU", "Penyakit Kronis (DM, CHF, CKD)", "Inkontinentia Uri/Alvi"); err != nil {
			return err
		}
		if err := Enum("penilaian_humptydumpty_skala3", d.PenilaianHumptydumptySkala3, "Kelainan Neurologi", "Perubahan Dalam Oksigen(Masalah Saluran Nafas, Dehidrasi, Anemia, Anoreksia / Sakit Kepala, Dll)", "Kelainan Psikis / Perilaku", "Diagnosa Lain"); err != nil {
			return err
		}
		if err := Enum("penilaian_humptydumpty_skala7", d.PenilaianHumptydumptySkala7, "Bermacam-macam Obat Yang Digunakan : Obat Sedative (Kecuali Pasien ICU Yang Menggunakan sedasi dan paralisis), Hipnotik, Barbiturat, Fenoti-Azin, Antidepresan, Laksans/Diuretika,Narkotik", "Salah Satu Dari Pengobatan Di Atas", "Pengobatan Lain"); err != nil {
			return err
		}
		if err := Enum("nyeri_aktifitas", d.NyeriAktifitas, "Tidur posisi normal, mudah bergerak", "Gerakan menggeliat/berguling, kaku", "Melengkungkan punggung/kaku menghentak"); err != nil {
			return err
		}
		if err := Enum("nyeri_menangis", d.NyeriMenangis, "Tidak menangis (mudah bergerak)", "Mengerang/merengek", "Menangis terus menerus, terisak, menjerit"); err != nil {
			return err
		}
		if err := Enum("nyeri_bersuara", d.NyeriBersuara, "Bersuara normal/tenang", "Tenang bila dipeluk, digendong/diajak bicara", "Sulit untuk menenangkan"); err != nil {
			return err
		}
		vPerencanaanPulang, err := support.ParseDate(d.PerencanaanPulang)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.KetInformasi = strings.TrimSpace(d.KetInformasi)
		m.TibaDiruangRawat = strings.TrimSpace(d.TibaDiruangRawat)
		m.Rps = strings.TrimSpace(d.Rps)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.TumbuhKembangTengkurap = support.Nullable(d.TumbuhKembangTengkurap)
		m.TumbuhKembangDuduk = support.Nullable(d.TumbuhKembangDuduk)
		m.TumbuhKembangBerdiri = support.Nullable(d.TumbuhKembangBerdiri)
		m.TumbuhKembangGigiPertama = support.Nullable(d.TumbuhKembangGigiPertama)
		m.TumbuhKembangBerjalan = support.Nullable(d.TumbuhKembangBerjalan)
		m.TumbuhKembangBicara = support.Nullable(d.TumbuhKembangBicara)
		m.TumbuhKembangMembaca = support.Nullable(d.TumbuhKembangMembaca)
		m.TumbuhKembangMenulis = support.Nullable(d.TumbuhKembangMenulis)
		m.TumbuhKembangGangguanEmosi = support.Nullable(d.TumbuhKembangGangguanEmosi)
		m.PersalinanAnakke = strings.TrimSpace(d.PersalinanAnakke)
		m.PersalinanDarisaudara = strings.TrimSpace(d.PersalinanDarisaudara)
		m.PersalinanKelahiran = strings.TrimSpace(d.PersalinanKelahiran)
		m.PersalinanKelahiranKeterangan = strings.TrimSpace(d.PersalinanKelahiranKeterangan)
		m.PersalinanUmurKelahiran = strings.TrimSpace(d.PersalinanUmurKelahiran)
		m.PersalinanKelainanBawaan = strings.TrimSpace(d.PersalinanKelainanBawaan)
		m.PersalinanKelainanBawaanKeterangan = strings.TrimSpace(d.PersalinanKelainanBawaanKeterangan)
		m.PersalinanBbLahir = strings.TrimSpace(d.PersalinanBbLahir)
		m.PersalinanPbLahir = strings.TrimSpace(d.PersalinanPbLahir)
		m.PersalinanLainnya = strings.TrimSpace(d.PersalinanLainnya)
		m.FisikKesadaran = strings.TrimSpace(d.FisikKesadaran)
		m.FisikGcs = strings.TrimSpace(d.FisikGcs)
		m.FisikTd = strings.TrimSpace(d.FisikTd)
		m.FisikRr = strings.TrimSpace(d.FisikRr)
		m.FisikSuhu = strings.TrimSpace(d.FisikSuhu)
		m.FisikNadi = strings.TrimSpace(d.FisikNadi)
		m.FisikBb = strings.TrimSpace(d.FisikBb)
		m.FisikTb = strings.TrimSpace(d.FisikTb)
		m.FisikLp = strings.TrimSpace(d.FisikLp)
		m.FisikLk = strings.TrimSpace(d.FisikLk)
		m.FisikLd = strings.TrimSpace(d.FisikLd)
		m.SarafPusatKepala = strings.TrimSpace(d.SarafPusatKepala)
		m.SarafPusatKepalaKeterangan = strings.TrimSpace(d.SarafPusatKepalaKeterangan)
		m.SarafPusatWajah = strings.TrimSpace(d.SarafPusatWajah)
		m.SarafPusatWajahKeterangan = strings.TrimSpace(d.SarafPusatWajahKeterangan)
		m.SarafPusatLeher = strings.TrimSpace(d.SarafPusatLeher)
		m.SarafPusatKejang = strings.TrimSpace(d.SarafPusatKejang)
		m.SarafPusatKejangKeterangan = strings.TrimSpace(d.SarafPusatKejangKeterangan)
		m.SarafPusatSensorik = strings.TrimSpace(d.SarafPusatSensorik)
		m.KardiovaskulerPulsasi = strings.TrimSpace(d.KardiovaskulerPulsasi)
		m.KardiovaskulerSirkulasi = strings.TrimSpace(d.KardiovaskulerSirkulasi)
		m.KardiovaskulerSirkulasiKeterangan = strings.TrimSpace(d.KardiovaskulerSirkulasiKeterangan)
		m.KardiovaskulerDenyutNadi = strings.TrimSpace(d.KardiovaskulerDenyutNadi)
		m.RespirasiRetraksi = strings.TrimSpace(d.RespirasiRetraksi)
		m.RespirasiPolaNafas = strings.TrimSpace(d.RespirasiPolaNafas)
		m.RespirasiSuaraNafas = strings.TrimSpace(d.RespirasiSuaraNafas)
		m.RespirasiBatuk = strings.TrimSpace(d.RespirasiBatuk)
		m.RespirasiVolume = strings.TrimSpace(d.RespirasiVolume)
		m.RespirasiJenisPernapasan = strings.TrimSpace(d.RespirasiJenisPernapasan)
		m.RespirasiJenisPernapasanKeterangan = strings.TrimSpace(d.RespirasiJenisPernapasanKeterangan)
		m.RespirasiIrama = strings.TrimSpace(d.RespirasiIrama)
		m.GastroMulut = strings.TrimSpace(d.GastroMulut)
		m.GastroMulutKeterangan = strings.TrimSpace(d.GastroMulutKeterangan)
		m.GastroTenggorakan = strings.TrimSpace(d.GastroTenggorakan)
		m.GastroTenggorakanKeterangan = strings.TrimSpace(d.GastroTenggorakanKeterangan)
		m.GastroLidah = strings.TrimSpace(d.GastroLidah)
		m.GastroLidahKeterangan = strings.TrimSpace(d.GastroLidahKeterangan)
		m.GastroAbdomen = strings.TrimSpace(d.GastroAbdomen)
		m.GastroAbdomenKeterangan = strings.TrimSpace(d.GastroAbdomenKeterangan)
		m.GastroGigi = strings.TrimSpace(d.GastroGigi)
		m.GastroGigiKeterangan = strings.TrimSpace(d.GastroGigiKeterangan)
		m.GastroUsus = strings.TrimSpace(d.GastroUsus)
		m.GastroAnus = strings.TrimSpace(d.GastroAnus)
		m.NeurologiSensorik = strings.TrimSpace(d.NeurologiSensorik)
		m.NeurologiPengilihatan = strings.TrimSpace(d.NeurologiPengilihatan)
		m.NeurologiPengilihatanKeterangan = strings.TrimSpace(d.NeurologiPengilihatanKeterangan)
		m.NeurologiAlatBantuPenglihatan = strings.TrimSpace(d.NeurologiAlatBantuPenglihatan)
		m.NeurologiMotorik = strings.TrimSpace(d.NeurologiMotorik)
		m.NeurologiPendengaran = strings.TrimSpace(d.NeurologiPendengaran)
		m.NeurologiBicara = strings.TrimSpace(d.NeurologiBicara)
		m.NeurologiBicaraKeterangan = strings.TrimSpace(d.NeurologiBicaraKeterangan)
		m.NeurologiOtot = strings.TrimSpace(d.NeurologiOtot)
		m.InteKulit = strings.TrimSpace(d.InteKulit)
		m.InteWarnaKulit = strings.TrimSpace(d.InteWarnaKulit)
		m.InteTugor = strings.TrimSpace(d.InteTugor)
		m.InteDecubi = strings.TrimSpace(d.InteDecubi)
		m.MuskuOdema = strings.TrimSpace(d.MuskuOdema)
		m.MuskuOdemaKeterangan = strings.TrimSpace(d.MuskuOdemaKeterangan)
		m.MuskuPegerakansendi = strings.TrimSpace(d.MuskuPegerakansendi)
		m.MuskuOtot = strings.TrimSpace(d.MuskuOtot)
		m.MuskuFraktur = strings.TrimSpace(d.MuskuFraktur)
		m.MuskuFrakturKeterangan = strings.TrimSpace(d.MuskuFrakturKeterangan)
		m.MuskuNyerisendi = strings.TrimSpace(d.MuskuNyerisendi)
		m.MuskuNyerisendiKeterangan = strings.TrimSpace(d.MuskuNyerisendiKeterangan)
		m.EliminasiBabFrekuensi = strings.TrimSpace(d.EliminasiBabFrekuensi)
		m.EliminasiBabFrekuensiPer = strings.TrimSpace(d.EliminasiBabFrekuensiPer)
		m.EliminasiBabKonsistesi = strings.TrimSpace(d.EliminasiBabKonsistesi)
		m.EliminasiBabWarna = strings.TrimSpace(d.EliminasiBabWarna)
		m.EliminasiBakFrekuensi = strings.TrimSpace(d.EliminasiBakFrekuensi)
		m.EliminasiBakFrekuensiPer = strings.TrimSpace(d.EliminasiBakFrekuensiPer)
		m.EliminasiBakWarna = strings.TrimSpace(d.EliminasiBakWarna)
		m.EliminasiBakLainlain = strings.TrimSpace(d.EliminasiBakLainlain)
		m.PsikoKondisi = strings.TrimSpace(d.PsikoKondisi)
		m.PsikoPerilaku = strings.TrimSpace(d.PsikoPerilaku)
		m.PsikoPerilakuKeterangan = strings.TrimSpace(d.PsikoPerilakuKeterangan)
		m.PsikoGangguanJiwa = strings.TrimSpace(d.PsikoGangguanJiwa)
		m.PsikoHubunganPasien = strings.TrimSpace(d.PsikoHubunganPasien)
		m.PsikoTinggalDengan = strings.TrimSpace(d.PsikoTinggalDengan)
		m.PsikoTinggalDenganKeterangan = strings.TrimSpace(d.PsikoTinggalDenganKeterangan)
		m.PsikoPekerjaanPj = strings.TrimSpace(d.PsikoPekerjaanPj)
		m.PsikoNilaiKepercayaan = strings.TrimSpace(d.PsikoNilaiKepercayaan)
		m.PsikoNilaiKepercayaanKeterangan = strings.TrimSpace(d.PsikoNilaiKepercayaanKeterangan)
		m.PsikoPendidikanPj = strings.TrimSpace(d.PsikoPendidikanPj)
		m.PsikoEdukasi = strings.TrimSpace(d.PsikoEdukasi)
		m.PsikoEdukasiKeterangan = strings.TrimSpace(d.PsikoEdukasiKeterangan)
		m.EdukasiBahasa = strings.TrimSpace(d.EdukasiBahasa)
		m.EdukasiBacaTulis = strings.TrimSpace(d.EdukasiBacaTulis)
		m.EdukasiPenerjemah = strings.TrimSpace(d.EdukasiPenerjemah)
		m.EdukasiPenerjemahKeterangan = strings.TrimSpace(d.EdukasiPenerjemahKeterangan)
		m.EdukasiTerdapatHambatan = strings.TrimSpace(d.EdukasiTerdapatHambatan)
		m.EdukasiHambatanBelajar = strings.TrimSpace(d.EdukasiHambatanBelajar)
		m.EdukasiHambatanBelajarKeterangan = strings.TrimSpace(d.EdukasiHambatanBelajarKeterangan)
		m.EdukasiHambatanBicara = strings.TrimSpace(d.EdukasiHambatanBicara)
		m.EdukasiBahasaIsyarat = strings.TrimSpace(d.EdukasiBahasaIsyarat)
		m.EdukasiCaraBelajar = strings.TrimSpace(d.EdukasiCaraBelajar)
		m.EdukasiMenerimaInformasi = strings.TrimSpace(d.EdukasiMenerimaInformasi)
		m.EdukasiMenerimaInformasiKeterangan = strings.TrimSpace(d.EdukasiMenerimaInformasiKeterangan)
		m.EdukasiNutrisi = strings.TrimSpace(d.EdukasiNutrisi)
		m.EdukasiPenyakit = strings.TrimSpace(d.EdukasiPenyakit)
		m.EdukasiPengobatan = strings.TrimSpace(d.EdukasiPengobatan)
		m.EdukasiPerawatan = strings.TrimSpace(d.EdukasiPerawatan)
		m.SkriningGizi1 = strings.TrimSpace(d.SkriningGizi1)
		m.NilaiGizi1 = strings.TrimSpace(d.NilaiGizi1)
		m.SkriningGizi2 = strings.TrimSpace(d.SkriningGizi2)
		m.NilaiGizi2 = strings.TrimSpace(d.NilaiGizi2)
		m.SkriningGizi3 = strings.TrimSpace(d.SkriningGizi3)
		m.NilaiGizi3 = strings.TrimSpace(d.NilaiGizi3)
		m.SkriningGizi4 = strings.TrimSpace(d.SkriningGizi4)
		m.NilaiGizi4 = strings.TrimSpace(d.NilaiGizi4)
		m.TotalNilai = strings.TrimSpace(d.TotalNilai)
		m.KeteranganSkriningGizi = strings.TrimSpace(d.KeteranganSkriningGizi)
		m.PenilaianHumptydumptySkala1 = support.Nullable(d.PenilaianHumptydumptySkala1)
		m.PenilaianHumptydumptyNilai1 = d.PenilaianHumptydumptyNilai1
		m.PenilaianHumptydumptySkala2 = support.Nullable(d.PenilaianHumptydumptySkala2)
		m.PenilaianHumptydumptyNilai2 = d.PenilaianHumptydumptyNilai2
		m.PenilaianHumptydumptySkala3 = support.Nullable(d.PenilaianHumptydumptySkala3)
		m.PenilaianHumptydumptyNilai3 = d.PenilaianHumptydumptyNilai3
		m.PenilaianHumptydumptySkala4 = support.Nullable(d.PenilaianHumptydumptySkala4)
		m.PenilaianHumptydumptyNilai4 = d.PenilaianHumptydumptyNilai4
		m.PenilaianHumptydumptySkala5 = support.Nullable(d.PenilaianHumptydumptySkala5)
		m.PenilaianHumptydumptyNilai5 = d.PenilaianHumptydumptyNilai5
		m.PenilaianHumptydumptySkala6 = support.Nullable(d.PenilaianHumptydumptySkala6)
		m.PenilaianHumptydumptyNilai6 = d.PenilaianHumptydumptyNilai6
		m.PenilaianHumptydumptySkala7 = support.Nullable(d.PenilaianHumptydumptySkala7)
		m.PenilaianHumptydumptyNilai7 = d.PenilaianHumptydumptyNilai7
		m.PenilaianHumptydumptyTotalnilai = d.PenilaianHumptydumptyTotalnilai
		m.HasilSkriningPenilaianHumptydumpty = support.Nullable(d.HasilSkriningPenilaianHumptydumpty)
		m.NyeriWajah = strings.TrimSpace(d.NyeriWajah)
		m.NyeriNilaiWajah = strings.TrimSpace(d.NyeriNilaiWajah)
		m.NyeriKaki = strings.TrimSpace(d.NyeriKaki)
		m.NyeriNilaiKaki = strings.TrimSpace(d.NyeriNilaiKaki)
		m.NyeriAktifitas = strings.TrimSpace(d.NyeriAktifitas)
		m.NyeriNilaiAktifitas = strings.TrimSpace(d.NyeriNilaiAktifitas)
		m.NyeriMenangis = strings.TrimSpace(d.NyeriMenangis)
		m.NyeriNilaiMenangis = strings.TrimSpace(d.NyeriNilaiMenangis)
		m.NyeriBersuara = strings.TrimSpace(d.NyeriBersuara)
		m.NyeriNilaiBersuara = strings.TrimSpace(d.NyeriNilaiBersuara)
		m.NyeriNilaiTotal = strings.TrimSpace(d.NyeriNilaiTotal)
		m.NyeriKondisi = strings.TrimSpace(d.NyeriKondisi)
		m.NyeriLokasi = strings.TrimSpace(d.NyeriLokasi)
		m.NyeriDurasi = strings.TrimSpace(d.NyeriDurasi)
		m.NyeriFrekuensi = strings.TrimSpace(d.NyeriFrekuensi)
		m.NyeriHilang = strings.TrimSpace(d.NyeriHilang)
		m.NyeriHilangKeterangan = strings.TrimSpace(d.NyeriHilangKeterangan)
		m.NyeriDiberitahukanPadaDokter = strings.TrimSpace(d.NyeriDiberitahukanPadaDokter)
		m.NyeriDiberitahukanPadaDokterKeterangan = strings.TrimSpace(d.NyeriDiberitahukanPadaDokterKeterangan)
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
	Refs: []Ref[model.PenilaianAwalKeperawatanRanapBayi]{
		{Column: "nip1", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRanapBayi) any { return Str(m.Nip1) }},
		{Column: "nip2", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRanapBayi) any { return Str(m.Nip2) }},
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianAwalKeperawatanRanapBayi) any { return Str(m.KdDokter) }},
	},
}
