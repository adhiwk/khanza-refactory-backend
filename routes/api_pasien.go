package routes

import (
	"goravel/app/http/controllers/pasien"
	"goravel/app/modules/rbac"

	"github.com/goravel/framework/contracts/route"
)

func registerPasienRoutes(router route.Router) {
	c := pasien.NewController()

	router.Middleware(rbac.RequirePermission("patients.view")).Get("/pasien", c.Index)
	router.Middleware(rbac.RequirePermission("patients.create")).Post("/pasien", c.Store)
	router.Middleware(rbac.RequirePermission("patients.view")).Get("/pasien/{no_rkm_medis}", c.Show)
	router.Middleware(rbac.RequirePermission("patients.update")).Put("/pasien/{no_rkm_medis}", c.Update)
	router.Middleware(rbac.RequirePermission("patients.delete")).Delete("/pasien/{no_rkm_medis}", c.Destroy)
}
