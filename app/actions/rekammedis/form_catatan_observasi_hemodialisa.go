package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiHemodialisa catatan observasi hemodialisa (RMDataCatatanObservasiHemodialisa).
var FormCatatanObservasiHemodialisa = &Form[model.CatatanObservasiHemodialisa, request.CatatanObservasiHemodialisaData]{
	Slug:  "catatan-observasi-hemodialisa",
	Label: "catatan observasi hemodialisa",
	Spec: repo.Spec{
		Table:  "catatan_observasi_hemodialisa",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiHemodialisa, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiHemodialisa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiHemodialisa) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiHemodialisa) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiHemodialisa, d request.CatatanObservasiHemodialisaData) error {
		m.Qb = support.Nullable(d.Qb)
		m.Qd = strings.TrimSpace(d.Qd)
		m.TekananArteri = support.Nullable(d.TekananArteri)
		m.TekananVena = support.Nullable(d.TekananVena)
		m.Tmp = support.Nullable(d.Tmp)
		m.Ufr = support.Nullable(d.Ufr)
		m.Tensi = support.Nullable(d.Tensi)
		m.Nadi = support.Nullable(d.Nadi)
		m.Suhu = support.Nullable(d.Suhu)
		m.Spo2 = support.Nullable(d.Spo2)
		m.Tindakan = support.Nullable(d.Tindakan)
		m.Ufg = support.Nullable(d.Ufg)
		m.BarcodeHf = strings.TrimSpace(d.BarcodeHf)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiHemodialisa]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiHemodialisa) any { return Str(m.Nip) }},
	},
}
