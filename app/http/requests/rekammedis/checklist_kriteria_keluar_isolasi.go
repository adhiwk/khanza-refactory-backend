package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaKeluarIsolasiData isian checklist kriteria keluar isolasi.
type ChecklistKriteriaKeluarIsolasiData struct {
	GejalaMembaik             string `form:"gejala_membaik" json:"gejala_membaik"`
	TidakAdaIndikasiTransmisi string `form:"tidak_ada_indikasi_transmisi" json:"tidak_ada_indikasi_transmisi"`
	HasilPenunjangMemenuhi    string `form:"hasil_penunjang_memenuhi" json:"hasil_penunjang_memenuhi"`
	KriteriaPedomanTerpenuhi  string `form:"kriteria_pedoman_terpenuhi" json:"kriteria_pedoman_terpenuhi"`
	PersetujuanDpjp           string `form:"persetujuan_dpjp" json:"persetujuan_dpjp"`
	Keputusan                 string `form:"keputusan" json:"keputusan"`
	Alasan                    string `form:"alasan" json:"alasan"`
	Nik                       string `form:"nik" json:"nik"`
}

func checklistKriteriaKeluarIsolasiRules() map[string]any {
	rules := map[string]any{
		"gejala_membaik":               "required|in:Ya,Tidak,Tidak Berlaku",
		"tidak_ada_indikasi_transmisi": "required|in:Ya,Tidak,Tidak Berlaku",
		"hasil_penunjang_memenuhi":     "required|in:Ya,Tidak,Tidak Berlaku",
		"kriteria_pedoman_terpenuhi":   "required|in:Ya,Tidak,Tidak Berlaku",
		"persetujuan_dpjp":             "required|in:Ya,Tidak",
		"keputusan":                    "required|in:Keluar Isolasi,Lanjut Isolasi",
		"alasan":                       "string|max_len:500",
		"nik":                          "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaKeluarIsolasiStore simpan checklist kriteria keluar isolasi; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaKeluarIsolasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaKeluarIsolasiData
}

func (r *ChecklistKriteriaKeluarIsolasiStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarIsolasiStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaKeluarIsolasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaKeluarIsolasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaKeluarIsolasiStore) Payload() ChecklistKriteriaKeluarIsolasiData {
	return r.ChecklistKriteriaKeluarIsolasiData
}

func (r *ChecklistKriteriaKeluarIsolasiStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaKeluarIsolasiUpdate ubah checklist kriteria keluar isolasi (PUT); kunci lewat query string.
type ChecklistKriteriaKeluarIsolasiUpdate struct {
	ChecklistKriteriaKeluarIsolasiData
}

func (r *ChecklistKriteriaKeluarIsolasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarIsolasiUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaKeluarIsolasiRules()
}

func (r *ChecklistKriteriaKeluarIsolasiUpdate) Payload() ChecklistKriteriaKeluarIsolasiData {
	return r.ChecklistKriteriaKeluarIsolasiData
}

func (r *ChecklistKriteriaKeluarIsolasiUpdate) DetailValues() map[string][]string { return nil }
