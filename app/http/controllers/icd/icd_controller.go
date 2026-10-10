package icd

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/icd"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/icd"
	model "goravel/app/models/icd"
	repo "goravel/app/repository/icd"
	"goravel/app/support"
)

// Controller diagnosa atau prosedur pasien.
type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController(jenis model.Jenis) *Controller {
	return &Controller{action: action.NewAction(jenis, repo.NewRepository())}
}

// Index ?no_rawat= atau ?no_rkm_medis= (riwayat seluruh kunjungan)
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"), ctx.Request().Query("no_rkm_medis"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data", err)
	}
	return c.ResponseSuccess(ctx, "Data berhasil diambil", data)
}

// Referensi ?search= kode ICD master.
func (c *Controller) Referensi(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	data, total, err := c.action.CariReferensi(ctx.Request().Query("search"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil referensi ICD", err)
	}
	return c.ResponsePaginated(ctx, "Referensi ICD berhasil diambil", data, page, limit, total)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	items := make([]model.Item, len(req.Items))
	for i, it := range req.Items {
		items[i] = model.Item{Kode: it.Kode, Prioritas: it.Prioritas, Jumlah: it.Jumlah}
	}
	data, err := c.action.Create(req.NoRawat, req.Status, items)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan", err)
	}
	return c.ResponseCreated(ctx, "Data berhasil disimpan", data)
}

// Destroy ?no_rawat=&kode=
func (c *Controller) Destroy(ctx http.Context) http.Response {
	if err := c.action.Delete(ctx.Request().Query("no_rawat"), ctx.Request().Query("kode")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus", err)
	}
	return c.ResponseSuccess(ctx, "Data berhasil dihapus", nil)
}

// StatusPenyakit PATCH (khusus diagnosa).
func (c *Controller) StatusPenyakit(ctx http.Context) http.Response {
	var req request.StatusPenyakitRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	if err := c.action.SetStatusPenyakit(req.NoRawat, req.Kode, req.StatusPenyakit); err != nil {
		return c.ResponseActionError(ctx, "Gagal mengubah status penyakit", err)
	}
	return c.ResponseSuccess(ctx, "Status penyakit berhasil diubah", nil)
}
