package industrifarmasi

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/industrifarmasi"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/industrifarmasi"
	repo "goravel/app/repository/industrifarmasi"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data industri farmasi", err)
	}
	return c.ResponsePaginated(ctx, "Data industri farmasi berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode_industri"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil industri farmasi", err)
	}
	return c.ResponseSuccess(ctx, "Detail industri farmasi berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan industri farmasi", err)
	}
	return c.ResponseCreated(ctx, "Industri farmasi berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode_industri"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui industri farmasi", err)
	}
	return c.ResponseSuccess(ctx, "Industri farmasi berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode_industri")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus industri farmasi", err)
	}
	return c.ResponseSuccess(ctx, "Industri farmasi berhasil dihapus", nil)
}
