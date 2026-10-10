// Package pemberianobat use case pemberian obat pasien rawat jalan (DlgCariObat) dan rawat inap (DlgCariObat2):
// harga dihitung server, stok depo dipotong, jurnal pendapatan & HPP obat diposting.
package pemberianobat

import (
	"fmt"
	"math"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/pemberianobat"
	"goravel/app/models/perawatan"
	akunrepo "goravel/app/repository/akun"
	repo "goravel/app/repository/pemberianobat"
	stokrepo "goravel/app/repository/stok"
	jurnalsvc "goravel/app/services/jurnal"
	stoksvc "goravel/app/services/stok"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrDepoNotFound       = support.NotFound("depo/lokasi obat tidak ditemukan, atur set_lokasi atau kirim kd_bangsal")
	ErrNotFound           = support.NotFound("data pemberian obat tidak ditemukan")
	ErrTerkunci           = support.Conflict("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
	ErrTanpaObat          = support.Invalid("masukkan minimal satu obat")
	ErrSebelumRegistrasi  = support.Invalid("waktu pemberian obat tidak boleh sebelum waktu registrasi")
	ErrJenisHarga         = support.Invalid("jenis harga tidak valid")
	ErrBatchWajib         = support.Invalid("no_batch & no_faktur wajib diisi saat stok per batch aktif")
)

// JenisHarga kolom harga yang boleh dipilih per jenis rawat.
var JenisHarga = map[perawatan.Rawat][]string{
	perawatan.Ralan: {"ralan", "karyawan", "beliluar", "utama"},
	perawatan.Ranap: {"kelas1", "kelas2", "kelas3", "utama", "vip", "vvip", "beliluar", "karyawan"},
}

// kelasHarga kolom harga default menurut kelas kamar inap aktif.
var kelasHarga = map[string]string{
	"Kelas 1": "kelas1", "Kelas 2": "kelas2", "Kelas 3": "kelas3", "Kelas Utama": "utama", "Kelas VIP": "vip", "Kelas VVIP": "vvip",
}

// Setting pengaturan farmasi dari config.
type Setting struct {
	Batch           bool
	Hpp             string
	PembulatanHarga bool
}

type Item struct {
	KodeBrng    string
	Jml         float64
	Embalase    float64
	Tuslah      float64
	AturanPakai string
	NoBatch     string
	NoFaktur    string
}

type Input struct {
	NoRawat    string
	KdBangsal  string
	Tgl, Jam   string
	JenisHarga string
	Items      []Item
	Operator   string
}

type Action struct {
	rawat  perawatan.Rawat
	set    Setting
	repo   repo.Repository
	akun   akunrepo.Repository
	stok   *stoksvc.Service
	jurnal *jurnalsvc.Service
}

func NewAction(rawat perawatan.Rawat, set Setting, repo repo.Repository, akun akunrepo.Repository, stok *stoksvc.Service, jurnal *jurnalsvc.Service) *Action {
	return &Action{rawat: rawat, set: set, repo: repo, akun: akun, stok: stok, jurnal: jurnal}
}

func (a *Action) konteks(noRawat string) (*perawatan.Konteks, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(noRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	return k, nil
}

func (a *Action) List(noRawat string) ([]model.Pemberian, error) {
	k, err := a.konteks(noRawat)
	if err != nil {
		return nil, err
	}
	return a.repo.List(k.NoRawat, a.rawat)
}

// harga konteks penentu harga jual: depo, kenaikan per penjab, dan kolom harga.
type harga struct {
	depo     string
	kenaikan float64
	kolom    string
}

func (a *Action) harga(k perawatan.Konteks, kdBangsal, jenis string) (*harga, error) {
	depo := strings.TrimSpace(kdBangsal)
	if depo == "" {
		d, err := a.repo.Depo(a.rawat, k)
		if err != nil {
			return nil, err
		}
		depo = d
	}
	if depo == "" {
		return nil, ErrDepoNotFound
	}
	if ok, err := a.repo.DepoExists(depo); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrDepoNotFound
	}

	kolom := strings.TrimSpace(jenis)
	if kolom == "" {
		kolom = "ralan"
		if a.rawat == perawatan.Ranap {
			kolom = support.OrDefault(kelasHarga[k.Kelas], "kelas3")
		}
	}
	valid := false
	for _, v := range JenisHarga[a.rawat] {
		valid = valid || v == kolom
	}
	if !valid {
		return nil, ErrJenisHarga
	}
	kenaikan, err := a.repo.Kenaikan(a.rawat, k)
	if err != nil {
		return nil, err
	}
	return &harga{depo: depo, kenaikan: kenaikan, kolom: kolom}, nil
}

// CariObat daftar obat berstok di depo beserta harga jual untuk pasien ini.
func (a *Action) CariObat(noRawat, kdBangsal, jenis, search string, page, limit int) ([]model.Barang, int64, error) {
	k, err := a.konteks(noRawat)
	if err != nil {
		return nil, 0, err
	}
	h, err := a.harga(*k, kdBangsal, jenis)
	if err != nil {
		return nil, 0, err
	}
	list, total, err := a.repo.CariBarang(h.depo, search, a.set.Batch, a.set.Hpp, page, limit)
	if err != nil {
		return nil, 0, err
	}
	for i := range list {
		list[i].HargaJual = HargaJual(list[i], h.kenaikan, h.kolom, a.set.PembulatanHarga)
	}
	return list, total, nil
}

// Create menyimpan pemberian obat, memotong stok, dan memposting jurnal dalam satu transaksi.
func (a *Action) Create(in Input) ([]model.Pemberian, error) {
	if len(in.Items) == 0 {
		return nil, ErrTanpaObat
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
	h, err := a.harga(*k, in.KdBangsal, in.JenisHarga)
	if err != nil {
		return nil, err
	}

	rows := make([]model.Pemberian, 0, len(in.Items))
	for _, it := range in.Items {
		kode := strings.TrimSpace(it.KodeBrng)
		if a.set.Batch && (it.NoBatch == "" || it.NoFaktur == "") {
			return nil, ErrBatchWajib
		}
		if !a.set.Batch {
			it.NoBatch, it.NoFaktur = "", ""
		}
		b, err := a.repo.Barang(kode, h.depo, it.NoBatch, it.NoFaktur, a.set.Batch, a.set.Hpp)
		if err != nil {
			return nil, err
		}
		if b == nil {
			return nil, support.NotFound(fmt.Sprintf("obat %s tidak tersedia di depo %s", kode, h.depo))
		}
		jual := HargaJual(*b, h.kenaikan, h.kolom, a.set.PembulatanHarga)
		rows = append(rows, model.Pemberian{
			TglPerawatan: tgl.Format(support.DateLayout), Jam: jam, NoRawat: k.NoRawat, KodeBrng: kode, NamaBrng: b.NamaBrng,
			HBeli: b.Hpp, BiayaObat: jual, Jml: it.Jml, Embalase: it.Embalase, Tuslah: it.Tuslah,
			Total:  math.Round(it.Embalase + it.Tuslah + jual*it.Jml),
			Status: string(a.rawat), KdBangsal: h.depo, NoBatch: it.NoBatch, NoFaktur: it.NoFaktur,
			AturanPakai: strings.TrimSpace(it.AturanPakai),
		})
	}

	err = support.Transaction(func(tx contractsorm.Query) error {
		for _, p := range rows {
			if err := a.stok.Keluar(tx, mutasi(p, in.Operator, false)); err != nil {
				return err
			}
			if err := a.repo.Insert(tx, p); err != nil {
				return err
			}
		}
		ket := fmt.Sprintf("PEMBERIAN OBAT %s PASIEN %s %s, DIPOSTING OLEH %s", a.label(), k.NoRkmMedis, k.NmPasien, in.Operator)
		return a.postJurnal(tx, k.NoRawat, ket, rows, false)
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Delete membatalkan pemberian obat: stok dikembalikan dan jurnal dibalik.
func (a *Action) Delete(key model.Key, operator string) error {
	k, err := a.konteks(key.NoRawat)
	if err != nil {
		return err
	}
	if k.Terkunci() {
		return ErrTerkunci
	}
	key.NoRawat = k.NoRawat
	return support.Transaction(func(tx contractsorm.Query) error {
		p, err := a.repo.Find(tx, key)
		if err != nil {
			return err
		}
		if p == nil || p.Status != string(a.rawat) {
			return ErrNotFound
		}
		if err := a.repo.Delete(tx, key); err != nil {
			return err
		}
		if err := a.stok.Masuk(tx, mutasi(*p, operator, true)); err != nil {
			return err
		}
		ket := fmt.Sprintf("PEMBATALAN PEMBERIAN OBAT %s PASIEN %s %s OLEH %s", a.label(), k.NoRkmMedis, k.NmPasien, operator)
		return a.postJurnal(tx, k.NoRawat, ket, []model.Pemberian{*p}, true)
	})
}

func (a *Action) postJurnal(tx contractsorm.Query, noRawat, ket string, rows []model.Pemberian, batal bool) error {
	akun, err := a.akun.Obat(tx, a.rawat)
	if err != nil {
		return err
	}
	var jual, hpp float64
	for _, p := range rows {
		jual += p.Total
		hpp += math.Round(p.HBeli * p.Jml)
	}
	e := &jurnalsvc.Entries{}
	e.Pair(akun.SuspenPiutang, akun.Pendapatan, jual)
	e.Pair(akun.HPP, akun.Persediaan, hpp)
	if batal {
		e.Reverse()
	}
	return a.jurnal.Post(tx, noRawat, jurnalsvc.JenisUmum, ket, e)
}

func (a *Action) label() string {
	if a.rawat == perawatan.Ranap {
		return "RAWAT INAP"
	}
	return "RAWAT JALAN"
}

func mutasi(p model.Pemberian, operator string, batal bool) stoksvc.Mutasi {
	return stoksvc.Mutasi{
		KodeBrng: p.KodeBrng, KdBangsal: p.KdBangsal, NoBatch: p.NoBatch, NoFaktur: p.NoFaktur, Jumlah: p.Jml,
		Posisi: stokrepo.PosisiPemberianObat, Petugas: operator, Keterangan: "Pemberian Obat " + p.NoRawat, Batal: batal,
	}
}

// HargaJual harga per satuan: harga beli + kenaikan (set_harga_obat_*) bila diatur, selain itu kolom harga jenis;
// dibulatkan ke atas kelipatan 100 bila pembulatan aktif, selain itu dibulatkan biasa (Valid.roundUp).
func HargaJual(b model.Barang, kenaikan float64, kolom string, pembulatan bool) float64 {
	h := b.Kolom(kolom)
	if kenaikan > 0 {
		h = b.HBeli + b.HBeli*kenaikan
	}
	if !pembulatan {
		return math.Round(h)
	}
	if math.Mod(h, 100) == 0 {
		return h
	}
	return (math.Floor(h/100) + 1) * 100
}
