package jabatan

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/jabatan"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/jabatan"
	repo "goravel/app/repository/jabatan"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data jabatan", err)
	}
	return c.ResponsePaginated(ctx, "Data jabatan berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kd_jbtn"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil jabatan", err)
	}
	return c.ResponseSuccess(ctx, "Detail jabatan berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan jabatan", err)
	}
	return c.ResponseCreated(ctx, "Jabatan berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kd_jbtn"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui jabatan", err)
	}
	return c.ResponseSuccess(ctx, "Jabatan berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kd_jbtn")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus jabatan", err)
	}
	return c.ResponseSuccess(ctx, "Jabatan berhasil dihapus", nil)
}
