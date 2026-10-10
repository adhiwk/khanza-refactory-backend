// Package cetak data pendukung dokumen cetak: kop instansi, identitas pasien, dan nama dari kode referensi.
package cetak

import (
	model "goravel/app/models/cetak"
	"goravel/app/support"
)

type Repository interface {
	Instansi() (*model.Instansi, []byte, error)
	IdentitasRawat(noRawat string) ([]model.Baris, error)
	IdentitasPasien(noRkmMedis string) ([]model.Baris, error)
	// Nama nama dari table.nameCol untuk kode (kosong bila tidak ada); table & kolom hanya dari daftar internal.
	Nama(table, keyCol, nameCol string, kode any) (string, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Instansi() (*model.Instansi, []byte, error) {
	var list []struct {
		model.Instansi
		Logo []byte `gorm:"column:logo"`
	}
	err := support.DB().Raw("select nama_instansi,ifnull(alamat_instansi,'') as alamat_instansi,ifnull(kabupaten,'') as kabupaten," +
		"ifnull(propinsi,'') as propinsi,kontak,email,logo from setting order by aktifkan='Yes' desc limit 1").Scan(&list)
	if err != nil || len(list) == 0 {
		return &model.Instansi{}, nil, err
	}
	return &list[0].Instansi, list[0].Logo, nil
}

type identitas struct {
	NoRawat    string `gorm:"column:no_rawat"`
	NoRkmMedis string `gorm:"column:no_rkm_medis"`
	NmPasien   string `gorm:"column:nm_pasien"`
	Jk         string `gorm:"column:jk"`
	TglLahir   string `gorm:"column:tgl_lahir"`
	Umur       string `gorm:"column:umur"`
	Alamat     string `gorm:"column:alamat"`
	TglReg     string `gorm:"column:tgl_reg"`
	Poli       string `gorm:"column:poli"`
	Dokter     string `gorm:"column:dokter"`
	Penjab     string `gorm:"column:penjab"`
}

const pasienCols = "pasien.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien,ifnull(pasien.jk,'') as jk," +
	"ifnull(date_format(pasien.tgl_lahir,'%d-%m-%Y'),'') as tgl_lahir,ifnull(pasien.umur,'') as umur," +
	"concat_ws(', ',pasien.alamat,kelurahan.nm_kel,kecamatan.nm_kec,kabupaten.nm_kab) as alamat"

const pasienJoin = " left join kelurahan on pasien.kd_kel=kelurahan.kd_kel left join kecamatan on pasien.kd_kec=kecamatan.kd_kec " +
	"left join kabupaten on pasien.kd_kab=kabupaten.kd_kab"

func pasienBaris(i identitas) []model.Baris {
	jk := map[string]string{"L": "Laki-laki", "P": "Perempuan"}[i.Jk]
	return []model.Baris{
		{Label: "No. Rekam Medis", Nilai: i.NoRkmMedis}, {Label: "Nama Pasien", Nilai: i.NmPasien},
		{Label: "Jenis Kelamin", Nilai: jk}, {Label: "Tanggal Lahir", Nilai: i.TglLahir + " (" + i.Umur + ")"},
		{Label: "Alamat", Nilai: i.Alamat},
	}
}

func (r *repository) IdentitasRawat(noRawat string) ([]model.Baris, error) {
	var list []identitas
	err := support.DB().Raw("select reg_periksa.no_rawat,"+pasienCols+",concat(date_format(reg_periksa.tgl_registrasi,'%d-%m-%Y'),' ',reg_periksa.jam_reg) as tgl_reg,"+
		"ifnull(poliklinik.nm_poli,'') as poli,ifnull(dokter.nm_dokter,'') as dokter,ifnull(penjab.png_jawab,'') as penjab "+
		"from reg_periksa inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis"+pasienJoin+
		" left join poliklinik on reg_periksa.kd_poli=poliklinik.kd_poli left join dokter on reg_periksa.kd_dokter=dokter.kd_dokter "+
		"left join penjab on reg_periksa.kd_pj=penjab.kd_pj where reg_periksa.no_rawat=?", noRawat).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	i := list[0]
	return append([]model.Baris{{Label: "No. Rawat", Nilai: i.NoRawat}}, append(pasienBaris(i),
		model.Baris{Label: "Tgl. Registrasi", Nilai: i.TglReg}, model.Baris{Label: "Poli / Dokter", Nilai: i.Poli + " / " + i.Dokter},
		model.Baris{Label: "Cara Bayar", Nilai: i.Penjab})...), nil
}

func (r *repository) IdentitasPasien(noRkmMedis string) ([]model.Baris, error) {
	var list []identitas
	err := support.DB().Raw("select "+pasienCols+" from pasien"+pasienJoin+" where pasien.no_rkm_medis=?", noRkmMedis).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return pasienBaris(list[0]), nil
}

func (r *repository) Nama(table, keyCol, nameCol string, kode any) (string, error) {
	var list []struct {
		V string `gorm:"column:v"`
	}
	if err := support.DB().Raw("select ifnull("+nameCol+",'') as v from "+table+" where "+keyCol+"=? limit 1", kode).Scan(&list); err != nil || len(list) == 0 {
		return "", err
	}
	return list[0].V, nil
}
