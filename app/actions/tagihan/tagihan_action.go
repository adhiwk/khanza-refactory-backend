// Package tagihan use case rincian tagihan pasien (perhitungan yang ditampilkan DlgBilingRalan / DlgBilingRanap sebelum nota disimpan).
package tagihan

import (
	"math"
	"strings"

	model "goravel/app/models/tagihan"
	repo "goravel/app/repository/tagihan"
	"goravel/app/support"
)

var ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) Hitung(noRawat string) (*model.Tagihan, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(noRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	t := &model.Tagihan{NoRawat: k.NoRawat, NoRkmMedis: k.NoRkmMedis, NmPasien: k.NmPasien, StatusLanjut: k.StatusLanjut, Rincian: []model.Kategori{}}
	if t.StatusBayar, err = a.repo.StatusBayar(k.NoRawat); err != nil {
		return nil, err
	}
	for _, kat := range repo.Kategori {
		items, err := a.repo.Items(kat.SQL, k.NoRawat)
		if err != nil {
			return nil, err
		}
		if len(items) == 0 {
			continue
		}
		g := model.Kategori{Kategori: kat.Nama, Items: items}
		for i := range items {
			items[i].Kategori = kat.Nama
			g.Total += items[i].Total
		}
		t.TotalBiaya += g.Total
		t.Rincian = append(t.Rincian, g)
	}
	potongan, err := a.repo.Potongan(k.NoRawat)
	if err != nil {
		return nil, err
	}
	if len(potongan) > 0 {
		g := model.Kategori{Kategori: "Potongan", Items: potongan}
		for i := range potongan {
			potongan[i].Kategori = "Potongan"
			g.Total += potongan[i].Total
		}
		t.Potongan = g.Total
		t.Rincian = append(t.Rincian, g)
	}
	if t.Deposit, err = a.repo.Deposit(k.NoRawat); err != nil {
		return nil, err
	}
	t.TotalBiaya = math.Round(t.TotalBiaya)
	t.SisaTagihan = math.Round(t.TotalBiaya - t.Potongan - t.Deposit)
	return t, nil
}
