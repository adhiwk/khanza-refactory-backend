package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// MppSkriningData isian skrining MPP.
type MppSkriningData struct {
	Param1  string `form:"param1" json:"param1"`
	Param2  string `form:"param2" json:"param2"`
	Param3  string `form:"param3" json:"param3"`
	Param4  string `form:"param4" json:"param4"`
	Param5  string `form:"param5" json:"param5"`
	Param6  string `form:"param6" json:"param6"`
	Param7  string `form:"param7" json:"param7"`
	Param8  string `form:"param8" json:"param8"`
	Param9  string `form:"param9" json:"param9"`
	Param10 string `form:"param10" json:"param10"`
	Param11 string `form:"param11" json:"param11"`
	Param12 string `form:"param12" json:"param12"`
	Param13 string `form:"param13" json:"param13"`
	Param14 string `form:"param14" json:"param14"`
	Param15 string `form:"param15" json:"param15"`
	Param16 string `form:"param16" json:"param16"`
	Nip     string `form:"nip" json:"nip"`
}

func mppSkriningRules() map[string]any {
	rules := map[string]any{
		"param1":  "in:Ya,Tidak",
		"param2":  "in:Ya,Tidak",
		"param3":  "in:Ya,Tidak",
		"param4":  "in:Ya,Tidak",
		"param5":  "in:Ya,Tidak",
		"param6":  "in:Ya,Tidak",
		"param7":  "in:Ya,Tidak",
		"param8":  "in:Ya,Tidak",
		"param9":  "in:Ya,Tidak",
		"param10": "in:Ya,Tidak",
		"param11": "in:Ya,Tidak",
		"param12": "in:Ya,Tidak",
		"param13": "in:Ya,Tidak",
		"param14": "in:Ya,Tidak",
		"param15": "in:Ya,Tidak",
		"param16": "in:Ya,Tidak",
		"nip":     "required|string|max_len:20",
	}
	return rules
}

// MppSkriningStore simpan skrining MPP; kolom waktu kunci kosong = sekarang.
type MppSkriningStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	MppSkriningData
}

func (r *MppSkriningStore) Authorize(ctx http.Context) error { return nil }

func (r *MppSkriningStore) Rules(ctx http.Context) map[string]any {
	rules := mppSkriningRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *MppSkriningStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *MppSkriningStore) Payload() MppSkriningData { return r.MppSkriningData }

func (r *MppSkriningStore) DetailValues() map[string][]string { return nil }

// MppSkriningUpdate ubah skrining MPP (PUT); kunci lewat query string.
type MppSkriningUpdate struct {
	MppSkriningData
}

func (r *MppSkriningUpdate) Authorize(ctx http.Context) error { return nil }

func (r *MppSkriningUpdate) Rules(ctx http.Context) map[string]any { return mppSkriningRules() }

func (r *MppSkriningUpdate) Payload() MppSkriningData { return r.MppSkriningData }

func (r *MppSkriningUpdate) DetailValues() map[string][]string { return nil }
