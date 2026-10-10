package resep

import (
	"fmt"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/resep"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

type Repository interface {
	perawatanrepo.KonteksRepository
	DokterExists(kdDokter string) (bool, error)
	BarangExists(kodeBrng string) (bool, error)
	List(noRawat string) ([]model.Resep, error)
	Find(tx contractsorm.Query, noResep string) (*model.Resep, error)
	NextNoResep(tx contractsorm.Query, tgl time.Time) (string, error)
	Insert(tx contractsorm.Query, r model.Resep) error
	ReplaceItems(tx contractsorm.Query, noResep string, items []model.Item) error
	Delete(tx contractsorm.Query, noResep string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) DokterExists(kdDokter string) (bool, error) {
	return support.DB().Table("dokter").Where("kd_dokter = ? and status = '1'", kdDokter).Exists()
}

func (r *repository) BarangExists(kodeBrng string) (bool, error) {
	return support.DB().Table("databarang").Where("kode_brng = ? and status = '1'", kodeBrng).Exists()
}

const headerColumns = "resep_obat.no_resep,resep_obat.no_rawat,resep_obat.kd_dokter,ifnull(dokter.nm_dokter,'') as nm_dokter," +
	"ifnull(cast(resep_obat.tgl_peresepan as char),'') as tgl_peresepan,ifnull(cast(resep_obat.jam_peresepan as char),'') as jam_peresepan," +
	"ifnull(cast(resep_obat.tgl_perawatan as char),'') as tgl_perawatan,cast(resep_obat.jam as char) as jam,ifnull(resep_obat.status,'') as status"

func header(q contractsorm.Query) contractsorm.Query {
	return q.Table("resep_obat").Join("left join dokter on resep_obat.kd_dokter=dokter.kd_dokter").Select(headerColumns)
}

func (r *repository) items(q contractsorm.Query, noResep string) ([]model.Item, error) {
	list := []model.Item{}
	err := q.Table("resep_dokter").Join("left join databarang on resep_dokter.kode_brng=databarang.kode_brng").
		Select("resep_dokter.kode_brng", "ifnull(databarang.nama_brng,'') as nama_brng", "ifnull(resep_dokter.jml,0) as jml",
			"ifnull(resep_dokter.aturan_pakai,'') as aturan_pakai").
		Where("resep_dokter.no_resep = ?", noResep).Scan(&list)
	return list, err
}

func (r *repository) List(noRawat string) ([]model.Resep, error) {
	list := []model.Resep{}
	if err := header(support.DB()).Where("resep_obat.no_rawat = ?", noRawat).Order("resep_obat.no_resep").Scan(&list); err != nil {
		return nil, err
	}
	for i := range list {
		items, err := r.items(support.DB(), list[i].NoResep)
		if err != nil {
			return nil, err
		}
		list[i].Items = items
	}
	return list, nil
}

// Find header resep dikunci for update beserta item; (nil, nil) bila tidak ada.
func (r *repository) Find(tx contractsorm.Query, noResep string) (*model.Resep, error) {
	list := []model.Resep{}
	if err := tx.Raw("select "+headerColumns+" from resep_obat left join dokter on resep_obat.kd_dokter=dokter.kd_dokter "+
		"where resep_obat.no_resep=? for update", noResep).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}
	items, err := r.items(tx, noResep)
	if err != nil {
		return nil, err
	}
	list[0].Items = items
	return &list[0], nil
}

// NextNoResep yyyyMMdd + 4 digit per tanggal peresepan/perawatan (DlgPeresepanDokter.emptTeksobat).
func (r *repository) NextNoResep(tx contractsorm.Query, tgl time.Time) (string, error) {
	var list []struct {
		Max int `gorm:"column:max"`
	}
	d := tgl.Format("2006-01-02")
	if err := tx.Raw("select ifnull(max(convert(right(no_resep,4),signed)),0) as max from resep_obat "+
		"where tgl_peresepan=? or tgl_perawatan=? for update", d, d).Scan(&list); err != nil {
		return "", err
	}
	max := 0
	if len(list) > 0 {
		max = list[0].Max
	}
	return fmt.Sprintf("%s%04d", tgl.Format("20060102"), max+1), nil
}

func (r *repository) Insert(tx contractsorm.Query, x model.Resep) error {
	_, err := tx.Exec("insert into resep_obat (no_resep,tgl_perawatan,jam,no_rawat,kd_dokter,tgl_peresepan,jam_peresepan,status,"+
		"tgl_penyerahan,jam_penyerahan) values (?,'0000-00-00','00:00:00',?,?,?,?,?,'0000-00-00','00:00:00')",
		x.NoResep, x.NoRawat, x.KdDokter, x.TglPeresepan, x.JamPeresepan, x.Status)
	if err != nil {
		return err
	}
	return r.ReplaceItems(tx, x.NoResep, x.Items)
}

func (r *repository) ReplaceItems(tx contractsorm.Query, noResep string, items []model.Item) error {
	if _, err := tx.Exec("delete from resep_dokter where no_resep=?", noResep); err != nil {
		return err
	}
	for _, it := range items {
		if _, err := tx.Exec("insert into resep_dokter (no_resep,kode_brng,jml,aturan_pakai) values (?,?,?,?)",
			noResep, it.KodeBrng, it.Jml, it.AturanPakai); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) Delete(tx contractsorm.Query, noResep string) error {
	if _, err := tx.Exec("delete from resep_dokter where no_resep=?", noResep); err != nil {
		return err
	}
	_, err := tx.Exec("delete from resep_obat where no_resep=?", noResep)
	return err
}
