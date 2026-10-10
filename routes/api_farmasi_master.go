package routes

import (
	"github.com/goravel/framework/contracts/route"

	"goravel/app/http/controllers/golonganobat"
	"goravel/app/http/controllers/industrifarmasi"
	"goravel/app/http/controllers/jenisobat"
	"goravel/app/http/controllers/kategoriobat"
	"goravel/app/http/controllers/metoderacik"
	"goravel/app/http/controllers/satuanobat"
	"goravel/app/modules/rbac"
)

func registerFarmasiMasterRoutes(router route.Router) {
	{
		c := jenisobat.NewController()
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/jenis", c.Index)
		router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/farmasi/jenis", c.Store)
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/jenis/{kdjns}", c.Show)
		router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/farmasi/jenis/{kdjns}", c.Update)
		router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/farmasi/jenis/{kdjns}", c.Destroy)
	}
	{
		c := kategoriobat.NewController()
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/kategori", c.Index)
		router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/farmasi/kategori", c.Store)
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/kategori/{kode}", c.Show)
		router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/farmasi/kategori/{kode}", c.Update)
		router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/farmasi/kategori/{kode}", c.Destroy)
	}
	{
		c := golonganobat.NewController()
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/golongan", c.Index)
		router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/farmasi/golongan", c.Store)
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/golongan/{kode}", c.Show)
		router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/farmasi/golongan/{kode}", c.Update)
		router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/farmasi/golongan/{kode}", c.Destroy)
	}
	{
		c := satuanobat.NewController()
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/satuan", c.Index)
		router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/farmasi/satuan", c.Store)
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/satuan/{kode_sat}", c.Show)
		router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/farmasi/satuan/{kode_sat}", c.Update)
		router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/farmasi/satuan/{kode_sat}", c.Destroy)
	}
	{
		c := industrifarmasi.NewController()
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/industri", c.Index)
		router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/farmasi/industri", c.Store)
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/industri/{kode_industri}", c.Show)
		router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/farmasi/industri/{kode_industri}", c.Update)
		router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/farmasi/industri/{kode_industri}", c.Destroy)
	}
	{
		c := metoderacik.NewController()
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/metode-racik", c.Index)
		router.Middleware(rbac.RequirePermission("farmasi_master.create")).Post("/farmasi/metode-racik", c.Store)
		router.Middleware(rbac.RequirePermission("farmasi_master.view")).Get("/farmasi/metode-racik/{kd_racik}", c.Show)
		router.Middleware(rbac.RequirePermission("farmasi_master.update")).Put("/farmasi/metode-racik/{kd_racik}", c.Update)
		router.Middleware(rbac.RequirePermission("farmasi_master.delete")).Delete("/farmasi/metode-racik/{kd_racik}", c.Destroy)
	}
}
