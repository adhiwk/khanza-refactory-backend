package tagihan

import (
	model "goravel/app/models/tagihan"
	perawatanrepo "goravel/app/repository/perawatan"
	"goravel/app/support"
)

// Kategori urutan kategori rincian beserta query-nya (kolom: nama, tanggal, biaya, jumlah, total). Parameter tunggal: no_rawat.
var Kategori = []struct {
	Nama string
	SQL  string
}{
	{"Registrasi", "select 'Registrasi' as nama,cast(tgl_registrasi as char) as tanggal,ifnull(biaya_reg,0) as biaya,1 as jumlah,ifnull(biaya_reg,0) as total from reg_periksa where no_rawat=? and ifnull(biaya_reg,0)<>0"},
	{"Ralan Dokter", tindakan("rawat_jl_dr", "jns_perawatan")},
	{"Ralan Paramedis", tindakan("rawat_jl_pr", "jns_perawatan")},
	{"Ralan Dokter Paramedis", tindakan("rawat_jl_drpr", "jns_perawatan")},
	{"Ranap Dokter", tindakan("rawat_inap_dr", "jns_perawatan_inap")},
	{"Ranap Paramedis", tindakan("rawat_inap_pr", "jns_perawatan_inap")},
	{"Ranap Dokter Paramedis", tindakan("rawat_inap_drpr", "jns_perawatan_inap")},
	{"Obat", "select ifnull(databarang.nama_brng,detail_pemberian_obat.kode_brng) as nama,cast(detail_pemberian_obat.tgl_perawatan as char) as tanggal," +
		"ifnull(detail_pemberian_obat.biaya_obat,0) as biaya,detail_pemberian_obat.jml as jumlah,detail_pemberian_obat.total as total " +
		"from detail_pemberian_obat left join databarang on detail_pemberian_obat.kode_brng=databarang.kode_brng where detail_pemberian_obat.no_rawat=?"},
	{"Obat Operasi", "select ifnull(obatbhp_ok.nm_obat,beri_obat_operasi.kd_obat) as nama,cast(date(beri_obat_operasi.tanggal) as char) as tanggal," +
		"beri_obat_operasi.hargasatuan as biaya,beri_obat_operasi.jumlah as jumlah,beri_obat_operasi.hargasatuan*beri_obat_operasi.jumlah as total " +
		"from beri_obat_operasi left join obatbhp_ok on beri_obat_operasi.kd_obat=obatbhp_ok.kd_obat where beri_obat_operasi.no_rawat=?"},
	{"Tagihan Obat Langsung", "select 'Obat Langsung' as nama,'' as tanggal,besar_tagihan as biaya,1 as jumlah,besar_tagihan as total from tagihan_obat_langsung where no_rawat=?"},
	{"Laborat", "select ifnull(jns_perawatan_lab.nm_perawatan,periksa_lab.kd_jenis_prw) as nama,cast(periksa_lab.tgl_periksa as char) as tanggal," +
		"periksa_lab.biaya as biaya,1 as jumlah,periksa_lab.biaya as total from periksa_lab " +
		"left join jns_perawatan_lab on periksa_lab.kd_jenis_prw=jns_perawatan_lab.kd_jenis_prw where periksa_lab.no_rawat=?"},
	{"Detail Laborat", "select ifnull(template_laboratorium.Pemeriksaan,detail_periksa_lab.kd_jenis_prw) as nama,cast(detail_periksa_lab.tgl_periksa as char) as tanggal," +
		"detail_periksa_lab.biaya_item as biaya,1 as jumlah,detail_periksa_lab.biaya_item as total from detail_periksa_lab " +
		"left join template_laboratorium on detail_periksa_lab.id_template=template_laboratorium.id_template where detail_periksa_lab.no_rawat=? and detail_periksa_lab.biaya_item<>0"},
	{"Radiologi", "select ifnull(jns_perawatan_radiologi.nm_perawatan,periksa_radiologi.kd_jenis_prw) as nama,cast(periksa_radiologi.tgl_periksa as char) as tanggal," +
		"periksa_radiologi.biaya as biaya,1 as jumlah,periksa_radiologi.biaya as total from periksa_radiologi " +
		"left join jns_perawatan_radiologi on periksa_radiologi.kd_jenis_prw=jns_perawatan_radiologi.kd_jenis_prw where periksa_radiologi.no_rawat=?"},
	{"Operasi", "select ifnull(paket_operasi.nm_perawatan,operasi.kode_paket) as nama,cast(date(operasi.tgl_operasi) as char) as tanggal," +
		biayaOperasi + " as biaya,1 as jumlah," + biayaOperasi + " as total from operasi " +
		"left join paket_operasi on operasi.kode_paket=paket_operasi.kode_paket where operasi.no_rawat=?"},
	{"Kamar", "select concat(kamar_inap.kd_kamar,' ',ifnull(bangsal.nm_bangsal,'')) as nama,cast(kamar_inap.tgl_masuk as char) as tanggal," +
		"ifnull(kamar_inap.trf_kamar,0) as biaya,ifnull(kamar_inap.lama,0) as jumlah,ifnull(kamar_inap.ttl_biaya,0) as total from kamar_inap " +
		"left join kamar on kamar_inap.kd_kamar=kamar.kd_kamar left join bangsal on kamar.kd_bangsal=bangsal.kd_bangsal where kamar_inap.no_rawat=?"},
	{"Tambahan", "select nama_biaya as nama,'' as tanggal,besar_biaya as biaya,1 as jumlah,besar_biaya as total from tambahan_biaya where no_rawat=?"},
}

// biayaOperasi penjumlahan seluruh komponen biaya operasi (sqlpsoperasi DlgBilingRalan).
const biayaOperasi = "(operasi.biayaoperator1+operasi.biayaoperator2+operasi.biayaoperator3+operasi.biayaasisten_operator1+" +
	"operasi.biayaasisten_operator2+operasi.biayaasisten_operator3+operasi.biayainstrumen+operasi.biayadokter_anak+" +
	"operasi.biayaperawaat_resusitas+operasi.biayadokter_anestesi+operasi.biayaasisten_anestesi+operasi.biayaasisten_anestesi2+" +
	"operasi.biayabidan+operasi.biayabidan2+operasi.biayabidan3+operasi.biayaperawat_luar+operasi.biayaalat+operasi.biayasewaok+" +
	"operasi.biaya_omloop+operasi.biaya_omloop2+operasi.biaya_omloop3+operasi.biaya_omloop4+operasi.biaya_omloop5+" +
	"operasi.biayasarpras+operasi.biaya_dokter_pjanak+operasi.biaya_dokter_umum)"

func tindakan(t, jns string) string {
	return "select ifnull(" + jns + ".nm_perawatan," + t + ".kd_jenis_prw) as nama,cast(" + t + ".tgl_perawatan as char) as tanggal," +
		"ifnull(" + t + ".biaya_rawat,0) as biaya,1 as jumlah,ifnull(" + t + ".biaya_rawat,0) as total from " + t +
		" left join " + jns + " on " + t + ".kd_jenis_prw=" + jns + ".kd_jenis_prw where " + t + ".no_rawat=?"
}

type Repository interface {
	perawatanrepo.KonteksRepository
	Items(sql, noRawat string) ([]model.Item, error)
	Potongan(noRawat string) ([]model.Item, error)
	Deposit(noRawat string) (float64, error)
	StatusBayar(noRawat string) (string, error)
}

type repository struct {
	perawatanrepo.KonteksRepository
}

func NewRepository() Repository {
	return &repository{perawatanrepo.NewKonteksRepository()}
}

func (r *repository) Items(sql, noRawat string) ([]model.Item, error) {
	list := []model.Item{}
	err := support.DB().Raw(sql, noRawat).Scan(&list)
	return list, err
}

func (r *repository) Potongan(noRawat string) ([]model.Item, error) {
	return r.Items("select nama_pengurangan as nama,'' as tanggal,ifnull(besar_pengurangan,0) as biaya,1 as jumlah,"+
		"ifnull(besar_pengurangan,0) as total from pengurangan_biaya where no_rawat=?", noRawat)
}

func (r *repository) Deposit(noRawat string) (float64, error) {
	var list []struct {
		Total float64 `gorm:"column:total"`
	}
	err := support.DB().Raw("select ifnull(sum(besar_deposit),0) as total from deposit where no_rawat=?", noRawat).Scan(&list)
	if err != nil || len(list) == 0 {
		return 0, err
	}
	return list[0].Total, nil
}

func (r *repository) StatusBayar(noRawat string) (string, error) {
	var list []struct {
		S string `gorm:"column:s"`
	}
	err := support.DB().Raw("select status_bayar as s from reg_periksa where no_rawat=?", noRawat).Scan(&list)
	if err != nil || len(list) == 0 {
		return "", err
	}
	return list[0].S, nil
}
