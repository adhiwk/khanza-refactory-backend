package routes

import (
	"github.com/goravel/framework/contracts/route"

	"goravel/app/http/controllers/biayalain"
	"goravel/app/http/controllers/deposit"
	"goravel/app/http/controllers/tagihan"
	model "goravel/app/models/biayalain"
	"goravel/app/modules/rbac"
)

func registerBillingRoutes(router route.Router) {
	perm := rbac.RequirePermission

	t := tagihan.NewController()
	router.Middleware(perm("billing.view")).Get("/billing/tagihan", t.Show)

	d := deposit.NewController()
	router.Middleware(perm("billing.view")).Get("/billing/deposit", d.Index)
	router.Middleware(perm("billing.create")).Post("/billing/deposit", d.Store)
	router.Middleware(perm("billing.delete")).Delete("/billing/deposit/{no_deposit}", d.Destroy)

	for _, j := range []model.Jenis{model.Tambahan, model.Potongan} {
		c := biayalain.NewController(j)
		router.Middleware(perm("billing.view")).Get("/billing/"+string(j), c.Index)
		router.Middleware(perm("billing.create")).Post("/billing/"+string(j), c.Store)
		router.Middleware(perm("billing.delete")).Delete("/billing/"+string(j), c.Destroy)
	}
}
