package controllers

import (
	"github.com/goravel/framework/contracts/http"
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
