package registrasi

import (
	"fmt"
	"strconv"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"

	regmodel "goravel/app/models/registrasi"
)

const connection = "mysql_kedua"

// Filter untuk list registrasi
type Filter struct {
	TglAwal  string // YYYY-MM-DD
	TglAkhir string // YYYY-MM-DD
	KdDokter string
	KdPoli   string
	Search   string
	Page     int
	Limit    int
}

// PasienInfo data pasien yang dipakai saat registrasi (isCekPasien di DlgReg).
type PasienInfo struct {
	NoRkmMedis   string     `gorm:"column:no_rkm_medis"`
	TglLahir     *time.Time `gorm:"column:tgl_lahir"`
	TglDaftar    *time.Time `gorm:"column:tgl_daftar"`
	Asal         string     `gorm:"column:asal"`
	NamaKeluarga string     `gorm:"column:namakeluarga"`
	Keluarga     string     `gorm:"column:keluarga"`
	KdPj         string     `gorm:"column:kd_pj"`
}

type Repository interface {
	Paginate(filter Filter) ([]regmodel.RegPeriksaView, int64, error)
	FindByNoRawat(noRawat string) (*regmodel.RegPeriksa, error)
	FindPasien(noRkmMedis string) (*PasienInfo, error)
	IsDirawatInap(noRkmMedis string) (bool, error)
	PernahKePoli(noRkmMedis, kdPoli string) (bool, error)
	DokterExists(kdDokter string) (bool, error)
	PenjabExists(kdPj string) (bool, error)
	BiayaPoli(kdPoli string, baru bool) (float64, bool, error)
	HasBilling(noRawat string) (bool, error)
	NextNoReg(urut, kdDokter, kdPoli string, tgl time.Time) (string, error)
	NextNoRawat(tgl time.Time) (string, error)
	Create(data *regmodel.RegPeriksa) error
	Save(data *regmodel.RegPeriksa) error
	Delete(noRawat string) error
	UpdateUmurPasien(noRkmMedis string) error
	EnsurePoli(kdPoli, nmPoli string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) query() contractsorm.Query {
	return facades.Orm().Connection(connection).Query()
}

func (r *repository) Paginate(filter Filter) ([]regmodel.RegPeriksaView, int64, error) {
	q := r.query().Table("reg_periksa").
		Join("inner join dokter on reg_periksa.kd_dokter=dokter.kd_dokter").
		Join("inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis").
		Join("inner join poliklinik on reg_periksa.kd_poli=poliklinik.kd_poli").
		Join("inner join penjab on reg_periksa.kd_pj=penjab.kd_pj").
		Where("reg_periksa.tgl_registrasi between ? and ?", filter.TglAwal, filter.TglAkhir)

	if filter.KdDokter != "" {
		q = q.Where("reg_periksa.kd_dokter = ?", filter.KdDokter)
	}
	if filter.KdPoli != "" {
		q = q.Where("reg_periksa.kd_poli = ?", filter.KdPoli)
	}
	if filter.Search != "" {
		like := "%" + filter.Search + "%"
		q = q.Where("(reg_periksa.no_reg like ? or reg_periksa.no_rawat like ? or reg_periksa.no_rkm_medis like ? or "+
			"pasien.nm_pasien like ? or dokter.nm_dokter like ? or poliklinik.nm_poli like ? or penjab.png_jawab like ? or reg_periksa.p_jawab like ?)",
			like, like, like, like, like, like, like, like)
	}

	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	list := []regmodel.RegPeriksaView{}
	err = q.Select("reg_periksa.*", "dokter.nm_dokter", "pasien.nm_pasien", "pasien.jk", "pasien.no_tlp", "poliklinik.nm_poli", "penjab.png_jawab").
		Order("reg_periksa.tgl_registrasi desc, reg_periksa.jam_reg desc").
		Offset((filter.Page - 1) * filter.Limit).Limit(filter.Limit).
		Scan(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// FindByNoRawat mengembalikan (nil, nil) bila data tidak ditemukan
func (r *repository) FindByNoRawat(noRawat string) (*regmodel.RegPeriksa, error) {
	var data regmodel.RegPeriksa
	if err := r.query().Where("no_rawat = ?", noRawat).First(&data); err != nil {
		return nil, err
	}
	if data.NoRawat == "" {
		return nil, nil
	}
	return &data, nil
}

// FindPasien mengembalikan (nil, nil) bila pasien tidak ditemukan
func (r *repository) FindPasien(noRkmMedis string) (*PasienInfo, error) {
	list := []PasienInfo{}
	err := r.query().Raw("select pasien.no_rkm_medis,pasien.tgl_lahir,pasien.tgl_daftar,"+
		"concat_ws(', ',pasien.alamat,kelurahan.nm_kel,kecamatan.nm_kec,kabupaten.nm_kab) as asal,"+
		"pasien.namakeluarga,ifnull(pasien.keluarga,'') as keluarga,pasien.kd_pj from pasien "+
		"left join kelurahan on pasien.kd_kel=kelurahan.kd_kel "+
		"left join kecamatan on pasien.kd_kec=kecamatan.kd_kec "+
		"left join kabupaten on pasien.kd_kab=kabupaten.kd_kab "+
		"where pasien.no_rkm_medis=?", noRkmMedis).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *repository) IsDirawatInap(noRkmMedis string) (bool, error) {
	count, err := r.query().Table("reg_periksa").
		Join("inner join kamar_inap on reg_periksa.no_rawat=kamar_inap.no_rawat").
		Where("kamar_inap.stts_pulang = '-' and reg_periksa.no_rkm_medis = ?", noRkmMedis).
		Count()
	return count > 0, err
}

func (r *repository) PernahKePoli(noRkmMedis, kdPoli string) (bool, error) {
	count, err := r.query().Table("reg_periksa").Where("no_rkm_medis = ? and kd_poli = ?", noRkmMedis, kdPoli).Count()
	return count > 0, err
}

func (r *repository) DokterExists(kdDokter string) (bool, error) {
	count, err := r.query().Table("dokter").Where("kd_dokter = ?", kdDokter).Count()
	return count > 0, err
}

func (r *repository) PenjabExists(kdPj string) (bool, error) {
	count, err := r.query().Table("penjab").Where("kd_pj = ?", kdPj).Count()
	return count > 0, err
}

// BiayaPoli: tarif registrasi pasien baru (registrasi) atau lama (registrasilama); found=false bila poli tidak ada.
func (r *repository) BiayaPoli(kdPoli string, baru bool) (float64, bool, error) {
	column := "registrasilama"
	if baru {
		column = "registrasi"
	}
	var list []struct {
		Biaya float64 `gorm:"column:biaya"`
	}
	err := r.query().Raw("select ifnull("+column+",0) as biaya from poliklinik where kd_poli=?", kdPoli).Scan(&list)
	if err != nil || len(list) == 0 {
		return 0, false, err
	}
	return list[0].Biaya, true, nil
}

func (r *repository) HasBilling(noRawat string) (bool, error) {
	count, err := r.query().Table("billing").Where("no_rawat = ?", noRawat).Count()
	return count > 0, err
}

// NextNoReg nomor antrian 3 digit per tanggal, diurutkan sesuai URUTNOREG (isNumber di DlgReg).
func (r *repository) NextNoReg(urut, kdDokter, kdPoli string, tgl time.Time) (string, error) {
	q := r.query().Table("reg_periksa").Where("tgl_registrasi = ?", tgl.Format("2006-01-02"))
	switch urut {
	case "poli":
		q = q.Where("kd_poli = ?", kdPoli)
	case "dokter + poli":
		q = q.Where("kd_dokter = ? and kd_poli = ?", kdDokter, kdPoli)
	default:
		q = q.Where("kd_dokter = ?", kdDokter)
	}
	max, err := r.maxNumber(q, "no_reg")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%03d", max+1), nil
}

// NextNoRawat format yyyy/MM/dd/NNNNNN per tanggal registrasi.
func (r *repository) NextNoRawat(tgl time.Time) (string, error) {
	q := r.query().Table("reg_periksa").Where("tgl_registrasi = ?", tgl.Format("2006-01-02"))
	max, err := r.maxNumber(q, "right(no_rawat,6)")
	if err != nil {
		return "", err
	}
	return tgl.Format("2006/01/02") + "/" + fmt.Sprintf("%06d", max+1), nil
}

func (r *repository) maxNumber(q contractsorm.Query, expr string) (int, error) {
	var list []struct {
		Max string `gorm:"column:max"`
	}
	if err := q.Select("ifnull(max(convert(" + expr + ",signed)),0) as max").Scan(&list); err != nil {
		return 0, err
	}
	if len(list) == 0 {
		return 0, nil
	}
	n, _ := strconv.Atoi(list[0].Max)
	return n, nil
}

func (r *repository) Create(data *regmodel.RegPeriksa) error {
	return r.query().Create(data)
}

func (r *repository) Save(data *regmodel.RegPeriksa) error {
	return r.query().Save(data)
}

func (r *repository) Delete(noRawat string) error {
	_, err := r.query().Where("no_rawat = ?", noRawat).Delete(&regmodel.RegPeriksa{})
	return err
}

// UpdateUmurPasien menyegarkan pasien.umur ("X Th Y Bl Z Hr"), sama dengan UpdateUmur di DlgReg.
func (r *repository) UpdateUmurPasien(noRkmMedis string) error {
	_, err := r.query().Exec("update pasien set umur=CONCAT(CONCAT(CONCAT(TIMESTAMPDIFF(YEAR, tgl_lahir, CURDATE()), ' Th '),"+
		"CONCAT(TIMESTAMPDIFF(MONTH, tgl_lahir, CURDATE()) - ((TIMESTAMPDIFF(MONTH, tgl_lahir, CURDATE()) div 12) * 12), ' Bl ')),"+
		"CONCAT(TIMESTAMPDIFF(DAY, DATE_ADD(DATE_ADD(tgl_lahir,INTERVAL TIMESTAMPDIFF(YEAR, tgl_lahir, CURDATE()) YEAR), "+
		"INTERVAL TIMESTAMPDIFF(MONTH, tgl_lahir, CURDATE()) - ((TIMESTAMPDIFF(MONTH, tgl_lahir, CURDATE()) div 12) * 12) MONTH), CURDATE()), ' Hr')) "+
		"where no_rkm_medis=?", noRkmMedis)
	return err
}

// EnsurePoli membuat poliklinik bila belum ada (DlgIGD membuat unit IGDK otomatis).
func (r *repository) EnsurePoli(kdPoli, nmPoli string) error {
	_, err := r.query().Exec("insert ignore into poliklinik (kd_poli,nm_poli,registrasi,registrasilama,status) values (?,?,0,0,'1')", kdPoli, nmPoli)
	return err
}
