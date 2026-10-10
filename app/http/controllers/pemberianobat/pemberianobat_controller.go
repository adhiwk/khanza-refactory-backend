package pemberianobat

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	action "goravel/app/actions/pemberianobat"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/pemberianobat"
	model "goravel/app/models/pemberianobat"
	"goravel/app/models/perawatan"
	akunrepo "goravel/app/repository/akun"
	jurnalrepo "goravel/app/repository/jurnal"
	repo "goravel/app/repository/pemberianobat"
	stokrepo "goravel/app/repository/stok"
	jurnalsvc "goravel/app/services/jurnal"
	stoksvc "goravel/app/services/stok"
	"goravel/app/support"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController(rawat perawatan.Rawat) *Controller {
	cfg := facades.Config()
	set := action.Setting{
		Batch:           cfg.GetBool("farmasi.aktifkan_batch"),
		Hpp:             cfg.GetString("farmasi.hpp", "dasar"),
		PembulatanHarga: cfg.GetBool("farmasi.pembulatan_harga"),
	}
	return &Controller{action: action.NewAction(rawat, set, repo.NewRepository(), akunrepo.NewRepository(),
		stoksvc.NewService(stokrepo.NewRepository(), set.Batch), jurnalsvc.NewService(jurnalrepo.NewRepository()))}
}

// Index ?no_rawat=
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data pemberian obat", err)
	}
	return c.ResponseSuccess(ctx, "Data pemberian obat berhasil diambil", data)
}

// Cari ?no_rawat=&kd_bangsal=&jenis_harga=&search= obat berstok di depo beserta harga jual pasien.
func (c *Controller) Cari(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.CariObat(q.Query("no_rawat"), q.Query("kd_bangsal"), q.Query("jenis_harga"), q.Query("search"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data obat", err)
	}
	return c.ResponsePaginated(ctx, "Data obat berhasil diambil", data, page, limit, total)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	items := make([]action.Item, len(req.Items))
	for i, it := range req.Items {
		items[i] = action.Item(it)
	}
	data, err := c.action.Create(action.Input{
		NoRawat: req.NoRawat, KdBangsal: req.KdBangsal, Tgl: req.TglPerawatan, Jam: req.Jam,
		JenisHarga: req.JenisHarga, Items: items, Operator: c.operator(ctx),
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan pemberian obat", err)
	}
	return c.ResponseCreated(ctx, "Pemberian obat berhasil disimpan", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	var req request.DeleteRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	err := c.action.Delete(model.Key{
		NoRawat: req.NoRawat, TglPerawatan: req.TglPerawatan, Jam: req.Jam, KodeBrng: req.KodeBrng, NoBatch: req.NoBatch, NoFaktur: req.NoFaktur,
	}, c.operator(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus pemberian obat", err)
	}
	return c.ResponseSuccess(ctx, "Pemberian obat berhasil dihapus", nil)
}

func (c *Controller) operator(ctx http.Context) string {
	return fmt.Sprintf("USER %d", c.UserID(ctx))
}
