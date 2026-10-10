package routes

import (
	"github.com/goravel/framework/contracts/route"

	"goravel/app/http/controllers/bahasapasien"
	"goravel/app/http/controllers/bangsal"
	"goravel/app/http/controllers/cacatfisik"
	"goravel/app/http/controllers/dokter"
	"goravel/app/http/controllers/jabatan"
	"goravel/app/http/controllers/kabupaten"
	"goravel/app/http/controllers/kamar"
	"goravel/app/http/controllers/kategoriperawatan"
	"goravel/app/http/controllers/kecamatan"
	"goravel/app/http/controllers/kelurahan"
	"goravel/app/http/controllers/penjab"
	"goravel/app/http/controllers/perusahaanpasien"
	"goravel/app/http/controllers/poliklinik"
	"goravel/app/http/controllers/propinsi"
	"goravel/app/http/controllers/spesialis"
	"goravel/app/http/controllers/sukubangsa"
	"goravel/app/modules/rbac"
)

func registerMasterRoutes(router route.Router) {
	{
		c := poliklinik.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/poliklinik", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/poliklinik", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/poliklinik/{kd_poli}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/poliklinik/{kd_poli}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/poliklinik/{kd_poli}", c.Destroy)
	}
	{
		c := bangsal.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/bangsal", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/bangsal", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/bangsal/{kd_bangsal}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/bangsal/{kd_bangsal}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/bangsal/{kd_bangsal}", c.Destroy)
	}
	{
		c := kamar.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kamar", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/kamar", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kamar/{kd_kamar}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/kamar/{kd_kamar}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/kamar/{kd_kamar}", c.Destroy)
	}
	{
		c := spesialis.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/spesialis", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/spesialis", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/spesialis/{kd_sps}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/spesialis/{kd_sps}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/spesialis/{kd_sps}", c.Destroy)
	}
	{
		c := penjab.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/penjab", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/penjab", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/penjab/{kd_pj}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/penjab/{kd_pj}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/penjab/{kd_pj}", c.Destroy)
	}
	{
		c := perusahaanpasien.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/perusahaan-pasien", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/perusahaan-pasien", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/perusahaan-pasien/{kode_perusahaan}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/perusahaan-pasien/{kode_perusahaan}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/perusahaan-pasien/{kode_perusahaan}", c.Destroy)
	}
	{
		c := sukubangsa.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/suku-bangsa", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/suku-bangsa", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/suku-bangsa/{id}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/suku-bangsa/{id}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/suku-bangsa/{id}", c.Destroy)
	}
	{
		c := bahasapasien.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/bahasa-pasien", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/bahasa-pasien", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/bahasa-pasien/{id}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/bahasa-pasien/{id}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/bahasa-pasien/{id}", c.Destroy)
	}
	{
		c := cacatfisik.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/cacat-fisik", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/cacat-fisik", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/cacat-fisik/{id}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/cacat-fisik/{id}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/cacat-fisik/{id}", c.Destroy)
	}
	{
		c := propinsi.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/propinsi", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/propinsi", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/propinsi/{kd_prop}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/propinsi/{kd_prop}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/propinsi/{kd_prop}", c.Destroy)
	}
	{
		c := kabupaten.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kabupaten", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/kabupaten", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kabupaten/{kd_kab}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/kabupaten/{kd_kab}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/kabupaten/{kd_kab}", c.Destroy)
	}
	{
		c := kecamatan.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kecamatan", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/kecamatan", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kecamatan/{kd_kec}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/kecamatan/{kd_kec}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/kecamatan/{kd_kec}", c.Destroy)
	}
	{
		c := kelurahan.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kelurahan", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/kelurahan", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kelurahan/{kd_kel}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/kelurahan/{kd_kel}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/kelurahan/{kd_kel}", c.Destroy)
	}
	{
		c := jabatan.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/jabatan", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/jabatan", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/jabatan/{kd_jbtn}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/jabatan/{kd_jbtn}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/jabatan/{kd_jbtn}", c.Destroy)
	}
	{
		c := kategoriperawatan.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kategori-perawatan", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/kategori-perawatan", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/kategori-perawatan/{kd_kategori}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/kategori-perawatan/{kd_kategori}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/kategori-perawatan/{kd_kategori}", c.Destroy)
	}
	{
		c := dokter.NewController()
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/dokter", c.Index)
		router.Middleware(rbac.RequirePermission("masterdata.create")).Post("/master/dokter", c.Store)
		router.Middleware(rbac.RequirePermission("masterdata.view")).Get("/master/dokter/{kd_dokter}", c.Show)
		router.Middleware(rbac.RequirePermission("masterdata.update")).Put("/master/dokter/{kd_dokter}", c.Update)
		router.Middleware(rbac.RequirePermission("masterdata.delete")).Delete("/master/dokter/{kd_dokter}", c.Destroy)
	}
}
