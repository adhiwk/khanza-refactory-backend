package templatedokter

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/templatedokter"
	"goravel/app/repository/crud"
	"goravel/app/support"
)

// anak seluruh tabel detail (urutan hapus: detail racikan sebelum racikan).
var anak = []string{
	"template_pemeriksaan_dokter_penyakit", "template_pemeriksaan_dokter_prosedur", "template_pemeriksaan_dokter_permintaan_radiologi",
	"template_pemeriksaan_dokter_detail_permintaan_lab", "template_pemeriksaan_dokter_permintaan_lab", "template_pemeriksaan_dokter_resep",
	"template_pemeriksaan_dokter_resep_racikan_detail", "template_pemeriksaan_dokter_resep_racikan", "template_pemeriksaan_dokter_tindakan",
}

type Repository interface {
	Paginate(search, kdDokter string, page, limit int) ([]model.Template, int64, error)
	Find(no string) (*model.Template, error)
	RefExists(table, column string, value any) (bool, error)
	LabTemplateExists(kdJenisPrw string, idTemplate int) (bool, error)
	NextKode(tx contractsorm.Query) (string, error)
	Insert(tx contractsorm.Query, t model.Template) error
	Update(tx contractsorm.Query, t model.Template) error
	Delete(tx contractsorm.Query, no string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

const headerSQL = "select t.no_template,ifnull(t.kd_dokter,'') as kd_dokter,ifnull(dokter.nm_dokter,'') as nm_dokter,ifnull(t.keluhan,'') as keluhan," +
	"ifnull(t.pemeriksaan,'') as pemeriksaan,ifnull(t.penilaian,'') as penilaian,t.rencana,t.instruksi,ifnull(t.evaluasi,'') as evaluasi " +
	"from template_pemeriksaan_dokter t left join dokter on t.kd_dokter=dokter.kd_dokter"

func (r *repository) Paginate(search, kdDokter string, page, limit int) ([]model.Template, int64, error) {
	q := support.DB().Table("(" + headerSQL + ") x")
	if kdDokter != "" {
		q = q.Where("kd_dokter = ?", kdDokter)
	}
	q = crud.WhereLike(q, search, "no_template", "nm_dokter", "keluhan", "penilaian")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Template{}
	err = q.Order("no_template desc").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

func scan[T any](sql string, args ...any) ([]T, error) {
	out := []T{}
	err := support.DB().Raw(sql, args...).Scan(&out)
	return out, err
}

// Find (nil, nil) bila tidak ada.
func (r *repository) Find(no string) (*model.Template, error) {
	list, err := scan[model.Template](headerSQL+" where t.no_template=?", no)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	t := &list[0]
	if t.Diagnosa, err = scan[model.Kode]("select d.kd_penyakit as kode,ifnull(p.nm_penyakit,'') as nama,d.urut from template_pemeriksaan_dokter_penyakit d "+
		"left join penyakit p on d.kd_penyakit=p.kd_penyakit where d.no_template=? order by d.urut", no); err != nil {
		return nil, err
	}
	if t.Prosedur, err = scan[model.Kode]("select d.kode,ifnull(i.deskripsi_panjang,'') as nama,d.urut,d.jumlah from template_pemeriksaan_dokter_prosedur d "+
		"left join icd9 i on d.kode=i.kode where d.no_template=? order by d.urut", no); err != nil {
		return nil, err
	}
	if t.Radiologi, err = scan[model.Kode]("select d.kd_jenis_prw as kode,ifnull(j.nm_perawatan,'') as nama from template_pemeriksaan_dokter_permintaan_radiologi d "+
		"left join jns_perawatan_radiologi j on d.kd_jenis_prw=j.kd_jenis_prw where d.no_template=?", no); err != nil {
		return nil, err
	}
	if t.Tindakan, err = scan[model.Kode]("select d.kd_jenis_prw as kode,ifnull(j.nm_perawatan,'') as nama from template_pemeriksaan_dokter_tindakan d "+
		"left join jns_perawatan j on d.kd_jenis_prw=j.kd_jenis_prw where d.no_template=?", no); err != nil {
		return nil, err
	}
	if t.Lab, err = scan[model.Lab]("select d.kd_jenis_prw,ifnull(j.nm_perawatan,'') as nama from template_pemeriksaan_dokter_permintaan_lab d "+
		"left join jns_perawatan_lab j on d.kd_jenis_prw=j.kd_jenis_prw where d.no_template=?", no); err != nil {
		return nil, err
	}
	for i := range t.Lab {
		ids := []int{}
		if err := support.DB().Table("template_pemeriksaan_dokter_detail_permintaan_lab").
			Where("no_template = ? and kd_jenis_prw = ?", no, t.Lab[i].KdJenisPrw).Order("id_template").Pluck("id_template", &ids); err != nil {
			return nil, err
		}
		t.Lab[i].IDTemplate = ids
	}
	if t.Resep, err = scan[model.Resep]("select d.kode_brng,ifnull(b.nama_brng,'') as nama_brng,ifnull(d.jml,0) as jml,ifnull(d.aturan_pakai,'') as aturan_pakai "+
		"from template_pemeriksaan_dokter_resep d left join databarang b on d.kode_brng=b.kode_brng where d.no_template=?", no); err != nil {
		return nil, err
	}
	if t.Racikan, err = scan[model.Racikan]("select no_racik,ifnull(nama_racik,'') as nama_racik,ifnull(kd_racik,'') as kd_racik,ifnull(jml_dr,0) as jml_dr,"+
		"ifnull(aturan_pakai,'') as aturan_pakai,ifnull(keterangan,'') as keterangan from template_pemeriksaan_dokter_resep_racikan where no_template=? order by no_racik", no); err != nil {
		return nil, err
	}
	for i := range t.Racikan {
		if t.Racikan[i].Detail, err = scan[model.RacikanDetail]("select d.kode_brng,ifnull(b.nama_brng,'') as nama_brng,ifnull(d.p1,0) as p1,ifnull(d.p2,0) as p2,"+
			"ifnull(d.kandungan,'') as kandungan,ifnull(d.jml,0) as jml from template_pemeriksaan_dokter_resep_racikan_detail d "+
			"left join databarang b on d.kode_brng=b.kode_brng where d.no_template=? and d.no_racik=?", no, t.Racikan[i].NoRacik); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (r *repository) RefExists(table, column string, value any) (bool, error) {
	return crud.ExistsIn(support.DB(), table, column, value)
}

func (r *repository) LabTemplateExists(kdJenisPrw string, idTemplate int) (bool, error) {
	return support.DB().Table("template_laboratorium").Where("kd_jenis_prw = ? and id_template = ?", kdJenisPrw, idTemplate).Exists()
}

// NextKode TPD + 16 digit (Valid.autoNomer("template_pemeriksaan_dokter","TPD",16)).
func (r *repository) NextKode(tx contractsorm.Query) (string, error) {
	return crud.NextCode(tx, "template_pemeriksaan_dokter", "no_template", "TPD", 16)
}

func (r *repository) Insert(tx contractsorm.Query, t model.Template) error {
	if _, err := tx.Exec("insert into template_pemeriksaan_dokter (no_template,kd_dokter,keluhan,pemeriksaan,penilaian,rencana,instruksi,evaluasi) "+
		"values (?,?,?,?,?,?,?,?)", t.NoTemplate, t.KdDokter, t.Keluhan, t.Pemeriksaan, t.Penilaian, t.Rencana, t.Instruksi, t.Evaluasi); err != nil {
		return err
	}
	return r.insertIsi(tx, t)
}

func (r *repository) Update(tx contractsorm.Query, t model.Template) error {
	if _, err := tx.Exec("update template_pemeriksaan_dokter set kd_dokter=?,keluhan=?,pemeriksaan=?,penilaian=?,rencana=?,instruksi=?,evaluasi=? "+
		"where no_template=?", t.KdDokter, t.Keluhan, t.Pemeriksaan, t.Penilaian, t.Rencana, t.Instruksi, t.Evaluasi, t.NoTemplate); err != nil {
		return err
	}
	if err := r.deleteIsi(tx, t.NoTemplate); err != nil {
		return err
	}
	return r.insertIsi(tx, t)
}

func (r *repository) Delete(tx contractsorm.Query, no string) error {
	if err := r.deleteIsi(tx, no); err != nil {
		return err
	}
	_, err := tx.Exec("delete from template_pemeriksaan_dokter where no_template=?", no)
	return err
}

func (r *repository) deleteIsi(tx contractsorm.Query, no string) error {
	for _, t := range anak {
		if _, err := tx.Exec("delete from "+t+" where no_template=?", no); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) insertIsi(tx contractsorm.Query, t model.Template) error {
	no := t.NoTemplate
	exec := func(sql string, args ...any) error {
		_, err := tx.Exec(sql, args...)
		return err
	}
	for _, d := range t.Diagnosa {
		if err := exec("insert into template_pemeriksaan_dokter_penyakit (no_template,kd_penyakit,urut) values (?,?,?)", no, d.Kode, d.Urut); err != nil {
			return err
		}
	}
	for _, p := range t.Prosedur {
		if err := exec("insert into template_pemeriksaan_dokter_prosedur (no_template,kode,urut,jumlah) values (?,?,?,?)", no, p.Kode, p.Urut, p.Jumlah); err != nil {
			return err
		}
	}
	for _, k := range t.Radiologi {
		if err := exec("insert into template_pemeriksaan_dokter_permintaan_radiologi (no_template,kd_jenis_prw) values (?,?)", no, k.Kode); err != nil {
			return err
		}
	}
	for _, l := range t.Lab {
		if err := exec("insert into template_pemeriksaan_dokter_permintaan_lab (no_template,kd_jenis_prw) values (?,?)", no, l.KdJenisPrw); err != nil {
			return err
		}
		for _, id := range l.IDTemplate {
			if err := exec("insert into template_pemeriksaan_dokter_detail_permintaan_lab (no_template,kd_jenis_prw,id_template) values (?,?,?)",
				no, l.KdJenisPrw, id); err != nil {
				return err
			}
		}
	}
	for _, o := range t.Resep {
		if err := exec("insert into template_pemeriksaan_dokter_resep (no_template,kode_brng,jml,aturan_pakai) values (?,?,?,?)",
			no, o.KodeBrng, o.Jml, o.AturanPakai); err != nil {
			return err
		}
	}
	for _, rc := range t.Racikan {
		if err := exec("insert into template_pemeriksaan_dokter_resep_racikan (no_template,no_racik,nama_racik,kd_racik,jml_dr,aturan_pakai,keterangan) "+
			"values (?,?,?,?,?,?,?)", no, rc.NoRacik, rc.NamaRacik, rc.KdRacik, rc.JmlDr, rc.AturanPakai, rc.Keterangan); err != nil {
			return err
		}
		for _, d := range rc.Detail {
			if err := exec("insert into template_pemeriksaan_dokter_resep_racikan_detail (no_template,no_racik,kode_brng,p1,p2,kandungan,jml) "+
				"values (?,?,?,?,?,?,?)", no, rc.NoRacik, d.KodeBrng, d.P1, d.P2, d.Kandungan, d.Jml); err != nil {
				return err
			}
		}
	}
	for _, k := range t.Tindakan {
		if err := exec("insert into template_pemeriksaan_dokter_tindakan (no_template,kd_jenis_prw) values (?,?)", no, k.Kode); err != nil {
			return err
		}
	}
	return nil
}
