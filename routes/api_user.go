package routes

import (
	"goravel/app/modules/rbac"
	"goravel/app/http/controllers/user"

	"github.com/goravel/framework/contracts/route"
)

func registerUserRoutes(router route.Router) {
	c := user.NewController()

	router.Middleware(rbac.RequirePermission("users.view")).Get("/users", c.Index)
	router.Middleware(rbac.RequirePermission("users.create")).Post("/users", c.Store)
	router.Middleware(rbac.RequirePermission("users.view")).Get("/users/{id}", c.Show)
	router.Middleware(rbac.RequirePermission("users.update")).Put("/users/{id}", c.Update)
	router.Middleware(rbac.RequirePermission("users.delete")).Delete("/users/{id}", c.Destroy)
}
