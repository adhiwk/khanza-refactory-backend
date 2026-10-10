// Package tindakan use case tindakan dokter/paramedis rawat jalan (DlgRawatJalan) dan rawat inap (DlgRawatInap),
// termasuk posting jurnal pendapatan & jasa medis.
package tindakan

import (
	"fmt"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/perawatan"
	akunrepo "goravel/app/repository/akun"
	repo "goravel/app/repository/tindakan"
	jurnalsvc "goravel/app/services/jurnal"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrNotFound           = support.NotFound("data tindakan tidak ditemukan")
	ErrDokterNotFound     = support.NotFound("data dokter tidak ditemukan")
	ErrPetugasNotFound    = support.NotFound("data petugas tidak ditemukan")
	ErrTerkunci           = support.Conflict("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
	ErrSebelumRegistrasi  = support.Invalid("waktu tindakan tidak boleh sebelum waktu registrasi")
	ErrKadaluarsa         = support.Forbidden("perubahan data / penghapusan data tidak boleh lebih dari 2 x 24 jam")
	ErrPelaksana          = support.Invalid("pelaksana harus dr, pr, atau drpr")
	ErrTanpaTindakan      = support.Invalid("pilih minimal satu tindakan")
)

const batasUbah = 48 * time.Hour

// Input satu kali simpan beberapa tindakan pada waktu yang sama (BtnSimpan DlgRawatJalan).
type Input struct {
	NoRawat    string
	Pelaksana  model.Pelaksana
	KdDokter   string
	Nip        string
	Tgl        string
	Jam        string
	KdJenisPrw []string
	Operator   string
}

type Action struct {
	rawat  model.Rawat
	repo   repo.Repository
	akun   akunrepo.Repository
	jurnal *jurnalsvc.Service
	now    func() time.Time
}

func NewAction(rawat model.Rawat, repo repo.Repository, akun akunrepo.Repository, jurnal *jurnalsvc.Service) *Action {
	return &Action{rawat: rawat, repo: repo, akun: akun, jurnal: jurnal, now: time.Now}
}

func (a *Action) konteks(noRawat string) (*model.Konteks, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(noRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	return k, nil
}

func (a *Action) List(noRawat string) ([]model.Tindakan, error) {
	k, err := a.konteks(noRawat)
	if err != nil {
		return nil, err
	}
	return a.repo.List(a.rawat, k.NoRawat)
}

// Tarif daftar tindakan yang boleh dipilih untuk registrasi (filter poli/bangsal, cara bayar, kelas).
func (a *Action) Tarif(noRawat string, pelaksana model.Pelaksana, search string, page, limit int) ([]model.Tarif, int64, error) {
	if !validPelaksana(pelaksana) {
		return nil, 0, ErrPelaksana
	}
	k, err := a.konteks(noRawat)
	if err != nil {
		return nil, 0, err
	}
	return a.repo.CariTarif(a.rawat, pelaksana, *k, search, page, limit)
}

// Create menyimpan tindakan + jurnal dalam satu transaksi; tarif selalu diambil dari master tarif.
func (a *Action) Create(in Input) ([]model.Tindakan, error) {
	if !validPelaksana(in.Pelaksana) {
		return nil, ErrPelaksana
	}
	if len(in.KdJenisPrw) == 0 {
		return nil, ErrTanpaTindakan
	}
	k, err := a.konteks(in.NoRawat)
	if err != nil {
		return nil, err
	}
	if k.Terkunci() {
		return nil, ErrTerkunci
	}

	tgl, err := support.DateOrToday(in.Tgl)
	if err != nil {
		return nil, err
	}
	jam, err := support.TimeOrNow(in.Jam)
	if err != nil {
		return nil, err
	}
	waktu, _ := time.ParseInLocation(support.DateTimeLayout, tgl.Format(support.DateLayout)+" "+jam, time.Local)
	if k.SebelumRegistrasi(waktu) {
		return nil, ErrSebelumRegistrasi
	}

	if err := a.cekPelaksana(in); err != nil {
		return nil, err
	}

	rows := make([]model.Tindakan, 0, len(in.KdJenisPrw))
	seen := map[string]bool{}
	for _, kode := range in.KdJenisPrw {
		kode = strings.TrimSpace(kode)
		if kode == "" || seen[kode] {
			continue
		}
		seen[kode] = true
		tarif, err := a.repo.Tarif(a.rawat, in.Pelaksana, *k, kode)
		if err != nil {
			return nil, err
		}
		if tarif == nil {
			return nil, support.NotFound(fmt.Sprintf("tindakan %s tidak tersedia untuk registrasi ini", kode))
		}
		rows = append(rows, baris(k.NoRawat, in, tgl.Format(support.DateLayout), jam, *tarif))
	}
	if len(rows) == 0 {
		return nil, ErrTanpaTindakan
	}

	err = support.Transaction(func(tx contractsorm.Query) error {
		for _, t := range rows {
			if err := a.repo.Insert(tx, a.rawat, t); err != nil {
				return err
			}
		}
		ket := fmt.Sprintf("TINDAKAN %s PASIEN %s %s, DIPOSTING OLEH %s", a.label(), k.NoRkmMedis, k.NmPasien, in.Operator)
		return a.postJurnal(tx, k.NoRawat, ket, rows, false)
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Delete membatalkan satu tindakan beserta jurnal balik; non super admin dibatasi 2 x 24 jam.
func (a *Action) Delete(key model.TindakanKey, operator string, superAdmin bool) error {
	if !validPelaksana(key.Pelaksana) {
		return ErrPelaksana
	}
	k, err := a.konteks(key.NoRawat)
	if err != nil {
		return err
	}
	if k.Terkunci() {
		return ErrTerkunci
	}
	key.NoRawat = k.NoRawat

	return support.Transaction(func(tx contractsorm.Query) error {
		t, err := a.repo.Find(tx, a.rawat, key)
		if err != nil {
			return err
		}
		if t == nil {
			return ErrNotFound
		}
		if !superAdmin {
			waktu, err := time.ParseInLocation(support.DateTimeLayout, t.TglPerawatan+" "+t.JamRawat, time.Local)
			if err == nil && a.now().Sub(waktu) > batasUbah {
				return ErrKadaluarsa
			}
		}
		if err := a.repo.Delete(tx, a.rawat, key); err != nil {
			return err
		}
		ket := fmt.Sprintf("PEMBATALAN TINDAKAN %s PASIEN %s %s OLEH %s", a.label(), k.NoRkmMedis, k.NmPasien, operator)
		return a.postJurnal(tx, k.NoRawat, ket, []model.Tindakan{*t}, true)
	})
}

func (a *Action) cekPelaksana(in Input) error {
	if in.Pelaksana.PakaiDokter() {
		ok, err := a.repo.DokterExists(strings.TrimSpace(in.KdDokter))
		if err != nil {
			return err
		}
		if !ok {
			return ErrDokterNotFound
		}
	}
	if in.Pelaksana.PakaiParamedis() {
		ok, err := a.repo.PetugasExists(strings.TrimSpace(in.Nip))
		if err != nil {
			return err
		}
		if !ok {
			return ErrPetugasNotFound
		}
	}
	return nil
}

func (a *Action) postJurnal(tx contractsorm.Query, noRawat, ket string, rows []model.Tindakan, batal bool) error {
	akun, err := a.akun.Tindakan(tx, a.rawat)
	if err != nil {
		return err
	}
	e := Jurnal(*akun, rows)
	if batal {
		e.Reverse()
	}
	return a.jurnal.Post(tx, noRawat, jurnalsvc.JenisUmum, ket, e)
}

func (a *Action) label() string {
	if a.rawat == model.Ranap {
		return "RAWAT INAP"
	}
	return "RAWAT JALAN"
}

// Jurnal baris jurnal tindakan: pendapatan, jasa medis dokter/paramedis, KSO, jasa sarana, BHP, menejemen.
func Jurnal(akun akunrepo.Tindakan, rows []model.Tindakan) *jurnalsvc.Entries {
	var pendapatan, jmdr, jmpr, kso, sarana, bhp, menejemen float64
	for _, t := range rows {
		pendapatan += t.BiayaRawat
		jmdr += t.TarifTindakandr
		jmpr += t.TarifTindakanpr
		kso += t.Kso
		sarana += t.Material
		bhp += t.Bhp
		menejemen += t.Menejemen
	}
	e := &jurnalsvc.Entries{}
	e.Pair(akun.SuspenPiutang, akun.Pendapatan, pendapatan)
	e.Pair(akun.BebanJMDokter, akun.UtangJMDokter, jmdr)
	e.Pair(akun.BebanJMParamedis, akun.UtangJMParamedis, jmpr)
	e.Pair(akun.BebanKSO, akun.UtangKSO, kso)
	e.Pair(akun.BebanJasaSarana, akun.UtangJasaSarana, sarana)
	e.Pair(akun.HPPBHP, akun.PersediaanBHP, bhp)
	e.Pair(akun.BebanMenejemen, akun.UtangMenejemen, menejemen)
	return e
}

// baris menyusun tarif per pelaksana: dr memakai total_byrdr, pr total_byrpr, drpr total_byrdrpr.
func baris(noRawat string, in Input, tgl, jam string, t model.Tarif) model.Tindakan {
	row := model.Tindakan{
		Pelaksana:    in.Pelaksana,
		NoRawat:      noRawat,
		KdJenisPrw:   t.KdJenisPrw,
		NmPerawatan:  t.NmPerawatan,
		TglPerawatan: tgl,
		JamRawat:     jam,
		Material:     t.Material,
		Bhp:          t.Bhp,
		Kso:          t.Kso,
		Menejemen:    t.Menejemen,
	}
	switch in.Pelaksana {
	case model.Dokter:
		row.KdDokter = strings.TrimSpace(in.KdDokter)
		row.TarifTindakandr = t.TarifTindakandr
		row.BiayaRawat = t.TotalByrdr
	case model.Paramedis:
		row.Nip = strings.TrimSpace(in.Nip)
		row.TarifTindakanpr = t.TarifTindakanpr
		row.BiayaRawat = t.TotalByrpr
	case model.DokterParamedis:
		row.KdDokter = strings.TrimSpace(in.KdDokter)
		row.Nip = strings.TrimSpace(in.Nip)
		row.TarifTindakandr = t.TarifTindakandr
		row.TarifTindakanpr = t.TarifTindakanpr
		row.BiayaRawat = t.TotalByrdrpr
	}
	return row
}

func validPelaksana(p model.Pelaksana) bool {
	return p == model.Dokter || p == model.Paramedis || p == model.DokterParamedis
}
