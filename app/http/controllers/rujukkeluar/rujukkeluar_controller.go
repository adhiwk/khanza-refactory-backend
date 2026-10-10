package rujukkeluar

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rujukkeluar"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/rujukkeluar"
	repo "goravel/app/repository/rujukkeluar"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data rujukan keluar", err)
	}
	return c.ResponsePaginated(ctx, "Data rujukan keluar berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_rujuk"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil rujukan keluar", err)
	}
	return c.ResponseSuccess(ctx, "Detail rujukan keluar berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan rujukan keluar", err)
	}
	return c.ResponseCreated(ctx, "Rujukan keluar berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_rujuk"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui rujukan keluar", err)
	}
	return c.ResponseSuccess(ctx, "Rujukan keluar berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_rujuk")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus rujukan keluar", err)
	}
	return c.ResponseSuccess(ctx, "Rujukan keluar berhasil dihapus", nil)
}
