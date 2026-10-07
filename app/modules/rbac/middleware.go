package rbac

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
)

// Semua middleware ini harus dipasang SETELAH middleware auth/JWT.
// Nama bisa dipisah "|" seperti Spatie: RequireRole("admin|dokter").

// namedMiddleware adalah adapter agar fungsi biasa memenuhi interface http.Middleware
// (butuh method Signature dan Handle).
type namedMiddleware struct {
	name string
	fn   func(ctx http.Context)
}

func (m namedMiddleware) Signature() string       { return m.name }
func (m namedMiddleware) Handle(ctx http.Context) { m.fn(ctx) }

func split(items []string) []string {
	var out []string
	for _, it := range items {
		for _, p := range strings.Split(it, "|") {
			if p = strings.TrimSpace(p); p != "" {
				out = append(out, p)
			}
		}
	}
	return out
}

func deny(ctx http.Context, status int, msg string) {
	ctx.Request().AbortWithStatusJson(status, http.Json{"message": msg})
}

// RequireRole: lolos bila user punya salah satu role (super-admin selalu lolos).
func RequireRole(roles ...string) http.Middleware {
	want := split(roles)
	return namedMiddleware{name: "rbac.role:" + strings.Join(want, "|"), fn: func(ctx http.Context) {
		uid, ok := UserID(ctx)
		if !ok {
			deny(ctx, http.StatusUnauthorized, "Unauthenticated")
			return
		}
		allowed, err := Default().HasAnyRole(uid, append(append([]string{}, want...), SuperAdminRole)...)
		if err != nil || !allowed {
			deny(ctx, http.StatusForbidden, "Anda tidak memiliki role yang dibutuhkan")
			return
		}
		ctx.Request().Next()
	}}
}

// RequirePermission: lolos bila user punya salah satu permission (langsung / via role).
func RequirePermission(perms ...string) http.Middleware {
	want := split(perms)
	return namedMiddleware{name: "rbac.permission:" + strings.Join(want, "|"), fn: func(ctx http.Context) {
		uid, ok := UserID(ctx)
		if !ok {
			deny(ctx, http.StatusUnauthorized, "Unauthenticated")
			return
		}
		for _, p := range want {
			if Default().Can(uid, p) {
				ctx.Request().Next()
				return
			}
		}
		deny(ctx, http.StatusForbidden, "Anda tidak memiliki izin untuk aksi ini")
	}}
}

// RequireRoleOrPermission: lolos bila salah satu role ATAU salah satu permission terpenuhi.
func RequireRoleOrPermission(items ...string) http.Middleware {
	want := split(items)
	return namedMiddleware{name: "rbac.role_or_permission:" + strings.Join(want, "|"), fn: func(ctx http.Context) {
		uid, ok := UserID(ctx)
		if !ok {
			deny(ctx, http.StatusUnauthorized, "Unauthenticated")
			return
		}
		if has, err := Default().HasAnyRole(uid, append(append([]string{}, want...), SuperAdminRole)...); err == nil && has {
			ctx.Request().Next()
			return
		}
		for _, p := range want {
			if Default().Can(uid, p) {
				ctx.Request().Next()
				return
			}
		}
		deny(ctx, http.StatusForbidden, "Akses ditolak")
	}}
}
