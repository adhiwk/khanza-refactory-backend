package pemberianobat

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/pemberianobat"
	"goravel/app/models/perawatan"
	"goravel/app/repository/crud"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

type Repository interface {
	perawatanrepo.KonteksRepository
	Depo(rawat perawatan.Rawat, k perawatan.Konteks) (string, error)
	DepoExists(kdBangsal string) (bool, error)
	Kenaikan(rawat perawatan.Rawat, k perawatan.Konteks) (float64, error)
	CariBarang(kdBangsal, search string, batch bool, hpp string, page, limit int) ([]model.Barang, int64, error)
	Barang(kodeBrng, kdBangsal, noBatch, noFaktur string, batch bool, hpp string) (*model.Barang, error)
	List(noRawat string, rawat perawatan.Rawat) ([]model.Pemberian, error)
	Find(tx contractsorm.Query, key model.Key) (*model.Pemberian, error)
	Insert(tx contractsorm.Query, p model.Pemberian) error
	Delete(tx contractsorm.Query, key model.Key) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) scalar(sql string, args ...any) (string, error) {
	var list []struct {
		V string `gorm:"column:v"`
	}
	if err := support.DB().Raw(sql, args...).Scan(&list); err != nil || len(list) == 0 {
		return "", err
	}
	return list[0].V, nil
}

// Depo gudang obat default: set_depo_ralan per poli / set_depo_ranap per bangsal, lalu set_lokasi.
func (r *repository) Depo(rawat perawatan.Rawat, k perawatan.Konteks) (string, error) {
	var depo string
	var err error
	if rawat == perawatan.Ralan {
		depo, err = r.scalar("select kd_bangsal as v from set_depo_ralan where kd_poli=? limit 1", k.KdPoli)
	} else {
		depo, err = r.scalar("select kd_depo as v from set_depo_ranap where kd_bangsal=? limit 1", k.KdBangsal)
	}
	if err != nil || depo != "" {
		return depo, err
	}
	return r.scalar("select kd_bangsal as v from set_lokasi limit 1")
}

func (r *repository) DepoExists(kdBangsal string) (bool, error) {
	return crud.ExistsIn(support.DB(), "bangsal", "kd_bangsal", kdBangsal)
}

// Kenaikan persentase harga jual dari harga beli (set_harga_obat_ralan / set_harga_obat_ranap), 0 bila tidak diatur.
func (r *repository) Kenaikan(rawat perawatan.Rawat, k perawatan.Konteks) (float64, error) {
	var list []struct {
		V float64 `gorm:"column:v"`
	}
	var err error
	if rawat == perawatan.Ralan {
		err = support.DB().Raw("select hargajual/100 as v from set_harga_obat_ralan where kd_pj=?", k.KdPj).Scan(&list)
	} else {
		err = support.DB().Raw("select hargajual/100 as v from set_harga_obat_ranap where kd_pj=? and kelas=?", k.KdPj, k.Kelas).Scan(&list)
	}
	if err != nil || len(list) == 0 {
		return 0, err
	}
	return list[0].V, nil
}

// barangQuery stok > 0 di depo; mode batch mengambil harga per data_batch, selain itu dari databarang (no_batch ”).
func barangQuery(kdBangsal string, batch bool, hpp string) contractsorm.Query {
	if hpp != "h_beli" {
		hpp = "dasar"
	}
	src := "databarang"
	q := support.DB().Table("gudangbarang").Join("inner join databarang on gudangbarang.kode_brng=databarang.kode_brng")
	if batch {
		src = "data_batch"
		q = q.Join("inner join data_batch on gudangbarang.kode_brng=data_batch.kode_brng and gudangbarang.no_batch=data_batch.no_batch and gudangbarang.no_faktur=data_batch.no_faktur").
			Where("gudangbarang.no_batch <> ''")
	} else {
		q = q.Where("gudangbarang.no_batch = '' and gudangbarang.no_faktur = ''")
	}
	cols := []string{"databarang.kode_brng", "ifnull(databarang.nama_brng,'') as nama_brng", "ifnull(databarang.kode_sat,'') as kode_sat",
		"gudangbarang.no_batch", "gudangbarang.no_faktur", "gudangbarang.stok", "ifnull(" + src + ".h_beli,0) as h_beli",
		"ifnull(" + src + "." + hpp + ",0) as hpp"}
	for _, c := range []string{"ralan", "kelas1", "kelas2", "kelas3", "utama", "vip", "vvip", "beliluar", "karyawan"} {
		cols = append(cols, "ifnull("+src+"."+c+",0) as "+c)
	}
	return q.Select(cols...).Where("gudangbarang.kd_bangsal = ? and databarang.status = '1'", kdBangsal)
}

func (r *repository) CariBarang(kdBangsal, search string, batch bool, hpp string, page, limit int) ([]model.Barang, int64, error) {
	q := crud.WhereLike(barangQuery(kdBangsal, batch, hpp).Where("gudangbarang.stok > 0"), search,
		"databarang.kode_brng", "databarang.nama_brng", "databarang.letak_barang", "gudangbarang.no_batch")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Barang{}
	err = q.Order("databarang.nama_brng").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

// Barang satu obat di depo/batch; (nil, nil) bila tidak ada.
func (r *repository) Barang(kodeBrng, kdBangsal, noBatch, noFaktur string, batch bool, hpp string) (*model.Barang, error) {
	q := barangQuery(kdBangsal, batch, hpp).Where("gudangbarang.kode_brng = ?", kodeBrng)
	if batch {
		q = q.Where("gudangbarang.no_batch = ? and gudangbarang.no_faktur = ?", noBatch, noFaktur)
	}
	list := []model.Barang{}
	if err := q.Limit(1).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

const pemberianColumns = "cast(detail_pemberian_obat.tgl_perawatan as char) as tgl_perawatan,cast(detail_pemberian_obat.jam as char) as jam," +
	"detail_pemberian_obat.no_rawat,detail_pemberian_obat.kode_brng,ifnull(databarang.nama_brng,'') as nama_brng," +
	"ifnull(detail_pemberian_obat.h_beli,0) as h_beli,ifnull(detail_pemberian_obat.biaya_obat,0) as biaya_obat,detail_pemberian_obat.jml," +
	"ifnull(detail_pemberian_obat.embalase,0) as embalase,ifnull(detail_pemberian_obat.tuslah,0) as tuslah,detail_pemberian_obat.total," +
	"ifnull(detail_pemberian_obat.status,'') as status,ifnull(detail_pemberian_obat.kd_bangsal,'') as kd_bangsal," +
	"detail_pemberian_obat.no_batch,detail_pemberian_obat.no_faktur,ifnull(aturan_pakai.aturan,'') as aturan"

func view(q contractsorm.Query) contractsorm.Query {
	return q.Table("detail_pemberian_obat").
		Join("left join databarang on detail_pemberian_obat.kode_brng=databarang.kode_brng").
		Join("left join aturan_pakai on aturan_pakai.no_rawat=detail_pemberian_obat.no_rawat and aturan_pakai.kode_brng=detail_pemberian_obat.kode_brng " +
			"and aturan_pakai.tgl_perawatan=detail_pemberian_obat.tgl_perawatan and aturan_pakai.jam=detail_pemberian_obat.jam").
		Select(pemberianColumns)
}

func (r *repository) List(noRawat string, rawat perawatan.Rawat) ([]model.Pemberian, error) {
	list := []model.Pemberian{}
	err := view(support.DB()).Where("detail_pemberian_obat.no_rawat = ? and detail_pemberian_obat.status = ?", noRawat, string(rawat)).
		Order("detail_pemberian_obat.tgl_perawatan, detail_pemberian_obat.jam").Scan(&list)
	return list, err
}

const keyWhere = "detail_pemberian_obat.no_rawat=? and detail_pemberian_obat.tgl_perawatan=? and detail_pemberian_obat.jam=? " +
	"and detail_pemberian_obat.kode_brng=? and detail_pemberian_obat.no_batch=? and detail_pemberian_obat.no_faktur=?"

func keyArgs(k model.Key) []any {
	return []any{k.NoRawat, k.TglPerawatan, k.Jam, k.KodeBrng, k.NoBatch, k.NoFaktur}
}

func (r *repository) Find(tx contractsorm.Query, key model.Key) (*model.Pemberian, error) {
	list := []model.Pemberian{}
	if err := view(tx).Where(keyWhere, keyArgs(key)...).Limit(1).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *repository) Insert(tx contractsorm.Query, p model.Pemberian) error {
	_, err := tx.Exec("insert into detail_pemberian_obat (tgl_perawatan,jam,no_rawat,kode_brng,h_beli,biaya_obat,jml,embalase,tuslah,total,"+
		"status,kd_bangsal,no_batch,no_faktur) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		p.TglPerawatan, p.Jam, p.NoRawat, p.KodeBrng, p.HBeli, p.BiayaObat, p.Jml, p.Embalase, p.Tuslah, p.Total,
		p.Status, p.KdBangsal, p.NoBatch, p.NoFaktur)
	if err != nil || p.AturanPakai == "" {
		return err
	}
	_, err = tx.Exec("insert into aturan_pakai (tgl_perawatan,jam,no_rawat,kode_brng,aturan) values (?,?,?,?,?) "+
		"on duplicate key update aturan=values(aturan)", p.TglPerawatan, p.Jam, p.NoRawat, p.KodeBrng, p.AturanPakai)
	return err
}

func (r *repository) Delete(tx contractsorm.Query, key model.Key) error {
	if _, err := tx.Exec("delete from detail_pemberian_obat where "+keyWhere, keyArgs(key)...); err != nil {
		return err
	}
	_, err := tx.Exec("delete from aturan_pakai where no_rawat=? and tgl_perawatan=? and jam=? and kode_brng=?",
		key.NoRawat, key.TglPerawatan, key.Jam, key.KodeBrng)
	return err
}
