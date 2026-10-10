package jurnal

import (
	"fmt"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"
)

// Repository persistence tabel jurnal & detailjurnal. Selalu dipanggil di dalam transaksi Action.
type Repository interface {
	NextNoJurnal(tx contractsorm.Query, tgl time.Time) (string, error)
	Insert(tx contractsorm.Query, noJurnal, noBukti string, waktu time.Time, jenis, keterangan string) error
	InsertDetail(tx contractsorm.Query, noJurnal, kdRek string, debet, kredit float64) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// NextNoJurnal format JRyyyyMMddNNNNNN per tanggal jurnal (Jurnal.simpanJurnal).
// Baris jurnal hari itu dikunci agar posting paralel tidak mendapat nomor yang sama.
func (r *repository) NextNoJurnal(tx contractsorm.Query, tgl time.Time) (string, error) {
	var list []struct {
		Max int `gorm:"column:max"`
	}
	err := tx.Raw("select ifnull(max(convert(right(no_jurnal,6),signed)),0) as max from jurnal where tgl_jurnal=? for update",
		tgl.Format("2006-01-02")).Scan(&list)
	if err != nil {
		return "", err
	}
	max := 0
	if len(list) > 0 {
		max = list[0].Max
	}
	return fmt.Sprintf("JR%s%06d", tgl.Format("20060102"), max+1), nil
}

func (r *repository) Insert(tx contractsorm.Query, noJurnal, noBukti string, waktu time.Time, jenis, keterangan string) error {
	_, err := tx.Exec("insert into jurnal (no_jurnal,no_bukti,tgl_jurnal,jam_jurnal,jenis,keterangan) values (?,?,?,?,?,?)",
		noJurnal, noBukti, waktu.Format("2006-01-02"), waktu.Format("15:04:05"), jenis, keterangan)
	return err
}

func (r *repository) InsertDetail(tx contractsorm.Query, noJurnal, kdRek string, debet, kredit float64) error {
	_, err := tx.Exec("insert into detailjurnal (no_jurnal,kd_rek,debet,kredit) values (?,?,?,?)", noJurnal, kdRek, debet, kredit)
	return err
}
