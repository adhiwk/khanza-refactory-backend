package rujukmasuk

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rujukmasuk"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/rujukmasuk"
	repo "goravel/app/repository/rujukmasuk"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data rujukan masuk", err)
	}
	return c.ResponsePaginated(ctx, "Data rujukan masuk berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil rujukan masuk", err)
	}
	return c.ResponseSuccess(ctx, "Detail rujukan masuk berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan rujukan masuk", err)
	}
	return c.ResponseCreated(ctx, "Rujukan masuk berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Query("no_rawat"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui rujukan masuk", err)
	}
	return c.ResponseSuccess(ctx, "Rujukan masuk berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus rujukan masuk", err)
	}
	return c.ResponseSuccess(ctx, "Rujukan masuk berhasil dihapus", nil)
}
