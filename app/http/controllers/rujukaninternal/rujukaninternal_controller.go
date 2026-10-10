package rujukaninternal

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rujukaninternal"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/rujukaninternal"
	repo "goravel/app/repository/rujukaninternal"
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
		return c.ResponseActionError(ctx, "Gagal mengambil rujukan internal", err)
	}
	return c.ResponseSuccess(ctx, "Data rujukan internal berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req.NoRawat, req.KdDokter, req.KdPoli)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan rujukan internal", err)
	}
	return c.ResponseCreated(ctx, "Rujukan internal berhasil disimpan", data)
}

// Destroy ?no_rawat=&kd_dokter=
func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), ctx.Request().Query("kd_dokter")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus rujukan internal", err)
	}
	return c.ResponseSuccess(ctx, "Rujukan internal berhasil dihapus", nil)
}
