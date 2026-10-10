package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiBayi catatan observasi bayi (RMDataCatatanObservasiBayi).
var FormCatatanObservasiBayi = &Form[model.CatatanObservasiBayi, request.CatatanObservasiBayiData]{
	Slug:  "catatan-observasi-bayi",
	Label: "catatan observasi bayi",
	Spec: repo.Spec{
		Table:  "catatan_observasi_bayi",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiBayi, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiBayi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiBayi) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiBayi) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.CatatanObservasiBayi, d request.CatatanObservasiBayiData) error {
		m.Gcs = support.Nullable(d.Gcs)
		m.Td = support.Nullable(d.Td)
		m.Hr = support.Nullable(d.Hr)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Spo2 = support.Nullable(d.Spo2)
		m.Nch = support.Nullable(d.Nch)
		m.IkterikStatus = support.Nullable(d.IkterikStatus)
		m.RetraksiDada = support.Nullable(d.RetraksiDada)
		m.OgtResidu = support.Nullable(d.OgtResidu)
		m.AsiJumlah = support.Nullable(d.AsiJumlah)
		m.PasiJumlah = support.Nullable(d.PasiJumlah)
		m.BakStatus = support.Nullable(d.BakStatus)
		m.BabStatus = support.Nullable(d.BabStatus)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiBayi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiBayi) any { return StrPtr(m.Nip) }},
	},
}
