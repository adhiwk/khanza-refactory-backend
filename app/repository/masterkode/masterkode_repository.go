package masterkode

import (
	"fmt"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/masterkode"
	"goravel/app/repository/crud"
	"goravel/app/support"
)

// Spec pemetaan kolom satu tabel master ke bentuk Kode.
type Spec struct {
	Slug      string
	Label     string
	Table     string
	KodeCol   string
	NamaCol   string
	NamaMax   int
	IndukCol  string // kosong = master tanpa induk
	Induk     *Spec  // master induk
	KodeWidth int    // panjang kode otomatis (Valid.autoNomer: 3 digit)
}

type Repository interface {
	Paginate(s *Spec, search, kodeInduk string, page, limit int) ([]model.Kode, int64, error)
	Find(s *Spec, kode string) (*model.Kode, error)
	Exists(s *Spec, kode string) (bool, error)
	NextKode(tx contractsorm.Query, s *Spec) (string, error)
	Create(tx contractsorm.Query, s *Spec, k model.Kode) error
	Update(s *Spec, k model.Kode) error
	Delete(s *Spec, kode string) error
	// Pemakai tabel yang masih mereferensikan kode (foreign key Khanza ON DELETE CASCADE akan ikut menghapusnya).
	Pemakai(s *Spec, kode string) ([]string, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) query(s *Spec) contractsorm.Query {
	cols := []string{"m." + s.KodeCol + " as kode", "ifnull(m." + s.NamaCol + ",'') as nama"}
	q := support.DB().Table(s.Table + " m")
	if s.IndukCol != "" {
		cols = append(cols, "m."+s.IndukCol+" as kode_induk", "ifnull(i."+s.Induk.NamaCol+",'') as nama_induk")
		q = q.Join("left join " + s.Induk.Table + " i on m." + s.IndukCol + "=i." + s.Induk.KodeCol)
	}
	return q.Select(cols...)
}

func (r *repository) Paginate(s *Spec, search, kodeInduk string, page, limit int) ([]model.Kode, int64, error) {
	q := crud.WhereLike(r.query(s), search, "m."+s.KodeCol, "m."+s.NamaCol)
	if s.IndukCol != "" && kodeInduk != "" {
		q = q.Where("m."+s.IndukCol+" = ?", kodeInduk)
	}
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Kode{}
	err = q.Order("m." + s.KodeCol).Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

func (r *repository) Find(s *Spec, kode string) (*model.Kode, error) {
	list := []model.Kode{}
	if err := r.query(s).Where("m."+s.KodeCol+" = ?", kode).Limit(1).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *repository) Exists(s *Spec, kode string) (bool, error) {
	return crud.ExistsIn(support.DB(), s.Table, s.KodeCol, kode)
}

// NextKode nomor berikutnya (angka terbesar + 1) dengan nol di depan; tabel dikunci selama transaksi.
func (r *repository) NextKode(tx contractsorm.Query, s *Spec) (string, error) {
	var list []struct {
		Max int `gorm:"column:max"`
	}
	if err := tx.Raw("select ifnull(max(cast(" + s.KodeCol + " as unsigned)),0) as max from " + s.Table + " for update").Scan(&list); err != nil {
		return "", err
	}
	max := 0
	if len(list) > 0 {
		max = list[0].Max
	}
	return fmt.Sprintf("%0*d", s.KodeWidth, max+1), nil
}

func (r *repository) Create(tx contractsorm.Query, s *Spec, k model.Kode) error {
	if s.IndukCol != "" {
		_, err := tx.Exec("insert into "+s.Table+" ("+s.IndukCol+","+s.KodeCol+","+s.NamaCol+") values (?,?,?)", k.KodeInduk, k.Kode, k.Nama)
		return err
	}
	_, err := tx.Exec("insert into "+s.Table+" ("+s.KodeCol+","+s.NamaCol+") values (?,?)", k.Kode, k.Nama)
	return err
}

func (r *repository) Update(s *Spec, k model.Kode) error {
	if s.IndukCol != "" {
		_, err := support.DB().Exec("update "+s.Table+" set "+s.NamaCol+"=?,"+s.IndukCol+"=? where "+s.KodeCol+"=?", k.Nama, k.KodeInduk, k.Kode)
		return err
	}
	_, err := support.DB().Exec("update "+s.Table+" set "+s.NamaCol+"=? where "+s.KodeCol+"=?", k.Nama, k.Kode)
	return err
}

func (r *repository) Delete(s *Spec, kode string) error {
	_, err := support.DB().Exec("delete from "+s.Table+" where "+s.KodeCol+"=?", kode)
	return err
}

func (r *repository) Pemakai(s *Spec, kode string) ([]string, error) {
	var refs []struct {
		Table  string `gorm:"column:table_name"`
		Column string `gorm:"column:column_name"`
	}
	err := support.DB().Raw("select table_name,column_name from information_schema.key_column_usage "+
		"where table_schema=database() and referenced_table_name=? and referenced_column_name=?", s.Table, s.KodeCol).Scan(&refs)
	if err != nil {
		return nil, err
	}
	out := []string{}
	for _, ref := range refs {
		// nama tabel & kolom berasal dari information_schema, bukan input pengguna.
		used, err := support.DB().Table(ref.Table).Where(ref.Column+" = ?", kode).Exists()
		if err != nil {
			return nil, err
		}
		if used {
			out = append(out, ref.Table)
		}
	}
	return out, nil
}
