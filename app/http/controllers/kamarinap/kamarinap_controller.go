package kamarinap

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/kamarinap"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/kamarinap"
	repo "goravel/app/repository/kamarinap"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Index ?status=dirawat|pulang|semua&tgl_awal=&tgl_akhir=&kd_bangsal=&search=
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.List(repo.Filter{
		Status:    q.Input("status", "dirawat"),
		TglAwal:   q.Input("tgl_awal"),
		TglAkhir:  q.Input("tgl_akhir"),
		KdBangsal: q.Input("kd_bangsal"),
		Search:    q.Input("search"),
		Page:      page,
		Limit:     limit,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data kamar inap", err)
	}
	return c.ResponsePaginated(ctx, "Data kamar inap berhasil diambil", data, page, limit, total)
}

// Riwayat ?no_rawat=
func (c *Controller) Riwayat(ctx http.Context) http.Response {
	data, err := c.action.Riwayat(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil riwayat kamar", err)
	}
	return c.ResponseSuccess(ctx, "Riwayat kamar berhasil diambil", data)
}

func (c *Controller) Masuk(ctx http.Context) http.Response {
	var req request.MasukRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Masuk(action.MasukInput{
		NoRawat: req.NoRawat, KdKamar: req.KdKamar, DiagnosaAwal: req.DiagnosaAwal, Tgl: req.TglMasuk, Jam: req.JamMasuk,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan pasien masuk kamar", err)
	}
	return c.ResponseCreated(ctx, "Pasien berhasil masuk kamar", data)
}

func (c *Controller) Pindah(ctx http.Context) http.Response {
	var req request.PindahRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Pindah(action.PindahInput{
		NoRawat: req.NoRawat, KdKamar: req.KdKamar, Mode: action.ModePindah(req.Mode), Tgl: req.TglPindah, Jam: req.JamPindah,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memindahkan kamar", err)
	}
	return c.ResponseSuccess(ctx, "Pasien berhasil pindah kamar", data)
}

func (c *Controller) Pulang(ctx http.Context) http.Response {
	var req request.PulangRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Pulang(action.PulangInput{
		NoRawat: req.NoRawat, SttsPulang: req.SttsPulang, DiagnosaAkhir: req.DiagnosaAkhir, Tgl: req.TglKeluar, Jam: req.JamKeluar,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memulangkan pasien", err)
	}
	return c.ResponseSuccess(ctx, "Pasien berhasil dipulangkan", data)
}

func (c *Controller) BatalPulang(ctx http.Context) http.Response {
	var req request.BatalPulangRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.BatalPulang(req.NoRawat)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal membatalkan pulang", err)
	}
	return c.ResponseSuccess(ctx, "Pulang berhasil dibatalkan", data)
}
