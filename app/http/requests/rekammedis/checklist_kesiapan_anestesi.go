package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKesiapanAnestesiData isian checklist kesiapan anestesi.
type ChecklistKesiapanAnestesiData struct {
	Nip               string `form:"nip" json:"nip"`
	KdDokter          string `form:"kd_dokter" json:"kd_dokter"`
	Tindakan          string `form:"tindakan" json:"tindakan"`
	TeknikAnestesi    string `form:"teknik_anestesi" json:"teknik_anestesi"`
	Listrik1          string `form:"listrik1" json:"listrik1"`
	Listrik2          string `form:"listrik2" json:"listrik2"`
	Listrik3          string `form:"listrik3" json:"listrik3"`
	Listrik4          string `form:"listrik4" json:"listrik4"`
	Gasmedis1         string `form:"gasmedis1" json:"gasmedis1"`
	Gasmedis2         string `form:"gasmedis2" json:"gasmedis2"`
	Gasmedis3         string `form:"gasmedis3" json:"gasmedis3"`
	Gasmedis4         string `form:"gasmedis4" json:"gasmedis4"`
	Gasmedis5         string `form:"gasmedis5" json:"gasmedis5"`
	Gasmedis6         string `form:"gasmedis6" json:"gasmedis6"`
	Mesinanes1        string `form:"mesinanes1" json:"mesinanes1"`
	Mesinanes2        string `form:"mesinanes2" json:"mesinanes2"`
	Mesinanes3        string `form:"mesinanes3" json:"mesinanes3"`
	Mesinanes4        string `form:"mesinanes4" json:"mesinanes4"`
	Mesinanes5        string `form:"mesinanes5" json:"mesinanes5"`
	Jalannapas1       string `form:"jalannapas1" json:"jalannapas1"`
	Jalannapas2       string `form:"jalannapas2" json:"jalannapas2"`
	Jalannapas3       string `form:"jalannapas3" json:"jalannapas3"`
	Jalannapas4       string `form:"jalannapas4" json:"jalannapas4"`
	Jalannapas5       string `form:"jalannapas5" json:"jalannapas5"`
	Jalannapas6       string `form:"jalannapas6" json:"jalannapas6"`
	Jalannapas7       string `form:"jalannapas7" json:"jalannapas7"`
	Jalannapas8       string `form:"jalannapas8" json:"jalannapas8"`
	Jalannapas9       string `form:"jalannapas9" json:"jalannapas9"`
	Lainlain1         string `form:"lainlain1" json:"lainlain1"`
	Lainlain2         string `form:"lainlain2" json:"lainlain2"`
	Lainlain3         string `form:"lainlain3" json:"lainlain3"`
	Lainlain4         string `form:"lainlain4" json:"lainlain4"`
	Lainlain5         string `form:"lainlain5" json:"lainlain5"`
	Lainlain6         string `form:"lainlain6" json:"lainlain6"`
	Lainlain7         string `form:"lainlain7" json:"lainlain7"`
	Lainlain8         string `form:"lainlain8" json:"lainlain8"`
	Obatobat1         string `form:"obatobat1" json:"obatobat1"`
	Obatobat2         string `form:"obatobat2" json:"obatobat2"`
	Obatobat3         string `form:"obatobat3" json:"obatobat3"`
	Obatobat4         string `form:"obatobat4" json:"obatobat4"`
	Obatobat5         string `form:"obatobat5" json:"obatobat5"`
	Obatobat6         string `form:"obatobat6" json:"obatobat6"`
	KeteranganLainnya string `form:"keterangan_lainnya" json:"keterangan_lainnya"`
}

func checklistKesiapanAnestesiRules() map[string]any {
	rules := map[string]any{
		"nip":                "string|max_len:20",
		"kd_dokter":          "string|max_len:20",
		"tindakan":           "string|max_len:100",
		"teknik_anestesi":    "string|max_len:30",
		"listrik1":           "in:Ya,Tidak",
		"listrik2":           "in:Ya,Tidak",
		"listrik3":           "in:Ya,Tidak",
		"listrik4":           "in:Ya,Tidak",
		"gasmedis1":          "in:Ya,Tidak",
		"gasmedis2":          "in:Ya,Tidak",
		"gasmedis3":          "in:Ya,Tidak",
		"gasmedis4":          "in:Ya,Tidak",
		"gasmedis5":          "in:Ya,Tidak",
		"gasmedis6":          "in:Ya,Tidak",
		"mesinanes1":         "in:Ya,Tidak",
		"mesinanes2":         "in:Ya,Tidak",
		"mesinanes3":         "in:Ya,Tidak",
		"mesinanes4":         "in:Ya,Tidak",
		"mesinanes5":         "in:Ya,Tidak",
		"jalannapas1":        "in:Ya,Tidak",
		"jalannapas2":        "in:Ya,Tidak",
		"jalannapas3":        "in:Ya,Tidak",
		"jalannapas4":        "in:Ya,Tidak",
		"jalannapas5":        "in:Ya,Tidak",
		"jalannapas6":        "in:Ya,Tidak",
		"jalannapas7":        "in:Ya,Tidak",
		"jalannapas8":        "in:Ya,Tidak",
		"jalannapas9":        "in:Ya,Tidak",
		"lainlain1":          "in:Ya,Tidak",
		"lainlain2":          "in:Ya,Tidak",
		"lainlain3":          "in:Ya,Tidak",
		"lainlain4":          "in:Ya,Tidak",
		"lainlain5":          "in:Ya,Tidak",
		"lainlain6":          "in:Ya,Tidak",
		"lainlain7":          "in:Ya,Tidak",
		"lainlain8":          "in:Ya,Tidak",
		"obatobat1":          "in:Ya,Tidak",
		"obatobat2":          "in:Ya,Tidak",
		"obatobat3":          "in:Ya,Tidak",
		"obatobat4":          "in:Ya,Tidak",
		"obatobat5":          "in:Ya,Tidak",
		"obatobat6":          "in:Ya,Tidak",
		"keterangan_lainnya": "string|max_len:1000",
	}
	return rules
}

// ChecklistKesiapanAnestesiStore simpan checklist kesiapan anestesi; kolom waktu kunci kosong = sekarang.
type ChecklistKesiapanAnestesiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKesiapanAnestesiData
}

func (r *ChecklistKesiapanAnestesiStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKesiapanAnestesiStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKesiapanAnestesiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKesiapanAnestesiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKesiapanAnestesiStore) Payload() ChecklistKesiapanAnestesiData {
	return r.ChecklistKesiapanAnestesiData
}

func (r *ChecklistKesiapanAnestesiStore) DetailValues() map[string][]string { return nil }

// ChecklistKesiapanAnestesiUpdate ubah checklist kesiapan anestesi (PUT); kunci lewat query string.
type ChecklistKesiapanAnestesiUpdate struct {
	ChecklistKesiapanAnestesiData
}

func (r *ChecklistKesiapanAnestesiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKesiapanAnestesiUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKesiapanAnestesiRules()
}

func (r *ChecklistKesiapanAnestesiUpdate) Payload() ChecklistKesiapanAnestesiData {
	return r.ChecklistKesiapanAnestesiData
}

func (r *ChecklistKesiapanAnestesiUpdate) DetailValues() map[string][]string { return nil }
