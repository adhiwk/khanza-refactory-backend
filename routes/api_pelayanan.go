package routes

import (
	"github.com/goravel/framework/contracts/route"

	"goravel/app/http/controllers/dpjp"
	"goravel/app/http/controllers/igd"
	"goravel/app/http/controllers/kamarinap"
	"goravel/app/http/controllers/pemberianobat"
	"goravel/app/http/controllers/pemeriksaan"
	"goravel/app/http/controllers/resep"
	"goravel/app/http/controllers/rujukaninternal"
	"goravel/app/http/controllers/tindakan"
	"goravel/app/models/perawatan"
	"goravel/app/modules/rbac"
)

// no_rawat (yyyy/MM/dd/NNNNNN) mengandung "/", jadi identitas baris dikirim lewat query string / body.
func registerPelayananRoutes(router route.Router) {
	perm := rbac.RequirePermission

	igdCtrl := igd.NewController()
	router.Middleware(perm("igd.view")).Get("/igd", igdCtrl.Index)
	router.Middleware(perm("igd.create")).Post("/igd", igdCtrl.Store)

	for _, r := range []struct {
		prefix, perm string
		rawat        perawatan.Rawat
	}{
		{"/rawat-jalan", "rawat_jalan", perawatan.Ralan},
		{"/rawat-inap", "rawat_inap", perawatan.Ranap},
	} {
		t := tindakan.NewController(r.rawat)
		router.Middleware(perm(r.perm+".view")).Get(r.prefix+"/tindakan", t.Index)
		router.Middleware(perm(r.perm+".view")).Get(r.prefix+"/tindakan/tarif", t.Tarif)
		router.Middleware(perm(r.perm+".create")).Post(r.prefix+"/tindakan", t.Store)
		router.Middleware(perm(r.perm+".delete")).Delete(r.prefix+"/tindakan", t.Destroy)

		p := pemeriksaan.NewController(r.rawat)
		router.Middleware(perm(r.perm+".view")).Get(r.prefix+"/pemeriksaan", p.Index)
		router.Middleware(perm(r.perm+".create")).Post(r.prefix+"/pemeriksaan", p.Store)
		router.Middleware(perm(r.perm+".update")).Put(r.prefix+"/pemeriksaan", p.Update)
		router.Middleware(perm(r.perm+".delete")).Delete(r.prefix+"/pemeriksaan", p.Destroy)

		o := pemberianobat.NewController(r.rawat)
		router.Middleware(perm(r.perm+".view")).Get(r.prefix+"/obat", o.Index)
		router.Middleware(perm(r.perm+".view")).Get(r.prefix+"/obat/cari", o.Cari)
		router.Middleware(perm(r.perm+".create")).Post(r.prefix+"/obat", o.Store)
		router.Middleware(perm(r.perm+".delete")).Delete(r.prefix+"/obat", o.Destroy)
	}

	ri := rujukaninternal.NewController()
	router.Middleware(perm("rawat_jalan.view")).Get("/rawat-jalan/rujukan-internal", ri.Index)
	router.Middleware(perm("rawat_jalan.create")).Post("/rawat-jalan/rujukan-internal", ri.Store)
	router.Middleware(perm("rawat_jalan.delete")).Delete("/rawat-jalan/rujukan-internal", ri.Destroy)

	k := kamarinap.NewController()
	router.Middleware(perm("rawat_inap.view")).Get("/rawat-inap/kamar", k.Index)
	router.Middleware(perm("rawat_inap.view")).Get("/rawat-inap/kamar/riwayat", k.Riwayat)
	router.Middleware(perm("rawat_inap.create")).Post("/rawat-inap/kamar/masuk", k.Masuk)
	router.Middleware(perm("rawat_inap.update")).Post("/rawat-inap/kamar/pindah", k.Pindah)
	router.Middleware(perm("rawat_inap.update")).Post("/rawat-inap/kamar/pulang", k.Pulang)
	router.Middleware(perm("rawat_inap.update")).Post("/rawat-inap/kamar/batal-pulang", k.BatalPulang)

	d := dpjp.NewController()
	router.Middleware(perm("rawat_inap.view")).Get("/rawat-inap/dpjp", d.Index)
	router.Middleware(perm("rawat_inap.create")).Post("/rawat-inap/dpjp", d.Store)
	router.Middleware(perm("rawat_inap.delete")).Delete("/rawat-inap/dpjp", d.Destroy)

	rs := resep.NewController()
	router.Middleware(perm("resep.view")).Get("/resep", rs.Index)
	router.Middleware(perm("resep.view")).Get("/resep/{no_resep}", rs.Show)
	router.Middleware(perm("resep.create")).Post("/resep", rs.Store)
	router.Middleware(perm("resep.update")).Put("/resep/{no_resep}", rs.Update)
	router.Middleware(perm("resep.delete")).Delete("/resep/{no_resep}", rs.Destroy)

	registerPelayananCrudRoutes(router)
}
