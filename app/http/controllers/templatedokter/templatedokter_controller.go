package templatedokter

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/templatedokter"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/templatedokter"
	model "goravel/app/models/templatedokter"
	repo "goravel/app/repository/templatedokter"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Index ?search=&kd_dokter=
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	data, total, err := c.action.List(ctx.Request().Query("search"), ctx.Request().Query("kd_dokter"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data template", err)
	}
	return c.ResponsePaginated(ctx, "Data template pemeriksaan dokter berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_template"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil template", err)
	}
	return c.ResponseSuccess(ctx, "Detail template pemeriksaan dokter berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.Request
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(input(req), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan template", err)
	}
	return c.ResponseCreated(ctx, "Template pemeriksaan dokter berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.Request
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_template"), input(req), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui template", err)
	}
	return c.ResponseSuccess(ctx, "Template pemeriksaan dokter berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_template"), c.actor(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus template", err)
	}
	return c.ResponseSuccess(ctx, "Template pemeriksaan dokter berhasil dihapus", nil)
}

func (c *Controller) actor(ctx http.Context) action.Actor {
	return action.Actor{KodePegawai: c.KodePegawai(ctx), SuperAdmin: c.IsSuperAdmin(ctx)}
}

func kode(in []request.Kode) []model.Kode {
	out := make([]model.Kode, len(in))
	for i, k := range in {
		out[i] = model.Kode{Kode: k.Kode, Urut: k.Urut, Jumlah: k.Jumlah}
	}
	return out
}

func input(r request.Request) model.Template {
	t := model.Template{KdDokter: r.KdDokter, Keluhan: r.Keluhan, Pemeriksaan: r.Pemeriksaan, Penilaian: r.Penilaian,
		Rencana: r.Rencana, Instruksi: r.Instruksi, Evaluasi: r.Evaluasi}
	t.Diagnosa, t.Prosedur, t.Radiologi, t.Tindakan = kode(r.Diagnosa), kode(r.Prosedur), kode(r.Radiologi), kode(r.Tindakan)
	for _, l := range r.Lab {
		t.Lab = append(t.Lab, model.Lab{KdJenisPrw: l.KdJenisPrw, IDTemplate: l.IDTemplate})
	}
	for _, o := range r.Resep {
		t.Resep = append(t.Resep, model.Resep{KodeBrng: o.KodeBrng, Jml: o.Jml, AturanPakai: o.AturanPakai})
	}
	for _, rc := range r.Racikan {
		m := model.Racikan{NoRacik: rc.NoRacik, NamaRacik: rc.NamaRacik, KdRacik: rc.KdRacik, JmlDr: rc.JmlDr, AturanPakai: rc.AturanPakai, Keterangan: rc.Keterangan}
		for _, d := range rc.Detail {
			m.Detail = append(m.Detail, model.RacikanDetail{KodeBrng: d.KodeBrng, P1: d.P1, P2: d.P2, Kandungan: d.Kandungan, Jml: d.Jml})
		}
		t.Racikan = append(t.Racikan, m)
	}
	return t
}
