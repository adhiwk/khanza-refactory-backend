package jenisobat

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/jenisobat"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/jenisobat"
	repo "goravel/app/repository/jenisobat"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	data, total, err := c.action.List(ctx.Request().Input("search"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data jenis obat", err)
	}
	return c.ResponsePaginated(ctx, "Data jenis obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kdjns"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil jenis obat", err)
	}
	return c.ResponseSuccess(ctx, "Detail jenis obat berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan jenis obat", err)
	}
	return c.ResponseCreated(ctx, "Jenis obat berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kdjns"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui jenis obat", err)
	}
	return c.ResponseSuccess(ctx, "Jenis obat berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kdjns")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus jenis obat", err)
	}
	return c.ResponseSuccess(ctx, "Jenis obat berhasil dihapus", nil)
}
