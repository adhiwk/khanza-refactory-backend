package rekammedis

import (
	"strings"

	"goravel/app/support"
)

// Kunjungan satu registrasi pasien pada riwayat rekam medis.
type Kunjungan struct {
	NoRawat       string `gorm:"column:no_rawat" json:"no_rawat"`
	TglRegistrasi string `gorm:"column:tgl_registrasi" json:"tgl_registrasi"`
	JamReg        string `gorm:"column:jam_reg" json:"jam_reg"`
	KdPoli        string `gorm:"column:kd_poli" json:"kd_poli"`
	NmPoli        string `gorm:"column:nm_poli" json:"nm_poli"`
	KdDokter      string `gorm:"column:kd_dokter" json:"kd_dokter"`
	NmDokter      string `gorm:"column:nm_dokter" json:"nm_dokter"`
	PngJawab      string `gorm:"column:png_jawab" json:"png_jawab"`
	StatusLanjut  string `gorm:"column:status_lanjut" json:"status_lanjut"`
	Stts          string `gorm:"column:stts" json:"stts"`
}

// Pasien identitas ringkas pemilik rekam medis.
type Pasien struct {
	NoRkmMedis string `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	NmPasien   string `gorm:"column:nm_pasien" json:"nm_pasien"`
	Jk         string `gorm:"column:jk" json:"jk"`
	TglLahir   string `gorm:"column:tgl_lahir" json:"tgl_lahir"`
	Alamat     string `gorm:"column:alamat" json:"alamat"`
	GolDarah   string `gorm:"column:gol_darah" json:"gol_darah"`
	Catatan    string `gorm:"column:catatan" json:"catatan"`
}

// RiwayatStore query riwayat rekam medis lintas modul.
type RiwayatStore interface {
	Pasien(noRkmMedis string) (*Pasien, error)
	Kunjungan(noRkmMedis, tglAwal, tglAkhir string, page, limit int) ([]Kunjungan, int64, error)
	// FormTerisi slug form asesmen yang memiliki data untuk no_rawat (tables: slug -> tabel).
	FormTerisi(noRawat string, tables [][2]string) ([]string, error)
}

type riwayatRepository struct{}

func NewRiwayatRepository() RiwayatStore {
	return &riwayatRepository{}
}

func (r *riwayatRepository) Pasien(noRkmMedis string) (*Pasien, error) {
	list := []Pasien{}
	err := support.DB().Raw("select pasien.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien,ifnull(pasien.jk,'') as jk,"+
		"ifnull(cast(pasien.tgl_lahir as char),'') as tgl_lahir,ifnull(pasien.alamat,'') as alamat,ifnull(pasien.gol_darah,'') as gol_darah,"+
		"ifnull(catatan_pasien.catatan,'') as catatan from pasien left join catatan_pasien on pasien.no_rkm_medis=catatan_pasien.no_rkm_medis "+
		"where pasien.no_rkm_medis=?", noRkmMedis).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *riwayatRepository) Kunjungan(noRkmMedis, tglAwal, tglAkhir string, page, limit int) ([]Kunjungan, int64, error) {
	q := support.DB().Table("reg_periksa").
		Join("left join poliklinik on reg_periksa.kd_poli=poliklinik.kd_poli").
		Join("left join dokter on reg_periksa.kd_dokter=dokter.kd_dokter").
		Join("left join penjab on reg_periksa.kd_pj=penjab.kd_pj").
		Where("reg_periksa.no_rkm_medis = ?", noRkmMedis)
	if tglAwal != "" && tglAkhir != "" {
		q = q.Where("reg_periksa.tgl_registrasi between ? and ?", tglAwal, tglAkhir)
	}
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []Kunjungan{}
	err = q.Select("reg_periksa.no_rawat", "cast(reg_periksa.tgl_registrasi as char) as tgl_registrasi", "cast(reg_periksa.jam_reg as char) as jam_reg",
		"ifnull(reg_periksa.kd_poli,'') as kd_poli", "ifnull(poliklinik.nm_poli,'') as nm_poli", "ifnull(reg_periksa.kd_dokter,'') as kd_dokter",
		"ifnull(dokter.nm_dokter,'') as nm_dokter", "ifnull(penjab.png_jawab,'') as png_jawab", "reg_periksa.status_lanjut",
		"ifnull(reg_periksa.stts,'') as stts").
		Order("reg_periksa.tgl_registrasi desc, reg_periksa.jam_reg desc").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

func (r *riwayatRepository) FormTerisi(noRawat string, tables [][2]string) ([]string, error) {
	if len(tables) == 0 {
		return []string{}, nil
	}
	parts := make([]string, len(tables))
	args := make([]any, len(tables))
	for i, t := range tables {
		// slug & tabel berasal dari daftar form hasil generator, bukan input pengguna.
		parts[i] = "select '" + t[0] + "' as slug from dual where exists (select 1 from " + t[1] + " where no_rawat=?)"
		args[i] = noRawat
	}
	var list []struct {
		Slug string `gorm:"column:slug"`
	}
	if err := support.DB().Raw(strings.Join(parts, " union all "), args...).Scan(&list); err != nil {
		return nil, err
	}
	out := make([]string, len(list))
	for i, l := range list {
		out[i] = l.Slug
	}
	return out, nil
}
