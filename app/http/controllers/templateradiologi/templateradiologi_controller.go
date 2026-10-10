package templateradiologi

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/templateradiologi"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/templateradiologi"
	repo "goravel/app/repository/templateradiologi"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data template hasil radiologi", err)
	}
	return c.ResponsePaginated(ctx, "Data template hasil radiologi berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_template"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil template hasil radiologi", err)
	}
	return c.ResponseSuccess(ctx, "Detail template hasil radiologi berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.Request
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan template hasil radiologi", err)
	}
	return c.ResponseCreated(ctx, "Template hasil radiologi berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.Request
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_template"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui template hasil radiologi", err)
	}
	return c.ResponseSuccess(ctx, "Template hasil radiologi berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_template")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus template hasil radiologi", err)
	}
	return c.ResponseSuccess(ctx, "Template hasil radiologi berhasil dihapus", nil)
}
