package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiRuangOk catatan observasi ruang operasi (RMDataCatatanObservasiRuangOperasi).
var FormCatatanObservasiRuangOk = &Form[model.CatatanObservasiRuangOk, request.CatatanObservasiRuangOkData]{
	Slug:  "catatan-observasi-ruang-ok",
	Label: "catatan observasi ruang operasi",
	Spec: repo.Spec{
		Table:  "catatan_observasi_ruang_ok",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiRuangOk, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiRuangOk) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiRuangOk) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiRuangOk) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiRuangOk, d request.CatatanObservasiRuangOkData) error {
		m.Gcs = support.Nullable(d.Gcs)
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = support.Nullable(d.Hr)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Spo2 = strings.TrimSpace(d.Spo2)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiRuangOk]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiRuangOk) any { return Str(m.Nip) }},
	},
}
