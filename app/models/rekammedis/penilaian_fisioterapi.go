package rekammedis

import "time"

// PenilaianFisioterapi tabel `penilaian_fisioterapi` (penilaian fisioterapi, RMPenilaianFisioterapi).
type PenilaianFisioterapi struct {
	NoRawat                    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                    *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Informasi                  string     `gorm:"column:informasi" json:"informasi"`
	KeluhanUtama               string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                        string     `gorm:"column:rps" json:"rps"`
	Rpd                        string     `gorm:"column:rpd" json:"rpd"`
	Td                         string     `gorm:"column:td" json:"td"`
	Hr                         string     `gorm:"column:hr" json:"hr"`
	Rr                         string     `gorm:"column:rr" json:"rr"`
	Suhu                       string     `gorm:"column:suhu" json:"suhu"`
	NyeriTekan                 string     `gorm:"column:nyeri_tekan" json:"nyeri_tekan"`
	NyeriGerak                 string     `gorm:"column:nyeri_gerak" json:"nyeri_gerak"`
	NyeriDiam                  string     `gorm:"column:nyeri_diam" json:"nyeri_diam"`
	Palpasi                    string     `gorm:"column:palpasi" json:"palpasi"`
	LuasGerakSendi             string     `gorm:"column:luas_gerak_sendi" json:"luas_gerak_sendi"`
	KekuatanOtot               string     `gorm:"column:kekuatan_otot" json:"kekuatan_otot"`
	Statis                     string     `gorm:"column:statis" json:"statis"`
	Dinamis                    string     `gorm:"column:dinamis" json:"dinamis"`
	Kognitif                   string     `gorm:"column:kognitif" json:"kognitif"`
	Auskultasi                 string     `gorm:"column:auskultasi" json:"auskultasi"`
	AlatBantu                  string     `gorm:"column:alat_bantu" json:"alat_bantu"`
	KetBantu                   string     `gorm:"column:ket_bantu" json:"ket_bantu"`
	Prothesa                   string     `gorm:"column:prothesa" json:"prothesa"`
	KetPro                     string     `gorm:"column:ket_pro" json:"ket_pro"`
	Deformitas                 string     `gorm:"column:deformitas" json:"deformitas"`
	KetDeformitas              string     `gorm:"column:ket_deformitas" json:"ket_deformitas"`
	Resikojatuh                string     `gorm:"column:resikojatuh" json:"resikojatuh"`
	KetResikojatuh             string     `gorm:"column:ket_resikojatuh" json:"ket_resikojatuh"`
	Adl                        string     `gorm:"column:adl" json:"adl"`
	LainlainFungsional         string     `gorm:"column:lainlain_fungsional" json:"lainlain_fungsional"`
	KetFisik                   string     `gorm:"column:ket_fisik" json:"ket_fisik"`
	PemeriksaanMusculoskeletal string     `gorm:"column:pemeriksaan_musculoskeletal" json:"pemeriksaan_musculoskeletal"`
	PemeriksaanNeuromuscular   string     `gorm:"column:pemeriksaan_neuromuscular" json:"pemeriksaan_neuromuscular"`
	PemeriksaanCardiopulmonal  string     `gorm:"column:pemeriksaan_cardiopulmonal" json:"pemeriksaan_cardiopulmonal"`
	PemeriksaanIntegument      string     `gorm:"column:pemeriksaan_integument" json:"pemeriksaan_integument"`
	PengukuranMusculoskeletal  string     `gorm:"column:pengukuran_musculoskeletal" json:"pengukuran_musculoskeletal"`
	PengukuranNeuromuscular    string     `gorm:"column:pengukuran_neuromuscular" json:"pengukuran_neuromuscular"`
	PengukuranCardiopulmonal   string     `gorm:"column:pengukuran_cardiopulmonal" json:"pengukuran_cardiopulmonal"`
	PengukuranIntegument       string     `gorm:"column:pengukuran_integument" json:"pengukuran_integument"`
	Penunjang                  string     `gorm:"column:penunjang" json:"penunjang"`
	DiagnosisFisio             string     `gorm:"column:diagnosis_fisio" json:"diagnosis_fisio"`
	RencanaTerapi              string     `gorm:"column:rencana_terapi" json:"rencana_terapi"`
	Nip                        string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianFisioterapi) TableName() string {
	return "penilaian_fisioterapi"
}
