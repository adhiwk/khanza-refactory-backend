// Package jurnal posting jurnal akuntansi otomatis (pengganti tampjurnal + Jurnal.simpanJurnal Khanza).
package jurnal

import (
	"math"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	jurnalrepo "goravel/app/repository/jurnal"
	"goravel/app/support"
)

const (
	JenisUmum         = "U"
	JenisPenyesuaian  = "P"
	maxKeteranganJrnl = 350
)

var ErrTidakBalance = support.Invalid("jurnal tidak balance: total debet dan kredit berbeda")

// Entries kumpulan baris jurnal; baris dengan kd_rek sama dijumlahkan (perilaku tampjurnal).
type Entries struct {
	order []string
	rows  map[string]*row
}

type row struct {
	debet, kredit float64
}

func (e *Entries) add(kdRek string, debet, kredit float64) {
	if e.rows == nil {
		e.rows = map[string]*row{}
	}
	r, ok := e.rows[kdRek]
	if !ok {
		r = &row{}
		e.rows[kdRek] = r
		e.order = append(e.order, kdRek)
	}
	r.debet += debet
	r.kredit += kredit
}

// Pair mencatat debet pada akun debet dan kredit pada akun kredit sebesar nilai (diabaikan bila <= 0).
func (e *Entries) Pair(debetRek, kreditRek string, nilai float64) {
	if nilai <= 0 {
		return
	}
	e.add(debetRek, nilai, 0)
	e.add(kreditRek, 0, nilai)
}

// Reverse membalik seluruh baris (dipakai saat pembatalan transaksi).
func (e *Entries) Reverse() {
	for _, r := range e.rows {
		r.debet, r.kredit = r.kredit, r.debet
	}
}

func (e *Entries) Empty() bool {
	return len(e.order) == 0
}

// Totals total debet dan kredit.
func (e *Entries) Totals() (float64, float64) {
	var d, k float64
	for _, r := range e.rows {
		d += r.debet
		k += r.kredit
	}
	return d, k
}

type Service struct {
	repo jurnalrepo.Repository
	now  func() time.Time
}

func NewService(repo jurnalrepo.Repository) *Service {
	return &Service{repo: repo, now: time.Now}
}

// Post menyimpan jurnal + detail di dalam transaksi tx. Tidak ada baris = tidak ada jurnal.
func (s *Service) Post(tx contractsorm.Query, noBukti, jenis, keterangan string, e *Entries) error {
	if e.Empty() {
		return nil
	}
	debet, kredit := e.Totals()
	if math.Abs(debet-kredit) >= 0.01 {
		return ErrTidakBalance
	}
	if len(keterangan) > maxKeteranganJrnl {
		keterangan = keterangan[:maxKeteranganJrnl]
	}

	waktu := s.now()
	noJurnal, err := s.repo.NextNoJurnal(tx, waktu)
	if err != nil {
		return err
	}
	if err := s.repo.Insert(tx, noJurnal, noBukti, waktu, jenis, strings.TrimSpace(keterangan)); err != nil {
		return err
	}
	for _, kdRek := range e.order {
		r := e.rows[kdRek]
		if err := s.repo.InsertDetail(tx, noJurnal, kdRek, r.debet, r.kredit); err != nil {
			return err
		}
	}
	return nil
}
