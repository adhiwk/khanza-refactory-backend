package tindakan

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/tindakan"
	"goravel/app/http/controllers"
	request "goravel/app/http/requests/tindakan"
	model "goravel/app/models/perawatan"
	akunrepo "goravel/app/repository/akun"
	jurnalrepo "goravel/app/repository/jurnal"
	repo "goravel/app/repository/tindakan"
	jurnalsvc "goravel/app/services/jurnal"
	"goravel/app/support"
)

// Controller tindakan untuk satu jenis rawat (Ralan / Ranap).
type Controller struct {
	controllers.BaseController
	action *action.Action
}

func NewController(rawat model.Rawat) *Controller {
	jurnal := jurnalsvc.NewService(jurnalrepo.NewRepository())
	return &Controller{action: action.NewAction(rawat, repo.NewRepository(), akunrepo.NewRepository(), jurnal)}
}

// Index GET ?no_rawat=
func (c *Controller) Index(ctx http.Context) http.Response {
	data, err := c.action.List(ctx.Request().Query("no_rawat"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil data tindakan", err)
	}
	return c.ResponseSuccess(ctx, "Data tindakan berhasil diambil", data)
}

// Tarif GET ?no_rawat=&pelaksana=dr|pr|drpr&search=
func (c *Controller) Tarif(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	data, total, err := c.action.Tarif(ctx.Request().Query("no_rawat"), model.Pelaksana(ctx.Request().Query("pelaksana", "dr")),
		ctx.Request().Query("search"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil tarif tindakan", err)
	}
	return c.ResponsePaginated(ctx, "Data tarif tindakan berhasil diambil", data, page, limit, total)
}

func (c *Controller) Store(ctx http.Context) http.Response {
	var req request.StoreRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	data, err := c.action.Create(action.Input{
		NoRawat:    req.NoRawat,
		Pelaksana:  model.Pelaksana(req.Pelaksana),
		KdDokter:   req.KdDokter,
		Nip:        req.Nip,
		Tgl:        req.Tgl,
		Jam:        req.Jam,
		KdJenisPrw: req.KdJenisPrw,
		Operator:   c.operator(ctx),
	})
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menyimpan tindakan", err)
	}
	return c.ResponseCreated(ctx, "Tindakan berhasil disimpan", data)
}

func (c *Controller) Destroy(ctx http.Context) http.Response {
	var req request.DeleteRequest
	if resp := c.Validate(ctx, &req); resp != nil {
		return resp
	}
	err := c.action.Delete(model.TindakanKey{
		Pelaksana:    model.Pelaksana(req.Pelaksana),
		NoRawat:      req.NoRawat,
		KdJenisPrw:   req.KdJenisPrw,
		KdDokter:     req.KdDokter,
		Nip:          req.Nip,
		TglPerawatan: req.Tgl,
		JamRawat:     req.Jam,
	}, c.operator(ctx), c.IsSuperAdmin(ctx))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal menghapus tindakan", err)
	}
	return c.ResponseSuccess(ctx, "Tindakan berhasil dihapus", nil)
}

func (c *Controller) operator(ctx http.Context) string {
	return fmt.Sprintf("USER %d", c.UserID(ctx))
}
