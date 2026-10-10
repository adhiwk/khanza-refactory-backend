package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ResumePasienRanapData isian resume pasien ranap.
type ResumePasienRanapData struct {
	KdDokter             string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaAwal         string `form:"diagnosa_awal" json:"diagnosa_awal"`
	Alasan               string `form:"alasan" json:"alasan"`
	KeluhanUtama         string `form:"keluhan_utama" json:"keluhan_utama"`
	PemeriksaanFisik     string `form:"pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	JalannyaPenyakit     string `form:"jalannya_penyakit" json:"jalannya_penyakit"`
	PemeriksaanPenunjang string `form:"pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	HasilLaborat         string `form:"hasil_laborat" json:"hasil_laborat"`
	TindakanDanOperasi   string `form:"tindakan_dan_operasi" json:"tindakan_dan_operasi"`
	ObatDiRs             string `form:"obat_di_rs" json:"obat_di_rs"`
	DiagnosaUtama        string `form:"diagnosa_utama" json:"diagnosa_utama"`
	KdDiagnosaUtama      string `form:"kd_diagnosa_utama" json:"kd_diagnosa_utama"`
	DiagnosaSekunder     string `form:"diagnosa_sekunder" json:"diagnosa_sekunder"`
	KdDiagnosaSekunder   string `form:"kd_diagnosa_sekunder" json:"kd_diagnosa_sekunder"`
	DiagnosaSekunder2    string `form:"diagnosa_sekunder2" json:"diagnosa_sekunder2"`
	KdDiagnosaSekunder2  string `form:"kd_diagnosa_sekunder2" json:"kd_diagnosa_sekunder2"`
	DiagnosaSekunder3    string `form:"diagnosa_sekunder3" json:"diagnosa_sekunder3"`
	KdDiagnosaSekunder3  string `form:"kd_diagnosa_sekunder3" json:"kd_diagnosa_sekunder3"`
	DiagnosaSekunder4    string `form:"diagnosa_sekunder4" json:"diagnosa_sekunder4"`
	KdDiagnosaSekunder4  string `form:"kd_diagnosa_sekunder4" json:"kd_diagnosa_sekunder4"`
	ProsedurUtama        string `form:"prosedur_utama" json:"prosedur_utama"`
	KdProsedurUtama      string `form:"kd_prosedur_utama" json:"kd_prosedur_utama"`
	ProsedurSekunder     string `form:"prosedur_sekunder" json:"prosedur_sekunder"`
	KdProsedurSekunder   string `form:"kd_prosedur_sekunder" json:"kd_prosedur_sekunder"`
	ProsedurSekunder2    string `form:"prosedur_sekunder2" json:"prosedur_sekunder2"`
	KdProsedurSekunder2  string `form:"kd_prosedur_sekunder2" json:"kd_prosedur_sekunder2"`
	ProsedurSekunder3    string `form:"prosedur_sekunder3" json:"prosedur_sekunder3"`
	KdProsedurSekunder3  string `form:"kd_prosedur_sekunder3" json:"kd_prosedur_sekunder3"`
	Alergi               string `form:"alergi" json:"alergi"`
	Diet                 string `form:"diet" json:"diet"`
	LabBelum             string `form:"lab_belum" json:"lab_belum"`
	Edukasi              string `form:"edukasi" json:"edukasi"`
	CaraKeluar           string `form:"cara_keluar" json:"cara_keluar"`
	KetKeluar            string `form:"ket_keluar" json:"ket_keluar"`
	Keadaan              string `form:"keadaan" json:"keadaan"`
	KetKeadaan           string `form:"ket_keadaan" json:"ket_keadaan"`
	Dilanjutkan          string `form:"dilanjutkan" json:"dilanjutkan"`
	KetDilanjutkan       string `form:"ket_dilanjutkan" json:"ket_dilanjutkan"`
	Kontrol              string `form:"kontrol" json:"kontrol"`
	ObatPulang           string `form:"obat_pulang" json:"obat_pulang"`
}

func resumePasienRanapRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":             "required|string|max_len:20",
		"diagnosa_awal":         "string|max_len:100",
		"alasan":                "string|max_len:100",
		"keluhan_utama":         "string",
		"pemeriksaan_fisik":     "string",
		"jalannya_penyakit":     "string",
		"pemeriksaan_penunjang": "string",
		"hasil_laborat":         "string",
		"tindakan_dan_operasi":  "string",
		"obat_di_rs":            "string",
		"diagnosa_utama":        "string|max_len:80",
		"kd_diagnosa_utama":     "string|max_len:10",
		"diagnosa_sekunder":     "string|max_len:80",
		"kd_diagnosa_sekunder":  "string|max_len:10",
		"diagnosa_sekunder2":    "string|max_len:80",
		"kd_diagnosa_sekunder2": "string|max_len:10",
		"diagnosa_sekunder3":    "string|max_len:80",
		"kd_diagnosa_sekunder3": "string|max_len:10",
		"diagnosa_sekunder4":    "string|max_len:80",
		"kd_diagnosa_sekunder4": "string|max_len:10",
		"prosedur_utama":        "string|max_len:80",
		"kd_prosedur_utama":     "string|max_len:8",
		"prosedur_sekunder":     "string|max_len:80",
		"kd_prosedur_sekunder":  "string|max_len:8",
		"prosedur_sekunder2":    "string|max_len:80",
		"kd_prosedur_sekunder2": "string|max_len:8",
		"prosedur_sekunder3":    "string|max_len:80",
		"kd_prosedur_sekunder3": "string|max_len:8",
		"alergi":                "string|max_len:100",
		"diet":                  "string",
		"lab_belum":             "string",
		"edukasi":               "string",
		"cara_keluar":           "required|in:Atas Izin Dokter,Pindah RS,Pulang Atas Permintaan Sendiri,Lainnya",
		"ket_keluar":            "string|max_len:50",
		"keadaan":               "required|in:Membaik,Sembuh,Keadaan Khusus,Meninggal",
		"ket_keadaan":           "string|max_len:50",
		"dilanjutkan":           "required|in:Kembali Ke RS,RS Lain,Dokter Luar,Puskesmes,Lainnya",
		"ket_dilanjutkan":       "string|max_len:50",
		"kontrol":               "date",
		"obat_pulang":           "string",
	}
	return rules
}

// ResumePasienRanapStore simpan resume pasien ranap; kolom waktu kunci kosong = sekarang.
type ResumePasienRanapStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	ResumePasienRanapData
}

func (r *ResumePasienRanapStore) Authorize(ctx http.Context) error { return nil }

func (r *ResumePasienRanapStore) Rules(ctx http.Context) map[string]any {
	rules := resumePasienRanapRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *ResumePasienRanapStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *ResumePasienRanapStore) Payload() ResumePasienRanapData { return r.ResumePasienRanapData }

func (r *ResumePasienRanapStore) DetailValues() map[string][]string { return nil }

// ResumePasienRanapUpdate ubah resume pasien ranap (PUT); kunci lewat query string.
type ResumePasienRanapUpdate struct {
	ResumePasienRanapData
}

func (r *ResumePasienRanapUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ResumePasienRanapUpdate) Rules(ctx http.Context) map[string]any {
	return resumePasienRanapRules()
}

func (r *ResumePasienRanapUpdate) Payload() ResumePasienRanapData { return r.ResumePasienRanapData }

func (r *ResumePasienRanapUpdate) DetailValues() map[string][]string { return nil }
