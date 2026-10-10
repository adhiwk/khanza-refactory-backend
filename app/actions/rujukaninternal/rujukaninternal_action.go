// Package rujukaninternal use case rujukan antar poli dalam satu kunjungan (DlgRujukanPoliInternal).
package rujukaninternal

import (
	"strings"

	model "goravel/app/models/rujukaninternal"
	repo "goravel/app/repository/rujukaninternal"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrDokterNotFound     = support.NotFound("data dokter tidak ditemukan")
	ErrPoliNotFound       = support.NotFound("data poliklinik tidak ditemukan")
	ErrNotFound           = support.NotFound("data rujukan internal tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("rujukan ke dokter tersebut sudah ada")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(noRawat string) ([]model.RujukanInternal, error) {
	return a.repo.List(strings.TrimSpace(noRawat))
}

func (a *Action) Create(noRawat, kdDokter, kdPoli string) ([]model.RujukanInternal, error) {
	d := model.RujukanInternal{NoRawat: strings.TrimSpace(noRawat), KdDokter: strings.TrimSpace(kdDokter), KdPoli: strings.TrimSpace(kdPoli)}
	checks := []struct {
		fn  func(string) (bool, error)
		val string
		err error
	}{
		{a.repo.RegistrasiExists, d.NoRawat, ErrRegistrasiNotFound},
		{a.repo.DokterExists, d.KdDokter, ErrDokterNotFound},
		{a.repo.PoliExists, d.KdPoli, ErrPoliNotFound},
	}
	for _, c := range checks {
		ok, err := c.fn(c.val)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, c.err
		}
	}
	exists, err := a.repo.Exists(d.NoRawat, d.KdDokter)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	if err := a.repo.Create(d); err != nil {
		return nil, err
	}
	return a.repo.List(d.NoRawat)
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
