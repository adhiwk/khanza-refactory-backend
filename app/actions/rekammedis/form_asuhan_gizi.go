package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormAsuhanGizi asuhan gizi (RMDataAsuhanGizi).
var FormAsuhanGizi = &Form[model.AsuhanGizi, request.AsuhanGiziData]{
	Slug:  "asuhan-gizi",
	Label: "asuhan gizi",
	Spec: repo.Spec{
		Table:  "asuhan_gizi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.AsuhanGizi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDate(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.AsuhanGizi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDate(m.Tanggal)}
	},
	Waktu:   func(m *model.AsuhanGizi) time.Time { return time.Time{} },
	Petugas: func(m *model.AsuhanGizi) []string { return []string{m.Nip} },
	Fill: func(m *model.AsuhanGizi, d request.AsuhanGiziData) error {
		m.AntropometriBb = support.Nullable(d.AntropometriBb)
		m.AntropometriTb = support.Nullable(d.AntropometriTb)
		m.AntropometriImt = support.Nullable(d.AntropometriImt)
		m.AntropometriLla = support.Nullable(d.AntropometriLla)
		m.AntropometriTl = support.Nullable(d.AntropometriTl)
		m.AntropometriUlna = strings.TrimSpace(d.AntropometriUlna)
		m.AntropometriBbideal = strings.TrimSpace(d.AntropometriBbideal)
		m.AntropometriBbperu = strings.TrimSpace(d.AntropometriBbperu)
		m.AntropometriTbperu = strings.TrimSpace(d.AntropometriTbperu)
		m.AntropometriBbpertb = strings.TrimSpace(d.AntropometriBbpertb)
		m.AntropometriLlaperu = strings.TrimSpace(d.AntropometriLlaperu)
		m.Biokimia = support.Nullable(d.Biokimia)
		m.FisikKlinis = support.Nullable(d.FisikKlinis)
		m.AlergiTelur = support.Nullable(d.AlergiTelur)
		m.AlergiSusuSapi = support.Nullable(d.AlergiSusuSapi)
		m.AlergiKacang = support.Nullable(d.AlergiKacang)
		m.AlergiGluten = support.Nullable(d.AlergiGluten)
		m.AlergiUdang = support.Nullable(d.AlergiUdang)
		m.AlergiIkan = support.Nullable(d.AlergiIkan)
		m.AlergiHazelnut = support.Nullable(d.AlergiHazelnut)
		m.PolaMakan = support.Nullable(d.PolaMakan)
		m.RiwayatPersonal = support.Nullable(d.RiwayatPersonal)
		m.Diagnosis = support.Nullable(d.Diagnosis)
		m.IntervensiGizi = support.Nullable(d.IntervensiGizi)
		m.MonitoringEvaluasi = support.Nullable(d.MonitoringEvaluasi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.AsuhanGizi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.AsuhanGizi) any { return Str(m.Nip) }},
	},
}
