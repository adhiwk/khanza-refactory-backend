package pemeriksaan

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/pemeriksaan"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/pemeriksaan"
	model "goravel/app/models/perawatan"
	repo "goravel/app/repository/pemeriksaan"
)

// Controller pemeriksaan untuk satu jenis rawat. no_rawat mengandung "/", jadi identitas baris memakai query string.
type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController(rawat model.Rawat) *Controller {
	return &Controller{action: action.NewAction(rawat, repo.NewRepository())}
}

func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data pemeriksaan", err)
	}
	return c.ResponseSuccess(ctx, "Data pemeriksaan berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan pemeriksaan", err)
	}
	return c.ResponseCreated(ctx, "Pemeriksaan berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	q := ctx.Request()
	data, err := c.action.Update(q.Query("no_rawat"), q.Query("tgl_perawatan"), q.Query("jam_rawat"), req.Data, c.IsSuperAdmin(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui pemeriksaan", err)
	}
	return c.ResponseSuccess(ctx, "Pemeriksaan berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	q := ctx.Request()
	if err := c.action.Delete(q.Query("no_rawat"), q.Query("tgl_perawatan"), q.Query("jam_rawat"), c.IsSuperAdmin(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus pemeriksaan", err)
	}
	return c.ResponseSuccess(ctx, "Pemeriksaan berhasil dihapus", nil)
}
