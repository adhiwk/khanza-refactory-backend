package dpjp

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/dpjp"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/dpjp"
	repo "goravel/app/repository/dpjp"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Index ?no_rawat=
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data DPJP", err)
	}
	return c.ResponseSuccess(ctx, "Data DPJP berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req.NoRawat, req.KdDokter)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan DPJP", err)
	}
	return c.ResponseCreated(ctx, "DPJP berhasil disimpan", data)
}

// Destroy ?no_rawat=&kd_dokter=
func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), ctx.Request().Query("kd_dokter")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus DPJP", err)
	}
	return c.ResponseSuccess(ctx, "DPJP berhasil dihapus", nil)
}
