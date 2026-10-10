package igd

import (
	"github.com/goravel/framework/contracts/http"

	igdaction "goravel/app/actions/igd"
	regaction "goravel/app/actions/registrasi"
	rujukaction "goravel/app/actions/rujukmasuk"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/igd"
	regrepo "goravel/app/repository/registrasi"
	rujukrepo "goravel/app/repository/rujukmasuk"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	registrasi *regaction.Action
	action     *igdaction.Action
}

func NewController() *Controller {
	reg := regaction.NewAction(regrepo.NewRepository())
	return &Controller{
		registrasi: reg,
		action:     igdaction.NewAction(reg, rujukaction.NewAction(rujukrepo.NewRepository())),
	}
}

// Index daftar registrasi IGD (filter sama dengan /registrasi, poli dikunci IGDK).
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.registrasi.List(regrepo.Filter{
		TglAwal:  q.Input("tgl_awal"),
		TglAkhir: q.Input("tgl_akhir"),
		KdDokter: q.Input("kd_dokter"),
		KdPoli:   regaction.KdPoliIGD,
		Search:   q.Input("search"),
		Page:     page,
		Limit:    limit,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data registrasi IGD", err)
	}
	return c.ResponsePaginated(ctx, "Data registrasi IGD berhasil diambil", data, page, limit, total)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Register(req.NoRkmMedis, req.RegistrasiData, igdaction.Rujukan{
		Perujuk: req.Perujuk,
		Alamat:  req.AlamatPerujuk,
		NoRujuk: req.NoRujuk,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan registrasi IGD", err)
	}
	return c.ResponseCreated(ctx, "Registrasi IGD berhasil disimpan", data)
}
