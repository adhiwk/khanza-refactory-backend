package biayalain

import (
	model "goravel/app/models/biayalain"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

// tabel nama tabel & kolom per jenis (whitelist).
var tabel = map[model.Jenis]struct{ table, nama, besar string }{
	model.Tambahan: {"tambahan_biaya", "nama_biaya", "besar_biaya"},
	model.Potongan: {"pengurangan_biaya", "nama_pengurangan", "besar_pengurangan"},
}

type Repository interface {
	perawatanrepo.KonteksRepository
	List(jenis model.Jenis, noRawat string) ([]model.BiayaLain, error)
	Exists(jenis model.Jenis, noRawat, nama string) (bool, error)
	Create(jenis model.Jenis, b model.BiayaLain) error
	Delete(jenis model.Jenis, noRawat, nama string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) List(jenis model.Jenis, noRawat string) ([]model.BiayaLain, error) {
	t := tabel[jenis]
	list := []model.BiayaLain{}
	err := support.DB().Raw("select no_rawat,"+t.nama+" as nama,ifnull("+t.besar+",0) as besar from "+t.table+" where no_rawat=?", noRawat).Scan(&list)
	return list, err
}

func (r *repository) Exists(jenis model.Jenis, noRawat, nama string) (bool, error) {
	t := tabel[jenis]
	return support.DB().Table(t.table).Where("no_rawat = ? and "+t.nama+" = ?", noRawat, nama).Exists()
}

func (r *repository) Create(jenis model.Jenis, b model.BiayaLain) error {
	t := tabel[jenis]
	_, err := support.DB().Exec("insert into "+t.table+" (no_rawat,"+t.nama+","+t.besar+") values (?,?,?)", b.NoRawat, b.Nama, b.Besar)
	return err
}

func (r *repository) Delete(jenis model.Jenis, noRawat, nama string) error {
	t := tabel[jenis]
	_, err := support.DB().Exec("delete from "+t.table+" where no_rawat=? and "+t.nama+"=?", noRawat, nama)
	return err
}
