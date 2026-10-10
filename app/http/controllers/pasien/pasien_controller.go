package pasien

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/pasien"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/pasien"
	repo "goravel/app/repository/pasien"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data pasien", err)
	}
	return c.ResponsePaginated(ctx, "Data pasien berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_rkm_medis"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil pasien", err)
	}
	return c.ResponseSuccess(ctx, "Detail pasien berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan pasien", err)
	}
	return c.ResponseCreated(ctx, "Pasien berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_rkm_medis"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui pasien", err)
	}
	return c.ResponseSuccess(ctx, "Pasien berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_rkm_medis")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus pasien", err)
	}
	return c.ResponseSuccess(ctx, "Pasien berhasil dihapus", nil)
}
