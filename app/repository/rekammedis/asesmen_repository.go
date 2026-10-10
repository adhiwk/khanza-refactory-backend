// Package rekammedis persistence form asesmen rekam medis Khanza (penilaian, skrining, checklist, hasil pemeriksaan, catatan).
// Setiap form satu tabel dengan kunci no_rawat (+ waktu); beberapa form punya tabel detail kode masalah/rencana.
package rekammedis

import (
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	"goravel/app/repository/crud"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

// Key nilai primary key form (kolom → nilai teks).
type Key map[string]string

// Detail tabel detail berisi daftar kode per no_rawat, mis. (no_rawat, kode_masalah).
type Detail struct {
	Name      string // nama field di API, mis. "masalah"
	Table     string
	Column    string
	RefTable  string // master kode, mis. master_masalah_keperawatan
	RefColumn string
}

// Spec metadata tabel form (dibangkitkan dari sik.sql).
type Spec struct {
	Table   string
	Keys    []string
	Waktu   string // ekspresi SQL waktu asesmen untuk filter & urutan; kosong bila tidak ada
	Search  []string
	Details []Detail
}

type Filter struct {
	NoRawat    string
	NoRkmMedis string
	TglAwal    string
	TglAkhir   string
	Search     string
}

// Store kontrak persistence satu form bermodel M.
type Store[M any] interface {
	perawatanrepo.KonteksRepository
	Paginate(f Filter, page, limit int) ([]M, int64, error)
	Find(key Key) (*M, error)
	Exists(key Key) (bool, error)
	Create(tx contractsorm.Query, m *M) error
	Save(tx contractsorm.Query, m *M) error
	Delete(tx contractsorm.Query, key Key) error
	Details(noRawat string) (map[string][]string, error)
	ReplaceDetails(tx contractsorm.Query, noRawat string, details map[string][]string) error
	RefExists(table, column string, value any) (bool, error)
}

type repository[M any] struct {
	perawatanrepo.KonteksRepository
	spec Spec
}

func NewRepository[M any](spec Spec) Store[M] {
	return &repository[M]{KonteksRepository: perawatanrepo.NewKonteksRepository(), spec: spec}
}

func (r *repository[M]) whereKey(q contractsorm.Query, key Key) contractsorm.Query {
	for _, k := range r.spec.Keys {
		q = q.Where(k+" = ?", key[k])
	}
	return q
}

func (r *repository[M]) Paginate(f Filter, page, limit int) ([]M, int64, error) {
	q := support.DB().Model(new(M))
	if f.NoRawat != "" {
		q = q.Where("no_rawat = ?", f.NoRawat)
	}
	if f.NoRkmMedis != "" {
		q = q.Where("no_rawat in (select no_rawat from reg_periksa where no_rkm_medis = ?)", f.NoRkmMedis)
	}
	if r.spec.Waktu != "" && f.TglAwal != "" && f.TglAkhir != "" {
		q = q.Where("date("+r.spec.Waktu+") between ? and ?", f.TglAwal, f.TglAkhir)
	}
	q = crud.WhereLike(q, f.Search, r.spec.Search...)
	order := "no_rawat desc"
	if r.spec.Waktu != "" {
		order = r.spec.Waktu + " desc"
	}
	list := []M{}
	var total int64
	if err := q.Order(order).Paginate(page, limit, &list, &total); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// Find (nil, nil) bila tidak ada.
func (r *repository[M]) Find(key Key) (*M, error) {
	list := []M{}
	if err := r.whereKey(support.DB().Model(new(M)), key).Limit(1).Find(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return &list[0], nil
}

func (r *repository[M]) Exists(key Key) (bool, error) {
	return r.whereKey(support.DB().Model(new(M)), key).Exists()
}

func (r *repository[M]) Create(tx contractsorm.Query, m *M) error {
	return tx.Create(m)
}

func (r *repository[M]) Save(tx contractsorm.Query, m *M) error {
	return tx.Save(m)
}

func (r *repository[M]) Delete(tx contractsorm.Query, key Key) error {
	if len(r.spec.Details) > 0 {
		if err := r.ReplaceDetails(tx, key["no_rawat"], map[string][]string{}); err != nil {
			return err
		}
	}
	_, err := r.whereKey(tx.Model(new(M)), key).Delete(new(M))
	return err
}

func (r *repository[M]) Details(noRawat string) (map[string][]string, error) {
	out := map[string][]string{}
	for _, d := range r.spec.Details {
		codes := []string{}
		if err := support.DB().Table(d.Table).Where("no_rawat = ?", noRawat).Order(d.Column).Pluck(d.Column, &codes); err != nil {
			return nil, err
		}
		out[d.Name] = codes
	}
	return out, nil
}

// ReplaceDetails mengganti seluruh kode detail; nama detail yang tidak ada di map dikosongkan.
func (r *repository[M]) ReplaceDetails(tx contractsorm.Query, noRawat string, details map[string][]string) error {
	for _, d := range r.spec.Details {
		if _, err := tx.Exec("delete from "+d.Table+" where no_rawat=?", noRawat); err != nil {
			return err
		}
		codes := details[d.Name]
		if len(codes) == 0 {
			continue
		}
		args := make([]any, 0, len(codes)*2)
		for _, c := range codes {
			args = append(args, noRawat, c)
		}
		sql := "insert into " + d.Table + " (no_rawat," + d.Column + ") values " + strings.TrimSuffix(strings.Repeat("(?,?),", len(codes)), ",")
		if _, err := tx.Exec(sql, args...); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository[M]) RefExists(table, column string, value any) (bool, error) {
	return crud.ExistsIn(support.DB(), table, column, value)
}
