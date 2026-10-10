package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiRestrainNonfarma catatan observasi restrain non farmakologi (RMDataCatatanObservasiRestrainNonFarmakologi).
var FormCatatanObservasiRestrainNonfarma = &Form[model.CatatanObservasiRestrainNonfarma, request.CatatanObservasiRestrainNonfarmaData]{
	Slug:  "catatan-observasi-restrain-nonfarma",
	Label: "catatan observasi restrain non farmakologi",
	Spec: repo.Spec{
		Table:  "catatan_observasi_restrain_nonfarma",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiRestrainNonfarma, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiRestrainNonfarma) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiRestrainNonfarma) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiRestrainNonfarma) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiRestrainNonfarma, d request.CatatanObservasiRestrainNonfarmaData) error {
		m.TanganKiri = support.Nullable(d.TanganKiri)
		m.TanganKanan = support.Nullable(d.TanganKanan)
		m.KakiKiri = support.Nullable(d.KakiKiri)
		m.KakiKanan = support.Nullable(d.KakiKanan)
		m.Badan = support.Nullable(d.Badan)
		m.Edema = support.Nullable(d.Edema)
		m.Iritasi = support.Nullable(d.Iritasi)
		m.Sirkulasi = support.Nullable(d.Sirkulasi)
		m.KondisiKeterangan = support.Nullable(d.KondisiKeterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiRestrainNonfarma]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiRestrainNonfarma) any { return Str(m.Nip) }},
	},
}
