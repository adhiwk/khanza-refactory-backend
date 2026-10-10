package deposit

import (
	"github.com/goravel/framework/contracts/http"
)

type StoreRequest struct {
	NoRawat      string  `form:"no_rawat" json:"no_rawat"`
	TglDeposit   string  `form:"tgl_deposit" json:"tgl_deposit"`
	JamDeposit   string  `form:"jam_deposit" json:"jam_deposit"`
	NamaBayar    string  `form:"nama_bayar" json:"nama_bayar"`
	BesarDeposit float64 `form:"besar_deposit" json:"besar_deposit"`
	Nip          string  `form:"nip" json:"nip"`
	Keterangan   string  `form:"keterangan" json:"keterangan"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":      "required|string|max_len:17",
		"tgl_deposit":   "date",
		"jam_deposit":   "string|len:8",
		"nama_bayar":    "required|string|max_len:50",
		"besar_deposit": "required|numeric|gt:0",
		"nip":           "string|max_len:20",
		"keterangan":    "string|max_len:70",
	}
}
