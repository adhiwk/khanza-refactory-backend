package rekammedis

import (
	"fmt"
	"strings"

	cetakmodel "goravel/app/models/cetak"
	icdmodel "goravel/app/models/icd"
	kamarmodel "goravel/app/models/kamarinap"
	obatmodel "goravel/app/models/pemberianobat"
	"goravel/app/models/perawatan"
	riwayatpasienmodel "goravel/app/models/riwayatpasien"
	icdrepo "goravel/app/repository/icd"
	kamarrepo "goravel/app/repository/kamarinap"
	obatrepo "goravel/app/repository/pemberianobat"
	pemeriksaanrepo "goravel/app/repository/pemeriksaan"
	repo "goravel/app/repository/rekammedis"
	riwayatpasienrepo "goravel/app/repository/riwayatpasien"
	tindakanrepo "goravel/app/repository/tindakan"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

var ErrPasienNotFound = support.NotFound("data pasien tidak ditemukan")

// formLain form rekam medis di luar engine asesmen yang ikut ditandai pada riwayat.
var formLain = []FormInfo{{Slug: "triase-igd", Label: "triase IGD", Table: "data_triase_igd"}}

// RiwayatKunjungan isi rekam medis satu kunjungan (RMRiwayatPerawatan).
type RiwayatKunjungan struct {
	repo.Kunjungan
	Diagnosa    []icdmodel.Kode         `json:"diagnosa"`
	Prosedur    []icdmodel.Kode         `json:"prosedur"`
	Pemeriksaan []perawatan.Pemeriksaan `json:"pemeriksaan"`
	Tindakan    []perawatan.Tindakan    `json:"tindakan"`
	Obat        []obatmodel.Pemberian   `json:"obat"`
	Kamar       []kamarmodel.KamarInap  `json:"kamar"`
	Asesmen     []FormInfo              `json:"asesmen"`
}

type Riwayat struct {
	Pasien     *repo.Pasien                    `json:"pasien"`
	Persalinan []riwayatpasienmodel.Persalinan `json:"riwayat_persalinan"`
	Imunisasi  []riwayatpasienmodel.Imunisasi  `json:"riwayat_imunisasi"`
	Kunjungan  []RiwayatKunjungan              `json:"kunjungan"`
}

// RiwayatAction menyusun riwayat rekam medis dari modul pelayanan yang sudah ada.
type RiwayatAction struct {
	riwayat     repo.RiwayatStore
	icd         icdrepo.Repository
	pemeriksaan pemeriksaanrepo.Repository
	tindakan    tindakanrepo.Repository
	obat        obatrepo.Repository
	kamar       kamarrepo.Repository
	pasien      riwayatpasienrepo.Repository
	cetak       *cetaksvc.Service
}

func NewRiwayatAction(riwayat repo.RiwayatStore, icd icdrepo.Repository, pemeriksaan pemeriksaanrepo.Repository,
	tindakan tindakanrepo.Repository, obat obatrepo.Repository, kamar kamarrepo.Repository, pasien riwayatpasienrepo.Repository,
	cetak *cetaksvc.Service) *RiwayatAction {
	return &RiwayatAction{riwayat: riwayat, icd: icd, pemeriksaan: pemeriksaan, tindakan: tindakan, obat: obat, kamar: kamar, pasien: pasien, cetak: cetak}
}

// Riwayat kunjungan pasien terbaru lebih dulu (paginasi per kunjungan).
func (a *RiwayatAction) Riwayat(noRkmMedis, tglAwal, tglAkhir string, page, limit int) (*Riwayat, int64, error) {
	p, err := a.riwayat.Pasien(strings.TrimSpace(noRkmMedis))
	if err != nil {
		return nil, 0, err
	}
	if p == nil {
		return nil, 0, ErrPasienNotFound
	}
	kunjungan, total, err := a.riwayat.Kunjungan(p.NoRkmMedis, tglAwal, tglAkhir, page, limit)
	if err != nil {
		return nil, 0, err
	}
	forms := append(append([]FormInfo{}, DaftarForm...), formLain...)
	tables := make([][2]string, len(forms))
	label := map[string]FormInfo{}
	for i, f := range forms {
		tables[i] = [2]string{f.Slug, f.Table}
		label[f.Slug] = f
	}

	out := &Riwayat{Pasien: p, Kunjungan: make([]RiwayatKunjungan, 0, len(kunjungan))}
	if out.Persalinan, err = a.pasien.Persalinan(p.NoRkmMedis); err != nil {
		return nil, 0, err
	}
	if out.Imunisasi, err = a.pasien.Imunisasi(p.NoRkmMedis); err != nil {
		return nil, 0, err
	}
	for _, k := range kunjungan {
		rk := RiwayatKunjungan{Kunjungan: k, Asesmen: []FormInfo{}}
		if rk.Diagnosa, err = a.icd.List(icdmodel.Diagnosa, k.NoRawat, ""); err != nil {
			return nil, 0, err
		}
		if rk.Prosedur, err = a.icd.List(icdmodel.Prosedur, k.NoRawat, ""); err != nil {
			return nil, 0, err
		}
		for _, rawat := range []perawatan.Rawat{perawatan.Ralan, perawatan.Ranap} {
			pem, err := a.pemeriksaan.List(rawat, k.NoRawat)
			if err != nil {
				return nil, 0, err
			}
			tin, err := a.tindakan.List(rawat, k.NoRawat)
			if err != nil {
				return nil, 0, err
			}
			obt, err := a.obat.List(k.NoRawat, rawat)
			if err != nil {
				return nil, 0, err
			}
			rk.Pemeriksaan = append(rk.Pemeriksaan, pem...)
			rk.Tindakan = append(rk.Tindakan, tin...)
			rk.Obat = append(rk.Obat, obt...)
		}
		if rk.Kamar, err = a.kamar.Riwayat(k.NoRawat); err != nil {
			return nil, 0, err
		}
		slugs, err := a.riwayat.FormTerisi(k.NoRawat, tables)
		if err != nil {
			return nil, 0, err
		}
		for _, s := range slugs {
			rk.Asesmen = append(rk.Asesmen, label[s])
		}
		out.Kunjungan = append(out.Kunjungan, rk)
	}
	return out, total, nil
}

// Cetak resume rekam medis pasien (RMRiwayatPerawatan): riwayat persalinan & imunisasi lalu per kunjungan
// diagnosa, prosedur, SOAP, tindakan, obat, kamar, dan asesmen yang terisi. Maksimal maxCetak kunjungan terbaru.
func (a *RiwayatAction) Cetak(noRkmMedis, tglAwal, tglAkhir string) (*cetakmodel.Dokumen, error) {
	r, _, err := a.Riwayat(noRkmMedis, tglAwal, tglAkhir, 1, maxCetak)
	if err != nil {
		return nil, err
	}
	dok, err := a.cetak.Dokumen("Riwayat Rekam Medis Pasien", r.Pasien.NoRkmMedis)
	if err != nil {
		return nil, err
	}
	if dok.Identitas, err = a.cetak.IdentitasPasien(r.Pasien.NoRkmMedis); err != nil {
		return nil, err
	}
	f := cetaksvc.Format
	if r.Pasien.Catatan != "" {
		dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Catatan Pasien", Baris: []cetakmodel.Baris{{Label: "Catatan", Nilai: r.Pasien.Catatan}}})
	}
	if len(r.Persalinan) > 0 {
		t := &cetakmodel.Tabel{Kolom: []string{"Tgl/Thn", "Tempat", "Usia Hamil", "Jenis", "Penolong", "Penyulit", "JK", "BB/PB", "Keadaan"}}
		for _, p := range r.Persalinan {
			t.Isi = append(t.Isi, []string{p.TglThn, f(p.TempatPersalinan), f(p.UsiaHamil), f(p.JenisPersalinan), f(p.Penolong), f(p.Penyulit), p.Jk, f(p.Bbpb), f(p.Keadaan)})
		}
		dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Riwayat Persalinan", Tabel: t})
	}
	if len(r.Imunisasi) > 0 {
		t := &cetakmodel.Tabel{Kolom: []string{"Kode", "Imunisasi", "Ke-"}}
		for _, i := range r.Imunisasi {
			t.Isi = append(t.Isi, []string{i.KodeImunisasi, i.NamaImunisasi, fmt.Sprint(i.NoImunisasi)})
		}
		dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Riwayat Imunisasi", Tabel: t})
	}
	for _, k := range r.Kunjungan {
		dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Kunjungan " + k.NoRawat, Baris: []cetakmodel.Baris{
			{Label: "Tanggal", Nilai: k.TglRegistrasi + " " + k.JamReg}, {Label: "Poli / Dokter", Nilai: k.NmPoli + " / " + k.NmDokter},
			{Label: "Status", Nilai: k.StatusLanjut + " (" + k.Stts + ")"}, {Label: "Cara Bayar", Nilai: k.PngJawab},
		}})
		if len(k.Diagnosa)+len(k.Prosedur) > 0 {
			t := &cetakmodel.Tabel{Kolom: []string{"Jenis", "Kode", "Uraian", "Status", "Prioritas"}}
			for _, d := range k.Diagnosa {
				t.Isi = append(t.Isi, []string{"Diagnosa", d.Kode, d.Nama, d.Status + " / " + d.StatusPenyakit, fmt.Sprint(d.Prioritas)})
			}
			for _, d := range k.Prosedur {
				t.Isi = append(t.Isi, []string{"Prosedur", d.Kode, d.Nama, d.Status, fmt.Sprint(d.Prioritas)})
			}
			dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Tabel: t})
		}
		if len(k.Pemeriksaan) > 0 {
			t := &cetakmodel.Tabel{Kolom: []string{"Waktu", "Subjek", "Objek", "Asesmen", "Plan", "Pemeriksa"}}
			for _, p := range k.Pemeriksaan {
				t.Isi = append(t.Isi, []string{p.TglPerawatan + " " + p.JamRawat, f(p.Keluhan), f(p.Pemeriksaan), f(p.Penilaian), f(p.Rtl), p.NmPegawai})
			}
			dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Tabel: t})
		}
		if len(k.Tindakan) > 0 {
			t := &cetakmodel.Tabel{Kolom: []string{"Waktu", "Tindakan", "Pelaksana"}}
			for _, x := range k.Tindakan {
				t.Isi = append(t.Isi, []string{x.TglPerawatan + " " + x.JamRawat, x.NmPerawatan, strings.Trim(x.NmDokter+" / "+x.NmPetugas, " /")})
			}
			dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Tabel: t})
		}
		if len(k.Obat) > 0 {
			t := &cetakmodel.Tabel{Kolom: []string{"Waktu", "Obat", "Jumlah", "Aturan Pakai"}}
			for _, o := range k.Obat {
				t.Isi = append(t.Isi, []string{o.TglPerawatan + " " + o.Jam, o.NamaBrng, f(o.Jml), f(o.AturanPakai)})
			}
			dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Tabel: t})
		}
		if len(k.Kamar) > 0 {
			t := &cetakmodel.Tabel{Kolom: []string{"Kamar", "Masuk", "Keluar", "Diagnosa Awal", "Diagnosa Akhir", "Status Pulang"}}
			for _, km := range k.Kamar {
				t.Isi = append(t.Isi, []string{km.KdKamar + " " + km.NmBangsal, km.TglMasuk + " " + km.JamMasuk, km.TglKeluar + " " + km.JamKeluar,
					f(km.DiagnosaAwal), f(km.DiagnosaAkhir), km.SttsPulang})
			}
			dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Tabel: t})
		}
		if len(k.Asesmen) > 0 {
			labels := make([]string, len(k.Asesmen))
			for i, as := range k.Asesmen {
				labels[i] = as.Label
			}
			dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Baris: []cetakmodel.Baris{{Label: "Asesmen terisi", Nilai: strings.Join(labels, ", ")}}})
		}
	}
	return dok, nil
}

// maxCetak batas kunjungan pada resume cetak (persempit dengan tgl_awal/tgl_akhir).
const maxCetak = 50
