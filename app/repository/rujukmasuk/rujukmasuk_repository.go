package rujukmasuk

import (
	"fmt"

	model "goravel/app/models/rujukmasuk"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.RujukMasuk, int64, error)
	Find(key string) (*model.RujukMasuk, error)
	Exists(key string) (bool, error)
	Create(data *model.RujukMasuk) error
	Save(data *model.RujukMasuk) error
	Delete(key string) error
	PenyakitExists(value any) (bool, error)
	RegPeriksaExists(value any) (bool, error)
	NextNoBalasan(noRawat string) (string, error)
}

type repository struct {
	crud.Table[model.RujukMasuk, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.RujukMasuk, string]{
		Key:    "no_rawat",
		Search: []string{"no_rawat", "perujuk", "no_rujuk", "dokter_perujuk", "no_balasan"},
		Active: "",
		Order:  "no_rawat desc",
	}}
}

func (r *repository) PenyakitExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "penyakit", "kd_penyakit", value)
}

func (r *repository) RegPeriksaExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "reg_periksa", "no_rawat", value)
}

// NextNoBalasan BR/yyyy/MM/dd/NNNN berurutan per tanggal registrasi (DlgRujukMasuk / DlgIGD).
func (r *repository) NextNoBalasan(noRawat string) (string, error) {
	var list []struct {
		Tgl string `gorm:"column:tgl"`
		Max int    `gorm:"column:max"`
	}
	err := r.Query().Raw("select date_format(reg.tgl_registrasi,'%Y/%m/%d') as tgl,"+
		"(select ifnull(max(convert(right(rujuk_masuk.no_balasan,4),signed)),0) from reg_periksa "+
		"inner join rujuk_masuk on reg_periksa.no_rawat=rujuk_masuk.no_rawat where reg_periksa.tgl_registrasi=reg.tgl_registrasi) as max "+
		"from reg_periksa reg where reg.no_rawat=?", noRawat).Scan(&list)
	if err != nil || len(list) == 0 {
		return "", err
	}
	return fmt.Sprintf("BR/%s/%04d", list[0].Tgl, list[0].Max+1), nil
}
