package obat

import (
	"errors"

	"github.com/goravel/framework/contracts/http"

	obatAction "goravel/app/actions/obat"
	"goravel/app/facades"
	"goravel/app/http/controllers"
	obatRequest "goravel/app/http/requests/obat"
	obatRepo "goravel/app/repository/obat"
	"goravel/app/support"
)

type ObatController struct {
	controllers.BaseController
	action *obatAction.Action
}

func NewObatController() *ObatController {
	return &ObatController{
		action: obatAction.NewAction(obatRepo.NewRepository()),
	}
}

// GET /obat?search=&status=&kdjns=&kode_kategori=&kode_golongan=&page=1&limit=10
func (c *ObatController) Index(ctx http.Context) http.Response {
	req := obatAction.ListRequest{
		Search:       ctx.Request().Query("search"),
		Status:       ctx.Request().Query("status"),
		Kdjns:        ctx.Request().Query("kdjns"),
		KodeKategori: ctx.Request().Query("kode_kategori"),
		KodeGolongan: ctx.Request().Query("kode_golongan"),
		Page:         ctx.Request().QueryInt("page", 1),
		Limit:        ctx.Request().QueryInt("limit", 10),
	}

	res, err := c.action.List(req)
	if err != nil {
		return c.handleError(ctx, err)
	}

	paginated := support.FormatLaravelPagination(ctx, res.Data, res.Meta.Page, res.Meta.Limit, res.Meta.Total)

	return c.ResponseSuccess(ctx, "Data obat berhasil diambil", paginated)
}

// GET /obat/{kode}
func (c *ObatController) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("kode"))
	if err != nil {
		return c.handleError(ctx, err)
	}

	return c.ResponseSuccess(ctx, "Detail obat berhasil diambil", data)
}

// POST /obat
func (c *ObatController) Store(ctx http.Context) http.Response {
	var req obatRequest.StoreObatRequest
	if resp := c.validateRequest(ctx, &req); resp != nil {
		return resp
	}

	data, err := c.action.Create(req.ObatRequest)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return c.ResponseCreated(ctx, "Obat berhasil ditambahkan", data)
}

// PUT /obat/{kode}
func (c *ObatController) Update(ctx http.Context) http.Response {
	var req obatRequest.UpdateObatRequest
	if resp := c.validateRequest(ctx, &req); resp != nil {
		return resp
	}

	data, err := c.action.Update(ctx.Request().Route("kode"), req.ObatRequest)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return c.ResponseSuccess(ctx, "Obat berhasil diperbarui", data)
}

// PATCH /obat/{kode}/status   body: {"status": "0" | "1"}
func (c *ObatController) UpdateStatus(ctx http.Context) http.Response {
	var req obatRequest.UpdateStatusRequest
	if resp := c.validateRequest(ctx, &req); resp != nil {
		return resp
	}

	data, err := c.action.UpdateStatus(ctx.Request().Route("kode"), req.Status)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return c.ResponseSuccess(ctx, "Status obat berhasil diperbarui", data)
}

// DELETE /obat/{kode}
func (c *ObatController) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("kode")); err != nil {
		return c.handleError(ctx, err)
	}

	return c.ResponseSuccess(ctx, "Obat berhasil dihapus", nil)
}

// ===== helper =====

// validateRequest menjalankan FormRequest (authorize + rules + bind).
// Mengembalikan response error bila gagal, atau nil bila valid.
func (c *ObatController) validateRequest(ctx http.Context, req http.FormRequest) http.Response {
	validationErrors, err := ctx.Request().ValidateRequest(req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "Format request tidak valid", err.Error())
	}
	if validationErrors != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", validationErrors.All())
	}
	return nil
}

// handleError memetakan error domain ke HTTP status; error tak dikenal
// dicatat ke log dan tidak dibocorkan ke client.
func (c *ObatController) handleError(ctx http.Context, err error) http.Response {
	switch {
	case errors.Is(err, obatAction.ErrNotFound):
		return c.ResponseError(ctx, http.StatusNotFound, err.Error(), nil)
	case errors.Is(err, obatAction.ErrAlreadyExists):
		return c.ResponseError(ctx, http.StatusConflict, err.Error(), nil)
	case errors.Is(err, obatAction.ErrInvalidStatus),
		errors.Is(err, obatAction.ErrInvalidExpire):
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
	default:
		facades.Log().WithContext(ctx).Errorf("obat: %v", err)
		return c.ResponseError(ctx, http.StatusInternalServerError, "Terjadi kesalahan pada server", nil)
	}
}
