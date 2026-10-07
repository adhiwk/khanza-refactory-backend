package rbac

import (
	"errors"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

type Controller struct{}

func NewController() *Controller { return &Controller{} }

type payload struct {
	Name        string   `form:"name" json:"name"`
	Roles       []string `form:"roles" json:"roles"`
	Permissions []string `form:"permissions" json:"permissions"`
}

func parseID(s string) (uint, bool) {
	n, err := strconv.ParseUint(s, 10, 64)
	return uint(n), err == nil && n > 0
}

func fail(ctx http.Context, err error) http.Response {
	if errors.Is(err, ErrRoleNotFound) || errors.Is(err, ErrPermissionNotFound) {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": err.Error()})
	}
	return ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Terjadi kesalahan"})
}

func isSuperAdmin(ctx http.Context) bool {
	uid, ok := UserID(ctx)
	if !ok {
		return false
	}
	has, err := Default().HasRole(uid, SuperAdminRole)
	return err == nil && has
}

// ---------- master data ----------

func (c *Controller) Roles(ctx http.Context) http.Response {
	roles, err := Default().Roles()
	if err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": roles})
}

func (c *Controller) Permissions(ctx http.Context) http.Response {
	perms, err := Default().Permissions()
	if err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": perms})
}

func (c *Controller) StoreRole(ctx http.Context) http.Response {
	var p payload
	if err := ctx.Request().Bind(&p); err != nil || p.Name == "" {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": "name wajib diisi"})
	}
	role, err := Default().CreateRole(p.Name)
	if err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Json(http.StatusCreated, http.Json{"data": role})
}

func (c *Controller) StorePermission(ctx http.Context) http.Response {
	var p payload
	if err := ctx.Request().Bind(&p); err != nil || p.Name == "" {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": "name wajib diisi"})
	}
	perm, err := Default().CreatePermission(p.Name)
	if err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Json(http.StatusCreated, http.Json{"data": perm})
}

// PUT /roles/{name}/permissions  body: {"permissions": ["users.view", ...]}
func (c *Controller) SyncRolePermissions(ctx http.Context) http.Response {
	var p payload
	if err := ctx.Request().Bind(&p); err != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": "payload tidak valid"})
	}
	if err := Default().SyncRolePermissions(ctx.Request().Route("name"), p.Permissions...); err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Permission role diperbarui"})
}

// ---------- akses per user ----------

// GET /users/{id}/access
func (c *Controller) UserAccess(ctx http.Context) http.Response {
	id, ok := parseID(ctx.Request().Route("id"))
	if !ok {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": "id tidak valid"})
	}
	roles, err := Default().GetRoleNames(id)
	if err != nil {
		return fail(ctx, err)
	}
	perms, err := Default().GetAllPermissions(id)
	if err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": http.Json{"roles": roles, "permissions": perms}})
}

// PUT /users/{id}/roles  body: {"roles": ["admin","dokter"]}
func (c *Controller) SyncUserRoles(ctx http.Context) http.Response {
	id, ok := parseID(ctx.Request().Route("id"))
	var p payload
	if !ok || ctx.Request().Bind(&p) != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": "payload tidak valid"})
	}
	// Cegah eskalasi hak akses: hanya super-admin yang boleh memberi role super-admin.
	for _, r := range p.Roles {
		if r == SuperAdminRole && !isSuperAdmin(ctx) {
			return ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Hanya super-admin yang boleh memberi role super-admin"})
		}
	}
	if err := Default().SyncRoles(id, p.Roles...); err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Role user diperbarui"})
}

// PUT /users/{id}/permissions  body: {"permissions": ["reports.view"]}
func (c *Controller) SyncUserPermissions(ctx http.Context) http.Response {
	id, ok := parseID(ctx.Request().Route("id"))
	var p payload
	if !ok || ctx.Request().Bind(&p) != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{"message": "payload tidak valid"})
	}
	if err := Default().SyncPermissions(id, p.Permissions...); err != nil {
		return fail(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Permission user diperbarui"})
}
