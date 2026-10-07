package routes

import (
	"goravel/app/modules/rbac"

	"github.com/goravel/framework/contracts/route"
)

func registerRbacRoutes(router route.Router) {
	c := rbac.NewController()

	// Master role & permission
	router.Middleware(rbac.RequirePermission("roles.view")).Get("/roles", c.Roles)
	router.Middleware(rbac.RequirePermission("roles.view")).Get("/permissions", c.Permissions)
	router.Middleware(rbac.RequirePermission("roles.manage")).Post("/roles", c.StoreRole)
	router.Middleware(rbac.RequirePermission("roles.manage")).Post("/permissions", c.StorePermission)
	router.Middleware(rbac.RequirePermission("roles.manage")).Put("/roles/{name}/permissions", c.SyncRolePermissions)

	// Akses per user
	router.Middleware(rbac.RequirePermission("users.view")).Get("/users/{id}/access", c.UserAccess)
	router.Middleware(rbac.RequirePermission("roles.manage")).Put("/users/{id}/roles", c.SyncUserRoles)
	router.Middleware(rbac.RequirePermission("roles.manage")).Put("/users/{id}/permissions", c.SyncUserPermissions)
}
