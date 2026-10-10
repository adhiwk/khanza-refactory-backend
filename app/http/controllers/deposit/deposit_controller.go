package deposit

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/deposit"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/deposit"
	repo "goravel/app/repository/deposit"
	jurnalrepo "goravel/app/repository/jurnal"
	jurnalsvc "goravel/app/services/jurnal"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository(), jurnalsvc.NewService(jurnalrepo.NewRepository()))}
}

// Index ?no_rawat=
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data deposit", err)
	}
	return c.ResponseSuccess(ctx, "Data deposit berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(action.Input{
		NoRawat: req.NoRawat, Tgl: req.TglDeposit, Jam: req.JamDeposit, NamaBayar: req.NamaBayar, Besar: req.BesarDeposit,
		Nip: req.Nip, Keterangan: req.Keterangan, Operator: c.operator(ctx),
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan deposit", err)
	}
	return c.ResponseCreated(ctx, "Deposit berhasil disimpan", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_deposit"), c.operator(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus deposit", err)
	}
	return c.ResponseSuccess(ctx, "Deposit berhasil dihapus", nil)
}

func (c *Controller) operator(ctx http.Context) string {
	return fmt.Sprintf("USER %d", c.UserID(ctx))
}
