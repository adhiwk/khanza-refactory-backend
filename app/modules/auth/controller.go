package auth

import (
	"errors"
	"goravel/app/http/controllers"
	userrepo "goravel/app/repository/user"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type Controller struct {
	controllers.BaseController
	action *Action
}

func NewController() *Controller {
	return &Controller{
		action: NewAction(userrepo.NewRepository()),
	}
}

func (c *Controller) Login(ctx http.Context) http.Response {
	var req LoginRequest

	// Tambahkan ctx sebagai parameter pertama pada facades.Validation().Make
	validator, err := facades.Validation().Make(ctx, ctx.Request().All(), req.Rules(ctx))
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memproses validasi", err.Error())
	}

	if validator.Fails() {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", validator.Errors().All())
	}

	if err := validator.Bind(&req); err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "Format request tidak valid", err.Error())
	}

	token, err := c.action.Login(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			return c.ResponseError(ctx, http.StatusUnauthorized, err.Error(), nil)
		}
		return c.ResponseError(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	return c.ResponseSuccess(ctx, "Login berhasil", map[string]any{
		"token": token,
	})
}

func (c *Controller) Register(ctx http.Context) http.Response {
	var req RegisterRequest

	// Tambahkan ctx sebagai parameter pertama pada facades.Validation().Make
	validator, err := facades.Validation().Make(ctx, ctx.Request().All(), req.Rules(ctx))
	if err != nil {
		return c.ResponseError(ctx, http.StatusInternalServerError, "Gagal memproses validasi", err.Error())
	}

	if validator.Fails() {
		return c.ResponseError(ctx, http.StatusUnprocessableEntity, "Validasi gagal", validator.Errors().All())
	}

	if err := validator.Bind(&req); err != nil {
		return c.ResponseError(ctx, http.StatusBadRequest, "Format request tidak valid", err.Error())
	}

	newUser, err := c.action.Register(req.Name, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, ErrEmailTaken) {
			return c.ResponseError(ctx, http.StatusConflict, err.Error(), nil)
		}
		return c.ResponseError(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	return c.ResponseSuccess(ctx, "Registrasi berhasil", newUser)
}

func (c *Controller) Me(ctx http.Context) http.Response {
	u, err := c.action.Me(ctx)
	if err != nil {
		return c.ResponseError(ctx, http.StatusUnauthorized, "Tidak terautentikasi", nil)
	}

	return c.ResponseSuccess(ctx, "Data user berhasil diambil", u)
}
