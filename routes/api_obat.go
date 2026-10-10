package routes

import (
	"goravel/app/http/controllers/obat"
	"goravel/app/modules/rbac"

	"github.com/goravel/framework/contracts/route"
)

func registerObatRoutes(router route.Router) {
	obatCtrl := obat.NewObatController()

	router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/obat", obatCtrl.Index)
	router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/obat", obatCtrl.Store)
	router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/obat/{kode}", obatCtrl.Show)
	router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/obat/{kode}", obatCtrl.Update)
	router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/obat/{kode}", obatCtrl.Destroy)
	router.Middleware(rbac.RequirePermission("farmasi_master.update")).Patch("/obat/{kode}/status", obatCtrl.UpdateStatus)
}
