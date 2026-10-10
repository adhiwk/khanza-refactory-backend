package kabupaten

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/kabupaten"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/kabupaten"
	repo "goravel/app/repository/kabupaten"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data kabupaten", err)
	}
	return c.ResponsePaginated(ctx, "Data kabupaten berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(c.key(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil kabupaten", err)
	}
	return c.ResponseSuccess(ctx, "Detail kabupaten berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan kabupaten", err)
	}
	return c.ResponseCreated(ctx, "Kabupaten berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(c.key(ctx), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui kabupaten", err)
	}
	return c.ResponseSuccess(ctx, "Kabupaten berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(c.key(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus kabupaten", err)
	}
	return c.ResponseSuccess(ctx, "Kabupaten berhasil dihapus", nil)
}

func (c *Controller) key(ctx http.Context) int {
	id, _ := strconv.Atoi(ctx.Request().Route("kd_kab"))
	return id
}
