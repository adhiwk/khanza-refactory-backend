package rujukkeluar

import (
	"fmt"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/rujukkeluar"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Rujuk, int64, error)
	Find(key string) (*model.Rujuk, error)
	Exists(key string) (bool, error)
	Create(data *model.Rujuk) error
	Save(data *model.Rujuk) error
	Delete(key string) error
	RegPeriksaExists(value any) (bool, error)
	DokterExists(value any) (bool, error)
	NextNoRujuk() (string, error)
	WithTx(tx contractsorm.Query) Repository
}

type repository struct {
	crud.Table[model.Rujuk, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Rujuk, string]{
		Key:    "no_rujuk",
		Search: []string{"no_rujuk", "no_rawat", "rujuk_ke"},
		Active: "",
		Order:  "tgl_rujuk desc, jam desc",
	}}
}

func (r *repository) RegPeriksaExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "reg_periksa", "no_rawat", value)
}

func (r *repository) DokterExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "dokter", "kd_dokter", value)
}

func (r *repository) WithTx(tx contractsorm.Query) Repository {
	return &repository{r.Table.WithTx(tx)}
}

// NextNoRujuk "R" + 9 digit (Valid.autoNomer("rujuk","R",9)); memakai nomor terbesar, bukan jumlah baris.
// Tabel rujuk tidak punya unique key, jadi panggil di dalam transaksi: baris dikunci for update.
func (r *repository) NextNoRujuk() (string, error) {
	var list []struct {
		Max int `gorm:"column:max"`
	}
	if err := r.Query().Raw("select ifnull(max(convert(substring(no_rujuk,2),signed)),0) as max from rujuk where no_rujuk like 'R%' for update").Scan(&list); err != nil {
		return "", err
	}
	max := 0
	if len(list) > 0 {
		max = list[0].Max
	}
	return fmt.Sprintf("R%09d", max+1), nil
}
