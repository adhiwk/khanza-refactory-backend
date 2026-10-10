package masterkode

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/masterkode"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/masterkode"
	model "goravel/app/models/masterkode"
	repo "goravel/app/repository/masterkode"
	"goravel/app/support"
)

// Controller satu tabel master kode.
type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController(spec *repo.Spec) *Controller {
	return &Controller{action: action.NewAction(spec, repo.NewRepository())}
}

// Index ?search=&kode_induk=
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	data, total, err := c.action.List(ctx.Request().Query("search"), ctx.Request().Query("kode_induk"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data "+c.action.Spec().Label, err)
	}
	return c.ResponsePaginated(ctx, "Data "+c.action.Spec().Label+" berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil "+c.action.Spec().Label, err)
	}
	return c.ResponseSuccess(ctx, "Detail "+c.action.Spec().Label+" berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(model.Kode{Kode: req.Kode, Nama: req.Nama, KodeInduk: req.KodeInduk})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan "+c.action.Spec().Label, err)
	}
	return c.ResponseCreated(ctx, c.action.Spec().Label+" berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode"), model.Kode{Nama: req.Nama, KodeInduk: req.KodeInduk})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui "+c.action.Spec().Label, err)
	}
	return c.ResponseSuccess(ctx, c.action.Spec().Label+" berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus "+c.action.Spec().Label, err)
	}
	return c.ResponseSuccess(ctx, c.action.Spec().Label+" berhasil dihapus", nil)
}
