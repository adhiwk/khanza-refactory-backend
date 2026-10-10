package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianLanjutanResikoJatuhDewasa penilaian lanjutan risiko jatuh dewasa (RMPenilaianLanjutanRisikoJatuhDewasa).
var FormPenilaianLanjutanResikoJatuhDewasa = &Form[model.PenilaianLanjutanResikoJatuhDewasa, request.PenilaianLanjutanResikoJatuhDewasaData]{
	Slug:  "penilaian-lanjutan-resiko-jatuh-dewasa",
	Label: "penilaian lanjutan risiko jatuh dewasa",
	Spec: repo.Spec{
		Table:  "penilaian_lanjutan_resiko_jatuh_dewasa",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianLanjutanResikoJatuhDewasa, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianLanjutanResikoJatuhDewasa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianLanjutanResikoJatuhDewasa) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianLanjutanResikoJatuhDewasa) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianLanjutanResikoJatuhDewasa, d request.PenilaianLanjutanResikoJatuhDewasaData) error {
		m.PenilaianJatuhmorseSkala1 = support.Nullable(d.PenilaianJatuhmorseSkala1)
		m.PenilaianJatuhmorseNilai1 = d.PenilaianJatuhmorseNilai1
		m.PenilaianJatuhmorseSkala2 = support.Nullable(d.PenilaianJatuhmorseSkala2)
		m.PenilaianJatuhmorseNilai2 = d.PenilaianJatuhmorseNilai2
		m.PenilaianJatuhmorseSkala3 = support.Nullable(d.PenilaianJatuhmorseSkala3)
		m.PenilaianJatuhmorseNilai3 = d.PenilaianJatuhmorseNilai3
		m.PenilaianJatuhmorseSkala4 = support.Nullable(d.PenilaianJatuhmorseSkala4)
		m.PenilaianJatuhmorseNilai4 = d.PenilaianJatuhmorseNilai4
		m.PenilaianJatuhmorseSkala5 = support.Nullable(d.PenilaianJatuhmorseSkala5)
		m.PenilaianJatuhmorseNilai5 = d.PenilaianJatuhmorseNilai5
		m.PenilaianJatuhmorseSkala6 = support.Nullable(d.PenilaianJatuhmorseSkala6)
		m.PenilaianJatuhmorseNilai6 = d.PenilaianJatuhmorseNilai6
		m.PenilaianJatuhmorseTotalnilai = d.PenilaianJatuhmorseTotalnilai
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Saran = support.Nullable(d.Saran)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianLanjutanResikoJatuhDewasa]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianLanjutanResikoJatuhDewasa) any { return Str(m.Nip) }},
	},
}
