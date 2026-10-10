package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rekammedis"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// Handler endpoint standar satu form asesmen.
type Handler interface {
	Index(ctx http.Context) http.Response
	Show(ctx http.Context) http.Response
	Store(ctx http.Context) http.Response
	Update(ctx http.Context) http.Response
	Destroy(ctx http.Context) http.Response
}

// Controller form asesmen; no_rawat mengandung "/", sehingga kunci dikirim lewat query string.
type Controller[M any, D any] struct {
	controllers.BaseController
	action    *action.Action[M, D]
	newStore  func() request.StoreRequest[D]
	newUpdate func() request.UpdateRequest[D]
}

func NewController[M any, D any](form *action.Form[M, D], newStore func() request.StoreRequest[D], newUpdate func() request.UpdateRequest[D]) *Controller[M, D] {
	return &Controller[M, D]{
		action:    action.NewAction(form, repo.NewRepository[M](form.Spec)),
		newStore:  newStore,
		newUpdate: newUpdate,
	}
}

// Index ?no_rawat=&no_rkm_medis=&tgl_awal=&tgl_akhir=&search=
func (c *Controller[M, D]) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.List(repo.Filter{
		NoRawat:    q.Query("no_rawat"),
		NoRkmMedis: q.Query("no_rkm_medis"),
		TglAwal:    q.Query("tgl_awal"),
		TglAkhir:   q.Query("tgl_akhir"),
		Search:     q.Query("search"),
	}, page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data "+c.action.Form().Label, err)
	}
	return c.ResponsePaginated(ctx, "Data "+c.action.Form().Label+" berhasil diambil", data, page, limit, total)
}

func (c *Controller[M, D]) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(c.key(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil "+c.action.Form().Label, err)
	}
	return c.ResponseSuccess(ctx, "Detail "+c.action.Form().Label+" berhasil diambil", data)
}

func (c *Controller[M, D]) Store(ctx http.Context) http.Response {
	req := c.newStore()
	if resp := c.Validate(ctx, req); resp != nil {
		return resp
	}
	data, err := c.action.Create(req.KeyValues(), req.Payload(), req.DetailValues(), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan "+c.action.Form().Label, err)
	}
	return c.ResponseCreated(ctx, c.action.Form().Label+" berhasil disimpan", data)
}

func (c *Controller[M, D]) Update(ctx http.Context) http.Response {
	req := c.newUpdate()
	if resp := c.Validate(ctx, req); resp != nil {
		return resp
	}
	data, err := c.action.Update(c.key(ctx), req.Payload(), req.DetailValues(), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui "+c.action.Form().Label, err)
	}
	return c.ResponseSuccess(ctx, c.action.Form().Label+" berhasil diperbarui", data)
}

func (c *Controller[M, D]) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(c.key(ctx), c.actor(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus "+c.action.Form().Label, err)
	}
	return c.ResponseSuccess(ctx, c.action.Form().Label+" berhasil dihapus", nil)
}

func (c *Controller[M, D]) key(ctx http.Context) repo.Key {
	key := repo.Key{}
	for _, k := range c.action.Form().Spec.Keys {
		key[k] = ctx.Request().Query(k)
	}
	return key
}

func (c *Controller[M, D]) actor(ctx http.Context) action.Actor {
	return action.Actor{KodePegawai: c.KodePegawai(ctx), SuperAdmin: c.IsSuperAdmin(ctx)}
}
