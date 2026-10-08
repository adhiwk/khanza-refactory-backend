package routes

import (
	"goravel/app/http/controllers/registrasi"
	"goravel/app/modules/rbac"

	"github.com/goravel/framework/contracts/route"
)

// no_rawat (yyyy/MM/dd/NNNNNN) mengandung "/", jadi detail/ubah/hapus memakai ?no_rawat=.
func registerRegistrasiRoutes(router route.Router) {
	c := registrasi.NewController()

	router.Middleware(rbac.RequirePermission("registrations.view")).Get("/registrasi", c.Index)
	router.Middleware(rbac.RequirePermission("registrations.create")).Post("/registrasi", c.Store)
	router.Middleware(rbac.RequirePermission("registrations.view")).Get("/registrasi/detail", c.Show)
	router.Middleware(rbac.RequirePermission("registrations.update")).Put("/registrasi", c.Update)
	router.Middleware(rbac.RequirePermission("registrations.delete")).Delete("/registrasi", c.Destroy)
}
