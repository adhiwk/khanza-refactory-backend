package routes

import (
	"goravel/app/modules/auth"

	"github.com/goravel/framework/contracts/route"
)

func registerAuthRoutes(router route.Router) {
	c := auth.NewController()

	router.Post("/login", c.Login)
	router.Post("/register", c.Register)
}
