package dokter

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/dokter"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/dokter"
	repo "goravel/app/repository/dokter"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data dokter", err)
	}
	return c.ResponsePaginated(ctx, "Data dokter berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kd_dokter"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil dokter", err)
	}
	return c.ResponseSuccess(ctx, "Detail dokter berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan dokter", err)
	}
	return c.ResponseCreated(ctx, "Dokter berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kd_dokter"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui dokter", err)
	}
	return c.ResponseSuccess(ctx, "Dokter berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kd_dokter")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus dokter", err)
	}
	return c.ResponseSuccess(ctx, "Dokter berhasil dihapus", nil)
}
