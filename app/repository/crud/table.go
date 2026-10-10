// Package crud menyediakan operasi persistence standar satu tabel Khanza yang di-embed
// oleh repository tiap modul. Aturan bisnis tetap berada di Action.
package crud

import (
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	"goravel/app/support"
)

// Table persistence satu tabel bermodel T dengan primary key bertipe K.
type Table[T any, K comparable] struct {
	Key    string   // kolom primary key
	Search []string // kolom untuk pencarian LIKE
	Active string   // kolom status aktif ('1'/'0'); kosong = hapus permanen
	Order  string   // urutan default list
	tx     contractsorm.Query
}

// WithTx salinan Table yang memakai transaksi tx.
func (t Table[T, K]) WithTx(tx contractsorm.Query) Table[T, K] {
	t.tx = tx
	return t
}

// Query query pada transaksi aktif atau koneksi Khanza.
func (t Table[T, K]) Query() contractsorm.Query {
	if t.tx != nil {
		return t.tx
	}
	return support.DB()
}

func (t Table[T, K]) active(q contractsorm.Query) contractsorm.Query {
	if t.Active != "" {
		return q.Where(t.Active+" = ?", "1")
	}
	return q
}

// Paginate list data aktif dengan pencarian LIKE pada kolom Search.
func (t Table[T, K]) Paginate(search string, page, limit int) ([]T, int64, error) {
	return t.PaginateWhere(search, nil, page, limit)
}

// PaginateWhere seperti Paginate ditambah filter kesamaan kolom (nilai kosong diabaikan).
// Filter eksplisit pada kolom Active menggantikan filter default "aktif saja".
func (t Table[T, K]) PaginateWhere(search string, eq map[string]string, page, limit int) ([]T, int64, error) {
	q := t.Query().Model(new(T))
	if _, ok := eq[t.Active]; !ok || eq[t.Active] == "" {
		q = t.active(q)
	}
	for col, v := range eq {
		if v != "" {
			q = q.Where(col+" = ?", v)
		}
	}
	q = WhereLike(q, search, t.Search...)
	if t.Order != "" {
		q = q.Order(t.Order)
	}
	list := []T{}
	var total int64
	if err := q.Paginate(page, limit, &list, &total); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// All seluruh data aktif (untuk pilihan/lookup kecil).
func (t Table[T, K]) All() ([]T, error) {
	q := t.active(t.Query().Model(new(T)))
	if t.Order != "" {
		q = q.Order(t.Order)
	}
	list := []T{}
	err := q.Get(&list)
	return list, err
}

// Find data aktif berdasarkan primary key; (nil, nil) bila tidak ada.
func (t Table[T, K]) Find(key K) (*T, error) {
	list := []T{}
	if err := t.active(t.Query().Where(t.Key+" = ?", key)).Limit(1).Find(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return &list[0], nil
}

// FindAny seperti Find tanpa memperhatikan status aktif; (nil, nil) bila tidak ada.
func (t Table[T, K]) FindAny(key K) (*T, error) {
	list := []T{}
	if err := t.Query().Where(t.Key+" = ?", key).Limit(1).Find(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return &list[0], nil
}

// SetActive mengubah kolom Active menjadi status ('1' aktif / '0' nonaktif).
func (t Table[T, K]) SetActive(key K, status string) error {
	_, err := t.Query().Model(new(T)).Where(t.Key+" = ?", key).Update(t.Active, status)
	return err
}

// Exists true bila primary key sudah dipakai, termasuk data yang dinonaktifkan.
func (t Table[T, K]) Exists(key K) (bool, error) {
	return t.Query().Model(new(T)).Where(t.Key+" = ?", key).Exists()
}

// ExistsActive true bila primary key ada dan aktif (dipakai untuk validasi referensi).
func (t Table[T, K]) ExistsActive(key K) (bool, error) {
	return t.active(t.Query().Model(new(T)).Where(t.Key+" = ?", key)).Exists()
}

func (t Table[T, K]) Create(m *T) error {
	return t.Query().Create(m)
}

func (t Table[T, K]) Save(m *T) error {
	return t.Query().Save(m)
}

// Delete menonaktifkan (Active diisi) atau menghapus permanen.
func (t Table[T, K]) Delete(key K) error {
	q := t.Query().Model(new(T)).Where(t.Key+" = ?", key)
	if t.Active != "" {
		_, err := q.Update(t.Active, "0")
		return err
	}
	_, err := q.Delete(new(T))
	return err
}

// WhereLike menambahkan (a like ? or b like ? ...) bila search tidak kosong.
func WhereLike(q contractsorm.Query, search string, columns ...string) contractsorm.Query {
	search = strings.TrimSpace(search)
	if search == "" || len(columns) == 0 {
		return q
	}
	like := "%" + search + "%"
	conds := make([]string, len(columns))
	args := make([]any, len(columns))
	for i, c := range columns {
		conds[i] = c + " like ?"
		args[i] = like
	}
	return q.Where("("+strings.Join(conds, " or ")+")", args...)
}

// ExistsIn true bila value ada pada table.column (cek referensi master lain).
func ExistsIn(q contractsorm.Query, table, column string, value any) (bool, error) {
	return q.Table(table).Where(column+" = ?", value).Exists()
}
