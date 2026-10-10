// Package rekammedis request form asesmen rekam medis; satu file per form dibangkitkan dari sik.sql.
package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// StoreRequest kontrak request simpan: kunci (no_rawat + waktu), isian, dan detail kode.
type StoreRequest[D any] interface {
	http.FormRequest
	KeyValues() repo.Key
	Payload() D
	DetailValues() map[string][]string
}

// UpdateRequest kontrak request ubah (PUT); kunci diambil dari query string.
type UpdateRequest[D any] interface {
	http.FormRequest
	Payload() D
	DetailValues() map[string][]string
}
