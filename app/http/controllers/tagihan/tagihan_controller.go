package tagihan

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/tagihan"
	"goravel/app/http/controllers"
	repo "goravel/app/repository/tagihan"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Show ?no_rawat=
func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Hitung(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menghitung tagihan", err)
	}
	return c.ResponseSuccess(ctx, "Rincian tagihan berhasil dihitung", data)
}
