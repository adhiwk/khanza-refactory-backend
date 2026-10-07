package rbac

import (
	"context"
	"fmt"
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type ctxKey string

// ContextUserIDKey: isi context.WithValue(ctx, rbac.ContextUserIDKey, uint(1))
// bila memanggil Gate di luar request HTTP (job, command, dll).
const ContextUserIDKey ctxKey = "rbac_user_id"

// UserID mengambil id user dari guard auth (JWT) yang sudah di-parse middleware auth.
func UserID(ctx http.Context) (uint, bool) {
	id, err := facades.Auth(ctx).ID()
	if err != nil {
		return 0, false
	}
	n, err := strconv.ParseUint(fmt.Sprint(id), 10, 64)
	if err != nil || n == 0 {
		return 0, false
	}
	return uint(n), true
}

func userIDFromContext(ctx context.Context) (uint, bool) {
	if v, ok := ctx.Value(ContextUserIDKey).(uint); ok {
		return v, true
	}
	if hc, ok := ctx.(http.Context); ok {
		return UserID(hc)
	}
	return 0, false
}
