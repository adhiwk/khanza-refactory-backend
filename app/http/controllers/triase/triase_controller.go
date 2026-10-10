package triase

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/triase"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/triase"
	model "goravel/app/models/triase"
	repo "goravel/app/repository/triase"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Index ?tgl_awal=&tgl_akhir=&search=
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.List(repo.Filter{TglAwal: q.Query("tgl_awal"), TglAkhir: q.Query("tgl_akhir"), Search: q.Query("search")}, page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data triase", err)
	}
	return c.ResponsePaginated(ctx, "Data triase berhasil diambil", data, page, limit, total)
}

// Show ?no_rawat=
func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil triase", err)
	}
	return c.ResponseSuccess(ctx, "Detail triase berhasil diambil", data)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	in := input(req.Data)
	in.NoRawat = req.NoRawat
	data, err := c.action.Create(in, c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan triase", err)
	}
	return c.ResponseCreated(ctx, "Triase berhasil disimpan", data)
}

// Update ?no_rawat=
func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Query("no_rawat"), input(req.Data), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui triase", err)
	}
	return c.ResponseSuccess(ctx, "Triase berhasil diperbarui", data)
}

// Destroy ?no_rawat=
func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), c.actor(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus triase", err)
	}
	return c.ResponseSuccess(ctx, "Triase berhasil dihapus", nil)
}

func (c *Controller) actor(ctx http.Context) action.Actor {
	return action.Actor{KodePegawai: c.KodePegawai(ctx), SuperAdmin: c.IsSuperAdmin(ctx)}
}

func input(d request.Data) model.Input {
	return model.Input{
		Utama: model.Utama{
			TglKunjungan: d.TglKunjungan, CaraMasuk: d.CaraMasuk, AlatTransportasi: d.AlatTransportasi,
			AlasanKedatangan: d.AlasanKedatangan, KeteranganKedatangan: d.KeteranganKedatangan, KodeKasus: d.KodeKasus,
			TekananDarah: d.TekananDarah, Nadi: d.Nadi, Pernapasan: d.Pernapasan, Suhu: d.Suhu, SaturasiO2: d.SaturasiO2, Nyeri: d.Nyeri,
		},
		Jenis: d.Jenis,
		Penilaian: model.Penilaian{
			Keluhan: d.Keluhan, KebutuhanKhusus: d.KebutuhanKhusus, Catatan: d.Catatan, Plan: d.Plan,
			TanggalTriase: d.TanggalTriase, Nik: d.Nik,
		},
		SkalaLevel: d.SkalaLevel,
		SkalaKode:  d.SkalaKode,
	}
}
