package templateradiologi

import (
	"github.com/goravel/framework/contracts/http"
)

// Data isian template hasil radiologi; no_template dibuat otomatis ("R" + 4 digit).
type Data struct {
	NamaPemeriksaan        string `form:"nama_pemeriksaan" json:"nama_pemeriksaan"`
	TemplateHasilRadiologi string `form:"template_hasil_radiologi" json:"template_hasil_radiologi"`
}

type Request struct {
	Data
}

func (r *Request) Authorize(ctx http.Context) error { return nil }

func (r *Request) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"nama_pemeriksaan":         "required|string|max_len:80",
		"template_hasil_radiologi": "required|string",
	}
}
