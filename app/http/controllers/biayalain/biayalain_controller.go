package biayalain

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/biayalain"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/biayalain"
	model "goravel/app/models/biayalain"
	repo "goravel/app/repository/biayalain"
)

// Controller tambahan atau potongan biaya.
type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController(jenis model.Jenis) *Controller {
	return &Controller{action: action.NewAction(jenis, repo.NewRepository())}
}

// Index ?no_rawat=
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data biaya", err)
	}
	return c.ResponseSuccess(ctx, "Data biaya berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req.NoRawat, req.Nama, req.Besar)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan biaya", err)
	}
	return c.ResponseCreated(ctx, "Biaya berhasil disimpan", data)
}

// Destroy ?no_rawat=&nama=
func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), ctx.Request().Query("nama")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus biaya", err)
	}
	return c.ResponseSuccess(ctx, "Biaya berhasil dihapus", nil)
}
