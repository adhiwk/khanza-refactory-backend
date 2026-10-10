// Package riwayatpasien use case riwayat persalinan & imunisasi pasien (bagian dari RMPenilaianBayiBaruLahir,
// RMPenilaianAwalKeperawatanKebidanan*, RMPenilaianAwalKeperawatan*BayiAnak, RMPenilaianAwalMedisRanapNeonatus).
package riwayatpasien

import (
	"strings"

	model "goravel/app/models/riwayatpasien"
	repo "goravel/app/repository/riwayatpasien"
	"goravel/app/support"
)

var (
	ErrPasienNotFound     = support.NotFound("data pasien tidak ditemukan")
	ErrImunisasiNotFound  = support.NotFound("data imunisasi tidak ditemukan")
	ErrPersalinanNotFound = support.NotFound("data riwayat persalinan tidak ditemukan")
	ErrPersalinanAda      = support.Conflict("riwayat persalinan dengan tgl/thn tersebut sudah ada")
	ErrImunisasiAda       = support.Conflict("imunisasi ke- tersebut sudah tercatat")
	ErrImunisasiTidakAda  = support.NotFound("riwayat imunisasi tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) pasien(noRkmMedis string) (string, error) {
	noRkmMedis = strings.TrimSpace(noRkmMedis)
	ok, err := a.repo.PasienExists(noRkmMedis)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrPasienNotFound
	}
	return noRkmMedis, nil
}

func (a *Action) Persalinan(noRkmMedis string) ([]model.Persalinan, error) {
	no, err := a.pasien(noRkmMedis)
	if err != nil {
		return nil, err
	}
	return a.repo.Persalinan(no)
}

func (a *Action) TambahPersalinan(p model.Persalinan) ([]model.Persalinan, error) {
	no, err := a.pasien(p.NoRkmMedis)
	if err != nil {
		return nil, err
	}
	p.NoRkmMedis, p.TglThn = no, strings.TrimSpace(p.TglThn)
	p.Jk = support.OrDefault(p.Jk, "-")
	exists, err := a.repo.PersalinanExists(no, p.TglThn)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrPersalinanAda
	}
	if err := a.repo.CreatePersalinan(p); err != nil {
		return nil, err
	}
	return a.repo.Persalinan(no)
}

func (a *Action) HapusPersalinan(noRkmMedis, tglThn string) error {
	exists, err := a.repo.PersalinanExists(strings.TrimSpace(noRkmMedis), strings.TrimSpace(tglThn))
	if err != nil {
		return err
	}
	if !exists {
		return ErrPersalinanNotFound
	}
	return a.repo.DeletePersalinan(strings.TrimSpace(noRkmMedis), strings.TrimSpace(tglThn))
}

func (a *Action) Imunisasi(noRkmMedis string) ([]model.Imunisasi, error) {
	no, err := a.pasien(noRkmMedis)
	if err != nil {
		return nil, err
	}
	return a.repo.Imunisasi(no)
}

func (a *Action) TambahImunisasi(i model.Imunisasi) ([]model.Imunisasi, error) {
	no, err := a.pasien(i.NoRkmMedis)
	if err != nil {
		return nil, err
	}
	i.NoRkmMedis, i.KodeImunisasi = no, strings.TrimSpace(i.KodeImunisasi)
	if ok, err := a.repo.ImunisasiExists(i.KodeImunisasi); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrImunisasiNotFound
	}
	if exists, err := a.repo.ImunisasiPasienExists(no, i.KodeImunisasi, i.NoImunisasi); err != nil {
		return nil, err
	} else if exists {
		return nil, ErrImunisasiAda
	}
	if err := a.repo.CreateImunisasi(i); err != nil {
		return nil, err
	}
	return a.repo.Imunisasi(no)
}

func (a *Action) HapusImunisasi(noRkmMedis, kode string, ke int) error {
	noRkmMedis, kode = strings.TrimSpace(noRkmMedis), strings.TrimSpace(kode)
	exists, err := a.repo.ImunisasiPasienExists(noRkmMedis, kode, ke)
	if err != nil {
		return err
	}
	if !exists {
		return ErrImunisasiTidakAda
	}
	return a.repo.DeleteImunisasi(noRkmMedis, kode, ke)
}
