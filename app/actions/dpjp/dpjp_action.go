// Package dpjp use case dokter penanggung jawab pelayanan rawat inap (DlgDpjp).
package dpjp

import (
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/dpjp"
	repo "goravel/app/repository/dpjp"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrDokterNotFound     = support.NotFound("data dokter tidak ditemukan")
	ErrNotFound           = support.NotFound("data DPJP tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("dokter sudah menjadi DPJP pada rawat ini")
	ErrTanpaDokter        = support.Invalid("pilih minimal satu dokter")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(noRawat string) ([]model.Dpjp, error) {
	return a.repo.List(strings.TrimSpace(noRawat))
}

// Create menambahkan satu/lebih dokter DPJP sekaligus (semua atau tidak sama sekali).
func (a *Action) Create(noRawat string, kdDokter []string) ([]model.Dpjp, error) {
	noRawat = strings.TrimSpace(noRawat)
	ok, err := a.repo.RegistrasiExists(noRawat)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrRegistrasiNotFound
	}
	var dokter []string
	seen := map[string]bool{}
	for _, kd := range kdDokter {
		kd = strings.TrimSpace(kd)
		if kd == "" || seen[kd] {
			continue
		}
		seen[kd] = true
		if ok, err := a.repo.DokterExists(kd); err != nil {
			return nil, err
		} else if !ok {
			return nil, ErrDokterNotFound
		}
		if ok, err := a.repo.Exists(noRawat, kd); err != nil {
			return nil, err
		} else if ok {
			return nil, ErrAlreadyExists
		}
		dokter = append(dokter, kd)
	}
	if len(dokter) == 0 {
		return nil, ErrTanpaDokter
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		for _, kd := range dokter {
			if err := a.repo.Create(tx, noRawat, kd); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return a.repo.List(noRawat)
}

func (a *Action) Delete(noRawat, kdDokter string) error {
	noRawat, kdDokter = strings.TrimSpace(noRawat), strings.TrimSpace(kdDokter)
	ok, err := a.repo.Exists(noRawat, kdDokter)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotFound
	}
	return a.repo.Delete(noRawat, kdDokter)
}
