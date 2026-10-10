// Package akun membaca pemetaan rekening akuntansi (set_akun_ralan / set_akun_ranap).
package akun

import (
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	"goravel/app/models/perawatan"
	"goravel/app/support"
)

// Tindakan rekening untuk posting jurnal tindakan rawat jalan / rawat inap.
type Tindakan struct {
	SuspenPiutang    string `gorm:"column:suspen_piutang"`
	Pendapatan       string `gorm:"column:pendapatan"`
	BebanJMDokter    string `gorm:"column:beban_jm_dokter"`
	UtangJMDokter    string `gorm:"column:utang_jm_dokter"`
	BebanJMParamedis string `gorm:"column:beban_jm_paramedis"`
	UtangJMParamedis string `gorm:"column:utang_jm_paramedis"`
	BebanKSO         string `gorm:"column:beban_kso"`
	UtangKSO         string `gorm:"column:utang_kso"`
	BebanJasaSarana  string `gorm:"column:beban_jasa_sarana"`
	UtangJasaSarana  string `gorm:"column:utang_jasa_sarana"`
	BebanMenejemen   string `gorm:"column:beban_menejemen"`
	UtangMenejemen   string `gorm:"column:utang_menejemen"`
	HPPBHP           string `gorm:"column:hpp_bhp"`
	PersediaanBHP    string `gorm:"column:persediaan_bhp"`
}

var ErrBelumDiatur = support.Invalid("pengaturan akun jurnal (set_akun_ralan / set_akun_ranap) belum diisi")

type Repository interface {
	Tindakan(tx contractsorm.Query, rawat perawatan.Rawat) (*Tindakan, error)
	Obat(tx contractsorm.Query, rawat perawatan.Rawat) (*Obat, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// Tindakan kolom set_akun_ralan / set_akun_ranap mengikuti akuntindakanralan / akuntindakanranap.
func (r *repository) Tindakan(tx contractsorm.Query, rawat perawatan.Rawat) (*Tindakan, error) {
	s := string(rawat)
	table := "set_akun_" + strings.ToLower(s)
	sql := "select " +
		"Suspen_Piutang_Tindakan_" + s + " as suspen_piutang," +
		"Tindakan_" + s + " as pendapatan," +
		"Beban_Jasa_Medik_Dokter_Tindakan_" + s + " as beban_jm_dokter," +
		"Utang_Jasa_Medik_Dokter_Tindakan_" + s + " as utang_jm_dokter," +
		"Beban_Jasa_Medik_Paramedis_Tindakan_" + s + " as beban_jm_paramedis," +
		"Utang_Jasa_Medik_Paramedis_Tindakan_" + s + " as utang_jm_paramedis," +
		"Beban_KSO_Tindakan_" + s + " as beban_kso," +
		"Utang_KSO_Tindakan_" + s + " as utang_kso," +
		"Beban_Jasa_Sarana_Tindakan_" + s + " as beban_jasa_sarana," +
		"Utang_Jasa_Sarana_Tindakan_" + s + " as utang_jasa_sarana," +
		"Beban_Jasa_Menejemen_Tindakan_" + s + " as beban_menejemen," +
		"Utang_Jasa_Menejemen_Tindakan_" + s + " as utang_menejemen," +
		"HPP_BHP_Tindakan_" + s + " as hpp_bhp," +
		"Persediaan_BHP_Tindakan_" + s + " as persediaan_bhp from " + table + " limit 1"
	list := []Tindakan{}
	if err := tx.Raw(sql).Scan(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, ErrBelumDiatur
	}
	return &list[0], nil
}

// Obat rekening jurnal pemberian obat (akunobatralan / akunobatranap).
type Obat struct {
	SuspenPiutang string `gorm:"column:suspen_piutang"`
	Pendapatan    string `gorm:"column:pendapatan"`
	HPP           string `gorm:"column:hpp"`
	Persediaan    string `gorm:"column:persediaan"`
}

// Obat kolom set_akun_ralan / set_akun_ranap untuk obat.
func (r *repository) Obat(tx contractsorm.Query, rawat perawatan.Rawat) (*Obat, error) {
	sql := "select Suspen_Piutang_Obat_Ralan as suspen_piutang,Obat_Ralan as pendapatan,HPP_Obat_Rawat_Jalan as hpp," +
		"Persediaan_Obat_Rawat_Jalan as persediaan from set_akun_ralan limit 1"
	if rawat == perawatan.Ranap {
		sql = "select Suspen_Piutang_Obat_Ranap as suspen_piutang,Obat_Ranap as pendapatan,HPP_Obat_Rawat_Inap as hpp," +
			"Persediaan_Obat_Rawat_Inap as persediaan from set_akun_ranap limit 1"
	}
	list := []Obat{}
	if err := tx.Raw(sql).Scan(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, ErrBelumDiatur
	}
	return &list[0], nil
}
