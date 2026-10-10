package kamarinap

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/kamarinap"
	"goravel/app/repository/crud"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

// TanpaTanggal nilai tgl_keluar/jam_keluar Khanza untuk pasien yang masih dirawat.
const (
	TanpaTanggal = "0000-00-00"
	TanpaJam     = "00:00:00"
)

type Filter struct {
	Status    string // dirawat | pulang | semua
	TglAwal   string
	TglAkhir  string
	KdBangsal string
	Search    string
	Page      int
	Limit     int
}

type Repository interface {
	perawatanrepo.KonteksRepository
	Paginate(f Filter) ([]model.KamarInap, int64, error)
	Riwayat(noRawat string) ([]model.KamarInap, error)
	Pengaturan() (model.Pengaturan, error)
	PasienDirawat(noRkmMedis string) (bool, error)
	LockKamar(tx contractsorm.Query, kdKamar string) (*model.Kamar, error)
	Aktif(tx contractsorm.Query, noRawat string) (*model.KamarInap, error)
	Terakhir(tx contractsorm.Query, noRawat string) (*model.KamarInap, error)
	Insert(tx contractsorm.Query, noRawat, kdKamar string, trf float64, diagnosaAwal, diagnosaAkhir, tgl, jam string) error
	Tutup(tx contractsorm.Query, key model.Key, trf float64, tgl, jam string, lama, biaya float64, stts, diagnosaAkhir string) error
	GantiKamar(tx contractsorm.Query, key model.Key, kdKamar string, trf float64) error
	BukaKembali(tx contractsorm.Query, key model.Key) error
	Hapus(tx contractsorm.Query, key model.Key) error
	SetStatusKamar(tx contractsorm.Query, kdKamar, status string) error
	SetStatusLanjut(tx contractsorm.Query, noRawat, status string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

const viewColumns = "kamar_inap.no_rawat,reg_periksa.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien,kamar_inap.kd_kamar," +
	"ifnull(kamar.kd_bangsal,'') as kd_bangsal,ifnull(bangsal.nm_bangsal,'') as nm_bangsal,ifnull(kamar.kelas,'') as kelas," +
	"ifnull(penjab.png_jawab,'') as png_jawab,ifnull(kamar_inap.trf_kamar,0) as trf_kamar,ifnull(kamar_inap.diagnosa_awal,'') as diagnosa_awal," +
	"ifnull(kamar_inap.diagnosa_akhir,'') as diagnosa_akhir,cast(kamar_inap.tgl_masuk as char) as tgl_masuk,cast(kamar_inap.jam_masuk as char) as jam_masuk," +
	"ifnull(cast(kamar_inap.tgl_keluar as char),'') as tgl_keluar,ifnull(cast(kamar_inap.jam_keluar as char),'') as jam_keluar," +
	"ifnull(kamar_inap.lama,0) as lama,ifnull(kamar_inap.ttl_biaya,0) as ttl_biaya,kamar_inap.stts_pulang"

func (r *repository) view(q contractsorm.Query) contractsorm.Query {
	return q.Table("kamar_inap").
		Join("inner join reg_periksa on kamar_inap.no_rawat=reg_periksa.no_rawat").
		Join("inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis").
		Join("inner join kamar on kamar_inap.kd_kamar=kamar.kd_kamar").
		Join("inner join bangsal on kamar.kd_bangsal=bangsal.kd_bangsal").
		Join("inner join penjab on reg_periksa.kd_pj=penjab.kd_pj")
}

// Paginate daftar pasien: dirawat (stts_pulang '-'), pulang (tgl_keluar pada rentang), atau semua (tgl_masuk pada rentang).
func (r *repository) Paginate(f Filter) ([]model.KamarInap, int64, error) {
	q := r.view(support.DB())
	switch f.Status {
	case "pulang":
		q = q.Where("kamar_inap.stts_pulang not in ('-','Pindah Kamar') and kamar_inap.tgl_keluar between ? and ?", f.TglAwal, f.TglAkhir)
	case "semua":
		q = q.Where("kamar_inap.tgl_masuk between ? and ?", f.TglAwal, f.TglAkhir)
	default:
		q = q.Where("kamar_inap.stts_pulang = '-'")
	}
	if f.KdBangsal != "" {
		q = q.Where("kamar.kd_bangsal = ?", f.KdBangsal)
	}
	q = crud.WhereLike(q, f.Search, "kamar_inap.no_rawat", "reg_periksa.no_rkm_medis", "pasien.nm_pasien", "kamar_inap.kd_kamar",
		"bangsal.nm_bangsal", "kamar_inap.diagnosa_awal", "penjab.png_jawab")

	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.KamarInap{}
	err = q.Select(viewColumns).Order("bangsal.nm_bangsal, kamar_inap.tgl_masuk, kamar_inap.jam_masuk").
		Offset((f.Page - 1) * f.Limit).Limit(f.Limit).Scan(&list)
	return list, total, err
}

func (r *repository) Riwayat(noRawat string) ([]model.KamarInap, error) {
	list := []model.KamarInap{}
	err := r.view(support.DB()).Where("kamar_inap.no_rawat = ?", noRawat).Select(viewColumns).
		Order("kamar_inap.tgl_masuk, kamar_inap.jam_masuk").Scan(&list)
	return list, err
}

func (r *repository) Pengaturan() (model.Pengaturan, error) {
	var list []struct {
		Lamajam  float64 `gorm:"column:lamajam"`
		Hariawal string  `gorm:"column:hariawal"`
	}
	if err := support.DB().Raw("select lamajam,hariawal from set_jam_minimal limit 1").Scan(&list); err != nil {
		return model.Pengaturan{}, err
	}
	if len(list) == 0 {
		return model.Pengaturan{}, nil
	}
	return model.Pengaturan{JamMinimal: list[0].Lamajam, HitungHariAwal: list[0].Hariawal == "Yes"}, nil
}

func (r *repository) PasienDirawat(noRkmMedis string) (bool, error) {
	return support.DB().Table("kamar_inap").
		Join("inner join reg_periksa on kamar_inap.no_rawat=reg_periksa.no_rawat").
		Where("kamar_inap.stts_pulang = '-' and reg_periksa.no_rkm_medis = ?", noRkmMedis).Exists()
}

// LockKamar mengunci baris kamar (for update) agar satu bed tidak diisi dua pasien; (nil, nil) bila tidak ada.
func (r *repository) LockKamar(tx contractsorm.Query, kdKamar string) (*model.Kamar, error) {
	list := []model.Kamar{}
	err := tx.Raw("select kd_kamar,ifnull(trf_kamar,0) as trf_kamar,ifnull(status,'') as status,ifnull(statusdata,'') as statusdata "+
		"from kamar where kd_kamar=? for update", kdKamar).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *repository) first(tx contractsorm.Query, where string, args ...any) (*model.KamarInap, error) {
	list := []model.KamarInap{}
	err := r.view(tx).Where(where, args...).Select(viewColumns).
		Order("kamar_inap.tgl_masuk desc, kamar_inap.jam_masuk desc").Limit(1).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

// Aktif baris kamar yang sedang ditempati (stts_pulang '-'); (nil, nil) bila tidak ada.
func (r *repository) Aktif(tx contractsorm.Query, noRawat string) (*model.KamarInap, error) {
	return r.first(tx, "kamar_inap.no_rawat = ? and kamar_inap.stts_pulang = '-'", noRawat)
}

// Terakhir baris kamar terakhir (untuk batal pulang); (nil, nil) bila tidak ada.
func (r *repository) Terakhir(tx contractsorm.Query, noRawat string) (*model.KamarInap, error) {
	return r.first(tx, "kamar_inap.no_rawat = ?", noRawat)
}

func (r *repository) Insert(tx contractsorm.Query, noRawat, kdKamar string, trf float64, diagnosaAwal, diagnosaAkhir, tgl, jam string) error {
	_, err := tx.Exec("insert into kamar_inap (no_rawat,kd_kamar,trf_kamar,diagnosa_awal,diagnosa_akhir,tgl_masuk,jam_masuk,"+
		"tgl_keluar,jam_keluar,lama,ttl_biaya,stts_pulang) values (?,?,?,?,?,?,?,?,?,0,0,'-')",
		noRawat, kdKamar, trf, diagnosaAwal, diagnosaAkhir, tgl, jam, TanpaTanggal, TanpaJam)
	return err
}

const keyWhere = " where no_rawat=? and tgl_masuk=? and jam_masuk=?"

func (r *repository) Tutup(tx contractsorm.Query, key model.Key, trf float64, tgl, jam string, lama, biaya float64, stts, diagnosaAkhir string) error {
	_, err := tx.Exec("update kamar_inap set trf_kamar=?,tgl_keluar=?,jam_keluar=?,lama=?,ttl_biaya=?,stts_pulang=?,diagnosa_akhir=?"+keyWhere,
		trf, tgl, jam, lama, biaya, stts, diagnosaAkhir, key.NoRawat, key.TglMasuk, key.JamMasuk)
	return err
}

func (r *repository) GantiKamar(tx contractsorm.Query, key model.Key, kdKamar string, trf float64) error {
	_, err := tx.Exec("update kamar_inap set kd_kamar=?,trf_kamar=?"+keyWhere, kdKamar, trf, key.NoRawat, key.TglMasuk, key.JamMasuk)
	return err
}

func (r *repository) BukaKembali(tx contractsorm.Query, key model.Key) error {
	_, err := tx.Exec("update kamar_inap set stts_pulang='-',tgl_keluar=?,jam_keluar=?,lama=0,ttl_biaya=0"+keyWhere,
		TanpaTanggal, TanpaJam, key.NoRawat, key.TglMasuk, key.JamMasuk)
	return err
}

func (r *repository) Hapus(tx contractsorm.Query, key model.Key) error {
	_, err := tx.Exec("delete from kamar_inap"+keyWhere, key.NoRawat, key.TglMasuk, key.JamMasuk)
	return err
}

func (r *repository) SetStatusKamar(tx contractsorm.Query, kdKamar, status string) error {
	_, err := tx.Exec("update kamar set status=? where kd_kamar=?", status, kdKamar)
	return err
}

func (r *repository) SetStatusLanjut(tx contractsorm.Query, noRawat, status string) error {
	_, err := tx.Exec("update reg_periksa set status_lanjut=? where no_rawat=?", status, noRawat)
	return err
}
