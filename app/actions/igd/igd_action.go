// Package igd use case registrasi Instalasi Gawat Darurat (DlgIGD).
package igd

import (
	"strings"

	"github.com/goravel/framework/facades"

	regaction "goravel/app/actions/registrasi"
	rujukaction "goravel/app/actions/rujukmasuk"
	regrequest "goravel/app/http/requests/registrasi"
	rujukrequest "goravel/app/http/requests/rujukmasuk"
	regmodel "goravel/app/models/registrasi"
	rujukmodel "goravel/app/models/rujukmasuk"
)

// Rujukan asal rujukan opsional saat pasien datang ke IGD.
type Rujukan struct {
	Perujuk string
	Alamat  string
	NoRujuk string
}

type Result struct {
	Registrasi *regmodel.RegPeriksa   `json:"registrasi"`
	Rujukan    *rujukmodel.RujukMasuk `json:"rujukan"`
}

type Action struct {
	registrasi *regaction.Action
	rujukan    *rujukaction.Action
}

func NewAction(registrasi *regaction.Action, rujukan *rujukaction.Action) *Action {
	return &Action{registrasi: registrasi, rujukan: rujukan}
}

// Register registrasi IGD lalu mencatat rujukan masuk bila asal rujukan diisi.
// Seperti DlgIGD, kegagalan pencatatan rujukan tidak membatalkan registrasi; kegagalan dicatat di log.
func (a *Action) Register(noRkmMedis string, data regrequest.RegistrasiData, rujukan Rujukan) (*Result, error) {
	reg, err := a.registrasi.CreateIGD(noRkmMedis, data)
	if err != nil {
		return nil, err
	}
	result := &Result{Registrasi: reg}

	perujuk := strings.TrimSpace(rujukan.Perujuk)
	if perujuk == "" {
		return result, nil
	}
	req := rujukrequest.StoreRequest{
		NoRawat: reg.NoRawat,
		Data: rujukrequest.Data{
			Perujuk:       perujuk,
			Alamat:        rujukan.Alamat,
			NoRujuk:       orDash(rujukan.NoRujuk),
			DokterPerujuk: perujuk,
			KategoriRujuk: "-",
			Keterangan:    "-",
		},
	}
	rujuk, err := a.rujukan.Create(req)
	if err != nil {
		facades.Log().Errorf("igd: gagal mencatat rujukan masuk %s: %v", reg.NoRawat, err)
		return result, nil
	}
	result.Rujukan = rujuk
	return result, nil
}

func orDash(s string) string {
	if s = strings.TrimSpace(s); s != "" {
		return s
	}
	return "-"
}
