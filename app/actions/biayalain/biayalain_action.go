// Package biayalain use case tambahan & potongan biaya pada tagihan (DlgDetailTambahan / DlgDetailPotongan).
package biayalain

import (
	"strings"

	model "goravel/app/models/biayalain"
	repo "goravel/app/repository/biayalain"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrNotFound           = support.NotFound("data biaya tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("nama biaya sudah ada pada rawat ini")
	ErrTerkunci           = support.Conflict("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
)

type Action struct {
	jenis model.Jenis
	repo  repo.Repository
}

func NewAction(jenis model.Jenis, repo repo.Repository) *Action {
	return &Action{jenis: jenis, repo: repo}
}

func (a *Action) List(noRawat string) ([]model.BiayaLain, error) {
	return a.repo.List(a.jenis, strings.TrimSpace(noRawat))
}

func (a *Action) Create(noRawat, nama string, besar float64) ([]model.BiayaLain, error) {
	b := model.BiayaLain{NoRawat: strings.TrimSpace(noRawat), Nama: strings.TrimSpace(nama), Besar: besar}
	if err := a.cekTerkunci(b.NoRawat); err != nil {
		return nil, err
	}
	exists, err := a.repo.Exists(a.jenis, b.NoRawat, b.Nama)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	if err := a.repo.Create(a.jenis, b); err != nil {
		return nil, err
	}
	return a.repo.List(a.jenis, b.NoRawat)
}

func (a *Action) Delete(noRawat, nama string) error {
	noRawat, nama = strings.TrimSpace(noRawat), strings.TrimSpace(nama)
	if err := a.cekTerkunci(noRawat); err != nil {
		return err
	}
	exists, err := a.repo.Exists(a.jenis, noRawat, nama)
	if err != nil {
		return err
	}
	if !exists {
		return ErrNotFound
	}
	return a.repo.Delete(a.jenis, noRawat, nama)
}

func (a *Action) cekTerkunci(noRawat string) error {
	k, err := a.repo.Konteks(noRawat)
	if err != nil {
		return err
	}
	if k == nil {
		return ErrRegistrasiNotFound
	}
	if k.Terkunci() {
		return ErrTerkunci
	}
	return nil
}
