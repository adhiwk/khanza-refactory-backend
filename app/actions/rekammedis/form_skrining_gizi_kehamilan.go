package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningGiziKehamilan skrining gizi kehamilan (RMDataSkriningGiziKehamilan).
var FormSkriningGiziKehamilan = &Form[model.SkriningGiziKehamilan, request.SkriningGiziKehamilanData]{
	Slug:  "skrining-gizi-kehamilan",
	Label: "skrining gizi kehamilan",
	Spec: repo.Spec{
		Table:  "skrining_gizi_kehamilan",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningGiziKehamilan, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.SkriningGiziKehamilan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.SkriningGiziKehamilan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningGiziKehamilan) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.SkriningGiziKehamilan, d request.SkriningGiziKehamilanData) error {
		m.Parameter1 = support.Nullable(d.Parameter1)
		m.Skor1 = strings.TrimSpace(d.Skor1)
		m.Parameter2 = support.Nullable(d.Parameter2)
		m.Skor2 = strings.TrimSpace(d.Skor2)
		m.Parameter3 = support.Nullable(d.Parameter3)
		m.Skor3 = strings.TrimSpace(d.Skor3)
		m.Parameter4 = support.Nullable(d.Parameter4)
		m.Skor4 = strings.TrimSpace(d.Skor4)
		m.NilaiSkor = support.Nullable(d.NilaiSkor)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningGiziKehamilan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningGiziKehamilan) any { return StrPtr(m.Nip) }},
	},
}
