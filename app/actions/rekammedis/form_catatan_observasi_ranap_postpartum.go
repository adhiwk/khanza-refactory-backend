package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiRanapPostpartum catatan observasi ranap post partum (RMDataCatatanObservasiRanapPostPartum).
var FormCatatanObservasiRanapPostpartum = &Form[model.CatatanObservasiRanapPostpartum, request.CatatanObservasiRanapPostpartumData]{
	Slug:  "catatan-observasi-ranap-postpartum",
	Label: "catatan observasi ranap post partum",
	Spec: repo.Spec{
		Table:  "catatan_observasi_ranap_postpartum",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiRanapPostpartum, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiRanapPostpartum) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiRanapPostpartum) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiRanapPostpartum) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiRanapPostpartum, d request.CatatanObservasiRanapPostpartumData) error {
		m.Gcs = support.Nullable(d.Gcs)
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = support.Nullable(d.Hr)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Spo2 = strings.TrimSpace(d.Spo2)
		m.Tfu = strings.TrimSpace(d.Tfu)
		m.Kontraksi = strings.TrimSpace(d.Kontraksi)
		m.Perdarahan = strings.TrimSpace(d.Perdarahan)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiRanapPostpartum]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiRanapPostpartum) any { return Str(m.Nip) }},
	},
}
