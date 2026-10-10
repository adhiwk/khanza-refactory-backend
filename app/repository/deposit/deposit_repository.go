package deposit

import (
	"fmt"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/deposit"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

type Repository interface {
	perawatanrepo.KonteksRepository
	AkunBayar(namaBayar string) (*model.AkunBayar, error)
	AkunUangMuka(tx contractsorm.Query) (string, error)
	List(noRawat string) ([]model.Deposit, error)
	Find(tx contractsorm.Query, noDeposit string) (*model.Deposit, error)
	NextNoDeposit(tx contractsorm.Query, tgl time.Time) (string, error)
	Insert(tx contractsorm.Query, d model.Deposit) error
	Delete(tx contractsorm.Query, noDeposit string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) AkunBayar(namaBayar string) (*model.AkunBayar, error) {
	list := []model.AkunBayar{}
	err := support.DB().Raw("select nama_bayar,ifnull(kd_rek,'') as kd_rek,ifnull(ppn,0) as ppn from akun_bayar where nama_bayar=?", namaBayar).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

// AkunUangMuka rekening Uang_Muka_Ranap (set_akun_ranap2); kosong bila belum diatur.
func (r *repository) AkunUangMuka(tx contractsorm.Query) (string, error) {
	var list []struct {
		V string `gorm:"column:v"`
	}
	if err := tx.Raw("select ifnull(Uang_Muka_Ranap,'') as v from set_akun_ranap2 limit 1").Scan(&list); err != nil || len(list) == 0 {
		return "", err
	}
	return list[0].V, nil
}

const columns = "no_deposit,no_rawat,cast(tgl_deposit as char) as tgl_deposit,nama_bayar,besarppn,ifnull(besar_deposit,0) as besar_deposit,nip,keterangan"

func (r *repository) List(noRawat string) ([]model.Deposit, error) {
	list := []model.Deposit{}
	err := support.DB().Raw("select "+columns+" from deposit where no_rawat=? order by tgl_deposit", noRawat).Scan(&list)
	return list, err
}

func (r *repository) Find(tx contractsorm.Query, noDeposit string) (*model.Deposit, error) {
	list := []model.Deposit{}
	if err := tx.Raw("select "+columns+" from deposit where no_deposit=? for update", noDeposit).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

// NextNoDeposit DP + yyyyMMdd + 4 digit per tanggal deposit (DlgDeposit).
func (r *repository) NextNoDeposit(tx contractsorm.Query, tgl time.Time) (string, error) {
	var list []struct {
		Max int `gorm:"column:max"`
	}
	if err := tx.Raw("select ifnull(max(convert(right(no_deposit,4),signed)),0) as max from deposit where date(tgl_deposit)=? for update",
		tgl.Format("2006-01-02")).Scan(&list); err != nil {
		return "", err
	}
	max := 0
	if len(list) > 0 {
		max = list[0].Max
	}
	return fmt.Sprintf("DP%s%04d", tgl.Format("20060102"), max+1), nil
}

func (r *repository) Insert(tx contractsorm.Query, d model.Deposit) error {
	_, err := tx.Exec("insert into deposit (no_deposit,no_rawat,tgl_deposit,nama_bayar,besarppn,besar_deposit,nip,keterangan) values (?,?,?,?,?,?,?,?)",
		d.NoDeposit, d.NoRawat, d.TglDeposit, d.NamaBayar, d.Besarppn, d.BesarDeposit, d.Nip, d.Keterangan)
	return err
}

func (r *repository) Delete(tx contractsorm.Query, noDeposit string) error {
	_, err := tx.Exec("delete from deposit where no_deposit=?", noDeposit)
	return err
}
