// Package deposit use case uang muka pasien (DlgDeposit) beserta jurnal kas/bank vs uang muka.
package deposit

import (
	"fmt"
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/deposit"
	repo "goravel/app/repository/deposit"
	jurnalsvc "goravel/app/services/jurnal"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrAkunBayarNotFound  = support.NotFound("akun bayar tidak ditemukan")
	ErrNotFound           = support.NotFound("data deposit tidak ditemukan")
	ErrTerkunci           = support.Conflict("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
	ErrAkunUangMuka       = support.Invalid("rekening Uang_Muka_Ranap (set_akun_ranap2) belum diatur")
)

type Input struct {
	NoRawat    string
	Tgl, Jam   string
	NamaBayar  string
	Besar      float64
	Nip        string
	Keterangan string
	Operator   string
}

type Action struct {
	repo   repo.Repository
	jurnal *jurnalsvc.Service
}

func NewAction(repo repo.Repository, jurnal *jurnalsvc.Service) *Action {
	return &Action{repo: repo, jurnal: jurnal}
}

func (a *Action) List(noRawat string) ([]model.Deposit, error) {
	return a.repo.List(strings.TrimSpace(noRawat))
}

func (a *Action) Create(in Input) (*model.Deposit, error) {
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
	akun, err := a.repo.AkunBayar(strings.TrimSpace(in.NamaBayar))
	if err != nil {
		return nil, err
	}
	if akun == nil {
		return nil, ErrAkunBayarNotFound
	}
	tgl, err := support.DateOrToday(in.Tgl)
	if err != nil {
		return nil, err
	}
	jam, err := support.TimeOrNow(in.Jam)
	if err != nil {
		return nil, err
	}

	d := model.Deposit{
		NoRawat: k.NoRawat, TglDeposit: tgl.Format(support.DateLayout) + " " + jam, NamaBayar: akun.NamaBayar,
		Besarppn: akun.Ppn * in.Besar / 100, BesarDeposit: in.Besar,
		Nip: support.OrDefault(in.Nip, "-"), Keterangan: strings.TrimSpace(in.Keterangan),
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		uangMuka, err := a.repo.AkunUangMuka(tx)
		if err != nil {
			return err
		}
		if uangMuka == "" {
			return ErrAkunUangMuka
		}
		no, err := a.repo.NextNoDeposit(tx, tgl)
		if err != nil {
			return err
		}
		d.NoDeposit = no
		if err := a.repo.Insert(tx, d); err != nil {
			return err
		}
		e := &jurnalsvc.Entries{}
		e.Pair(akun.KdRek, uangMuka, d.BesarDeposit)
		ket := fmt.Sprintf("DEPOSIT PASIEN %s %s %s, OLEH %s", k.NoRawat, k.NoRkmMedis, k.NmPasien, in.Operator)
		return a.jurnal.Post(tx, d.NoDeposit, jurnalsvc.JenisUmum, ket, e)
	})
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// Delete membatalkan deposit dan membalik jurnalnya.
func (a *Action) Delete(noDeposit, operator string) error {
	return support.Transaction(func(tx contractsorm.Query) error {
		d, err := a.repo.Find(tx, strings.TrimSpace(noDeposit))
		if err != nil {
			return err
		}
		if d == nil {
			return ErrNotFound
		}
		k, err := a.repo.Konteks(d.NoRawat)
		if err != nil {
			return err
		}
		if k != nil && k.Terkunci() {
			return ErrTerkunci
		}
		akun, err := a.repo.AkunBayar(d.NamaBayar)
		if err != nil {
			return err
		}
		if akun == nil {
			return ErrAkunBayarNotFound
		}
		uangMuka, err := a.repo.AkunUangMuka(tx)
		if err != nil {
			return err
		}
		if uangMuka == "" {
			return ErrAkunUangMuka
		}
		if err := a.repo.Delete(tx, d.NoDeposit); err != nil {
			return err
		}
		e := &jurnalsvc.Entries{}
		e.Pair(uangMuka, akun.KdRek, d.BesarDeposit)
		nama := ""
		if k != nil {
			nama = k.NoRkmMedis + " " + k.NmPasien
		}
		ket := fmt.Sprintf("PEMBATALAN DEPOSIT PASIEN %s %s, OLEH %s", d.NoRawat, nama, operator)
		return a.jurnal.Post(tx, d.NoDeposit, jurnalsvc.JenisUmum, ket, e)
	})
}
