package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"

	masterkodeaction "goravel/app/actions/masterkode"
	"goravel/app/http/controllers/icd"
	"goravel/app/http/controllers/masterkode"
	rmctrl "goravel/app/http/controllers/rekammedis"
	"goravel/app/http/controllers/rekonsiliasi"
	"goravel/app/http/controllers/riwayatpasien"
	"goravel/app/http/controllers/skriningralan"
	"goravel/app/http/controllers/templatedokter"
	"goravel/app/http/controllers/templateedukasi"
	"goravel/app/http/controllers/templateoperasi"
	"goravel/app/http/controllers/templateradiologi"
	"goravel/app/http/controllers/triase"
	icdmodel "goravel/app/models/icd"
	"goravel/app/modules/rbac"
)

func registerRekamMedisRoutes(router route.Router) {
	perm := rbac.RequirePermission

	rw := rmctrl.NewRiwayatController()
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/riwayat", rw.Show)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/riwayat/cetak", rw.Cetak)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/forms", rw.Forms)

	sk := skriningralan.NewController()
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/skrining-rawat-jalan", sk.Index)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/skrining-rawat-jalan/detail", sk.Show)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/skrining-rawat-jalan/cetak", sk.Cetak)
	router.Middleware(perm("rekam_medis.create")).Post("/rekam-medis/skrining-rawat-jalan", sk.Store)
	router.Middleware(perm("rekam_medis.update")).Put("/rekam-medis/skrining-rawat-jalan", sk.Update)
	router.Middleware(perm("rekam_medis.delete")).Delete("/rekam-medis/skrining-rawat-jalan", sk.Destroy)

	rp := riwayatpasien.NewController()
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/riwayat-persalinan", rp.Persalinan)
	router.Middleware(perm("rekam_medis.create")).Post("/rekam-medis/riwayat-persalinan", rp.TambahPersalinan)
	router.Middleware(perm("rekam_medis.delete")).Delete("/rekam-medis/riwayat-persalinan", rp.HapusPersalinan)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/riwayat-imunisasi", rp.Imunisasi)
	router.Middleware(perm("rekam_medis.create")).Post("/rekam-medis/riwayat-imunisasi", rp.TambahImunisasi)
	router.Middleware(perm("rekam_medis.delete")).Delete("/rekam-medis/riwayat-imunisasi", rp.HapusImunisasi)

	for _, j := range []icdmodel.Jenis{icdmodel.Diagnosa, icdmodel.Prosedur} {
		c := icd.NewController(j)
		p := "/rekam-medis/" + string(j)
		router.Middleware(perm("rekam_medis.view")).Get(p, c.Index)
		router.Middleware(perm("rekam_medis.view")).Get(p+"/referensi", c.Referensi)
		router.Middleware(perm("rekam_medis.create")).Post(p, c.Store)
		router.Middleware(perm("rekam_medis.delete")).Delete(p, c.Destroy)
		if j == icdmodel.Diagnosa {
			router.Middleware(perm("rekam_medis.update")).Patch(p+"/status-penyakit", c.StatusPenyakit)
		}
	}

	t := triase.NewController()
	router.Middleware(perm("igd.view")).Get("/igd/triase", t.Index)
	router.Middleware(perm("igd.view")).Get("/igd/triase/detail", t.Show)
	router.Middleware(perm("igd.view")).Get("/igd/triase/cetak", t.Cetak)
	router.Middleware(perm("igd.create")).Post("/igd/triase", t.Store)
	router.Middleware(perm("igd.update")).Put("/igd/triase", t.Update)
	router.Middleware(perm("igd.delete")).Delete("/igd/triase", t.Destroy)

	rk := rekonsiliasi.NewController()
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/rekonsiliasi-obat", rk.Index)
	router.Middleware(perm("rekam_medis.create")).Post("/rekam-medis/rekonsiliasi-obat", rk.Store)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/rekonsiliasi-obat/{no_rekonsiliasi}", rk.Show)
	router.Middleware(perm("rekam_medis.view")).Get("/rekam-medis/rekonsiliasi-obat/{no_rekonsiliasi}/cetak", rk.Cetak)
	router.Middleware(perm("rekam_medis.update")).Put("/rekam-medis/rekonsiliasi-obat/{no_rekonsiliasi}", rk.Update)
	router.Middleware(perm("rekam_medis.delete")).Delete("/rekam-medis/rekonsiliasi-obat/{no_rekonsiliasi}", rk.Destroy)
	router.Middleware(perm(rekonsiliasi.PermissionKonfirmasi)).Put("/rekam-medis/rekonsiliasi-obat/{no_rekonsiliasi}/konfirmasi", rk.Konfirmasi)

	for _, tpl := range []struct {
		slug string
		c    templateHandler
	}{
		{"hasil-radiologi", templateradiologi.NewController()},
		{"laporan-operasi", templateoperasi.NewController()},
		{"informasi-edukasi", templateedukasi.NewController()},
		{"pemeriksaan-dokter", templatedokter.NewController()},
	} {
		p := "/rekam-medis/template/" + tpl.slug
		router.Middleware(perm("rekam_medis.view")).Get(p, tpl.c.Index)
		router.Middleware(perm("rekam_medis.create")).Post(p, tpl.c.Store)
		router.Middleware(perm("rekam_medis.view")).Get(p+"/{no_template}", tpl.c.Show)
		router.Middleware(perm("rekam_medis.update")).Put(p+"/{no_template}", tpl.c.Update)
		router.Middleware(perm("rekam_medis.delete")).Delete(p+"/{no_template}", tpl.c.Destroy)
	}

	for _, spec := range masterkodeaction.Specs {
		c := masterkode.NewController(spec)
		p := "/rekam-medis/master/" + spec.Slug
		router.Middleware(perm("masterdata.view")).Get(p, c.Index)
		router.Middleware(perm("masterdata.create")).Post(p, c.Store)
		router.Middleware(perm("masterdata.view")).Get(p+"/{kode}", c.Show)
		router.Middleware(perm("masterdata.update")).Put(p+"/{kode}", c.Update)
		router.Middleware(perm("masterdata.delete")).Delete(p+"/{kode}", c.Destroy)
	}

	registerRekamMedisAsesmenRoutes(router)
}

// templateHandler endpoint standar master template.
type templateHandler interface {
	Index(ctx http.Context) http.Response
	Show(ctx http.Context) http.Response
	Store(ctx http.Context) http.Response
	Update(ctx http.Context) http.Response
	Destroy(ctx http.Context) http.Response
}
