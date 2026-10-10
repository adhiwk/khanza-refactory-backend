package icd

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/icd"
	"goravel/app/repository/crud"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

// tabel metadata per jenis (whitelist).
type tabel struct {
	table, kode, master, masterKode, masterNama string
	resume                                      []string
}

var tabels = map[model.Jenis]tabel{
	model.Diagnosa: {"diagnosa_pasien", "kd_penyakit", "penyakit", "kd_penyakit", "nm_penyakit",
		[]string{"kd_diagnosa_utama", "kd_diagnosa_sekunder", "kd_diagnosa_sekunder2", "kd_diagnosa_sekunder3", "kd_diagnosa_sekunder4"}},
	model.Prosedur: {"prosedur_pasien", "kode", "icd9", "kode", "deskripsi_panjang",
		[]string{"kd_prosedur_utama", "kd_prosedur_sekunder", "kd_prosedur_sekunder2", "kd_prosedur_sekunder3"}},
}

type Repository interface {
	perawatanrepo.KonteksRepository
	List(jenis model.Jenis, noRawat, noRkmMedis string) ([]model.Kode, error)
	CariReferensi(jenis model.Jenis, search string, page, limit int) ([]model.Referensi, int64, error)
	ReferensiExists(jenis model.Jenis, kode string) (bool, error)
	Exists(jenis model.Jenis, noRawat, kode string) (bool, error)
	PernahDiderita(noRkmMedis, kdPenyakit string) (bool, error)
	MaxPrioritas(tx contractsorm.Query, jenis model.Jenis, noRawat, status string) (int, error)
	Insert(tx contractsorm.Query, jenis model.Jenis, noRawat, status string, k model.Item, statusPenyakit string) error
	Delete(tx contractsorm.Query, jenis model.Jenis, noRawat, kode string) error
	SetStatusPenyakit(noRawat, kode, status string) error
	SyncResume(tx contractsorm.Query, jenis model.Jenis, noRawat, status string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) List(jenis model.Jenis, noRawat, noRkmMedis string) ([]model.Kode, error) {
	t := tabels[jenis]
	var extra string
	if jenis == model.Diagnosa {
		extra = "ifnull(x.status_penyakit,'') as status_penyakit,'' as jumlah"
	} else {
		extra = "'' as status_penyakit,x.jumlah"
	}
	sql := "select x.no_rawat,cast(reg_periksa.tgl_registrasi as char) as tgl_registrasi,x." + t.kode + " as kode," +
		"ifnull(m." + t.masterNama + ",'') as nama,x.status,x.prioritas," + extra +
		" from " + t.table + " x inner join reg_periksa on x.no_rawat=reg_periksa.no_rawat " +
		"left join " + t.master + " m on x." + t.kode + "=m." + t.masterKode + " where "
	var args []any
	if noRawat != "" {
		sql += "x.no_rawat=?"
		args = append(args, noRawat)
	} else {
		sql += "reg_periksa.no_rkm_medis=?"
		args = append(args, noRkmMedis)
	}
	list := []model.Kode{}
	err := support.DB().Raw(sql+" order by reg_periksa.tgl_registrasi desc,x.no_rawat,x.status,x.prioritas", args...).Scan(&list)
	return list, err
}

func (r *repository) CariReferensi(jenis model.Jenis, search string, page, limit int) ([]model.Referensi, int64, error) {
	t := tabels[jenis]
	q := crud.WhereLike(support.DB().Table(t.master), search, t.masterKode, t.masterNama)
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Referensi{}
	err = q.Select(t.masterKode+" as kode", t.masterNama+" as nama").Order(t.masterKode).
		Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

func (r *repository) ReferensiExists(jenis model.Jenis, kode string) (bool, error) {
	t := tabels[jenis]
	return crud.ExistsIn(support.DB(), t.master, t.masterKode, kode)
}

func (r *repository) Exists(jenis model.Jenis, noRawat, kode string) (bool, error) {
	t := tabels[jenis]
	return support.DB().Table(t.table).Where("no_rawat = ? and "+t.kode+" = ?", noRawat, kode).Exists()
}

func (r *repository) PernahDiderita(noRkmMedis, kdPenyakit string) (bool, error) {
	return support.DB().Table("diagnosa_pasien").
		Join("inner join reg_periksa on diagnosa_pasien.no_rawat=reg_periksa.no_rawat").
		Where("reg_periksa.no_rkm_medis = ? and diagnosa_pasien.kd_penyakit = ?", noRkmMedis, kdPenyakit).Exists()
}

func (r *repository) MaxPrioritas(tx contractsorm.Query, jenis model.Jenis, noRawat, status string) (int, error) {
	t := tabels[jenis]
	var list []struct {
		Max int `gorm:"column:max"`
	}
	if err := tx.Raw("select ifnull(max(prioritas),0) as max from "+t.table+" where no_rawat=? and status=? for update", noRawat, status).Scan(&list); err != nil {
		return 0, err
	}
	if len(list) == 0 {
		return 0, nil
	}
	return list[0].Max, nil
}

func (r *repository) Insert(tx contractsorm.Query, jenis model.Jenis, noRawat, status string, k model.Item, statusPenyakit string) error {
	var err error
	if jenis == model.Diagnosa {
		_, err = tx.Exec("insert into diagnosa_pasien (no_rawat,kd_penyakit,status,prioritas,status_penyakit) values (?,?,?,?,?)",
			noRawat, k.Kode, status, k.Prioritas, statusPenyakit)
	} else {
		_, err = tx.Exec("insert into prosedur_pasien (no_rawat,kode,status,prioritas,jumlah) values (?,?,?,?,?)",
			noRawat, k.Kode, status, k.Prioritas, k.Jumlah)
	}
	return err
}

func (r *repository) Delete(tx contractsorm.Query, jenis model.Jenis, noRawat, kode string) error {
	t := tabels[jenis]
	_, err := tx.Exec("delete from "+t.table+" where no_rawat=? and "+t.kode+"=?", noRawat, kode)
	return err
}

func (r *repository) SetStatusPenyakit(noRawat, kode, status string) error {
	_, err := support.DB().Exec("update diagnosa_pasien set status_penyakit=? where no_rawat=? and kd_penyakit=?", status, noRawat, kode)
	return err
}

// SyncResume mengisi kode utama/sekunder resume_pasien (Ralan) / resume_pasien_ranap (Ranap) sesuai urutan prioritas.
func (r *repository) SyncResume(tx contractsorm.Query, jenis model.Jenis, noRawat, status string) error {
	t := tabels[jenis]
	resume := "resume_pasien"
	if status == "Ranap" {
		resume = "resume_pasien_ranap"
	}
	codes := []string{}
	if err := tx.Table(t.table).Where("no_rawat = ? and status = ?", noRawat, status).Order("prioritas").Pluck(t.kode, &codes); err != nil {
		return err
	}
	set, args := "", []any{}
	for i, col := range t.resume {
		v := ""
		if i < len(codes) {
			v = codes[i]
		}
		if set != "" {
			set += ","
		}
		set += col + "=?"
		args = append(args, v)
	}
	_, err := tx.Exec("update "+resume+" set "+set+" where no_rawat=?", append(args, noRawat)...)
	return err
}
