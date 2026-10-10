package routes

import (
	"github.com/goravel/framework/contracts/route"

	"goravel/app/http/controllers/catatanpasien"
	"goravel/app/http/controllers/pasienmati"
	"goravel/app/http/controllers/rujukkeluar"
	"goravel/app/http/controllers/rujukmasuk"
	"goravel/app/modules/rbac"
)

func registerPelayananCrudRoutes(router route.Router) {
	{
		c := rujukmasuk.NewController()
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/rujuk-masuk", c.Index)
		router.Middleware(rbac.RequirePermission("pelayanan.create")).Post("/rujuk-masuk", c.Store)
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/rujuk-masuk/detail", c.Show)
		router.Middleware(rbac.RequirePermission("pelayanan.update")).Put("/rujuk-masuk", c.Update)
		router.Middleware(rbac.RequirePermission("pelayanan.delete")).Delete("/rujuk-masuk", c.Destroy)
	}
	{
		c := rujukkeluar.NewController()
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/rujuk-keluar", c.Index)
		router.Middleware(rbac.RequirePermission("pelayanan.create")).Post("/rujuk-keluar", c.Store)
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/rujuk-keluar/{no_rujuk}", c.Show)
		router.Middleware(rbac.RequirePermission("pelayanan.update")).Put("/rujuk-keluar/{no_rujuk}", c.Update)
		router.Middleware(rbac.RequirePermission("pelayanan.delete")).Delete("/rujuk-keluar/{no_rujuk}", c.Destroy)
	}
	{
		c := pasienmati.NewController()
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/pasien-meninggal", c.Index)
		router.Middleware(rbac.RequirePermission("pelayanan.create")).Post("/pasien-meninggal", c.Store)
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/pasien-meninggal/{no_rkm_medis}", c.Show)
		router.Middleware(rbac.RequirePermission("pelayanan.update")).Put("/pasien-meninggal/{no_rkm_medis}", c.Update)
		router.Middleware(rbac.RequirePermission("pelayanan.delete")).Delete("/pasien-meninggal/{no_rkm_medis}", c.Destroy)
	}
	{
		c := catatanpasien.NewController()
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/catatan-pasien", c.Index)
		router.Middleware(rbac.RequirePermission("pelayanan.create")).Post("/catatan-pasien", c.Store)
		router.Middleware(rbac.RequirePermission("pelayanan.view")).Get("/catatan-pasien/{no_rkm_medis}", c.Show)
		router.Middleware(rbac.RequirePermission("pelayanan.update")).Put("/catatan-pasien/{no_rkm_medis}", c.Update)
		router.Middleware(rbac.RequirePermission("pelayanan.delete")).Delete("/catatan-pasien/{no_rkm_medis}", c.Destroy)
	}
}
