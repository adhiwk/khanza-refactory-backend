package metoderacik

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/metoderacik"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/metoderacik"
	repo "goravel/app/repository/metoderacik"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data metode racik", err)
	}
	return c.ResponsePaginated(ctx, "Data metode racik berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kd_racik"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil metode racik", err)
	}
	return c.ResponseSuccess(ctx, "Detail metode racik berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan metode racik", err)
	}
	return c.ResponseCreated(ctx, "Metode racik berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kd_racik"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui metode racik", err)
	}
	return c.ResponseSuccess(ctx, "Metode racik berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kd_racik")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus metode racik", err)
	}
	return c.ResponseSuccess(ctx, "Metode racik berhasil dihapus", nil)
}
