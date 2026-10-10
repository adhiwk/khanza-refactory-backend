package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningGiziKehamilanData isian skrining gizi kehamilan.
type SkriningGiziKehamilanData struct {
	Parameter1 string `form:"parameter1" json:"parameter1"`
	Skor1      string `form:"skor1" json:"skor1"`
	Parameter2 string `form:"parameter2" json:"parameter2"`
	Skor2      string `form:"skor2" json:"skor2"`
	Parameter3 string `form:"parameter3" json:"parameter3"`
	Skor3      string `form:"skor3" json:"skor3"`
	Parameter4 string `form:"parameter4" json:"parameter4"`
	Skor4      string `form:"skor4" json:"skor4"`
	NilaiSkor  string `form:"nilai_skor" json:"nilai_skor"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nip        string `form:"nip" json:"nip"`
}

func skriningGiziKehamilanRules() map[string]any {
	rules := map[string]any{
		"parameter1": "in:Ya,Tidak",
		"skor1":      "string|max_len:1",
		"parameter2": "in:Ya,Tidak",
		"skor2":      "string|max_len:1",
		"parameter3": "in:Ya,Tidak",
		"skor3":      "string|max_len:1",
		"parameter4": "in:Ya,Tidak",
		"skor4":      "string|max_len:1",
		"nilai_skor": "string|max_len:1",
		"keterangan": "string|max_len:50",
		"nip":        "string|max_len:20",
	}
	return rules
}

// SkriningGiziKehamilanStore simpan skrining gizi kehamilan; kolom waktu kunci kosong = sekarang.
type SkriningGiziKehamilanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SkriningGiziKehamilanData
}

func (r *SkriningGiziKehamilanStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningGiziKehamilanStore) Rules(ctx http.Context) map[string]any {
	rules := skriningGiziKehamilanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SkriningGiziKehamilanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SkriningGiziKehamilanStore) Payload() SkriningGiziKehamilanData {
	return r.SkriningGiziKehamilanData
}

func (r *SkriningGiziKehamilanStore) DetailValues() map[string][]string { return nil }

// SkriningGiziKehamilanUpdate ubah skrining gizi kehamilan (PUT); kunci lewat query string.
type SkriningGiziKehamilanUpdate struct {
	SkriningGiziKehamilanData
}

func (r *SkriningGiziKehamilanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningGiziKehamilanUpdate) Rules(ctx http.Context) map[string]any {
	return skriningGiziKehamilanRules()
}

func (r *SkriningGiziKehamilanUpdate) Payload() SkriningGiziKehamilanData {
	return r.SkriningGiziKehamilanData
}

func (r *SkriningGiziKehamilanUpdate) DetailValues() map[string][]string { return nil }
