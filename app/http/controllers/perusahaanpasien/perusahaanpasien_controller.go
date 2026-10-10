package perusahaanpasien

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/perusahaanpasien"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/perusahaanpasien"
	repo "goravel/app/repository/perusahaanpasien"
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
		return c.ResponseActionError(ctx, "Gagal mengambil data instansi/perusahaan pasien", err)
	}
	return c.ResponsePaginated(ctx, "Data instansi/perusahaan pasien berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode_perusahaan"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil instansi/perusahaan pasien", err)
	}
	return c.ResponseSuccess(ctx, "Detail instansi/perusahaan pasien berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan instansi/perusahaan pasien", err)
	}
	return c.ResponseCreated(ctx, "Instansi/perusahaan pasien berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode_perusahaan"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui instansi/perusahaan pasien", err)
	}
	return c.ResponseSuccess(ctx, "Instansi/perusahaan pasien berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode_perusahaan")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus instansi/perusahaan pasien", err)
	}
	return c.ResponseSuccess(ctx, "Instansi/perusahaan pasien berhasil dihapus", nil)
}
