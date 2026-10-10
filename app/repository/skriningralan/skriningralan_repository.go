package skriningralan

import (
	model "goravel/app/models/skriningralan"
	"goravel/app/repository/crud"
	"goravel/app/support"
)

type Filter struct {
	NoRkmMedis string
	TglAwal    string
	TglAkhir   string
	Search     string
}

type Repository interface {
	PasienExists(noRkmMedis string) (bool, error)
	PetugasExists(nip string) (bool, error)
	Paginate(f Filter, page, limit int) ([]model.Skrining, int64, error)
	Find(key model.Key) (*model.Skrining, error)
	Create(s model.Skrining) error
	Update(s model.Skrining) error
	Delete(key model.Key) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) PasienExists(noRkmMedis string) (bool, error) {
	return crud.ExistsIn(support.DB(), "pasien", "no_rkm_medis", noRkmMedis)
}

func (r *repository) PetugasExists(nip string) (bool, error) {
	return crud.ExistsIn(support.DB(), "petugas", "nip", nip)
}

const selectSQL = "select cast(s.tanggal as char) as tanggal,cast(s.jam as char) as jam,s.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien," +
	"ifnull(s.geriatri,'') as geriatri,ifnull(s.kesadaran,'') as kesadaran,ifnull(s.pernapasan,'') as pernapasan,ifnull(s.nyeri_dada,'') as nyeri_dada," +
	"ifnull(s.skala_nyeri,'') as skala_nyeri,s.batuk,s.risiko_jatuh,ifnull(s.keputusan,'') as keputusan,ifnull(s.nip,'') as nip," +
	"ifnull(petugas.nama,'') as nm_petugas from skrining_rawat_jalan s inner join pasien on s.no_rkm_medis=pasien.no_rkm_medis " +
	"left join petugas on s.nip=petugas.nip"

func (r *repository) Paginate(f Filter, page, limit int) ([]model.Skrining, int64, error) {
	q := support.DB().Table("(" + selectSQL + ") x")
	if f.NoRkmMedis != "" {
		q = q.Where("no_rkm_medis = ?", f.NoRkmMedis)
	}
	if f.TglAwal != "" && f.TglAkhir != "" {
		q = q.Where("tanggal between ? and ?", f.TglAwal, f.TglAkhir)
	}
	q = crud.WhereLike(q, f.Search, "no_rkm_medis", "nm_pasien", "keputusan", "nm_petugas")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Skrining{}
	err = q.Order("tanggal desc, jam desc").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

// Find (nil, nil) bila tidak ada.
func (r *repository) Find(key model.Key) (*model.Skrining, error) {
	list := []model.Skrining{}
	err := support.DB().Raw(selectSQL+" where s.tanggal=? and s.jam=? and s.no_rkm_medis=?", key.Tanggal, key.Jam, key.NoRkmMedis).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *repository) Create(s model.Skrining) error {
	_, err := support.DB().Exec("insert into skrining_rawat_jalan (tanggal,jam,no_rkm_medis,geriatri,kesadaran,pernapasan,nyeri_dada,skala_nyeri,"+
		"batuk,risiko_jatuh,keputusan,nip) values (?,?,?,?,?,?,?,?,?,?,?,?)", s.Tanggal, s.Jam, s.NoRkmMedis, s.Geriatri, s.Kesadaran, s.Pernapasan,
		s.NyeriDada, s.SkalaNyeri, s.Batuk, s.RisikoJatuh, s.Keputusan, s.Nip)
	return err
}

func (r *repository) Update(s model.Skrining) error {
	_, err := support.DB().Exec("update skrining_rawat_jalan set geriatri=?,kesadaran=?,pernapasan=?,nyeri_dada=?,skala_nyeri=?,batuk=?,"+
		"risiko_jatuh=?,keputusan=?,nip=? where tanggal=? and jam=? and no_rkm_medis=?", s.Geriatri, s.Kesadaran, s.Pernapasan, s.NyeriDada,
		s.SkalaNyeri, s.Batuk, s.RisikoJatuh, s.Keputusan, s.Nip, s.Tanggal, s.Jam, s.NoRkmMedis)
	return err
}

func (r *repository) Delete(key model.Key) error {
	_, err := support.DB().Exec("delete from skrining_rawat_jalan where tanggal=? and jam=? and no_rkm_medis=?", key.Tanggal, key.Jam, key.NoRkmMedis)
	return err
}
