package sukubangsa

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/sukubangsa"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/sukubangsa"
	repo "goravel/app/repository/sukubangsa"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data suku bangsa", err)
	}
	return c.ResponsePaginated(ctx, "Data suku bangsa berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(c.key(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil suku bangsa", err)
	}
	return c.ResponseSuccess(ctx, "Detail suku bangsa berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan suku bangsa", err)
	}
	return c.ResponseCreated(ctx, "Suku bangsa berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(c.key(ctx), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui suku bangsa", err)
	}
	return c.ResponseSuccess(ctx, "Suku bangsa berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(c.key(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus suku bangsa", err)
	}
	return c.ResponseSuccess(ctx, "Suku bangsa berhasil dihapus", nil)
}

func (c *Controller) key(ctx http.Context) int {
	id, _ := strconv.Atoi(ctx.Request().Route("id"))
	return id
}
