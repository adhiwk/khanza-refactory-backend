package resep

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/resep"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/resep"
	model "goravel/app/models/resep"
	repo "goravel/app/repository/resep"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Index ?no_rawat=
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data resep", err)
	}
	return c.ResponseSuccess(ctx, "Data resep berhasil diambil", data)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_resep"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil resep", err)
	}
	return c.ResponseSuccess(ctx, "Detail resep berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(action.Input{
		NoRawat: req.NoRawat, KdDokter: req.KdDokter, Tgl: req.TglPeresepan, Jam: req.JamPeresepan, Items: items(req.Items),
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan resep", err)
	}
	return c.ResponseCreated(ctx, "Resep berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_resep"), items(req.Items))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui resep", err)
	}
	return c.ResponseSuccess(ctx, "Resep berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_resep")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus resep", err)
	}
	return c.ResponseSuccess(ctx, "Resep berhasil dihapus", nil)
}

func items(in []request.Item) []model.Item {
	out := make([]model.Item, len(in))
	for i, it := range in {
		out[i] = model.Item{KodeBrng: it.KodeBrng, Jml: it.Jml, AturanPakai: it.AturanPakai}
	}
	return out
}
