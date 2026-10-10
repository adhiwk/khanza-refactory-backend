package user

import (
	useraction "goravel/app/actions/user"
	userrequest "goravel/app/http/requests/user"
	userrepo "goravel/app/repository/user"

	"goravel/app/http/controllers"
	"goravel/app/support"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

type Controller struct {
	controllers.BaseController
	action *useraction.Action
}

func NewController() *Controller {
	repo := userrepo.NewRepository()
	return &Controller{
		action: useraction.NewAction(repo),
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

	data, total, err := c.action.List(page, limit)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
	}

	// Format pagination ala Laravel
	paginatedData := support.FormatLaravelPagination(ctx, data, page, limit, total)

	return c.ResponseSuccess(ctx, "Data users berhasil diambil", paginatedData)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	id, err := strconv.Atoi(ctx.Request().Route("id"))
	if err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "ID tidak valid", err.Error())
	}

	data, err := c.action.Detail(uint(id))
	if err != nil {
		return c.ResponseError(ctx, http.StatusNotFound, "User tidak ditemukan", nil)
	}

	return c.ResponseSuccess(ctx, "Detail user berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req userrequest.StoreUserRequest
	errors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memvalidasi request", err.Error())
	}
	if errors != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errors.All())
	}

	result, err := c.action.Create(req.Name, req.Email, req.Password)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal membuat user", err.Error())
	}

	return c.ResponseSuccess(ctx, "User berhasil dibuat", result)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	id, err := strconv.Atoi(ctx.Request().Route("id"))
	if err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var req userrequest.UpdateUserRequest
	errors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memvalidasi request", err.Error())
	}
	if errors != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errors.All())
	}

	result, err := c.action.Update(uint(id), req.Name, req.Email, req.Password)
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memperbarui user", err.Error())
	}

	return c.ResponseSuccess(ctx, "User berhasil diperbarui", result)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	id, err := strconv.Atoi(ctx.Request().Route("id"))
	if err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "ID tidak valid", err.Error())
	}

	err = c.action.Delete(uint(id))
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal menghapus user", err.Error())
	}

	return c.ResponseSuccess(ctx, "User berhasil dihapus", nil)
}

// SetPegawai PUT /users/{id}/pegawai
func (c *Controller) SetPegawai(ctx http.Context) http.Response {
	var req userrequest.PegawaiRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	id, _ := strconv.Atoi(ctx.Request().Route("id"))
	data, err := c.action.SetPegawai(uint(id), req.KdPegawai)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menghubungkan pegawai", err)
	}
	return c.ResponseSuccess(ctx, "Pegawai user berhasil diperbarui", data)
}
