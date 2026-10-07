package middleware

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type auth struct{}

func (r *auth) Signature() string {
	return "auth"
}

func (r *auth) Handle(ctx http.Context) {
	token := ctx.Request().Header("Authorization", "")
	if token == "" {
		ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if _, err := facades.Auth(ctx).Parse(token); err != nil {
		ctx.Request().AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Request().Next()
}

func Auth() http.Middleware {
	return &auth{}
}
