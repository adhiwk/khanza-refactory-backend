package templateoperasi

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/templateoperasi"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/templateoperasi"
	repo "goravel/app/repository/templateoperasi"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data template laporan operasi", err)
	}
	return c.ResponsePaginated(ctx, "Data template laporan operasi berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_template"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil template laporan operasi", err)
	}
	return c.ResponseSuccess(ctx, "Detail template laporan operasi berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.Request
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan template laporan operasi", err)
	}
	return c.ResponseCreated(ctx, "Template laporan operasi berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.Request
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_template"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui template laporan operasi", err)
	}
	return c.ResponseSuccess(ctx, "Template laporan operasi berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_template")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus template laporan operasi", err)
	}
	return c.ResponseSuccess(ctx, "Template laporan operasi berhasil dihapus", nil)
}
