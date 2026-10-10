package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanRanap penilaian awal keperawatan ranap (RMPenilaianAwalKeperawatanRanap).
var FormPenilaianAwalKeperawatanRanap = &Form[model.PenilaianAwalKeperawatanRanap, request.PenilaianAwalKeperawatanRanapData]{
	Slug:  "penilaian-awal-keperawatan-ranap",
	Label: "penilaian awal keperawatan ranap",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_ranap",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip1", "nip2", "kd_dokter"},
		Details: []repo.Detail{
			{Name: "kode_masalah", Table: "penilaian_awal_keperawatan_ranap_masalah", Column: "kode_masalah", RefTable: "master_masalah_keperawatan", RefColumn: "kode_masalah"},
			{Name: "kode_rencana", Table: "penilaian_awal_keperawatan_ranap_rencana", Column: "kode_rencana", RefTable: "master_rencana_keperawatan", RefColumn: "kode_rencana"},
		},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanRanap, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanRanap) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianAwalKeperawatanRanap) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanRanap) []string { return []string{m.Nip1, m.Nip2, m.KdDokter} },
	Fill: func(m *model.PenilaianAwalKeperawatanRanap, d request.PenilaianAwalKeperawatanRanapData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("pemeriksaan_integument_dekubitas", d.PemeriksaanIntegumentDekubitas, "Tidak Ada", "Usia > 65 tahun", "Obesitas", "Imobilisasi", "Paraplegi/Vegetative State", "Dirawat Di HCU", "Penyakit Kronis (DM, CHF, CKD)", "Inkontinentia Uri/Alvi"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.KetInformasi = strings.TrimSpace(d.KetInformasi)
		m.TibaDiruangRawat = strings.TrimSpace(d.TibaDiruangRawat)
		m.KasusTrauma = support.Nullable(d.KasusTrauma)
		m.CaraMasuk = strings.TrimSpace(d.CaraMasuk)
		m.Rps = strings.TrimSpace(d.Rps)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Rpo = strings.TrimSpace(d.Rpo)
		m.RiwayatPembedahan = strings.TrimSpace(d.RiwayatPembedahan)
		m.RiwayatDirawatDirs = strings.TrimSpace(d.RiwayatDirawatDirs)
		m.AlatBantuDipakai = strings.TrimSpace(d.AlatBantuDipakai)
		m.RiwayatKehamilan = strings.TrimSpace(d.RiwayatKehamilan)
		m.RiwayatKehamilanPerkiraan = strings.TrimSpace(d.RiwayatKehamilanPerkiraan)
		m.RiwayatTranfusi = strings.TrimSpace(d.RiwayatTranfusi)
		m.RiwayatAlergi = strings.TrimSpace(d.RiwayatAlergi)
		m.RiwayatMerokok = strings.TrimSpace(d.RiwayatMerokok)
		m.RiwayatMerokokJumlah = strings.TrimSpace(d.RiwayatMerokokJumlah)
		m.RiwayatAlkohol = strings.TrimSpace(d.RiwayatAlkohol)
		m.RiwayatAlkoholJumlah = strings.TrimSpace(d.RiwayatAlkoholJumlah)
		m.RiwayatNarkoba = strings.TrimSpace(d.RiwayatNarkoba)
		m.RiwayatOlahraga = strings.TrimSpace(d.RiwayatOlahraga)
		m.PemeriksaanMental = strings.TrimSpace(d.PemeriksaanMental)
		m.PemeriksaanKeadaanUmum = strings.TrimSpace(d.PemeriksaanKeadaanUmum)
		m.PemeriksaanGcs = strings.TrimSpace(d.PemeriksaanGcs)
		m.PemeriksaanTd = strings.TrimSpace(d.PemeriksaanTd)
		m.PemeriksaanNadi = strings.TrimSpace(d.PemeriksaanNadi)
		m.PemeriksaanRr = strings.TrimSpace(d.PemeriksaanRr)
		m.PemeriksaanSuhu = strings.TrimSpace(d.PemeriksaanSuhu)
		m.PemeriksaanSpo2 = strings.TrimSpace(d.PemeriksaanSpo2)
		m.PemeriksaanBb = strings.TrimSpace(d.PemeriksaanBb)
		m.PemeriksaanTb = strings.TrimSpace(d.PemeriksaanTb)
		m.PemeriksaanSusunanKepala = strings.TrimSpace(d.PemeriksaanSusunanKepala)
		m.PemeriksaanSusunanKepalaKeterangan = strings.TrimSpace(d.PemeriksaanSusunanKepalaKeterangan)
		m.PemeriksaanSusunanWajah = strings.TrimSpace(d.PemeriksaanSusunanWajah)
		m.PemeriksaanSusunanWajahKeterangan = strings.TrimSpace(d.PemeriksaanSusunanWajahKeterangan)
		m.PemeriksaanSusunanLeher = strings.TrimSpace(d.PemeriksaanSusunanLeher)
		m.PemeriksaanSusunanKejang = strings.TrimSpace(d.PemeriksaanSusunanKejang)
		m.PemeriksaanSusunanKejangKeterangan = strings.TrimSpace(d.PemeriksaanSusunanKejangKeterangan)
		m.PemeriksaanSusunanSensorik = strings.TrimSpace(d.PemeriksaanSusunanSensorik)
		m.PemeriksaanKardiovaskulerDenyutNadi = strings.TrimSpace(d.PemeriksaanKardiovaskulerDenyutNadi)
		m.PemeriksaanKardiovaskulerSirkulasi = strings.TrimSpace(d.PemeriksaanKardiovaskulerSirkulasi)
		m.PemeriksaanKardiovaskulerSirkulasiKeterangan = strings.TrimSpace(d.PemeriksaanKardiovaskulerSirkulasiKeterangan)
		m.PemeriksaanKardiovaskulerPulsasi = strings.TrimSpace(d.PemeriksaanKardiovaskulerPulsasi)
		m.PemeriksaanRespirasiPolaNafas = strings.TrimSpace(d.PemeriksaanRespirasiPolaNafas)
		m.PemeriksaanRespirasiRetraksi = strings.TrimSpace(d.PemeriksaanRespirasiRetraksi)
		m.PemeriksaanRespirasiSuaraNafas = strings.TrimSpace(d.PemeriksaanRespirasiSuaraNafas)
		m.PemeriksaanRespirasiVolumePernafasan = strings.TrimSpace(d.PemeriksaanRespirasiVolumePernafasan)
		m.PemeriksaanRespirasiJenisPernafasan = strings.TrimSpace(d.PemeriksaanRespirasiJenisPernafasan)
		m.PemeriksaanRespirasiJenisPernafasanKeterangan = strings.TrimSpace(d.PemeriksaanRespirasiJenisPernafasanKeterangan)
		m.PemeriksaanRespirasiIramaNafas = strings.TrimSpace(d.PemeriksaanRespirasiIramaNafas)
		m.PemeriksaanRespirasiBatuk = strings.TrimSpace(d.PemeriksaanRespirasiBatuk)
		m.PemeriksaanGastrointestinalMulut = strings.TrimSpace(d.PemeriksaanGastrointestinalMulut)
		m.PemeriksaanGastrointestinalMulutKeterangan = strings.TrimSpace(d.PemeriksaanGastrointestinalMulutKeterangan)
		m.PemeriksaanGastrointestinalGigi = strings.TrimSpace(d.PemeriksaanGastrointestinalGigi)
		m.PemeriksaanGastrointestinalGigiKeterangan = strings.TrimSpace(d.PemeriksaanGastrointestinalGigiKeterangan)
		m.PemeriksaanGastrointestinalLidah = strings.TrimSpace(d.PemeriksaanGastrointestinalLidah)
		m.PemeriksaanGastrointestinalLidahKeterangan = strings.TrimSpace(d.PemeriksaanGastrointestinalLidahKeterangan)
		m.PemeriksaanGastrointestinalTenggorokan = strings.TrimSpace(d.PemeriksaanGastrointestinalTenggorokan)
		m.PemeriksaanGastrointestinalTenggorokanKeterangan = strings.TrimSpace(d.PemeriksaanGastrointestinalTenggorokanKeterangan)
		m.PemeriksaanGastrointestinalAbdomen = strings.TrimSpace(d.PemeriksaanGastrointestinalAbdomen)
		m.PemeriksaanGastrointestinalAbdomenKeterangan = strings.TrimSpace(d.PemeriksaanGastrointestinalAbdomenKeterangan)
		m.PemeriksaanGastrointestinalPeistatikUsus = strings.TrimSpace(d.PemeriksaanGastrointestinalPeistatikUsus)
		m.PemeriksaanGastrointestinalAnus = strings.TrimSpace(d.PemeriksaanGastrointestinalAnus)
		m.PemeriksaanNeurologiPengelihatan = strings.TrimSpace(d.PemeriksaanNeurologiPengelihatan)
		m.PemeriksaanNeurologiPengelihatanKeterangan = strings.TrimSpace(d.PemeriksaanNeurologiPengelihatanKeterangan)
		m.PemeriksaanNeurologiAlatBantuPenglihatan = strings.TrimSpace(d.PemeriksaanNeurologiAlatBantuPenglihatan)
		m.PemeriksaanNeurologiPendengaran = strings.TrimSpace(d.PemeriksaanNeurologiPendengaran)
		m.PemeriksaanNeurologiBicara = strings.TrimSpace(d.PemeriksaanNeurologiBicara)
		m.PemeriksaanNeurologiBicaraKeterangan = strings.TrimSpace(d.PemeriksaanNeurologiBicaraKeterangan)
		m.PemeriksaanNeurologiSensorik = strings.TrimSpace(d.PemeriksaanNeurologiSensorik)
		m.PemeriksaanNeurologiMotorik = strings.TrimSpace(d.PemeriksaanNeurologiMotorik)
		m.PemeriksaanNeurologiKekuatanOtot = strings.TrimSpace(d.PemeriksaanNeurologiKekuatanOtot)
		m.PemeriksaanIntegumentWarnakulit = strings.TrimSpace(d.PemeriksaanIntegumentWarnakulit)
		m.PemeriksaanIntegumentTurgor = strings.TrimSpace(d.PemeriksaanIntegumentTurgor)
		m.PemeriksaanIntegumentKulit = strings.TrimSpace(d.PemeriksaanIntegumentKulit)
		m.PemeriksaanIntegumentDekubitas = strings.TrimSpace(d.PemeriksaanIntegumentDekubitas)
		m.PemeriksaanMuskuloskletalPergerakanSendi = strings.TrimSpace(d.PemeriksaanMuskuloskletalPergerakanSendi)
		m.PemeriksaanMuskuloskletalKekauatanOtot = strings.TrimSpace(d.PemeriksaanMuskuloskletalKekauatanOtot)
		m.PemeriksaanMuskuloskletalNyeriSendi = strings.TrimSpace(d.PemeriksaanMuskuloskletalNyeriSendi)
		m.PemeriksaanMuskuloskletalNyeriSendiKeterangan = strings.TrimSpace(d.PemeriksaanMuskuloskletalNyeriSendiKeterangan)
		m.PemeriksaanMuskuloskletalOedema = strings.TrimSpace(d.PemeriksaanMuskuloskletalOedema)
		m.PemeriksaanMuskuloskletalOedemaKeterangan = strings.TrimSpace(d.PemeriksaanMuskuloskletalOedemaKeterangan)
		m.PemeriksaanMuskuloskletalFraktur = strings.TrimSpace(d.PemeriksaanMuskuloskletalFraktur)
		m.PemeriksaanMuskuloskletalFrakturKeterangan = strings.TrimSpace(d.PemeriksaanMuskuloskletalFrakturKeterangan)
		m.PemeriksaanEliminasiBabFrekuensiJumlah = strings.TrimSpace(d.PemeriksaanEliminasiBabFrekuensiJumlah)
		m.PemeriksaanEliminasiBabFrekuensiDurasi = strings.TrimSpace(d.PemeriksaanEliminasiBabFrekuensiDurasi)
		m.PemeriksaanEliminasiBabKonsistensi = strings.TrimSpace(d.PemeriksaanEliminasiBabKonsistensi)
		m.PemeriksaanEliminasiBabWarna = strings.TrimSpace(d.PemeriksaanEliminasiBabWarna)
		m.PemeriksaanEliminasiBakFrekuensiJumlah = strings.TrimSpace(d.PemeriksaanEliminasiBakFrekuensiJumlah)
		m.PemeriksaanEliminasiBakFrekuensiDurasi = strings.TrimSpace(d.PemeriksaanEliminasiBakFrekuensiDurasi)
		m.PemeriksaanEliminasiBakWarna = strings.TrimSpace(d.PemeriksaanEliminasiBakWarna)
		m.PemeriksaanEliminasiBakLainlain = strings.TrimSpace(d.PemeriksaanEliminasiBakLainlain)
		m.PolaAktifitasMakanminum = strings.TrimSpace(d.PolaAktifitasMakanminum)
		m.PolaAktifitasMandi = strings.TrimSpace(d.PolaAktifitasMandi)
		m.PolaAktifitasEliminasi = strings.TrimSpace(d.PolaAktifitasEliminasi)
		m.PolaAktifitasBerpakaian = strings.TrimSpace(d.PolaAktifitasBerpakaian)
		m.PolaAktifitasBerpindah = strings.TrimSpace(d.PolaAktifitasBerpindah)
		m.PolaNutrisiFrekuesiMakan = strings.TrimSpace(d.PolaNutrisiFrekuesiMakan)
		m.PolaNutrisiJenisMakanan = strings.TrimSpace(d.PolaNutrisiJenisMakanan)
		m.PolaNutrisiPorsiMakan = strings.TrimSpace(d.PolaNutrisiPorsiMakan)
		m.PolaTidurLamaTidur = strings.TrimSpace(d.PolaTidurLamaTidur)
		m.PolaTidurGangguan = strings.TrimSpace(d.PolaTidurGangguan)
		m.PengkajianFungsiKemampuanSehari = strings.TrimSpace(d.PengkajianFungsiKemampuanSehari)
		m.PengkajianFungsiAktifitas = strings.TrimSpace(d.PengkajianFungsiAktifitas)
		m.PengkajianFungsiBerjalan = strings.TrimSpace(d.PengkajianFungsiBerjalan)
		m.PengkajianFungsiBerjalanKeterangan = strings.TrimSpace(d.PengkajianFungsiBerjalanKeterangan)
		m.PengkajianFungsiAmbulasi = strings.TrimSpace(d.PengkajianFungsiAmbulasi)
		m.PengkajianFungsiEkstrimitasAtas = strings.TrimSpace(d.PengkajianFungsiEkstrimitasAtas)
		m.PengkajianFungsiEkstrimitasAtasKeterangan = strings.TrimSpace(d.PengkajianFungsiEkstrimitasAtasKeterangan)
		m.PengkajianFungsiEkstrimitasBawah = strings.TrimSpace(d.PengkajianFungsiEkstrimitasBawah)
		m.PengkajianFungsiEkstrimitasBawahKeterangan = strings.TrimSpace(d.PengkajianFungsiEkstrimitasBawahKeterangan)
		m.PengkajianFungsiMenggenggam = strings.TrimSpace(d.PengkajianFungsiMenggenggam)
		m.PengkajianFungsiMenggenggamKeterangan = strings.TrimSpace(d.PengkajianFungsiMenggenggamKeterangan)
		m.PengkajianFungsiKoordinasi = strings.TrimSpace(d.PengkajianFungsiKoordinasi)
		m.PengkajianFungsiKoordinasiKeterangan = strings.TrimSpace(d.PengkajianFungsiKoordinasiKeterangan)
		m.PengkajianFungsiKesimpulan = strings.TrimSpace(d.PengkajianFungsiKesimpulan)
		m.RiwayatPsikoKondisiPsiko = strings.TrimSpace(d.RiwayatPsikoKondisiPsiko)
		m.RiwayatPsikoGangguanJiwa = strings.TrimSpace(d.RiwayatPsikoGangguanJiwa)
		m.RiwayatPsikoPerilaku = strings.TrimSpace(d.RiwayatPsikoPerilaku)
		m.RiwayatPsikoPerilakuKeterangan = strings.TrimSpace(d.RiwayatPsikoPerilakuKeterangan)
		m.RiwayatPsikoHubunganKeluarga = strings.TrimSpace(d.RiwayatPsikoHubunganKeluarga)
		m.RiwayatPsikoTinggal = strings.TrimSpace(d.RiwayatPsikoTinggal)
		m.RiwayatPsikoTinggalKeterangan = strings.TrimSpace(d.RiwayatPsikoTinggalKeterangan)
		m.RiwayatPsikoNilaiKepercayaan = strings.TrimSpace(d.RiwayatPsikoNilaiKepercayaan)
		m.RiwayatPsikoNilaiKepercayaanKeterangan = strings.TrimSpace(d.RiwayatPsikoNilaiKepercayaanKeterangan)
		m.RiwayatPsikoPendidikanPj = strings.TrimSpace(d.RiwayatPsikoPendidikanPj)
		m.RiwayatPsikoEdukasiDiberikan = strings.TrimSpace(d.RiwayatPsikoEdukasiDiberikan)
		m.RiwayatPsikoEdukasiDiberikanKeterangan = strings.TrimSpace(d.RiwayatPsikoEdukasiDiberikanKeterangan)
		m.PenilaianNyeri = strings.TrimSpace(d.PenilaianNyeri)
		m.PenilaianNyeriPenyebab = strings.TrimSpace(d.PenilaianNyeriPenyebab)
		m.PenilaianNyeriKetPenyebab = strings.TrimSpace(d.PenilaianNyeriKetPenyebab)
		m.PenilaianNyeriKualitas = strings.TrimSpace(d.PenilaianNyeriKualitas)
		m.PenilaianNyeriKetKualitas = strings.TrimSpace(d.PenilaianNyeriKetKualitas)
		m.PenilaianNyeriLokasi = strings.TrimSpace(d.PenilaianNyeriLokasi)
		m.PenilaianNyeriMenyebar = strings.TrimSpace(d.PenilaianNyeriMenyebar)
		m.PenilaianNyeriSkala = strings.TrimSpace(d.PenilaianNyeriSkala)
		m.PenilaianNyeriWaktu = strings.TrimSpace(d.PenilaianNyeriWaktu)
		m.PenilaianNyeriHilang = strings.TrimSpace(d.PenilaianNyeriHilang)
		m.PenilaianNyeriKetHilang = strings.TrimSpace(d.PenilaianNyeriKetHilang)
		m.PenilaianNyeriDiberitahukanDokter = strings.TrimSpace(d.PenilaianNyeriDiberitahukanDokter)
		m.PenilaianNyeriJamDiberitahukanDokter = strings.TrimSpace(d.PenilaianNyeriJamDiberitahukanDokter)
		m.PenilaianJatuhmorseSkala1 = support.Nullable(d.PenilaianJatuhmorseSkala1)
		m.PenilaianJatuhmorseNilai1 = d.PenilaianJatuhmorseNilai1
		m.PenilaianJatuhmorseSkala2 = support.Nullable(d.PenilaianJatuhmorseSkala2)
		m.PenilaianJatuhmorseNilai2 = d.PenilaianJatuhmorseNilai2
		m.PenilaianJatuhmorseSkala3 = support.Nullable(d.PenilaianJatuhmorseSkala3)
		m.PenilaianJatuhmorseNilai3 = d.PenilaianJatuhmorseNilai3
		m.PenilaianJatuhmorseSkala4 = support.Nullable(d.PenilaianJatuhmorseSkala4)
		m.PenilaianJatuhmorseNilai4 = d.PenilaianJatuhmorseNilai4
		m.PenilaianJatuhmorseSkala5 = support.Nullable(d.PenilaianJatuhmorseSkala5)
		m.PenilaianJatuhmorseNilai5 = d.PenilaianJatuhmorseNilai5
		m.PenilaianJatuhmorseSkala6 = support.Nullable(d.PenilaianJatuhmorseSkala6)
		m.PenilaianJatuhmorseNilai6 = d.PenilaianJatuhmorseNilai6
		m.PenilaianJatuhmorseTotalnilai = d.PenilaianJatuhmorseTotalnilai
		m.PenilaianJatuhsydneySkala1 = support.Nullable(d.PenilaianJatuhsydneySkala1)
		m.PenilaianJatuhsydneyNilai1 = d.PenilaianJatuhsydneyNilai1
		m.PenilaianJatuhsydneySkala2 = support.Nullable(d.PenilaianJatuhsydneySkala2)
		m.PenilaianJatuhsydneyNilai2 = d.PenilaianJatuhsydneyNilai2
		m.PenilaianJatuhsydneySkala3 = support.Nullable(d.PenilaianJatuhsydneySkala3)
		m.PenilaianJatuhsydneyNilai3 = d.PenilaianJatuhsydneyNilai3
		m.PenilaianJatuhsydneySkala4 = support.Nullable(d.PenilaianJatuhsydneySkala4)
		m.PenilaianJatuhsydneyNilai4 = d.PenilaianJatuhsydneyNilai4
		m.PenilaianJatuhsydneySkala5 = support.Nullable(d.PenilaianJatuhsydneySkala5)
		m.PenilaianJatuhsydneyNilai5 = d.PenilaianJatuhsydneyNilai5
		m.PenilaianJatuhsydneySkala6 = support.Nullable(d.PenilaianJatuhsydneySkala6)
		m.PenilaianJatuhsydneyNilai6 = d.PenilaianJatuhsydneyNilai6
		m.PenilaianJatuhsydneySkala7 = support.Nullable(d.PenilaianJatuhsydneySkala7)
		m.PenilaianJatuhsydneyNilai7 = d.PenilaianJatuhsydneyNilai7
		m.PenilaianJatuhsydneySkala8 = support.Nullable(d.PenilaianJatuhsydneySkala8)
		m.PenilaianJatuhsydneyNilai8 = d.PenilaianJatuhsydneyNilai8
		m.PenilaianJatuhsydneySkala9 = support.Nullable(d.PenilaianJatuhsydneySkala9)
		m.PenilaianJatuhsydneyNilai9 = d.PenilaianJatuhsydneyNilai9
		m.PenilaianJatuhsydneySkala10 = support.Nullable(d.PenilaianJatuhsydneySkala10)
		m.PenilaianJatuhsydneyNilai10 = d.PenilaianJatuhsydneyNilai10
		m.PenilaianJatuhsydneySkala11 = support.Nullable(d.PenilaianJatuhsydneySkala11)
		m.PenilaianJatuhsydneyNilai11 = d.PenilaianJatuhsydneyNilai11
		m.PenilaianJatuhsydneyTotalnilai = d.PenilaianJatuhsydneyTotalnilai
		m.SkriningGizi1 = support.Nullable(d.SkriningGizi1)
		m.NilaiGizi1 = d.NilaiGizi1
		m.SkriningGizi2 = support.Nullable(d.SkriningGizi2)
		m.NilaiGizi2 = d.NilaiGizi2
		m.NilaiTotalGizi = d.NilaiTotalGizi
		m.SkriningGiziDiagnosaKhusus = support.Nullable(d.SkriningGiziDiagnosaKhusus)
		m.SkriningGiziKetDiagnosaKhusus = support.Nullable(d.SkriningGiziKetDiagnosaKhusus)
		m.SkriningGiziDiketahuiDietisen = support.Nullable(d.SkriningGiziDiketahuiDietisen)
		m.SkriningGiziJamDiketahuiDietisen = support.Nullable(d.SkriningGiziJamDiketahuiDietisen)
		m.Rencana = support.Nullable(d.Rencana)
		m.Nip1 = strings.TrimSpace(d.Nip1)
		m.Nip2 = strings.TrimSpace(d.Nip2)
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanRanap]{
		{Column: "nip1", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRanap) any { return Str(m.Nip1) }},
		{Column: "nip2", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanRanap) any { return Str(m.Nip2) }},
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianAwalKeperawatanRanap) any { return Str(m.KdDokter) }},
	},
}
