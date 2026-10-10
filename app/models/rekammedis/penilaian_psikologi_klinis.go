package rekammedis

import "time"

// PenilaianPsikologiKlinis tabel `penilaian_psikologi_klinis` (penilaian psikologi klinis, RMPenilaianPsikologiKlinis).
type PenilaianPsikologiKlinis struct {
	NoRawat                            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                                string     `gorm:"column:nip" json:"nip"`
	Anamnesis                          string     `gorm:"column:anamnesis" json:"anamnesis"`
	DikirimDari                        string     `gorm:"column:dikirim_dari" json:"dikirim_dari"`
	TujuanPemeriksaan                  string     `gorm:"column:tujuan_pemeriksaan" json:"tujuan_pemeriksaan"`
	KetAnamnesis                       *string    `gorm:"column:ket_anamnesis" json:"ket_anamnesis"`
	KeluhanUtama                       *string    `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	RiwayatPenyakit                    *string    `gorm:"column:riwayat_penyakit" json:"riwayat_penyakit"`
	RiwayatKeluhan                     *string    `gorm:"column:riwayat_keluhan" json:"riwayat_keluhan"`
	PermasalahanSaatIni                *string    `gorm:"column:permasalahan_saat_ini" json:"permasalahan_saat_ini"`
	PermasalahanAlasan                 *string    `gorm:"column:permasalahan_alasan" json:"permasalahan_alasan"`
	PermasalahanEkspektasi             *string    `gorm:"column:permasalahan_ekspektasi" json:"permasalahan_ekspektasi"`
	RiwayatHidupSingkat                *string    `gorm:"column:riwayat_hidup_singkat" json:"riwayat_hidup_singkat"`
	KondisiPsikologisPenampilan        *string    `gorm:"column:kondisi_psikologis_penampilan" json:"kondisi_psikologis_penampilan"`
	KondisiPsikologisEkspresiWajah     *string    `gorm:"column:kondisi_psikologis_ekspresi_wajah" json:"kondisi_psikologis_ekspresi_wajah"`
	KondisiPsikologisSuasanaHati       *string    `gorm:"column:kondisi_psikologis_suasana_hati" json:"kondisi_psikologis_suasana_hati"`
	KondisiPsikologisTingkahLaku       *string    `gorm:"column:kondisi_psikologis_tingkah_laku" json:"kondisi_psikologis_tingkah_laku"`
	KondisiPsikologisFungsiUmum        *string    `gorm:"column:kondisi_psikologis_fungsi_umum" json:"kondisi_psikologis_fungsi_umum"`
	KondisiPsikologisFungsiIntelektual *string    `gorm:"column:kondisi_psikologis_fungsi_intelektual" json:"kondisi_psikologis_fungsi_intelektual"`
	KondisiPsikologisPengalaman        *string    `gorm:"column:kondisi_psikologis_pengalaman" json:"kondisi_psikologis_pengalaman"`
	KondisiPsikologisLainnya           *string    `gorm:"column:kondisi_psikologis_lainnya" json:"kondisi_psikologis_lainnya"`
	KondisiPatologisDelusi             *string    `gorm:"column:kondisi_patologis_delusi" json:"kondisi_patologis_delusi"`
	KondisiPatologisProsesPikiran      *string    `gorm:"column:kondisi_patologis_proses_pikiran" json:"kondisi_patologis_proses_pikiran"`
	KondisiPatologisHalusinasi         *string    `gorm:"column:kondisi_patologis_halusinasi" json:"kondisi_patologis_halusinasi"`
	KondisiPatologisAfek               *string    `gorm:"column:kondisi_patologis_afek" json:"kondisi_patologis_afek"`
	KondisiPatologisInsight            *string    `gorm:"column:kondisi_patologis_insight" json:"kondisi_patologis_insight"`
	KondisiPatologisKesadaran          *string    `gorm:"column:kondisi_patologis_kesadaran" json:"kondisi_patologis_kesadaran"`
	KondisiPatologisOrientasi          *string    `gorm:"column:kondisi_patologis_orientasi" json:"kondisi_patologis_orientasi"`
	KondisiPatologisAtensi             *string    `gorm:"column:kondisi_patologis_atensi" json:"kondisi_patologis_atensi"`
	KondisiPatologisKontrolImpuls      *string    `gorm:"column:kondisi_patologis_kontrol_impuls" json:"kondisi_patologis_kontrol_impuls"`
	PsikotesTanggalPelaksanaan         *time.Time `gorm:"column:psikotes_tanggal_pelaksanaan" json:"psikotes_tanggal_pelaksanaan"`
	PsikotesNamaTes                    *string    `gorm:"column:psikotes_nama_tes" json:"psikotes_nama_tes"`
	PsikotesHasil                      *string    `gorm:"column:psikotes_hasil" json:"psikotes_hasil"`
	DinamikaPsikologis                 *string    `gorm:"column:dinamika_psikologis" json:"dinamika_psikologis"`
	DiagnosaPsikologis                 *string    `gorm:"column:diagnosa_psikologis" json:"diagnosa_psikologis"`
	ManifestasiFungsiPsikologis        *string    `gorm:"column:manifestasi_fungsi_psikologis" json:"manifestasi_fungsi_psikologis"`
	RencanaIntervensi                  *string    `gorm:"column:rencana_intervensi" json:"rencana_intervensi"`
	TahapanIntervensi1                 *string    `gorm:"column:tahapan_intervensi1" json:"tahapan_intervensi1"`
	TargetTerapi1                      *string    `gorm:"column:target_terapi1" json:"target_terapi1"`
	TahapanIntervensi2                 *string    `gorm:"column:tahapan_intervensi2" json:"tahapan_intervensi2"`
	TargetTerapi2                      *string    `gorm:"column:target_terapi2" json:"target_terapi2"`
	TahapanIntervensi3                 *string    `gorm:"column:tahapan_intervensi3" json:"tahapan_intervensi3"`
	TargetTerapi3                      *string    `gorm:"column:target_terapi3" json:"target_terapi3"`
	TahapanIntervensi4                 *string    `gorm:"column:tahapan_intervensi4" json:"tahapan_intervensi4"`
	TargetTerapi4                      *string    `gorm:"column:target_terapi4" json:"target_terapi4"`
	TahapanIntervensi5                 *string    `gorm:"column:tahapan_intervensi5" json:"tahapan_intervensi5"`
	TargetTerapi5                      *string    `gorm:"column:target_terapi5" json:"target_terapi5"`
	TahapanIntervensi6                 *string    `gorm:"column:tahapan_intervensi6" json:"tahapan_intervensi6"`
	TargetTerapi6                      *string    `gorm:"column:target_terapi6" json:"target_terapi6"`
	TahapanIntervensi7                 *string    `gorm:"column:tahapan_intervensi7" json:"tahapan_intervensi7"`
	TargetTerapi7                      *string    `gorm:"column:target_terapi7" json:"target_terapi7"`
	Evaluasi                           *string    `gorm:"column:evaluasi" json:"evaluasi"`
}

func (PenilaianPsikologiKlinis) TableName() string {
	return "penilaian_psikologi_klinis"
}
