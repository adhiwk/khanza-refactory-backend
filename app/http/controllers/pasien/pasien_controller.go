package pasien

import (
	pasienaction "goravel/app/actions/pasien"
	pasienrequest "goravel/app/http/requests/pasien"
	pasienrepo "goravel/app/repository/pasien"

	"goravel/app/http/controllers"
	"goravel/app/support"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

type Controller struct {
	controllers.BaseController
	action *pasienaction.Action
}

func NewController() *Controller {
	repo := pasienrepo.NewRepository()
	return &Controller{
		action: pasienaction.NewAction(repo),
	}
}

func (c *Controller) Index(ctx http.Context) http.Response {
	page, _ := strconv.Atoi(ctx.Request().Input("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Input("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	data, total, err := c.action.List(ctx.Request().Input("search"), page, limit)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
	}

	paginatedData := support.FormatLaravelPagination(ctx, data, page, limit, total)

	return c.ResponseSuccess(ctx, "Data pasien berhasil diambil", paginatedData)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_rkm_medis"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil pasien", err)
	}

	return c.ResponseSuccess(ctx, "Detail pasien berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req pasienrequest.StorePasienRequest
	errs, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memvalidasi request", err.Error())
	}
	if errs != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errs.All())
	}

	result, err := c.action.Create(req.NoRkmMedis, req.PasienData)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal membuat pasien", err)
	}

	return c.ResponseCreated(ctx, "Pasien berhasil dibuat", result)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req pasienrequest.UpdatePasienRequest
	errs, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memvalidasi request", err.Error())
	}
	if errs != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errs.All())
	}

	result, err := c.action.Update(ctx.Request().Route("no_rkm_medis"), req.PasienData)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui pasien", err)
	}

	return c.ResponseSuccess(ctx, "Pasien berhasil diperbarui", result)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_rkm_medis")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus pasien", err)
	}

	return c.ResponseSuccess(ctx, "Pasien berhasil dihapus", nil)
}
