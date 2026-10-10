package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormMppEvaluasiCatatan skrining MPP form b (RMSkriningMPPFormB).
var FormMppEvaluasiCatatan = &Form[model.MppEvaluasiCatatan, request.MppEvaluasiCatatanData]{
	Slug:  "mpp-evaluasi-catatan",
	Label: "skrining MPP form b",
	Spec: repo.Spec{
		Table:  "mpp_evaluasi_catatan",
		Keys:   []string{"no_rawat", "tgl_implementasi"},
		Waktu:  "tgl_implementasi",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.MppEvaluasiCatatan, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.TglImplementasi, err = KeyDateTime(key["tgl_implementasi"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.MppEvaluasiCatatan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_implementasi": FmtDateTime(m.TglImplementasi)}
	},
	Waktu:   func(m *model.MppEvaluasiCatatan) time.Time { return Tm(m.TglImplementasi) },
	Petugas: func(m *model.MppEvaluasiCatatan) []string { return []string{m.Nip} },
	Fill: func(m *model.MppEvaluasiCatatan, d request.MppEvaluasiCatatanData) error {
		m.Masalah = support.Nullable(d.Masalah)
		m.Tinjut = support.Nullable(d.Tinjut)
		m.Evaluasi = support.Nullable(d.Evaluasi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.MppEvaluasiCatatan]{
		{Column: "nip", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.MppEvaluasiCatatan) any { return Str(m.Nip) }},
	},
}
