package tindakan

import (
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/perawatan"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

// tabel nama tabel tindakan per jenis rawat & pelaksana (whitelist, tidak pernah dari input).
var tabel = map[model.Rawat]map[model.Pelaksana]string{
	model.Ralan: {model.Dokter: "rawat_jl_dr", model.Paramedis: "rawat_jl_pr", model.DokterParamedis: "rawat_jl_drpr"},
	model.Ranap: {model.Dokter: "rawat_inap_dr", model.Paramedis: "rawat_inap_pr", model.DokterParamedis: "rawat_inap_drpr"},
}

var kolomTotal = map[model.Pelaksana]string{
	model.Dokter: "total_byrdr", model.Paramedis: "total_byrpr", model.DokterParamedis: "total_byrdrpr",
}

type Repository interface {
	perawatanrepo.KonteksRepository
	CariTarif(rawat model.Rawat, pelaksana model.Pelaksana, k model.Konteks, search string, page, limit int) ([]model.Tarif, int64, error)
	Tarif(rawat model.Rawat, pelaksana model.Pelaksana, k model.Konteks, kdJenisPrw string) (*model.Tarif, error)
	DokterExists(kdDokter string) (bool, error)
	PetugasExists(nip string) (bool, error)
	List(rawat model.Rawat, noRawat string) ([]model.Tindakan, error)
	Find(tx contractsorm.Query, rawat model.Rawat, key model.TindakanKey) (*model.Tindakan, error)
	Insert(tx contractsorm.Query, rawat model.Rawat, t model.Tindakan) error
	Delete(tx contractsorm.Query, rawat model.Rawat, key model.TindakanKey) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

// tarifQuery filter tarif mengikuti set_tarif (tarifralan/tarifranap) dan DlgCariPerawatanRanap.
func (r *repository) tarifQuery(rawat model.Rawat, pelaksana model.Pelaksana, k model.Konteks) (contractsorm.Query, error) {
	var list []struct {
		PoliRalan      string `gorm:"column:poli_ralan"`
		CaraBayarRalan string `gorm:"column:cara_bayar_ralan"`
		RuangRanap     string `gorm:"column:ruang_ranap"`
		CaraBayarRanap string `gorm:"column:cara_bayar_ranap"`
		KelasRanap     string `gorm:"column:kelas_ranap"`
	}
	if err := support.DB().Raw("select poli_ralan,cara_bayar_ralan,ruang_ranap,cara_bayar_ranap,kelas_ranap from set_tarif limit 1").Scan(&list); err != nil {
		return nil, err
	}
	var cfg struct{ PoliRalan, CaraBayarRalan, RuangRanap, CaraBayarRanap, KelasRanap bool }
	if len(list) > 0 {
		cfg.PoliRalan = list[0].PoliRalan == "Yes"
		cfg.CaraBayarRalan = list[0].CaraBayarRalan == "Yes"
		cfg.RuangRanap = list[0].RuangRanap == "Yes"
		cfg.CaraBayarRanap = list[0].CaraBayarRanap == "Yes"
		cfg.KelasRanap = list[0].KelasRanap == "Yes"
	}

	table := "jns_perawatan"
	if rawat == model.Ranap {
		table = "jns_perawatan_inap"
	}
	q := support.DB().Table(table).Where("status = '1' and " + kolomTotal[pelaksana] + " > 0")
	if rawat == model.Ralan {
		if cfg.PoliRalan {
			q = q.Where("kd_poli in (?, '-')", k.KdPoli)
		}
		if cfg.CaraBayarRalan {
			q = q.Where("kd_pj in (?, '-')", k.KdPj)
		}
		return q, nil
	}
	if cfg.RuangRanap {
		q = q.Where("kd_bangsal in (?, '-')", k.KdBangsal)
	}
	if cfg.CaraBayarRanap {
		q = q.Where("kd_pj in (?, '-')", k.KdPj)
	}
	if cfg.KelasRanap {
		q = q.Where("kelas in (?, '-')", k.Kelas)
	}
	return q, nil
}

const tarifColumns = "kd_jenis_prw,ifnull(nm_perawatan,'') as nm_perawatan,ifnull(material,0) as material,bhp," +
	"ifnull(tarif_tindakandr,0) as tarif_tindakandr,ifnull(tarif_tindakanpr,0) as tarif_tindakanpr,ifnull(kso,0) as kso," +
	"ifnull(menejemen,0) as menejemen,ifnull(total_byrdr,0) as total_byrdr,ifnull(total_byrpr,0) as total_byrpr,total_byrdrpr"

func (r *repository) CariTarif(rawat model.Rawat, pelaksana model.Pelaksana, k model.Konteks, search string, page, limit int) ([]model.Tarif, int64, error) {
	q, err := r.tarifQuery(rawat, pelaksana, k)
	if err != nil {
		return nil, 0, err
	}
	if s := strings.TrimSpace(search); s != "" {
		like := "%" + s + "%"
		q = q.Where("(kd_jenis_prw like ? or nm_perawatan like ?)", like, like)
	}
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Tarif{}
	err = q.Select(tarifColumns).Order("nm_perawatan").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

// Tarif satu tindakan yang berlaku untuk konteks; (nil, nil) bila tidak tersedia.
func (r *repository) Tarif(rawat model.Rawat, pelaksana model.Pelaksana, k model.Konteks, kdJenisPrw string) (*model.Tarif, error) {
	q, err := r.tarifQuery(rawat, pelaksana, k)
	if err != nil {
		return nil, err
	}
	list := []model.Tarif{}
	if err := q.Where("kd_jenis_prw = ?", kdJenisPrw).Select(tarifColumns).Limit(1).Scan(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return &list[0], nil
}

func (r *repository) DokterExists(kdDokter string) (bool, error) {
	return support.DB().Table("dokter").Where("kd_dokter = ? and status = '1'", kdDokter).Exists()
}

func (r *repository) PetugasExists(nip string) (bool, error) {
	return support.DB().Table("petugas").Where("nip = ? and status = '1'", nip).Exists()
}

func selectTindakan(rawat model.Rawat, p model.Pelaksana) string {
	t := tabel[rawat][p]
	dr, pr := "''", "''"
	nmdr, nmpr := "''", "''"
	tdr, tpr := "0", "0"
	join := ""
	if p.PakaiDokter() {
		dr, nmdr, tdr = t+".kd_dokter", "ifnull(dokter.nm_dokter,'')", "ifnull("+t+".tarif_tindakandr,0)"
		join += " left join dokter on " + t + ".kd_dokter=dokter.kd_dokter"
	}
	if p.PakaiParamedis() {
		pr, nmpr, tpr = t+".nip", "ifnull(petugas.nama,'')", "ifnull("+t+".tarif_tindakanpr,0)"
		join += " left join petugas on " + t + ".nip=petugas.nip"
	}
	jns := "jns_perawatan"
	if rawat == model.Ranap {
		jns = "jns_perawatan_inap"
	}
	return "select '" + string(p) + "' as pelaksana," + t + ".no_rawat," + t + ".kd_jenis_prw,ifnull(" + jns + ".nm_perawatan,'') as nm_perawatan," +
		dr + " as kd_dokter," + nmdr + " as nm_dokter," + pr + " as nip," + nmpr + " as nm_petugas," +
		"date_format(" + t + ".tgl_perawatan,'%Y-%m-%d') as tgl_perawatan,time_format(" + t + ".jam_rawat,'%H:%i:%s') as jam_rawat," +
		"ifnull(" + t + ".material,0) as material," + t + ".bhp," + tdr + " as tarif_tindakandr," + tpr + " as tarif_tindakanpr," +
		"ifnull(" + t + ".kso,0) as kso,ifnull(" + t + ".menejemen,0) as menejemen,ifnull(" + t + ".biaya_rawat,0) as biaya_rawat " +
		"from " + t + " left join " + jns + " on " + t + ".kd_jenis_prw=" + jns + ".kd_jenis_prw" + join
}

// List seluruh tindakan dokter, paramedis, dan dokter+paramedis untuk satu no_rawat.
func (r *repository) List(rawat model.Rawat, noRawat string) ([]model.Tindakan, error) {
	out := []model.Tindakan{}
	for _, p := range []model.Pelaksana{model.Dokter, model.Paramedis, model.DokterParamedis} {
		list := []model.Tindakan{}
		sql := selectTindakan(rawat, p) + " where " + tabel[rawat][p] + ".no_rawat=? order by tgl_perawatan,jam_rawat"
		if err := support.DB().Raw(sql, noRawat).Scan(&list); err != nil {
			return nil, err
		}
		out = append(out, list...)
	}
	return out, nil
}

func keyWhere(t string, key model.TindakanKey) (string, []any) {
	cond := t + ".no_rawat=? and " + t + ".kd_jenis_prw=? and " + t + ".tgl_perawatan=? and " + t + ".jam_rawat=?"
	args := []any{key.NoRawat, key.KdJenisPrw, key.TglPerawatan, key.JamRawat}
	if key.Pelaksana.PakaiDokter() {
		cond += " and " + t + ".kd_dokter=?"
		args = append(args, key.KdDokter)
	}
	if key.Pelaksana.PakaiParamedis() {
		cond += " and " + t + ".nip=?"
		args = append(args, key.Nip)
	}
	return cond, args
}

// Find baris tindakan (dikunci for update di dalam transaksi); (nil, nil) bila tidak ada.
func (r *repository) Find(tx contractsorm.Query, rawat model.Rawat, key model.TindakanKey) (*model.Tindakan, error) {
	t := tabel[rawat][key.Pelaksana]
	cond, args := keyWhere(t, key)
	list := []model.Tindakan{}
	if err := tx.Raw(selectTindakan(rawat, key.Pelaksana)+" where "+cond+" for update", args...).Scan(&list); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return &list[0], nil
}

func (r *repository) Insert(tx contractsorm.Query, rawat model.Rawat, t model.Tindakan) error {
	table := tabel[rawat][t.Pelaksana]
	cols := []string{"no_rawat", "kd_jenis_prw"}
	args := []any{t.NoRawat, t.KdJenisPrw}
	if t.Pelaksana.PakaiDokter() {
		cols = append(cols, "kd_dokter")
		args = append(args, t.KdDokter)
	}
	if t.Pelaksana.PakaiParamedis() {
		cols = append(cols, "nip")
		args = append(args, t.Nip)
	}
	cols = append(cols, "tgl_perawatan", "jam_rawat", "material", "bhp")
	args = append(args, t.TglPerawatan, t.JamRawat, t.Material, t.Bhp)
	if t.Pelaksana.PakaiDokter() {
		cols = append(cols, "tarif_tindakandr")
		args = append(args, t.TarifTindakandr)
	}
	if t.Pelaksana.PakaiParamedis() {
		cols = append(cols, "tarif_tindakanpr")
		args = append(args, t.TarifTindakanpr)
	}
	cols = append(cols, "kso", "menejemen", "biaya_rawat")
	args = append(args, t.Kso, t.Menejemen, t.BiayaRawat)
	if rawat == model.Ralan {
		cols = append(cols, "stts_bayar")
		args = append(args, "Belum")
	}
	_, err := tx.Exec("insert into "+table+" ("+strings.Join(cols, ",")+") values (?"+strings.Repeat(",?", len(cols)-1)+")", args...)
	return err
}

func (r *repository) Delete(tx contractsorm.Query, rawat model.Rawat, key model.TindakanKey) error {
	t := tabel[rawat][key.Pelaksana]
	cond, args := keyWhere(t, key)
	_, err := tx.Exec("delete from "+t+" where "+cond, args...)
	return err
}
