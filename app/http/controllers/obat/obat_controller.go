package obat

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/obat"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/obat"
	repo "goravel/app/repository/obat"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Index ?search=&status=&kdjns=&kode_kategori=&kode_golongan=&page=&limit= (tanpa status = hanya aktif)
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.List(q.Input("search"), action.Filter{
		Status:       q.Input("status"),
		Kdjns:        q.Input("kdjns"),
		KodeKategori: q.Input("kode_kategori"),
		KodeGolongan: q.Input("kode_golongan"),
	}, page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data obat", err)
	}
	return c.ResponsePaginated(ctx, "Data obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode_brng"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil obat", err)
	}
	return c.ResponseSuccess(ctx, "Detail obat berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan obat", err)
	}
	return c.ResponseCreated(ctx, "Obat berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("kode_brng"), req.Data)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui obat", err)
	}
	return c.ResponseSuccess(ctx, "Obat berhasil diperbarui", data)
}

func (c *Controller) UpdateStatus(ctx http.Context) http.Response {
	var req request.StatusRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.UpdateStatus(ctx.Request().Route("kode_brng"), req.Status)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui status obat", err)
	}
	return c.ResponseSuccess(ctx, "Status obat berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode_brng")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus obat", err)
	}
	return c.ResponseSuccess(ctx, "Obat berhasil dihapus", nil)
}
