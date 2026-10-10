package riwayatpasien

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/riwayatpasien"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/riwayatpasien"
	model "goravel/app/models/riwayatpasien"
	repo "goravel/app/repository/riwayatpasien"
)

type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController() *Controller {
	return &Controller{action: action.NewAction(repo.NewRepository())}
}

// Persalinan ?no_rkm_medis=
func (c *Controller) Persalinan(ctx http.Context) http.Response {
	data, err := c.action.Persalinan(ctx.Request().Query("no_rkm_medis"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil riwayat persalinan", err)
	}
	return c.ResponseSuccess(ctx, "Riwayat persalinan berhasil diambil", data)
}

func (c *Controller) TambahPersalinan(ctx http.Context) http.Response {
	var req request.PersalinanRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.TambahPersalinan(model.Persalinan{
		NoRkmMedis: req.NoRkmMedis, TglThn: req.TglThn, TempatPersalinan: req.TempatPersalinan, UsiaHamil: req.UsiaHamil,
		JenisPersalinan: req.JenisPersalinan, Penolong: req.Penolong, Penyulit: req.Penyulit, Jk: req.Jk, Bbpb: req.Bbpb, Keadaan: req.Keadaan,
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan riwayat persalinan", err)
	}
	return c.ResponseCreated(ctx, "Riwayat persalinan berhasil disimpan", data)
}

// HapusPersalinan ?no_rkm_medis=&tgl_thn=
func (c *Controller) HapusPersalinan(ctx http.Context) http.Response {
	if err := c.action.HapusPersalinan(ctx.Request().Query("no_rkm_medis"), ctx.Request().Query("tgl_thn")); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus riwayat persalinan", err)
	}
	return c.ResponseSuccess(ctx, "Riwayat persalinan berhasil dihapus", nil)
}

// Imunisasi ?no_rkm_medis=
func (c *Controller) Imunisasi(ctx http.Context) http.Response {
	data, err := c.action.Imunisasi(ctx.Request().Query("no_rkm_medis"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil riwayat imunisasi", err)
	}
	return c.ResponseSuccess(ctx, "Riwayat imunisasi berhasil diambil", data)
}

func (c *Controller) TambahImunisasi(ctx http.Context) http.Response {
	var req request.ImunisasiRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.TambahImunisasi(model.Imunisasi{NoRkmMedis: req.NoRkmMedis, KodeImunisasi: req.KodeImunisasi, NoImunisasi: req.NoImunisasi})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan riwayat imunisasi", err)
	}
	return c.ResponseCreated(ctx, "Riwayat imunisasi berhasil disimpan", data)
}

// HapusImunisasi ?no_rkm_medis=&kode_imunisasi=&no_imunisasi=
func (c *Controller) HapusImunisasi(ctx http.Context) http.Response {
	ke, _ := strconv.Atoi(ctx.Request().Query("no_imunisasi"))
	if err := c.action.HapusImunisasi(ctx.Request().Query("no_rkm_medis"), ctx.Request().Query("kode_imunisasi"), ke); err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus riwayat imunisasi", err)
	}
	return c.ResponseSuccess(ctx, "Riwayat imunisasi berhasil dihapus", nil)
}
