package rekonsiliasi

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rekonsiliasi"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/rekonsiliasi"
	model "goravel/app/models/rekonsiliasi"
	"goravel/app/modules/rbac"
	cetakrepo "goravel/app/repository/cetak"
	repo "goravel/app/repository/rekonsiliasi"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

// PermissionKonfirmasi hak konfirmasi rekonsiliasi obat (akses.getkonfirmasi_rekonsiliasi_obat).
const PermissionKonfirmasi = "rekonsiliasi.konfirmasi"

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository(), cetaksvc.NewService(cetakrepo.NewRepository()))}
}

// Index ?no_rawat=&no_rkm_medis=&tgl_awal=&tgl_akhir=&search=
func (c *Controller) Index(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.List(repo.Filter{NoRawat: q.Query("no_rawat"), NoRkmMedis: q.Query("no_rkm_medis"),
		TglAwal: q.Query("tgl_awal"), TglAkhir: q.Query("tgl_akhir"), Search: q.Query("search")}, page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data rekonsiliasi obat", err)
	}
	return c.ResponsePaginated(ctx, "Data rekonsiliasi obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Show(ctx http.Context) http.Response {
	data, err := c.action.Detail(ctx.Request().Route("no_rekonsiliasi"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil rekonsiliasi obat", err)
	}
	return c.ResponseSuccess(ctx, "Detail rekonsiliasi obat berhasil diambil", data)
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
		return c.ResponseActionError(ctx, "Gagal menyimpan rekonsiliasi obat", err)
	}
	return c.ResponseCreated(ctx, "Rekonsiliasi obat berhasil disimpan", data)
}

func (c *Controller) Update(ctx http.Context) http.Response {
	var req request.UpdateRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Update(ctx.Request().Route("no_rekonsiliasi"), input(req.Data), c.actor(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal memperbarui rekonsiliasi obat", err)
	}
	return c.ResponseSuccess(ctx, "Rekonsiliasi obat berhasil diperbarui", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Route("no_rekonsiliasi"), c.actor(ctx)); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus rekonsiliasi obat", err)
	}
	return c.ResponseSuccess(ctx, "Rekonsiliasi obat berhasil dihapus", nil)
}

func (c *Controller) Konfirmasi(ctx http.Context) http.Response {
	var req request.KonfirmasiRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Konfirmasi(ctx.Request().Route("no_rekonsiliasi"), model.Konfirmasi{
		DiterimaFarmasi: req.DiterimaFarmasi, DikonfirmasiApoteker: req.DikonfirmasiApoteker, DiserahkanPasien: req.DiserahkanPasien, Nip: req.Nip,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan konfirmasi", err)
	}
	return c.ResponseSuccess(ctx, "Konfirmasi rekonsiliasi obat berhasil disimpan", data)
}

func (c *Controller) Cetak(ctx http.Context) http.Response {
	dok, err := c.action.Cetak(ctx.Request().Route("no_rekonsiliasi"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mencetak rekonsiliasi obat", err)
	}
	return c.ResponseCetak(ctx, dok)
}

func (c *Controller) actor(ctx http.Context) action.Actor {
	a := action.Actor{KodePegawai: c.KodePegawai(ctx), SuperAdmin: c.IsSuperAdmin(ctx)}
	if uid, ok := rbac.UserID(ctx); ok {
		a.BolehKonfirmasi = rbac.Default().Can(uid, PermissionKonfirmasi)
	}
	return a
}

func input(d request.Data) model.Rekonsiliasi {
	r := model.Rekonsiliasi{
		TanggalWawancara: d.TanggalWawancara, RekonsiliasiObatSaat: d.RekonsiliasiObatSaat, AlergiObat: d.AlergiObat,
		ManifestasiAlergi: d.ManifestasiAlergi, DampakAlergi: d.DampakAlergi, Nip: d.Nip,
	}
	for _, o := range d.Obat {
		r.Obat = append(r.Obat, model.Obat(o))
	}
	return r
}
