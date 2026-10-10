package templateedukasi

import (
	"github.com/goravel/framework/contracts/http"
)

// Data isian template informasi edukasi; no_template dibuat otomatis ("E" + 3 digit).
type Data struct {
	MateriEdukasi string `form:"materi_edukasi" json:"materi_edukasi"`
	LamaEdukasi   string `form:"lama_edukasi" json:"lama_edukasi"`
	MetodeEdukasi string `form:"metode_edukasi" json:"metode_edukasi"`
}

type Request struct {
	Data
}

func (r *Request) Authorize(ctx http.Context) error { return nil }

func (r *Request) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"materi_edukasi": "required|string|max_len:1000",
		"lama_edukasi":   "string|max_len:10",
		"metode_edukasi": "required|in:Ceramah,Diskusi,Demonstrasi",
	}
}
