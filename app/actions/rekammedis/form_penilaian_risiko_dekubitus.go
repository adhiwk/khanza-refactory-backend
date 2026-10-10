package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianRisikoDekubitus penilaian risiko dekubitus (RMPenilaianRisikoDekubitus).
var FormPenilaianRisikoDekubitus = &Form[model.PenilaianRisikoDekubitus, request.PenilaianRisikoDekubitusData]{
	Slug:  "penilaian-risiko-dekubitus",
	Label: "penilaian risiko dekubitus",
	Spec: repo.Spec{
		Table:  "penilaian_risiko_dekubitus",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianRisikoDekubitus, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianRisikoDekubitus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianRisikoDekubitus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianRisikoDekubitus) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianRisikoDekubitus, d request.PenilaianRisikoDekubitusData) error {
		m.KondisiFisik = support.Nullable(d.KondisiFisik)
		m.KondisiFisikNilai = d.KondisiFisikNilai
		m.StatusMental = support.Nullable(d.StatusMental)
		m.StatusMentalNilai = d.StatusMentalNilai
		m.Aktifitas = support.Nullable(d.Aktifitas)
		m.AktifitasNilai = d.AktifitasNilai
		m.Mobilitas = support.Nullable(d.Mobilitas)
		m.MobilitasNilai = d.MobilitasNilai
		m.Inkontinensia = support.Nullable(d.Inkontinensia)
		m.InkontinensiaNilai = d.InkontinensiaNilai
		m.Totalnilai = d.Totalnilai
		m.Kategorinilai = support.Nullable(d.Kategorinilai)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianRisikoDekubitus]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianRisikoDekubitus) any { return Str(m.Nip) }},
	},
}
