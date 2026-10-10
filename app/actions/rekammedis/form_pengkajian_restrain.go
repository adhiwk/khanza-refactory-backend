package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPengkajianRestrain pengkajian restrain (RMPengkajianRestrain).
var FormPengkajianRestrain = &Form[model.PengkajianRestrain, request.PengkajianRestrainData]{
	Slug:  "pengkajian-restrain",
	Label: "pengkajian restrain",
	Spec: repo.Spec{
		Table:  "pengkajian_restrain",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PengkajianRestrain, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PengkajianRestrain) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PengkajianRestrain) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PengkajianRestrain) []string { return []string{m.Nip} },
	Fill: func(m *model.PengkajianRestrain, d request.PengkajianRestrainData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Gcs = support.Nullable(d.Gcs)
		m.ReflekaCahayaKa = support.Nullable(d.ReflekaCahayaKa)
		m.ReflekaCahayaKi = support.Nullable(d.ReflekaCahayaKi)
		m.UkuranPupilKa = support.Nullable(d.UkuranPupilKa)
		m.UkuranPupilKi = support.Nullable(d.UkuranPupilKi)
		m.Td = support.Nullable(d.Td)
		m.Suhu = support.Nullable(d.Suhu)
		m.Rr = support.Nullable(d.Rr)
		m.Nadi = support.Nullable(d.Nadi)
		m.HasilObservasi = support.Nullable(d.HasilObservasi)
		m.PertimbanganKlinis = support.Nullable(d.PertimbanganKlinis)
		m.RestrainNonFarmakologi = support.Nullable(d.RestrainNonFarmakologi)
		m.RestrainNonFarmakologiKeterangan = support.Nullable(d.RestrainNonFarmakologiKeterangan)
		m.RestrainFarmakologi = support.Nullable(d.RestrainFarmakologi)
		m.SudahDijelaskanKeluarga = strings.TrimSpace(d.SudahDijelaskanKeluarga)
		m.KeluargaYangMenyetujui = support.Nullable(d.KeluargaYangMenyetujui)
		return nil
	},
	Refs: []Ref[model.PengkajianRestrain]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PengkajianRestrain) any { return Str(m.Nip) }},
	},
}
