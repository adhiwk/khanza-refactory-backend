package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianAwalKeperawatanKebidananRanap laporan pemantauan anastesi (RMLaporanPemantauanAnastesi).
var FormPenilaianAwalKeperawatanKebidananRanap = &Form[model.PenilaianAwalKeperawatanKebidananRanap, request.PenilaianAwalKeperawatanKebidananRanapData]{
	Slug:  "penilaian-awal-keperawatan-kebidanan-ranap",
	Label: "laporan pemantauan anastesi",
	Spec: repo.Spec{
		Table:  "penilaian_awal_keperawatan_kebidanan_ranap",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip1", "nip2", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianAwalKeperawatanKebidananRanap, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianAwalKeperawatanKebidananRanap) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu: func(m *model.PenilaianAwalKeperawatanKebidananRanap) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianAwalKeperawatanKebidananRanap) []string {
		return []string{m.Nip1, m.Nip2, m.KdDokter}
	},
	Fill: func(m *model.PenilaianAwalKeperawatanKebidananRanap, d request.PenilaianAwalKeperawatanKebidananRanapData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		vRiwayatHamilHpht, err := support.ParseDate(d.RiwayatHamilHpht)
		if err != nil {
			return err
		}
		vRiwayatHamilTp, err := support.ParseDate(d.RiwayatHamilTp)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.TibaDiruangRawat = strings.TrimSpace(d.TibaDiruangRawat)
		m.CaraMasuk = strings.TrimSpace(d.CaraMasuk)
		m.Keluhan = strings.TrimSpace(d.Keluhan)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Psk = strings.TrimSpace(d.Psk)
		m.Rp = strings.TrimSpace(d.Rp)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.KomplikasiSebelumnya = strings.TrimSpace(d.KomplikasiSebelumnya)
		m.KeteranganKomplikasiSebelumnya = strings.TrimSpace(d.KeteranganKomplikasiSebelumnya)
		m.RiwayatMensUmur = strings.TrimSpace(d.RiwayatMensUmur)
		m.RiwayatMensLamanya = strings.TrimSpace(d.RiwayatMensLamanya)
		m.RiwayatMensBanyaknya = strings.TrimSpace(d.RiwayatMensBanyaknya)
		m.RiwayatMensSiklus = strings.TrimSpace(d.RiwayatMensSiklus)
		m.RiwayatMensKetSiklus = strings.TrimSpace(d.RiwayatMensKetSiklus)
		m.RiwayatMensDirasakan = strings.TrimSpace(d.RiwayatMensDirasakan)
		m.RiwayatPerkawinanStatus = strings.TrimSpace(d.RiwayatPerkawinanStatus)
		m.RiwayatPerkawinanKetStatus = strings.TrimSpace(d.RiwayatPerkawinanKetStatus)
		m.RiwayatPerkawinanUsia1 = strings.TrimSpace(d.RiwayatPerkawinanUsia1)
		m.RiwayatPerkawinanKetUsia1 = strings.TrimSpace(d.RiwayatPerkawinanKetUsia1)
		m.RiwayatPerkawinanUsia2 = strings.TrimSpace(d.RiwayatPerkawinanUsia2)
		m.RiwayatPerkawinanKetUsia2 = strings.TrimSpace(d.RiwayatPerkawinanKetUsia2)
		m.RiwayatPerkawinanUsia3 = strings.TrimSpace(d.RiwayatPerkawinanUsia3)
		m.RiwayatPerkawinanKetUsia3 = strings.TrimSpace(d.RiwayatPerkawinanKetUsia3)
		m.RiwayatPersalinanG = strings.TrimSpace(d.RiwayatPersalinanG)
		m.RiwayatPersalinanP = strings.TrimSpace(d.RiwayatPersalinanP)
		m.RiwayatPersalinanA = strings.TrimSpace(d.RiwayatPersalinanA)
		m.RiwayatPersalinanHidup = strings.TrimSpace(d.RiwayatPersalinanHidup)
		m.RiwayatHamilHpht = vRiwayatHamilHpht
		m.RiwayatHamilUsiahamil = strings.TrimSpace(d.RiwayatHamilUsiahamil)
		m.RiwayatHamilTp = vRiwayatHamilTp
		m.RiwayatHamilImunisasi = strings.TrimSpace(d.RiwayatHamilImunisasi)
		m.RiwayatHamilAnc = strings.TrimSpace(d.RiwayatHamilAnc)
		m.RiwayatHamilAncke = strings.TrimSpace(d.RiwayatHamilAncke)
		m.RiwayatHamilKetAncke = strings.TrimSpace(d.RiwayatHamilKetAncke)
		m.RiwayatHamilKeluhanHamilMuda = strings.TrimSpace(d.RiwayatHamilKeluhanHamilMuda)
		m.RiwayatHamilKeluhanHamilTua = strings.TrimSpace(d.RiwayatHamilKeluhanHamilTua)
		m.RiwayatKb = strings.TrimSpace(d.RiwayatKb)
		m.RiwayatKbLamanya = strings.TrimSpace(d.RiwayatKbLamanya)
		m.RiwayatKbKomplikasi = strings.TrimSpace(d.RiwayatKbKomplikasi)
		m.RiwayatKbKetKomplikasi = strings.TrimSpace(d.RiwayatKbKetKomplikasi)
		m.RiwayatKbKapaberhenti = strings.TrimSpace(d.RiwayatKbKapaberhenti)
		m.RiwayatKbAlasanberhenti = strings.TrimSpace(d.RiwayatKbAlasanberhenti)
		m.RiwayatGenekologi = strings.TrimSpace(d.RiwayatGenekologi)
		m.RiwayatKebiasaanObat = strings.TrimSpace(d.RiwayatKebiasaanObat)
		m.RiwayatKebiasaanKetObat = strings.TrimSpace(d.RiwayatKebiasaanKetObat)
		m.RiwayatKebiasaanMerokok = strings.TrimSpace(d.RiwayatKebiasaanMerokok)
		m.RiwayatKebiasaanKetMerokok = strings.TrimSpace(d.RiwayatKebiasaanKetMerokok)
		m.RiwayatKebiasaanAlkohol = strings.TrimSpace(d.RiwayatKebiasaanAlkohol)
		m.RiwayatKebiasaanKetAlkohol = strings.TrimSpace(d.RiwayatKebiasaanKetAlkohol)
		m.RiwayatKebiasaanNarkoba = strings.TrimSpace(d.RiwayatKebiasaanNarkoba)
		m.PemeriksaanKebidananMental = strings.TrimSpace(d.PemeriksaanKebidananMental)
		m.PemeriksaanKebidananKeadaanUmum = strings.TrimSpace(d.PemeriksaanKebidananKeadaanUmum)
		m.PemeriksaanKebidananGcs = strings.TrimSpace(d.PemeriksaanKebidananGcs)
		m.PemeriksaanKebidananTd = strings.TrimSpace(d.PemeriksaanKebidananTd)
		m.PemeriksaanKebidananNadi = strings.TrimSpace(d.PemeriksaanKebidananNadi)
		m.PemeriksaanKebidananRr = strings.TrimSpace(d.PemeriksaanKebidananRr)
		m.PemeriksaanKebidananSuhu = strings.TrimSpace(d.PemeriksaanKebidananSuhu)
		m.PemeriksaanKebidananSpo2 = strings.TrimSpace(d.PemeriksaanKebidananSpo2)
		m.PemeriksaanKebidananBb = strings.TrimSpace(d.PemeriksaanKebidananBb)
		m.PemeriksaanKebidananTb = strings.TrimSpace(d.PemeriksaanKebidananTb)
		m.PemeriksaanKebidananLila = strings.TrimSpace(d.PemeriksaanKebidananLila)
		m.PemeriksaanKebidananTfu = strings.TrimSpace(d.PemeriksaanKebidananTfu)
		m.PemeriksaanKebidananTbj = strings.TrimSpace(d.PemeriksaanKebidananTbj)
		m.PemeriksaanKebidananLetak = strings.TrimSpace(d.PemeriksaanKebidananLetak)
		m.PemeriksaanKebidananPresentasi = strings.TrimSpace(d.PemeriksaanKebidananPresentasi)
		m.PemeriksaanKebidananPenurunan = strings.TrimSpace(d.PemeriksaanKebidananPenurunan)
		m.PemeriksaanKebidananHis = strings.TrimSpace(d.PemeriksaanKebidananHis)
		m.PemeriksaanKebidananKekuatan = strings.TrimSpace(d.PemeriksaanKebidananKekuatan)
		m.PemeriksaanKebidananLamanya = strings.TrimSpace(d.PemeriksaanKebidananLamanya)
		m.PemeriksaanKebidananDjj = strings.TrimSpace(d.PemeriksaanKebidananDjj)
		m.PemeriksaanKebidananKetDjj = strings.TrimSpace(d.PemeriksaanKebidananKetDjj)
		m.PemeriksaanKebidananPortio = strings.TrimSpace(d.PemeriksaanKebidananPortio)
		m.PemeriksaanKebidananPembukaan = strings.TrimSpace(d.PemeriksaanKebidananPembukaan)
		m.PemeriksaanKebidananKetuban = strings.TrimSpace(d.PemeriksaanKebidananKetuban)
		m.PemeriksaanKebidananHodge = strings.TrimSpace(d.PemeriksaanKebidananHodge)
		m.PemeriksaanKebidananPanggul = strings.TrimSpace(d.PemeriksaanKebidananPanggul)
		m.PemeriksaanKebidananInspekulo = strings.TrimSpace(d.PemeriksaanKebidananInspekulo)
		m.PemeriksaanKebidananKetInspekulo = strings.TrimSpace(d.PemeriksaanKebidananKetInspekulo)
		m.PemeriksaanKebidananLakmus = strings.TrimSpace(d.PemeriksaanKebidananLakmus)
		m.PemeriksaanKebidananKetLakmus = strings.TrimSpace(d.PemeriksaanKebidananKetLakmus)
		m.PemeriksaanKebidananCtg = strings.TrimSpace(d.PemeriksaanKebidananCtg)
		m.PemeriksaanKebidananKetCtg = strings.TrimSpace(d.PemeriksaanKebidananKetCtg)
		m.PemeriksaanUmumKepala = strings.TrimSpace(d.PemeriksaanUmumKepala)
		m.PemeriksaanUmumMuka = strings.TrimSpace(d.PemeriksaanUmumMuka)
		m.PemeriksaanUmumMata = strings.TrimSpace(d.PemeriksaanUmumMata)
		m.PemeriksaanUmumHidung = strings.TrimSpace(d.PemeriksaanUmumHidung)
		m.PemeriksaanUmumTelinga = strings.TrimSpace(d.PemeriksaanUmumTelinga)
		m.PemeriksaanUmumMulut = strings.TrimSpace(d.PemeriksaanUmumMulut)
		m.PemeriksaanUmumLeher = strings.TrimSpace(d.PemeriksaanUmumLeher)
		m.PemeriksaanUmumDada = strings.TrimSpace(d.PemeriksaanUmumDada)
		m.PemeriksaanUmumPerut = strings.TrimSpace(d.PemeriksaanUmumPerut)
		m.PemeriksaanUmumGenitalia = strings.TrimSpace(d.PemeriksaanUmumGenitalia)
		m.PemeriksaanUmumEkstrimitas = strings.TrimSpace(d.PemeriksaanUmumEkstrimitas)
		m.PengkajianFungsiKemampuanAktifitas = strings.TrimSpace(d.PengkajianFungsiKemampuanAktifitas)
		m.PengkajianFungsiBerjalan = strings.TrimSpace(d.PengkajianFungsiBerjalan)
		m.PengkajianFungsiKetBerjalan = strings.TrimSpace(d.PengkajianFungsiKetBerjalan)
		m.PengkajianFungsiAktivitas = strings.TrimSpace(d.PengkajianFungsiAktivitas)
		m.PengkajianFungsiAmbulasi = strings.TrimSpace(d.PengkajianFungsiAmbulasi)
		m.PengkajianFungsiEkstrimitasAtas = strings.TrimSpace(d.PengkajianFungsiEkstrimitasAtas)
		m.PengkajianFungsiKetEkstrimitasAtas = strings.TrimSpace(d.PengkajianFungsiKetEkstrimitasAtas)
		m.PengkajianFungsiEkstrimitasBawah = strings.TrimSpace(d.PengkajianFungsiEkstrimitasBawah)
		m.PengkajianFungsiKetEkstrimitasBawah = strings.TrimSpace(d.PengkajianFungsiKetEkstrimitasBawah)
		m.PengkajianFungsiKemampuanMenggenggam = strings.TrimSpace(d.PengkajianFungsiKemampuanMenggenggam)
		m.PengkajianFungsiKetKemampuanMenggenggam = strings.TrimSpace(d.PengkajianFungsiKetKemampuanMenggenggam)
		m.PengkajianFungsiKoordinasi = strings.TrimSpace(d.PengkajianFungsiKoordinasi)
		m.PengkajianFungsiKetKoordinasi = strings.TrimSpace(d.PengkajianFungsiKetKoordinasi)
		m.PengkajianFungsiGangguanFungsi = strings.TrimSpace(d.PengkajianFungsiGangguanFungsi)
		m.RiwayatPsikoKondisipsiko = strings.TrimSpace(d.RiwayatPsikoKondisipsiko)
		m.RiwayatPsikoAdakahPrilaku = strings.TrimSpace(d.RiwayatPsikoAdakahPrilaku)
		m.RiwayatPsikoKetAdakahPrilaku = strings.TrimSpace(d.RiwayatPsikoKetAdakahPrilaku)
		m.RiwayatPsikoGangguanJiwa = strings.TrimSpace(d.RiwayatPsikoGangguanJiwa)
		m.RiwayatPsikoHubunganPasien = strings.TrimSpace(d.RiwayatPsikoHubunganPasien)
		m.RiwayatPsikoTinggalDengan = strings.TrimSpace(d.RiwayatPsikoTinggalDengan)
		m.RiwayatPsikoKetTinggalDengan = strings.TrimSpace(d.RiwayatPsikoKetTinggalDengan)
		m.RiwayatPsikoBudaya = strings.TrimSpace(d.RiwayatPsikoBudaya)
		m.RiwayatPsikoKetBudaya = strings.TrimSpace(d.RiwayatPsikoKetBudaya)
		m.RiwayatPsikoPendPj = strings.TrimSpace(d.RiwayatPsikoPendPj)
		m.RiwayatPsikoEdukasiPada = strings.TrimSpace(d.RiwayatPsikoEdukasiPada)
		m.RiwayatPsikoKetEdukasiPada = strings.TrimSpace(d.RiwayatPsikoKetEdukasiPada)
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
		m.PenilaianJatuhSkala1 = strings.TrimSpace(d.PenilaianJatuhSkala1)
		m.PenilaianJatuhNilai1 = d.PenilaianJatuhNilai1
		m.PenilaianJatuhSkala2 = strings.TrimSpace(d.PenilaianJatuhSkala2)
		m.PenilaianJatuhNilai2 = d.PenilaianJatuhNilai2
		m.PenilaianJatuhSkala3 = strings.TrimSpace(d.PenilaianJatuhSkala3)
		m.PenilaianJatuhNilai3 = d.PenilaianJatuhNilai3
		m.PenilaianJatuhSkala4 = strings.TrimSpace(d.PenilaianJatuhSkala4)
		m.PenilaianJatuhNilai4 = d.PenilaianJatuhNilai4
		m.PenilaianJatuhSkala5 = strings.TrimSpace(d.PenilaianJatuhSkala5)
		m.PenilaianJatuhNilai5 = d.PenilaianJatuhNilai5
		m.PenilaianJatuhSkala6 = strings.TrimSpace(d.PenilaianJatuhSkala6)
		m.PenilaianJatuhNilai6 = d.PenilaianJatuhNilai6
		m.PenilaianJatuhTotalnilai = d.PenilaianJatuhTotalnilai
		m.SkriningGizi1 = strings.TrimSpace(d.SkriningGizi1)
		m.NilaiGizi1 = d.NilaiGizi1
		m.SkriningGizi2 = strings.TrimSpace(d.SkriningGizi2)
		m.NilaiGizi2 = d.NilaiGizi2
		m.NilaiTotalGizi = d.NilaiTotalGizi
		m.SkriningGiziDiagnosaKhusus = strings.TrimSpace(d.SkriningGiziDiagnosaKhusus)
		m.SkriningGiziKetDiagnosaKhusus = strings.TrimSpace(d.SkriningGiziKetDiagnosaKhusus)
		m.SkriningGiziDiketahuiDietisen = strings.TrimSpace(d.SkriningGiziDiketahuiDietisen)
		m.SkriningGiziJamDiketahuiDietisen = strings.TrimSpace(d.SkriningGiziJamDiketahuiDietisen)
		m.Masalah = strings.TrimSpace(d.Masalah)
		m.Rencana = strings.TrimSpace(d.Rencana)
		m.Nip1 = strings.TrimSpace(d.Nip1)
		m.Nip2 = strings.TrimSpace(d.Nip2)
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.PenilaianAwalKeperawatanKebidananRanap]{
		{Column: "nip1", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanKebidananRanap) any { return Str(m.Nip1) }},
		{Column: "nip2", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianAwalKeperawatanKebidananRanap) any { return Str(m.Nip2) }},
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianAwalKeperawatanKebidananRanap) any { return Str(m.KdDokter) }},
	},
}
