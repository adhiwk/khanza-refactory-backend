package routes

import (
	"github.com/goravel/framework/contracts/route"

	"goravel/app/http/controllers/obat"
	"goravel/app/modules/rbac"
)

func registerObatRoutes(router route.Router) {
	c := obat.NewController()

	router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/obat", c.Index)
	router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/obat", c.Store)
	router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/obat/{kode_brng}", c.Show)
	router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/obat/{kode_brng}", c.Update)
	router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/obat/{kode_brng}", c.Destroy)
	router.Middleware(rbac.RequirePermission("farmasi_master.update")).Patch("/obat/{kode_brng}/status", c.UpdateStatus)
}
