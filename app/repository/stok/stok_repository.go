// Package stok persistence stok obat per depo/gudang (gudangbarang, data_batch, riwayat_barang_medis).
package stok

import (
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"
)

// Posisi nilai enum riwayat_barang_medis.posisi yang dipakai.
type Posisi string

const (
	PosisiPemberianObat Posisi = "Pemberian Obat"
	PosisiPenjualan     Posisi = "Penjualan"
	PosisiPenerimaan    Posisi = "Penerimaan"
	PosisiMutasi        Posisi = "Mutasi"
	PosisiOpname        Posisi = "Opname"
)

type Repository interface {
	// Lock stok saat ini (dikunci for update); found=false bila belum ada baris gudangbarang.
	Lock(tx contractsorm.Query, kodeBrng, kdBangsal, noBatch, noFaktur string) (stok float64, found bool, err error)
	Upsert(tx contractsorm.Query, kodeBrng, kdBangsal, noBatch, noFaktur string, delta float64) error
	CatatRiwayat(tx contractsorm.Query, r Riwayat) error
	UbahSisaBatch(tx contractsorm.Query, kodeBrng, noBatch, noFaktur string, delta float64) error
}

// Riwayat satu baris riwayat_barang_medis.
type Riwayat struct {
	KodeBrng   string
	StokAwal   float64
	Masuk      float64
	Keluar     float64
	StokAkhir  float64
	Posisi     Posisi
	Waktu      time.Time
	Petugas    string
	KdBangsal  string
	Status     string // Simpan | Hapus
	NoBatch    string
	NoFaktur   string
	Keterangan string
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Lock(tx contractsorm.Query, kodeBrng, kdBangsal, noBatch, noFaktur string) (float64, bool, error) {
	var list []struct {
		Stok float64 `gorm:"column:stok"`
	}
	err := tx.Raw("select stok from gudangbarang where kode_brng=? and kd_bangsal=? and no_batch=? and no_faktur=? for update",
		kodeBrng, kdBangsal, noBatch, noFaktur).Scan(&list)
	if err != nil || len(list) == 0 {
		return 0, false, err
	}
	return list[0].Stok, true, nil
}

func (r *repository) Upsert(tx contractsorm.Query, kodeBrng, kdBangsal, noBatch, noFaktur string, delta float64) error {
	_, err := tx.Exec("insert into gudangbarang (kode_brng,kd_bangsal,stok,no_batch,no_faktur) values (?,?,?,?,?) "+
		"on duplicate key update stok=stok+?", kodeBrng, kdBangsal, delta, noBatch, noFaktur, delta)
	return err
}

func (r *repository) CatatRiwayat(tx contractsorm.Query, x Riwayat) error {
	_, err := tx.Exec("insert into riwayat_barang_medis (kode_brng,stok_awal,masuk,keluar,stok_akhir,posisi,tanggal,jam,petugas,"+
		"kd_bangsal,status,no_batch,no_faktur,keterangan) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		x.KodeBrng, x.StokAwal, x.Masuk, x.Keluar, x.StokAkhir, string(x.Posisi), x.Waktu.Format("2006-01-02"), x.Waktu.Format("15:04:05"),
		x.Petugas, x.KdBangsal, x.Status, x.NoBatch, x.NoFaktur, x.Keterangan)
	return err
}

func (r *repository) UbahSisaBatch(tx contractsorm.Query, kodeBrng, noBatch, noFaktur string, delta float64) error {
	_, err := tx.Exec("update data_batch set sisa=sisa+? where kode_brng=? and no_batch=? and no_faktur=?", delta, kodeBrng, noBatch, noFaktur)
	return err
}
