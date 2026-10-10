package kategoriobat

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/kategoriobat"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/kategoriobat"
	repo "goravel/app/repository/kategoriobat"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data kategori obat", err)
	}
	return c.ResponsePaginated(ctx, "Data kategori obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil kategori obat", err)
	}
	return c.ResponseSuccess(ctx, "Detail kategori obat berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan kategori obat", err)
	}
	return c.ResponseCreated(ctx, "Kategori obat berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui kategori obat", err)
	}
	return c.ResponseSuccess(ctx, "Kategori obat berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus kategori obat", err)
	}
	return c.ResponseSuccess(ctx, "Kategori obat berhasil dihapus", nil)
}
