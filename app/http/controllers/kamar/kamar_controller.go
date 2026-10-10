package kamar

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/kamar"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/kamar"
	repo "goravel/app/repository/kamar"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data kamar", err)
	}
	return c.ResponsePaginated(ctx, "Data kamar berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kd_kamar"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil kamar", err)
	}
	return c.ResponseSuccess(ctx, "Detail kamar berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan kamar", err)
	}
	return c.ResponseCreated(ctx, "Kamar berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kd_kamar"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui kamar", err)
	}
	return c.ResponseSuccess(ctx, "Kamar berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kd_kamar")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus kamar", err)
	}
	return c.ResponseSuccess(ctx, "Kamar berhasil dihapus", nil)
}
