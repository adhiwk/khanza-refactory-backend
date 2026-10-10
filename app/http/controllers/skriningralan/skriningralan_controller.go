package skriningralan

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/skriningralan"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/skriningralan"
	model "goravel/app/models/skriningralan"
	cetakrepo "goravel/app/repository/cetak"
	repo "goravel/app/repository/skriningralan"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

// Controller skrining rawat jalan; kunci ?tanggal=&jam=&no_rkm_medis=.
type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository(), cetaksvc.NewService(cetakrepo.NewRepository()))}
}

// Index ?no_rkm_medis=&tgl_awal=&tgl_akhir=&search=
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.List(repo.Filter{NoRkmMedis: q.Query("no_rkm_medis"), TglAwal: q.Query("tgl_awal"),
		TglAkhir: q.Query("tgl_akhir"), Search: q.Query("search")}, page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data skrining", err)
	}
	return c.ResponsePaginated(ctx, "Data skrining rawat jalan berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(c.key(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil skrining", err)
	}
	return c.ResponseSuccess(ctx, "Detail skrining rawat jalan berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	in := data(req.Data)
	in.NoRkmMedis, in.Tanggal, in.Jam = req.NoRkmMedis, req.Tanggal, req.Jam
	out, err := c.action.Create(in, c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan skrining", err)
	}
	return c.ResponseCreated(ctx, "Skrining rawat jalan berhasil disimpan", out)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	out, err := c.action.Update(c.key(ctx), data(req.Data), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui skrining", err)
	}
	return c.ResponseSuccess(ctx, "Skrining rawat jalan berhasil diperbarui", out)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(c.key(ctx), c.actor(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus skrining", err)
	}
	return c.ResponseSuccess(ctx, "Skrining rawat jalan berhasil dihapus", nil)
}

func (c *Controller) Cetak(ctx http.Context) http.Response {
	dok, err := c.action.Cetak(c.key(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mencetak skrining", err)
	}
	return c.ResponseCetak(ctx, dok)
}

func (c *Controller) key(ctx http.Context) model.Key {
	q := ctx.Request()
	return model.Key{Tanggal: q.Query("tanggal"), Jam: q.Query("jam"), NoRkmMedis: q.Query("no_rkm_medis")}
}

func (c *Controller) actor(ctx http.Context) action.Actor {
	return action.Actor{KodePegawai: c.KodePegawai(ctx), SuperAdmin: c.IsSuperAdmin(ctx)}
}

func data(d request.Data) model.Skrining {
	return model.Skrining{Geriatri: d.Geriatri, Kesadaran: d.Kesadaran, Pernapasan: d.Pernapasan, NyeriDada: d.NyeriDada,
		SkalaNyeri: d.SkalaNyeri, Batuk: d.Batuk, RisikoJatuh: d.RisikoJatuh, Keputusan: d.Keputusan, Nip: d.Nip}
}
