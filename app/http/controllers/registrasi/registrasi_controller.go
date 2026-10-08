package registrasi

import (
	"errors"
	"strconv"

	"github.com/goravel/framework/contracts/http"

	regaction "goravel/app/actions/registrasi"
	"goravel/app/http/controllers"
	regrequest "goravel/app/http/requests/registrasi"
	"goravel/app/modules/rbac"
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
		return c.actionError(ctx, "Gagal mengambil registrasi", err)
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
		return c.actionError(ctx, "Gagal menyimpan registrasi", err)
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

	result, err := c.action.Update(ctx.Request().Query("no_rawat"), req.RegistrasiData, c.isSuperAdmin(ctx))
	if err != nil {
		return c.actionError(ctx, "Gagal memperbarui registrasi", err)
	}

	return c.ResponseSuccess(ctx, "Registrasi berhasil diperbarui", result)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), c.isSuperAdmin(ctx)); err != nil {
		return c.actionError(ctx, "Gagal menghapus registrasi", err)
	}

	return c.ResponseSuccess(ctx, "Registrasi berhasil dihapus", nil)
}

// isSuperAdmin: super admin bebas dari batas 2 x 24 jam (setara "Admin Utama" di Khanza).
func (c *Controller) isSuperAdmin(ctx http.Context) bool {
	uid, ok := rbac.UserID(ctx)
	if !ok {
		return false
	}
	has, err := rbac.Default().HasAnyRole(uid, rbac.SuperAdminRole)
	return err == nil && has
}

// actionError memetakan error domain dari action ke HTTP status.
func (c *Controller) actionError(ctx http.Context, message string, err error) http.Response {
	switch {
	case errors.Is(err, regaction.ErrNotFound),
		errors.Is(err, regaction.ErrPasienNotFound),
		errors.Is(err, regaction.ErrDokterNotFound),
		errors.Is(err, regaction.ErrPoliNotFound),
		errors.Is(err, regaction.ErrPenjabNotFound):
		return c.ResponseError(ctx, http.StatusNotFound, err.Error(), nil)
	case errors.Is(err, regaction.ErrDirawatInap),
		errors.Is(err, regaction.ErrLocked),
		errors.Is(err, regaction.ErrNomorTidakTersedia):
		return c.ResponseError(ctx, http.StatusConflict, err.Error(), nil)
	case errors.Is(err, regaction.ErrExpired):
		return c.ResponseError(ctx, http.StatusForbidden, err.Error(), nil)
	case errors.Is(err, regaction.ErrInvalidDate),
		errors.Is(err, regaction.ErrInvalidTime):
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
	default:
		return c.ResponseError(ctx, http.StatusInternalServerError, message, err.Error())
	}
}
