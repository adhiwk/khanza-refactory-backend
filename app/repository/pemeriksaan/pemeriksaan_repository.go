package pemeriksaan

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/perawatan"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

var tabel = map[model.Rawat]string{model.Ralan: "pemeriksaan_ralan", model.Ranap: "pemeriksaan_ranap"}

type Repository interface {
	perawatanrepo.KonteksRepository
	List(rawat model.Rawat, noRawat string) ([]model.Pemeriksaan, error)
	Find(rawat model.Rawat, noRawat, tgl, jam string) (*model.Pemeriksaan, error)
	Create(rawat model.Rawat, p *model.Pemeriksaan) error
	Update(rawat model.Rawat, p *model.Pemeriksaan) error
	Delete(rawat model.Rawat, noRawat, tgl, jam string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

var kolom = []string{"suhu_tubuh", "tensi", "nadi", "respirasi", "tinggi", "berat", "spo2", "gcs", "kesadaran",
	"keluhan", "pemeriksaan", "alergi", "rtl", "penilaian", "instruksi", "evaluasi", "nip"}

func (r *repository) query(rawat model.Rawat) contractsorm.Query {
	t := tabel[rawat]
	cols := []string{t + ".no_rawat", "date_format(" + t + ".tgl_perawatan,'%Y-%m-%d') as tgl_perawatan",
		"time_format(" + t + ".jam_rawat,'%H:%i:%s') as jam_rawat", "ifnull(pegawai.nama,'') as nm_pegawai"}
	for _, c := range kolom {
		cols = append(cols, t+"."+c)
	}
	if rawat == model.Ralan {
		cols = append(cols, t+".lingkar_perut")
	}
	return support.DB().Table(t).Select(cols...).Join("left join pegawai on " + t + ".nip=pegawai.nik")
}

func (r *repository) List(rawat model.Rawat, noRawat string) ([]model.Pemeriksaan, error) {
	t := tabel[rawat]
	list := []model.Pemeriksaan{}
	err := r.query(rawat).Where(t+".no_rawat = ?", noRawat).Order(t + ".tgl_perawatan, " + t + ".jam_rawat").Scan(&list)
	return list, err
}

// Find (nil, nil) bila tidak ada.
func (r *repository) Find(rawat model.Rawat, noRawat, tgl, jam string) (*model.Pemeriksaan, error) {
	t := tabel[rawat]
	list := []model.Pemeriksaan{}
	err := r.query(rawat).Where(t+".no_rawat = ? and "+t+".tgl_perawatan = ? and "+t+".jam_rawat = ?", noRawat, tgl, jam).Limit(1).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func values(rawat model.Rawat, p *model.Pemeriksaan) map[string]any {
	v := map[string]any{
		"suhu_tubuh": p.SuhuTubuh, "tensi": p.Tensi, "nadi": p.Nadi, "respirasi": p.Respirasi,
		"tinggi": p.Tinggi, "berat": p.Berat, "spo2": p.Spo2, "gcs": p.Gcs, "kesadaran": p.Kesadaran,
		"keluhan": p.Keluhan, "pemeriksaan": p.Pemeriksaan, "alergi": p.Alergi, "rtl": p.Rtl,
		"penilaian": p.Penilaian, "instruksi": p.Instruksi, "evaluasi": p.Evaluasi, "nip": p.Nip,
	}
	if rawat == model.Ralan {
		v["lingkar_perut"] = p.LingkarPerut
	}
	return v
}

func (r *repository) Create(rawat model.Rawat, p *model.Pemeriksaan) error {
	v := values(rawat, p)
	v["no_rawat"], v["tgl_perawatan"], v["jam_rawat"] = p.NoRawat, p.TglPerawatan, p.JamRawat
	return support.DB().Table(tabel[rawat]).Create(v)
}

func (r *repository) Update(rawat model.Rawat, p *model.Pemeriksaan) error {
	_, err := support.DB().Table(tabel[rawat]).
		Where("no_rawat = ? and tgl_perawatan = ? and jam_rawat = ?", p.NoRawat, p.TglPerawatan, p.JamRawat).
		Update(values(rawat, p))
	return err
}

func (r *repository) Delete(rawat model.Rawat, noRawat, tgl, jam string) error {
	_, err := support.DB().Exec("delete from "+tabel[rawat]+" where no_rawat=? and tgl_perawatan=? and jam_rawat=?", noRawat, tgl, jam)
	return err
}
