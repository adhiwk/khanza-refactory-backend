package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanAdimeGizi catatan ADIME gizi (RMCatatanADIMEGizi).
var FormCatatanAdimeGizi = &Form[model.CatatanAdimeGizi, request.CatatanAdimeGiziData]{
	Slug:  "catatan-adime-gizi",
	Label: "catatan ADIME gizi",
	Spec: repo.Spec{
		Table:  "catatan_adime_gizi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanAdimeGizi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.CatatanAdimeGizi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.CatatanAdimeGizi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.CatatanAdimeGizi) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.CatatanAdimeGizi, d request.CatatanAdimeGiziData) error {
		m.Asesmen = support.Nullable(d.Asesmen)
		m.Diagnosis = support.Nullable(d.Diagnosis)
		m.Intervensi = support.Nullable(d.Intervensi)
		m.Monitoring = support.Nullable(d.Monitoring)
		m.Evaluasi = support.Nullable(d.Evaluasi)
		m.Instruksi = support.Nullable(d.Instruksi)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanAdimeGizi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanAdimeGizi) any { return StrPtr(m.Nip) }},
	},
}
