package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rekammedis"
	"goravel/app/http/controllers"
	icdrepo "goravel/app/repository/icd"
	kamarrepo "goravel/app/repository/kamarinap"
	obatrepo "goravel/app/repository/pemberianobat"
	pemeriksaanrepo "goravel/app/repository/pemeriksaan"
	repo "goravel/app/repository/rekammedis"
	tindakanrepo "goravel/app/repository/tindakan"
	"goravel/app/support"
)

type RiwayatController struct {
	controllers.BaseController
	action *action.RiwayatAction
}

func NewRiwayatController() *RiwayatController {
	return &RiwayatController{action: action.NewRiwayatAction(repo.NewRiwayatRepository(), icdrepo.NewRepository(),
		pemeriksaanrepo.NewRepository(), tindakanrepo.NewRepository(), obatrepo.NewRepository(), kamarrepo.NewRepository())}
}

// Show ?no_rkm_medis=&tgl_awal=&tgl_akhir=&page=&limit= (paginasi per kunjungan, default 10)
func (c *RiwayatController) Show(ctx http.Context) http.Response {
	page, limit := support.PageParams(ctx)
	q := ctx.Request()
	data, total, err := c.action.Riwayat(q.Query("no_rkm_medis"), q.Query("tgl_awal"), q.Query("tgl_akhir"), page, limit)
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mengambil riwayat rekam medis", err)
	}
	return c.ResponsePaginated(ctx, "Riwayat rekam medis berhasil diambil", data, page, limit, total)
}

// Forms daftar form asesmen yang tersedia.
func (c *RiwayatController) Forms(ctx http.Context) http.Response {
	return c.ResponseSuccess(ctx, "Daftar form asesmen", action.DaftarForm)
}
