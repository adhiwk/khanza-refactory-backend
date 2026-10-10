package triase

import (
	"strconv"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/triase"
	"goravel/app/repository/crud"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

type Filter struct {
	TglAwal  string
	TglAkhir string
	Search   string
}

type Repository interface {
	perawatanrepo.KonteksRepository
	Paginate(f Filter, page, limit int) ([]model.Ringkas, int64, error)
	Find(noRawat string) (*model.Triase, error)
	Exists(noRawat string) (bool, error)
	KasusExists(kode string) (bool, error)
	SkalaExists(level int, kode string) (bool, error)
	Insert(tx contractsorm.Query, in model.Input) error
	Update(tx contractsorm.Query, in model.Input, jenisLama string) error
	Delete(tx contractsorm.Query, noRawat string) error
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

const ringkasSQL = "select t.no_rawat,reg_periksa.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien,cast(t.tgl_kunjungan as char) as tgl_kunjungan," +
	"t.kode_kasus,ifnull(k.macam_kasus,'') as macam_kasus,if(p.no_rawat is not null,'primer',if(s.no_rawat is not null,'sekunder','')) as jenis," +
	"ifnull(p.plan,ifnull(s.plan,'')) as plan,ifnull(p.nik,ifnull(s.nik,'')) as nik,ifnull(pegawai.nama,'') as nm_petugas " +
	"from data_triase_igd t inner join reg_periksa on t.no_rawat=reg_periksa.no_rawat " +
	"inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis " +
	"left join master_triase_macam_kasus k on t.kode_kasus=k.kode_kasus " +
	"left join data_triase_igdprimer p on t.no_rawat=p.no_rawat left join data_triase_igdsekunder s on t.no_rawat=s.no_rawat " +
	"left join pegawai on pegawai.nik=ifnull(p.nik,s.nik)"

func (r *repository) Paginate(f Filter, page, limit int) ([]model.Ringkas, int64, error) {
	q := support.DB().Table("(" + ringkasSQL + ") x")
	if f.TglAwal != "" && f.TglAkhir != "" {
		q = q.Where("date(tgl_kunjungan) between ? and ?", f.TglAwal, f.TglAkhir)
	}
	q = crud.WhereLike(q, f.Search, "no_rawat", "no_rkm_medis", "nm_pasien", "macam_kasus", "nm_petugas")
	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}
	list := []model.Ringkas{}
	err = q.Order("tgl_kunjungan desc").Offset((page - 1) * limit).Limit(limit).Scan(&list)
	return list, total, err
}

// Find (nil, nil) bila tidak ada.
func (r *repository) Find(noRawat string) (*model.Triase, error) {
	var head []struct {
		model.Utama
		NoRkmMedis string `gorm:"column:no_rkm_medis"`
		NmPasien   string `gorm:"column:nm_pasien"`
	}
	err := support.DB().Raw("select t.no_rawat,cast(t.tgl_kunjungan as char) as tgl_kunjungan,t.cara_masuk,t.alat_transportasi,t.alasan_kedatangan,"+
		"t.keterangan_kedatangan,t.kode_kasus,ifnull(k.macam_kasus,'') as macam_kasus,t.tekanan_darah,t.nadi,t.pernapasan,t.suhu,t.saturasi_o2,t.nyeri,"+
		"reg_periksa.no_rkm_medis,ifnull(pasien.nm_pasien,'') as nm_pasien from data_triase_igd t "+
		"inner join reg_periksa on t.no_rawat=reg_periksa.no_rawat inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis "+
		"left join master_triase_macam_kasus k on t.kode_kasus=k.kode_kasus where t.no_rawat=?", noRawat).Scan(&head)
	if err != nil || len(head) == 0 {
		return nil, err
	}
	t := &model.Triase{Utama: head[0].Utama, NoRkmMedis: head[0].NoRkmMedis, NmPasien: head[0].NmPasien, Skala: model.Skala{Items: []model.SkalaItem{}}}

	var pen []model.Penilaian
	if err := support.DB().Raw("select p.keluhan_utama as keluhan,p.kebutuhan_khusus,p.catatan,p.plan,cast(p.tanggaltriase as char) as tanggaltriase,"+
		"p.nik,ifnull(pegawai.nama,'') as nm_petugas from data_triase_igdprimer p left join pegawai on p.nik=pegawai.nik where p.no_rawat=?", noRawat).Scan(&pen); err != nil {
		return nil, err
	}
	if len(pen) > 0 {
		t.Jenis, t.Penilaian = model.Primer, &pen[0]
	} else {
		if err := support.DB().Raw("select s.anamnesa_singkat as keluhan,'' as kebutuhan_khusus,s.catatan,s.plan,cast(s.tanggaltriase as char) as tanggaltriase,"+
			"s.nik,ifnull(pegawai.nama,'') as nm_petugas from data_triase_igdsekunder s left join pegawai on s.nik=pegawai.nik where s.no_rawat=?", noRawat).Scan(&pen); err != nil {
			return nil, err
		}
		if len(pen) > 0 {
			t.Jenis, t.Penilaian = model.Sekunder, &pen[0]
		}
	}

	for level := 1; level <= 5; level++ {
		n := strconv.Itoa(level)
		items := []model.SkalaItem{}
		if err := support.DB().Raw("select d.kode_skala"+n+" as kode,ifnull(m.pengkajian_skala"+n+",'') as pengkajian,ifnull(m.kode_pemeriksaan,'') as kode_pemeriksaan,"+
			"ifnull(p.nama_pemeriksaan,'') as nama_pemeriksaan from data_triase_igddetail_skala"+n+" d "+
			"left join master_triase_skala"+n+" m on d.kode_skala"+n+"=m.kode_skala"+n+" "+
			"left join master_triase_pemeriksaan p on m.kode_pemeriksaan=p.kode_pemeriksaan where d.no_rawat=? order by d.kode_skala"+n, noRawat).Scan(&items); err != nil {
			return nil, err
		}
		if len(items) > 0 {
			t.Skala = model.Skala{Level: level, Items: items}
			break
		}
	}
	return t, nil
}

func (r *repository) Exists(noRawat string) (bool, error) {
	return crud.ExistsIn(support.DB(), "data_triase_igd", "no_rawat", noRawat)
}

func (r *repository) KasusExists(kode string) (bool, error) {
	return crud.ExistsIn(support.DB(), "master_triase_macam_kasus", "kode_kasus", kode)
}

func (r *repository) SkalaExists(level int, kode string) (bool, error) {
	n := strconv.Itoa(level)
	return crud.ExistsIn(support.DB(), "master_triase_skala"+n, "kode_skala"+n, kode)
}

func (r *repository) Insert(tx contractsorm.Query, in model.Input) error {
	u := in.Utama
	if _, err := tx.Exec("insert into data_triase_igd (no_rawat,tgl_kunjungan,cara_masuk,alat_transportasi,alasan_kedatangan,keterangan_kedatangan,"+
		"kode_kasus,tekanan_darah,nadi,pernapasan,suhu,saturasi_o2,nyeri,id_observation_cara_masuk,id_observation_alat_transportasi,"+
		"id_observation_alasan_kedatangan,id_observation_macam_kasus,id_observation_tekanan_darah,id_observation_nadi,id_observation_pernapasan,"+
		"id_observation_suhu,id_observation_saturasi_o2,id_observation_nyeri,id_composition) values (?,?,?,?,?,?,?,?,?,?,?,?,?,'','','','','','','','','','','')",
		u.NoRawat, u.TglKunjungan, u.CaraMasuk, u.AlatTransportasi, u.AlasanKedatangan, u.KeteranganKedatangan, u.KodeKasus,
		u.TekananDarah, u.Nadi, u.Pernapasan, u.Suhu, u.SaturasiO2, u.Nyeri); err != nil {
		return err
	}
	if err := r.insertPenilaian(tx, in); err != nil {
		return err
	}
	return r.replaceSkala(tx, in)
}

func (r *repository) insertPenilaian(tx contractsorm.Query, in model.Input) error {
	p := in.Penilaian
	if in.Jenis == model.Primer {
		_, err := tx.Exec("insert into data_triase_igdprimer (no_rawat,keluhan_utama,kebutuhan_khusus,catatan,plan,tanggaltriase,nik,"+
			"id_observation_keluhan_utama,id_observation_kebutuhan_khusus,id_observation_catatan,id_careplan_keputusan,id_observation_skala) "+
			"values (?,?,?,?,?,?,?,'','','','','')", in.NoRawat, p.Keluhan, p.KebutuhanKhusus, p.Catatan, p.Plan, p.TanggalTriase, p.Nik)
		return err
	}
	_, err := tx.Exec("insert into data_triase_igdsekunder (no_rawat,anamnesa_singkat,catatan,plan,tanggaltriase,nik,"+
		"id_clinicalimpression_anamnesa,id_observation_catatan,id_careplan_keputusan,id_observation_skala) values (?,?,?,?,?,?,'','','','')",
		in.NoRawat, p.Keluhan, p.Catatan, p.Plan, p.TanggalTriase, p.Nik)
	return err
}

func (r *repository) replaceSkala(tx contractsorm.Query, in model.Input) error {
	for level := 1; level <= 5; level++ {
		if _, err := tx.Exec("delete from data_triase_igddetail_skala"+strconv.Itoa(level)+" where no_rawat=?", in.NoRawat); err != nil {
			return err
		}
	}
	n := strconv.Itoa(in.SkalaLevel)
	for _, kode := range in.SkalaKode {
		if _, err := tx.Exec("insert into data_triase_igddetail_skala"+n+" (no_rawat,kode_skala"+n+",id_observation_skala"+n+") values (?,?,'')",
			in.NoRawat, kode); err != nil {
			return err
		}
	}
	return nil
}

// Update kolom id SatuSehat tidak disentuh; bila jenis berubah, bagian lama dihapus dan bagian baru dibuat.
func (r *repository) Update(tx contractsorm.Query, in model.Input, jenisLama string) error {
	u := in.Utama
	if _, err := tx.Exec("update data_triase_igd set tgl_kunjungan=?,cara_masuk=?,alat_transportasi=?,alasan_kedatangan=?,keterangan_kedatangan=?,"+
		"kode_kasus=?,tekanan_darah=?,nadi=?,pernapasan=?,suhu=?,saturasi_o2=?,nyeri=? where no_rawat=?",
		u.TglKunjungan, u.CaraMasuk, u.AlatTransportasi, u.AlasanKedatangan, u.KeteranganKedatangan, u.KodeKasus,
		u.TekananDarah, u.Nadi, u.Pernapasan, u.Suhu, u.SaturasiO2, u.Nyeri, u.NoRawat); err != nil {
		return err
	}
	p := in.Penilaian
	switch {
	case jenisLama == in.Jenis && in.Jenis == model.Primer:
		if _, err := tx.Exec("update data_triase_igdprimer set keluhan_utama=?,kebutuhan_khusus=?,catatan=?,plan=?,tanggaltriase=?,nik=? where no_rawat=?",
			p.Keluhan, p.KebutuhanKhusus, p.Catatan, p.Plan, p.TanggalTriase, p.Nik, in.NoRawat); err != nil {
			return err
		}
	case jenisLama == in.Jenis:
		if _, err := tx.Exec("update data_triase_igdsekunder set anamnesa_singkat=?,catatan=?,plan=?,tanggaltriase=?,nik=? where no_rawat=?",
			p.Keluhan, p.Catatan, p.Plan, p.TanggalTriase, p.Nik, in.NoRawat); err != nil {
			return err
		}
	default:
		if err := r.deletePenilaian(tx, in.NoRawat); err != nil {
			return err
		}
		if err := r.insertPenilaian(tx, in); err != nil {
			return err
		}
	}
	return r.replaceSkala(tx, in)
}

func (r *repository) deletePenilaian(tx contractsorm.Query, noRawat string) error {
	for _, t := range []string{"data_triase_igdprimer", "data_triase_igdsekunder"} {
		if _, err := tx.Exec("delete from "+t+" where no_rawat=?", noRawat); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) Delete(tx contractsorm.Query, noRawat string) error {
	if err := r.replaceSkala(tx, model.Input{Utama: model.Utama{NoRawat: noRawat}}); err != nil {
		return err
	}
	if err := r.deletePenilaian(tx, noRawat); err != nil {
		return err
	}
	_, err := tx.Exec("delete from data_triase_igd where no_rawat=?", noRawat)
	return err
}
