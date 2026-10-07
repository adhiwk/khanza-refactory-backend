package routes

import (
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	// Public Routes
	facades.Route().Prefix("api/v1").Group(func(router route.Router) {
		registerAuthRoutes(router)
	})

	// Protected Routes (wajib login, lalu dicek RBAC per route)
	facades.Route().Prefix("api/v1").Middleware(middleware.Auth()).Group(func(router route.Router) {
		registerUserRoutes(router)
		registerRbacRoutes(router)
		registerObatRoutes(router)
		registerPasienRoutes(router)
	})
}
