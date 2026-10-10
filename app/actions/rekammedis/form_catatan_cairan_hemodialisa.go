package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanCairanHemodialisa catatan cairan hemodialisa (RMDataCatatanCairanHemodialisa).
var FormCatatanCairanHemodialisa = &Form[model.CatatanCairanHemodialisa, request.CatatanCairanHemodialisaData]{
	Slug:  "catatan-cairan-hemodialisa",
	Label: "catatan cairan hemodialisa",
	Spec: repo.Spec{
		Table:  "catatan_cairan_hemodialisa",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanCairanHemodialisa, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.TglPerawatan, err = KeyDate(key["tgl_perawatan"]); err != nil {
			return err
		}
		if m.JamRawat, err = KeyTime(key["jam_rawat"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.CatatanCairanHemodialisa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanCairanHemodialisa) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanCairanHemodialisa) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.CatatanCairanHemodialisa, d request.CatatanCairanHemodialisaData) error {
		m.Minum = support.Nullable(d.Minum)
		m.Infus = support.Nullable(d.Infus)
		m.Tranfusi = support.Nullable(d.Tranfusi)
		m.SisaPriming = support.Nullable(d.SisaPriming)
		m.WashOut = support.Nullable(d.WashOut)
		m.Urine = support.Nullable(d.Urine)
		m.Pendarahan = support.Nullable(d.Pendarahan)
		m.Muntah = support.Nullable(d.Muntah)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanCairanHemodialisa]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanCairanHemodialisa) any { return StrPtr(m.Nip) }},
	},
}
