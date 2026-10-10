package golonganobat

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/golonganobat"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/golonganobat"
	repo "goravel/app/repository/golonganobat"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data golongan obat", err)
	}
	return c.ResponsePaginated(ctx, "Data golongan obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil golongan obat", err)
	}
	return c.ResponseSuccess(ctx, "Detail golongan obat berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan golongan obat", err)
	}
	return c.ResponseCreated(ctx, "Golongan obat berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui golongan obat", err)
	}
	return c.ResponseSuccess(ctx, "Golongan obat berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus golongan obat", err)
	}
	return c.ResponseSuccess(ctx, "Golongan obat berhasil dihapus", nil)
}
