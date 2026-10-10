package satuanobat

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/satuanobat"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/satuanobat"
	repo "goravel/app/repository/satuanobat"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data satuan obat", err)
	}
	return c.ResponsePaginated(ctx, "Data satuan obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode_sat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil satuan obat", err)
	}
	return c.ResponseSuccess(ctx, "Detail satuan obat berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan satuan obat", err)
	}
	return c.ResponseCreated(ctx, "Satuan obat berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode_sat"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui satuan obat", err)
	}
	return c.ResponseSuccess(ctx, "Satuan obat berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode_sat")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus satuan obat", err)
	}
	return c.ResponseSuccess(ctx, "Satuan obat berhasil dihapus", nil)
}
