// Package resep use case peresepan elektronik dokter (DlgPeresepanDokter, obat non-racikan).
package resep

import (
	"fmt"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/resep"
	repo "goravel/app/repository/resep"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrDokterNotFound     = support.NotFound("data dokter tidak ditemukan")
	ErrNotFound           = support.NotFound("data resep tidak ditemukan")
	ErrTanpaObat          = support.Invalid("masukkan minimal satu obat")
	ErrTerkunci           = support.Conflict("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
	ErrDilayani           = support.Conflict("resep sudah dilayani farmasi, tidak boleh diubah/dihapus")
	ErrSebelumRegistrasi  = support.Invalid("waktu peresepan tidak boleh sebelum waktu registrasi")
)

type Input struct {
	NoRawat  string
	KdDokter string
	Tgl, Jam string
	Items    []model.Item
}

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(noRawat string) ([]model.Resep, error) {
	return a.repo.List(strings.TrimSpace(noRawat))
}

func (a *Action) Detail(noResep string) (*model.Resep, error) {
	r, err := a.repo.Find(support.DB(), strings.TrimSpace(noResep))
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, ErrNotFound
	}
	return r, nil
}

// Create nomor resep dibuat otomatis; status ralan/ranap mengikuti status lanjut registrasi.
func (a *Action) Create(in Input) (*model.Resep, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(in.NoRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	if k.Terkunci() {
		return nil, ErrTerkunci
	}
	kdDokter := strings.TrimSpace(in.KdDokter)
	if ok, err := a.repo.DokterExists(kdDokter); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrDokterNotFound
	}
	tgl, err := support.DateOrToday(in.Tgl)
	if err != nil {
		return nil, err
	}
	jam, err := support.TimeOrNow(in.Jam)
	if err != nil {
		return nil, err
	}
	waktu, _ := time.ParseInLocation(support.DateTimeLayout, tgl.Format(support.DateLayout)+" "+jam, time.Local)
	if k.SebelumRegistrasi(waktu) {
		return nil, ErrSebelumRegistrasi
	}
	items, err := a.items(in.Items)
	if err != nil {
		return nil, err
	}

	r := model.Resep{NoRawat: k.NoRawat, KdDokter: kdDokter, TglPeresepan: tgl.Format(support.DateLayout), JamPeresepan: jam,
		Status: strings.ToLower(k.StatusLanjut), Items: items}
	err = support.Transaction(func(tx contractsorm.Query) error {
		no, err := a.repo.NextNoResep(tx, tgl)
		if err != nil {
			return err
		}
		r.NoResep = no
		return a.repo.Insert(tx, r)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(r.NoResep)
}

// Update mengganti daftar obat resep yang belum dilayani.
func (a *Action) Update(noResep string, items []model.Item) (*model.Resep, error) {
	list, err := a.items(items)
	if err != nil {
		return nil, err
	}
	err = a.editable(noResep, func(tx contractsorm.Query, r *model.Resep) error {
		return a.repo.ReplaceItems(tx, r.NoResep, list)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(noResep)
}

func (a *Action) Delete(noResep string) error {
	return a.editable(noResep, func(tx contractsorm.Query, r *model.Resep) error {
		return a.repo.Delete(tx, r.NoResep)
	})
}

func (a *Action) editable(noResep string, fn func(tx contractsorm.Query, r *model.Resep) error) error {
	return support.Transaction(func(tx contractsorm.Query) error {
		r, err := a.repo.Find(tx, strings.TrimSpace(noResep))
		if err != nil {
			return err
		}
		if r == nil {
			return ErrNotFound
		}
		if r.Dilayani() {
			return ErrDilayani
		}
		k, err := a.repo.Konteks(r.NoRawat)
		if err != nil {
			return err
		}
		if k != nil && k.Terkunci() {
			return ErrTerkunci
		}
		return fn(tx, r)
	})
}

// items memvalidasi obat aktif; kode yang sama digabung jumlahnya.
func (a *Action) items(in []model.Item) ([]model.Item, error) {
	idx := map[string]int{}
	var out []model.Item
	for _, it := range in {
		it.KodeBrng = strings.TrimSpace(it.KodeBrng)
		if it.KodeBrng == "" || it.Jml <= 0 {
			continue
		}
		if i, ok := idx[it.KodeBrng]; ok {
			out[i].Jml += it.Jml
			continue
		}
		ok, err := a.repo.BarangExists(it.KodeBrng)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, support.NotFound(fmt.Sprintf("obat %s tidak ditemukan", it.KodeBrng))
		}
		idx[it.KodeBrng] = len(out)
		out = append(out, model.Item{KodeBrng: it.KodeBrng, Jml: it.Jml, AturanPakai: strings.TrimSpace(it.AturanPakai)})
	}
	if len(out) == 0 {
		return nil, ErrTanpaObat
	}
	return out, nil
}
