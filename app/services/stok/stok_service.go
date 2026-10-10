// Package stok kemampuan mutasi stok obat (pengganti riwayatobat.catatRiwayat + update gudangbarang Khanza).
package stok

import (
	"fmt"
	"math"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	stokrepo "goravel/app/repository/stok"
	"goravel/app/support"
)

// Mutasi satu perubahan stok pada satu depo + batch.
type Mutasi struct {
	KodeBrng   string
	KdBangsal  string
	NoBatch    string
	NoFaktur   string
	Jumlah     float64
	Posisi     stokrepo.Posisi
	Petugas    string
	Keterangan string
	Batal      bool // true = mutasi pembatalan (riwayat berstatus Hapus)
}

type Service struct {
	repo  stokrepo.Repository
	batch bool
	now   func() time.Time
}

// NewService batch=true juga memperbarui data_batch.sisa (AKTIFKANBATCHOBAT).
func NewService(repo stokrepo.Repository, batch bool) *Service {
	return &Service{repo: repo, batch: batch, now: time.Now}
}

// Keluar mengurangi stok; ditolak bila stok depo tidak mencukupi.
func (s *Service) Keluar(tx contractsorm.Query, m Mutasi) error {
	return s.mutasi(tx, m, -m.Jumlah)
}

// Masuk menambah stok (penerimaan, pembatalan pemberian/penjualan).
func (s *Service) Masuk(tx contractsorm.Query, m Mutasi) error {
	return s.mutasi(tx, m, m.Jumlah)
}

func (s *Service) mutasi(tx contractsorm.Query, m Mutasi, delta float64) error {
	if m.Jumlah <= 0 {
		return support.Invalid("jumlah obat harus lebih dari 0")
	}
	awal, _, err := s.repo.Lock(tx, m.KodeBrng, m.KdBangsal, m.NoBatch, m.NoFaktur)
	if err != nil {
		return err
	}
	akhir := awal + delta
	if delta < 0 && akhir < -1e-9 {
		return support.Conflict(fmt.Sprintf("stok %s di depo %s tidak mencukupi (sisa %s)", m.KodeBrng, m.KdBangsal, angka(awal)))
	}
	status := "Simpan"
	if m.Batal {
		status = "Hapus"
	}
	if err := s.repo.CatatRiwayat(tx, stokrepo.Riwayat{
		KodeBrng: m.KodeBrng, StokAwal: awal, Masuk: math.Max(delta, 0), Keluar: math.Max(-delta, 0), StokAkhir: akhir,
		Posisi: m.Posisi, Waktu: s.now(), Petugas: m.Petugas, KdBangsal: m.KdBangsal, Status: status,
		NoBatch: m.NoBatch, NoFaktur: m.NoFaktur, Keterangan: m.Keterangan,
	}); err != nil {
		return err
	}
	if err := s.repo.Upsert(tx, m.KodeBrng, m.KdBangsal, m.NoBatch, m.NoFaktur, delta); err != nil {
		return err
	}
	if s.batch && m.NoBatch != "" {
		return s.repo.UbahSisaBatch(tx, m.KodeBrng, m.NoBatch, m.NoFaktur, delta)
	}
	return nil
}

func angka(f float64) string {
	return fmt.Sprintf("%g", math.Round(f*100)/100)
}
