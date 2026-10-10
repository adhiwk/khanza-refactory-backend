package rekonsiliasi

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/rekonsiliasi"
	"goravel/app/repository/crud"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

type Filter struct {
	NoRawat    string
	NoRkmMedis string
	TglAwal    string
	TglAkhir   string
	Search     string
}

type Repository interface {
	perawatanrepo.KonteksRepository
	PetugasExists(nip string) (bool, error)
	Paginate(f Filter, page, limit int) ([]model.Rekonsiliasi, int64, error)
	Find(noRekonsiliasi string) (*model.Rekonsiliasi, error)
	NextNomor(tx contractsorm.Query, tanggal string) (string, error)
	Insert(tx contractsorm.Query, r model.Rekonsiliasi) error
	Update(tx contractsorm.Query, r model.Rekonsiliasi) error
	ReplaceObat(tx contractsorm.Query, noRekonsiliasi string, obat []model.Obat) error
	Delete(tx contractsorm.Query, noRekonsiliasi string) error
	SimpanKonfirmasi(tx contractsorm.Query, noRekonsiliasi string, k model.Konfirmasi) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) PetugasExists(nip string) (bool, error) {
	return crud.ExistsIn(support.DB(), "petugas", "nip", nip)
}

const headerSQL = "select r.no_rekonsiliasi,r.no_rawat,reg_periksa.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien," +
	"cast(r.tanggal_wawancara as char) as tanggal_wawancara,ifnull(r.rekonsiliasi_obat_saat,'') as rekonsiliasi_obat_saat," +
	"ifnull(r.alergi_obat,'') as alergi_obat,ifnull(r.manifestasi_alergi,'') as manifestasi_alergi,ifnull(r.dampak_alergi,'') as dampak_alergi," +
	"ifnull(r.nip,'') as nip,ifnull(petugas.nama,'') as nm_petugas," +
	"exists(select 1 from rekonsiliasi_obat_konfirmasi k where k.no_rekonsiliasi=r.no_rekonsiliasi) as dikonfirmasi " +
	"from rekonsiliasi_obat r inner join reg_periksa on r.no_rawat=reg_periksa.no_rawat " +
	"inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis left join petugas on r.nip=petugas.nip"

func (r *repository) Paginate(f Filter, page, limit int) ([]model.Rekonsiliasi, int64, error) {
	q := support.DB().Table("(" + headerSQL + ") x")
	if f.NoRawat != "" {
		q = q.Where("no_rawat = ?", f.NoRawat)
	}
	if f.NoRkmMedis != "" {
		q = q.Where("no_rkm_medis = ?", f.NoRkmMedis)
	}
	if f.TglAwal != "" && f.TglAkhir != "" {
		q = q.Where("date(tanggal_wawancara) between ? and ?", f.TglAwal, f.TglAkhir)
	}
	q = crud.WhereLike(q, f.Search, "no_rekonsiliasi", "no_rawat", "no_rkm_medis", "nm_pasien", "nm_petugas")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Rekonsiliasi{}
	err = q.Order("tanggal_wawancara desc").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

// Find (nil, nil) bila tidak ada.
func (r *repository) Find(noRekonsiliasi string) (*model.Rekonsiliasi, error) {
	list := []model.Rekonsiliasi{}
	if err := support.DB().Raw(headerSQL+" where r.no_rekonsiliasi=?", noRekonsiliasi).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}
	x := &list[0]
	x.Obat = []model.Obat{}
	if err := support.DB().Raw("select ifnull(nama_obat,'') as nama_obat,ifnull(dosis_obat,'') as dosis_obat,ifnull(frekuensi,'') as frekuensi,"+
		"ifnull(cara_pemberian,'') as cara_pemberian,ifnull(waktu_pemberian_terakhir,'') as waktu_pemberian_terakhir,"+
		"ifnull(tindak_lanjut,'') as tindak_lanjut,ifnull(perubahan_aturan_pakai,'') as perubahan_aturan_pakai "+
		"from rekonsiliasi_obat_detail_obat where no_rekonsiliasi=?", noRekonsiliasi).Scan(&x.Obat); err != nil {
		return nil, err
	}
	var k []model.Konfirmasi
	if err := support.DB().Raw("select ifnull(cast(k.diterima_farmasi as char),'') as diterima_farmasi,"+
		"ifnull(cast(k.dikonfirmasi_apoteker as char),'') as dikonfirmasi_apoteker,ifnull(cast(k.diserahkan_pasien as char),'') as diserahkan_pasien,"+
		"ifnull(k.nip,'') as nip,ifnull(petugas.nama,'') as nm_petugas from rekonsiliasi_obat_konfirmasi k "+
		"left join petugas on k.nip=petugas.nip where k.no_rekonsiliasi=?", noRekonsiliasi).Scan(&k); err != nil {
		return nil, err
	}
	if len(k) > 0 {
		x.Konfirmasi = &k[0]
	}
	return x, nil
}

// NextNomor RO + yyyyMMdd + 4 digit per tanggal wawancara (RMRekonsiliasiObat.autoNomor).
func (r *repository) NextNomor(tx contractsorm.Query, tanggal string) (string, error) {
	return crud.NextCode(tx, "rekonsiliasi_obat", "no_rekonsiliasi", "RO"+tanggal, 4)
}

func (r *repository) Insert(tx contractsorm.Query, x model.Rekonsiliasi) error {
	_, err := tx.Exec("insert into rekonsiliasi_obat (no_rekonsiliasi,no_rawat,tanggal_wawancara,rekonsiliasi_obat_saat,alergi_obat,"+
		"manifestasi_alergi,dampak_alergi,nip) values (?,?,?,?,?,?,?,?)", x.NoRekonsiliasi, x.NoRawat, x.TanggalWawancara,
		x.RekonsiliasiObatSaat, x.AlergiObat, x.ManifestasiAlergi, x.DampakAlergi, x.Nip)
	return err
}

func (r *repository) Update(tx contractsorm.Query, x model.Rekonsiliasi) error {
	_, err := tx.Exec("update rekonsiliasi_obat set tanggal_wawancara=?,rekonsiliasi_obat_saat=?,alergi_obat=?,manifestasi_alergi=?,"+
		"dampak_alergi=?,nip=? where no_rekonsiliasi=?", x.TanggalWawancara, x.RekonsiliasiObatSaat, x.AlergiObat, x.ManifestasiAlergi,
		x.DampakAlergi, x.Nip, x.NoRekonsiliasi)
	return err
}

func (r *repository) ReplaceObat(tx contractsorm.Query, noRekonsiliasi string, obat []model.Obat) error {
	if _, err := tx.Exec("delete from rekonsiliasi_obat_detail_obat where no_rekonsiliasi=?", noRekonsiliasi); err != nil {
		return err
	}
	for _, o := range obat {
		if _, err := tx.Exec("insert into rekonsiliasi_obat_detail_obat (no_rekonsiliasi,nama_obat,dosis_obat,frekuensi,cara_pemberian,"+
			"waktu_pemberian_terakhir,tindak_lanjut,perubahan_aturan_pakai) values (?,?,?,?,?,?,?,?)", noRekonsiliasi, o.NamaObat,
			o.DosisObat, o.Frekuensi, o.CaraPemberian, o.WaktuPemberianTerakhir, o.TindakLanjut, o.PerubahanAturanPakai); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) Delete(tx contractsorm.Query, noRekonsiliasi string) error {
	for _, t := range []string{"rekonsiliasi_obat_detail_obat", "rekonsiliasi_obat_konfirmasi", "rekonsiliasi_obat"} {
		if _, err := tx.Exec("delete from "+t+" where no_rekonsiliasi=?", noRekonsiliasi); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) SimpanKonfirmasi(tx contractsorm.Query, noRekonsiliasi string, k model.Konfirmasi) error {
	if _, err := tx.Exec("delete from rekonsiliasi_obat_konfirmasi where no_rekonsiliasi=?", noRekonsiliasi); err != nil {
		return err
	}
	_, err := tx.Exec("insert into rekonsiliasi_obat_konfirmasi (no_rekonsiliasi,diterima_farmasi,dikonfirmasi_apoteker,nip,diserahkan_pasien) "+
		"values (?,?,?,?,?)", noRekonsiliasi, k.DiterimaFarmasi, k.DikonfirmasiApoteker, k.Nip, k.DiserahkanPasien)
	return err
}
