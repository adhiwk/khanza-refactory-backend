package rekammedis

import "time"

// PenilaianTerapiWicara tabel `penilaian_terapi_wicara` (penilaian terapi wicara, RMPenilaianTerapiWicara).
type PenilaianTerapiWicara struct {
	NoRawat                              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                              *time.Time `gorm:"column:tanggal" json:"tanggal"`
	DiagnosaTerapiWicara                 *string    `gorm:"column:diagnosa_terapi_wicara" json:"diagnosa_terapi_wicara"`
	DiagnosaMedis                        *string    `gorm:"column:diagnosa_medis" json:"diagnosa_medis"`
	Anamnesa                             string     `gorm:"column:anamnesa" json:"anamnesa"`
	Suhu                                 string     `gorm:"column:suhu" json:"suhu"`
	Rr                                   string     `gorm:"column:rr" json:"rr"`
	Nadi                                 string     `gorm:"column:nadi" json:"nadi"`
	Td                                   string     `gorm:"column:td" json:"td"`
	PerilakuAdaptifKontakMata            *string    `gorm:"column:perilaku_adaptif_kontak_mata" json:"perilaku_adaptif_kontak_mata"`
	PerilakuAdaptifAtensi                *string    `gorm:"column:perilaku_adaptif_atensi" json:"perilaku_adaptif_atensi"`
	PerilakuAdaptifPerilaku              *string    `gorm:"column:perilaku_adaptif_perilaku" json:"perilaku_adaptif_perilaku"`
	KemampuanBahasaBicaraSpontan         *string    `gorm:"column:kemampuan_bahasa_bicara_spontan" json:"kemampuan_bahasa_bicara_spontan"`
	KemampuanBahasaPemahamanBahasa       *string    `gorm:"column:kemampuan_bahasa_pemahaman_bahasa" json:"kemampuan_bahasa_pemahaman_bahasa"`
	KemampuanBahasaPengujaran            *string    `gorm:"column:kemampuan_bahasa_pengujaran" json:"kemampuan_bahasa_pengujaran"`
	KemampuanBahasaMembaca               *string    `gorm:"column:kemampuan_bahasa_membaca" json:"kemampuan_bahasa_membaca"`
	KemampuanBahasaPenamaan              *string    `gorm:"column:kemampuan_bahasa_penamaan" json:"kemampuan_bahasa_penamaan"`
	OrganWicaraAnatomisLip               string     `gorm:"column:organ_wicara_anatomis_lip" json:"organ_wicara_anatomis_lip"`
	OrganWicaraAnatomisTongue            string     `gorm:"column:organ_wicara_anatomis_tongue" json:"organ_wicara_anatomis_tongue"`
	OrganWicaraAnatomisHardPalate        string     `gorm:"column:organ_wicara_anatomis_hard_palate" json:"organ_wicara_anatomis_hard_palate"`
	OrganWicaraAnatomisSoftPalate        string     `gorm:"column:organ_wicara_anatomis_soft_palate" json:"organ_wicara_anatomis_soft_palate"`
	OrganWicaraAnatomisUvula             string     `gorm:"column:organ_wicara_anatomis_uvula" json:"organ_wicara_anatomis_uvula"`
	OrganWicaraAnatomisMandibula         string     `gorm:"column:organ_wicara_anatomis_mandibula" json:"organ_wicara_anatomis_mandibula"`
	OrganWicaraAnatomisMaxila            string     `gorm:"column:organ_wicara_anatomis_maxila" json:"organ_wicara_anatomis_maxila"`
	OrganWicaraAnatomisDental            string     `gorm:"column:organ_wicara_anatomis_dental" json:"organ_wicara_anatomis_dental"`
	OrganWicaraAnatomisFaring            string     `gorm:"column:organ_wicara_anatomis_faring" json:"organ_wicara_anatomis_faring"`
	OrganWicaraFisiologisLip             string     `gorm:"column:organ_wicara_fisiologis_lip" json:"organ_wicara_fisiologis_lip"`
	OrganWicaraFisiologisTongue          string     `gorm:"column:organ_wicara_fisiologis_tongue" json:"organ_wicara_fisiologis_tongue"`
	OrganWicaraFisiologisHardPalate      string     `gorm:"column:organ_wicara_fisiologis_hard_palate" json:"organ_wicara_fisiologis_hard_palate"`
	OrganWicaraFisiologisSoftPalate      string     `gorm:"column:organ_wicara_fisiologis_soft_palate" json:"organ_wicara_fisiologis_soft_palate"`
	OrganWicaraFisiologisUvula           string     `gorm:"column:organ_wicara_fisiologis_uvula" json:"organ_wicara_fisiologis_uvula"`
	OrganWicaraFisiologisMandibula       string     `gorm:"column:organ_wicara_fisiologis_mandibula" json:"organ_wicara_fisiologis_mandibula"`
	OrganWicaraFisiologisMaxilla         string     `gorm:"column:organ_wicara_fisiologis_maxilla" json:"organ_wicara_fisiologis_maxilla"`
	OrganWicaraFisiologisDental          string     `gorm:"column:organ_wicara_fisiologis_dental" json:"organ_wicara_fisiologis_dental"`
	OrganWicaraFisiologisFaring          string     `gorm:"column:organ_wicara_fisiologis_faring" json:"organ_wicara_fisiologis_faring"`
	AktifitasOralMenghisap               string     `gorm:"column:aktifitas_oral_menghisap" json:"aktifitas_oral_menghisap"`
	AktifitasOralMengunyah               string     `gorm:"column:aktifitas_oral_mengunyah" json:"aktifitas_oral_mengunyah"`
	AktifitasOralMeniup                  string     `gorm:"column:aktifitas_oral_meniup" json:"aktifitas_oral_meniup"`
	KemampuanArtikulasiSubtitusi         string     `gorm:"column:kemampuan_artikulasi_subtitusi" json:"kemampuan_artikulasi_subtitusi"`
	KemampuanArtikulasiOmisi             string     `gorm:"column:kemampuan_artikulasi_omisi" json:"kemampuan_artikulasi_omisi"`
	KemampuanArtikulasiDistorsi          string     `gorm:"column:kemampuan_artikulasi_distorsi" json:"kemampuan_artikulasi_distorsi"`
	KemampuanArtikulasiAdisi             string     `gorm:"column:kemampuan_artikulasi_adisi" json:"kemampuan_artikulasi_adisi"`
	Resonasi                             string     `gorm:"column:resonasi" json:"resonasi"`
	KemampuanSuaraNada                   string     `gorm:"column:kemampuan_suara_nada" json:"kemampuan_suara_nada"`
	KemampuanSuaraKualitas               string     `gorm:"column:kemampuan_suara_kualitas" json:"kemampuan_suara_kualitas"`
	KemampuanSuaraKenyaringan            string     `gorm:"column:kemampuan_suara_kenyaringan" json:"kemampuan_suara_kenyaringan"`
	KemampuanIramaKelancaran             string     `gorm:"column:kemampuan_irama_kelancaran" json:"kemampuan_irama_kelancaran"`
	KemampuanMenelan                     string     `gorm:"column:kemampuan_menelan" json:"kemampuan_menelan"`
	Pernafasan                           string     `gorm:"column:pernafasan" json:"pernafasan"`
	TingkatKomunikasiDekodingPendengaran string     `gorm:"column:tingkat_komunikasi_dekoding_pendengaran" json:"tingkat_komunikasi_dekoding_pendengaran"`
	TingkatKomunikasiDekodingPenglihatan string     `gorm:"column:tingkat_komunikasi_dekoding_penglihatan" json:"tingkat_komunikasi_dekoding_penglihatan"`
	TingkatKomunikasiDekodingKinesik     string     `gorm:"column:tingkat_komunikasi_dekoding_kinesik" json:"tingkat_komunikasi_dekoding_kinesik"`
	TingkatKomunikasiEnkodingBicara      string     `gorm:"column:tingkat_komunikasi_enkoding_bicara" json:"tingkat_komunikasi_enkoding_bicara"`
	TingkatKomunikasiEnkodingTulisan     string     `gorm:"column:tingkat_komunikasi_enkoding_tulisan" json:"tingkat_komunikasi_enkoding_tulisan"`
	TingkatKomunikasiEnkodingMimik       string     `gorm:"column:tingkat_komunikasi_enkoding_mimik" json:"tingkat_komunikasi_enkoding_mimik"`
	TingkatKomunikasiEnkodingGesture     string     `gorm:"column:tingkat_komunikasi_enkoding_gesture" json:"tingkat_komunikasi_enkoding_gesture"`
	PenunjangMedis                       string     `gorm:"column:penunjang_medis" json:"penunjang_medis"`
	PerencanaanTerapiTujuan              string     `gorm:"column:perencanaan_terapi_tujuan" json:"perencanaan_terapi_tujuan"`
	PerencanaanTerapiProgram             string     `gorm:"column:perencanaan_terapi_program" json:"perencanaan_terapi_program"`
	Edukasi                              string     `gorm:"column:edukasi" json:"edukasi"`
	TindakLanjut                         string     `gorm:"column:tindak_lanjut" json:"tindak_lanjut"`
	Nip                                  string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianTerapiWicara) TableName() string {
	return "penilaian_terapi_wicara"
}
