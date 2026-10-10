package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	action "goravel/app/actions/rekammedis"
	"goravel/app/http/controllers"
	cetakrepo "goravel/app/repository/cetak"
	icdrepo "goravel/app/repository/icd"
	kamarrepo "goravel/app/repository/kamarinap"
	obatrepo "goravel/app/repository/pemberianobat"
	pemeriksaanrepo "goravel/app/repository/pemeriksaan"
	repo "goravel/app/repository/rekammedis"
	riwayatpasienrepo "goravel/app/repository/riwayatpasien"
	tindakanrepo "goravel/app/repository/tindakan"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

type RiwayatController struct {
	controllers.BaseController
	action *action.RiwayatAction
}

func NewRiwayatController() *RiwayatController {
	return &RiwayatController{action: action.NewRiwayatAction(repo.NewRiwayatRepository(), icdrepo.NewRepository(),
		pemeriksaanrepo.NewRepository(), tindakanrepo.NewRepository(), obatrepo.NewRepository(), kamarrepo.NewRepository(),
		riwayatpasienrepo.NewRepository(), cetaksvc.NewService(cetakrepo.NewRepository()))}
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

// Cetak resume rekam medis ?no_rkm_medis=&tgl_awal=&tgl_akhir=
func (c *RiwayatController) Cetak(ctx http.Context) http.Response {
	q := ctx.Request()
	dok, err := c.action.Cetak(q.Query("no_rkm_medis"), q.Query("tgl_awal"), q.Query("tgl_akhir"))
	if err != nil {
		return c.ResponseActionError(ctx, "Gagal mencetak riwayat rekam medis", err)
	}
	return c.ResponseCetak(ctx, dok)
}

// Forms daftar form asesmen yang tersedia.
func (c *RiwayatController) Forms(ctx http.Context) http.Response {
	return c.ResponseSuccess(ctx, "Daftar form asesmen", action.DaftarForm)
}
