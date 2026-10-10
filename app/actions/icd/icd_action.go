// Package icd use case diagnosa & prosedur pasien (PanelDiagnosa Khanza).
package icd

import (
	"fmt"
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/icd"
	repo "goravel/app/repository/icd"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrTanpaKode          = support.Invalid("pilih minimal satu kode")
	ErrFilter             = support.Invalid("isi no_rawat atau no_rkm_medis")
	ErrStatus             = support.Invalid("status harus Ralan atau Ranap")
)

type Action struct {
	jenis model.Jenis
	repo  repo.Repository
}

func NewAction(jenis model.Jenis, repo repo.Repository) *Action {
	return &Action{jenis: jenis, repo: repo}
}

func (a *Action) label() string {
	return string(a.jenis)
}

// List per no_rawat, atau seluruh riwayat pasien bila no_rkm_medis diisi.
func (a *Action) List(noRawat, noRkmMedis string) ([]model.Kode, error) {
	noRawat, noRkmMedis = strings.TrimSpace(noRawat), strings.TrimSpace(noRkmMedis)
	if noRawat == "" && noRkmMedis == "" {
		return nil, ErrFilter
	}
	return a.repo.List(a.jenis, noRawat, noRkmMedis)
}

func (a *Action) CariReferensi(search string, page, limit int) ([]model.Referensi, int64, error) {
	return a.repo.CariReferensi(a.jenis, search, page, limit)
}

// Create menambah kode; status default mengikuti status lanjut registrasi, prioritas default berurutan.
// Diagnosa yang pernah diderita pasien pada kunjungan lain berstatus "Lama", selain itu "Baru".
func (a *Action) Create(noRawat, status string, items []model.Item) ([]model.Kode, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(noRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	status = support.OrDefault(status, k.StatusLanjut)
	if status != "Ralan" && status != "Ranap" {
		return nil, ErrStatus
	}

	var clean []model.Item
	seen := map[string]bool{}
	for _, it := range items {
		it.Kode = strings.TrimSpace(it.Kode)
		if it.Kode == "" || seen[it.Kode] {
			continue
		}
		seen[it.Kode] = true
		ok, err := a.repo.ReferensiExists(a.jenis, it.Kode)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, support.NotFound(fmt.Sprintf("kode %s %s tidak ditemukan", a.label(), it.Kode))
		}
		exists, err := a.repo.Exists(a.jenis, k.NoRawat, it.Kode)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, support.Conflict(fmt.Sprintf("kode %s %s sudah ada pada rawat ini", a.label(), it.Kode))
		}
		it.Jumlah = support.OrDefault(it.Jumlah, "1")
		clean = append(clean, it)
	}
	if len(clean) == 0 {
		return nil, ErrTanpaKode
	}

	err = support.Transaction(func(tx contractsorm.Query) error {
		next, err := a.repo.MaxPrioritas(tx, a.jenis, k.NoRawat, status)
		if err != nil {
			return err
		}
		for _, it := range clean {
			if it.Prioritas <= 0 {
				next++
				it.Prioritas = next
			}
			statusPenyakit := ""
			if a.jenis == model.Diagnosa {
				statusPenyakit = "Baru"
				lama, err := a.repo.PernahDiderita(k.NoRkmMedis, it.Kode)
				if err != nil {
					return err
				}
				if lama {
					statusPenyakit = "Lama"
				}
			}
			if err := a.repo.Insert(tx, a.jenis, k.NoRawat, status, it, statusPenyakit); err != nil {
				return err
			}
		}
		return a.repo.SyncResume(tx, a.jenis, k.NoRawat, status)
	})
	if err != nil {
		return nil, err
	}
	return a.repo.List(a.jenis, k.NoRawat, "")
}

func (a *Action) Delete(noRawat, kode string) error {
	noRawat, kode = strings.TrimSpace(noRawat), strings.TrimSpace(kode)
	exists, err := a.repo.Exists(a.jenis, noRawat, kode)
	if err != nil {
		return err
	}
	if !exists {
		return support.NotFound("data " + a.label() + " tidak ditemukan")
	}
	return support.Transaction(func(tx contractsorm.Query) error {
		if err := a.repo.Delete(tx, a.jenis, noRawat, kode); err != nil {
			return err
		}
		for _, status := range []string{"Ralan", "Ranap"} {
			if err := a.repo.SyncResume(tx, a.jenis, noRawat, status); err != nil {
				return err
			}
		}
		return nil
	})
}

// SetStatusPenyakit menandai diagnosa Lama/Baru (menu Status Penyakit di PanelDiagnosa).
func (a *Action) SetStatusPenyakit(noRawat, kode, status string) error {
	exists, err := a.repo.Exists(model.Diagnosa, strings.TrimSpace(noRawat), strings.TrimSpace(kode))
	if err != nil {
		return err
	}
	if !exists {
		return support.NotFound("data diagnosa tidak ditemukan")
	}
	return a.repo.SetStatusPenyakit(strings.TrimSpace(noRawat), strings.TrimSpace(kode), status)
}
