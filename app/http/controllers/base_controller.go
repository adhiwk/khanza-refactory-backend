package controllers

import (
	"errors"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/modules/rbac"
	"goravel/app/support"
)

// BaseController menyamakan format response dengan ResponseTrait di backend Laravel lama:
// selalu mengirim 4 key: status, message, errors, data.
type BaseController struct{}

// ResponseSuccess format standard untuk sukses (HTTP 200).
func (c *BaseController) ResponseSuccess(ctx http.Context, message string, data any) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"status":  true,
		"message": message,
		"errors":  nil,
		"data":    data,
	})
}

// ResponseCreated format standard untuk resource baru (HTTP 201).
func (c *BaseController) ResponseCreated(ctx http.Context, message string, data any) http.Response {
	return ctx.Response().Json(http.StatusCreated, http.Json{
		"status":  true,
		"message": message,
		"errors":  nil,
		"data":    data,
	})
}

// ResponseError format standard untuk error.
func (c *BaseController) ResponseError(ctx http.Context, code int, message string, errors any) http.Response {
	return ctx.Response().Json(code, http.Json{
		"status":  false,
		"message": message,
		"errors":  errors,
		"data":    nil,
	})
}

// ResponsePaginated membungkus list dengan format pagination ala Laravel.
func (c *BaseController) ResponsePaginated(ctx http.Context, message string, data any, page, limit int, total int64) http.Response {
	return c.ResponseSuccess(ctx, message, support.FormatLaravelPagination(ctx, data, page, limit, total))
}

// Validate bind + validasi FormRequest; mengembalikan response error bila gagal, nil bila valid.
func (c *BaseController) Validate(ctx http.Context, request http.FormRequest) http.Response {
	errs, err := ctx.Request().ValidateRequest(request)
	if err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "Request tidak valid", err.Error())
	}
	if errs != nil {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", errs.All())
	}
	return nil
}

// ResponseActionError memetakan error dari Action ke HTTP status; error tak dikenal menjadi 500 dengan pesan fallback.
func (c *BaseController) ResponseActionError(ctx http.Context, fallback string, err error) http.Response {
	var appErr *support.AppError
	if errors.As(err, &appErr) {
		status := http.StatusInternalServerError
		switch appErr.Kind {
		case support.KindNotFound:
			status = http.StatusNotFound
		case support.KindConflict:
			status = http.StatusConflict
		case support.KindInvalid:
			status = http.StatusUnprocessableEntity
		case support.KindForbidden:
			status = http.StatusForbidden
		}
		return c.ResponseError(ctx, status, appErr.Message, nil)
	}
	if support.IsDuplicate(err) {
		return c.ResponseError(ctx, http.StatusConflict, "Data sudah ada", nil)
	}
	if support.IsForeignKey(err) {
		return c.ResponseError(ctx, http.StatusConflict, "Data masih dipakai atau referensi tidak ditemukan", nil)
	}
	return c.ResponseError(ctx, http.StatusInternalServerError, fallback, err.Error())
}

// IsSuperAdmin super admin bebas dari batasan waktu ubah/hapus (setara "Admin Utama" di Khanza).
func (c *BaseController) IsSuperAdmin(ctx http.Context) bool {
	uid, ok := rbac.UserID(ctx)
	if !ok {
		return false
	}
	has, err := rbac.Default().HasAnyRole(uid, rbac.SuperAdminRole)
	return err == nil && has
}

// UserID id user login (0 bila tidak ada).
func (c *BaseController) UserID(ctx http.Context) uint {
	uid, _ := rbac.UserID(ctx)
	return uid
}

// KodePegawai pegawai.nik akun login (users.kd_pegawai); kosong bila belum dihubungkan.
func (c *BaseController) KodePegawai(ctx http.Context) string {
	uid, ok := rbac.UserID(ctx)
	if !ok {
		return ""
	}
	var list []struct {
		Kd string `gorm:"column:kd"`
	}
	if err := facades.Orm().Query().Raw("select ifnull(kd_pegawai,'') as kd from users where id=?", uid).Scan(&list); err != nil || len(list) == 0 {
		return ""
	}
	return list[0].Kd
}

// ResponseCetak merender dokumen cetak HTML (resources/views/cetak/dokumen.tmpl).
func (c *BaseController) ResponseCetak(ctx http.Context, dokumen any) http.Response {
	return ctx.Response().View().Make("cetak/dokumen.tmpl", dokumen)
}
