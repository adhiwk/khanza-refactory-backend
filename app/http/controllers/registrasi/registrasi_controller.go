package registrasi

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"

	regaction "goravel/app/actions/registrasi"
	"goravel/app/http/controllers"
	regrequest "goravel/app/http/requests/registrasi"
	regrepo "goravel/app/repository/registrasi"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *regaction.Action
}

func NewController() *Controller {
	repo := regrepo.NewRepository()
	return &Controller{
		action: regaction.NewAction(repo),
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

	filter := regrepo.Filter{
		TglAwal:  ctx.Request().Input("tgl_awal"),
		TglAkhir: ctx.Request().Input("tgl_akhir"),
		KdDokter: ctx.Request().Input("kd_dokter"),
		KdPoli:   ctx.Request().Input("kd_poli"),
		Search:   ctx.Request().Input("search"),
		Page:     page,
		Limit:    limit,
	}

	data, total, err := c.action.List(filter)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
	}

	paginatedData := support.FormatLaravelPagination(ctx, data, page, limit, total)

	return c.ResponseSuccess(ctx, "Data registrasi berhasil diambil", paginatedData)
}

// no_rawat berformat yyyy/MM/dd/NNNNNN sehingga dikirim lewat query string, bukan path.
func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil registrasi", err)
	}

	return c.ResponseSuccess(ctx, "Detail registrasi berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req regrequest.StoreRegistrasiRequest
	errs, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memvalidasi request", err.Error())
	}
	if errs != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errs.All())
	}

	result, err := c.action.Create(req.NoRkmMedis, req.RegistrasiData)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan registrasi", err)
	}

	return c.ResponseCreated(ctx, "Registrasi berhasil disimpan", result)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req regrequest.UpdateRegistrasiRequest
	errs, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memvalidasi request", err.Error())
	}
	if errs != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errs.All())
	}

	result, err := c.action.Update(ctx.Request().Query("no_rawat"), req.RegistrasiData, c.IsSuperAdmin(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui registrasi", err)
	}

	return c.ResponseSuccess(ctx, "Registrasi berhasil diperbarui", result)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), c.IsSuperAdmin(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus registrasi", err)
	}

	return c.ResponseSuccess(ctx, "Registrasi berhasil dihapus", nil)
}
