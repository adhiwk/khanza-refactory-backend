package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningIndraPendengaran skrining indra pendengaran (RMSkriningIndraPendengaran).
var FormSkriningIndraPendengaran = &Form[model.SkriningIndraPendengaran, request.SkriningIndraPendengaranData]{
	Slug:  "skrining-indra-pendengaran",
	Label: "skrining indra pendengaran",
	Spec: repo.Spec{
		Table:  "skrining_indra_pendengaran",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningIndraPendengaran, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningIndraPendengaran) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningIndraPendengaran) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningIndraPendengaran) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningIndraPendengaran, d request.SkriningIndraPendengaranData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.CurigaTuliTelingaKiri = support.Nullable(d.CurigaTuliTelingaKiri)
		m.CurigaTuliTelingaKanan = support.Nullable(d.CurigaTuliTelingaKanan)
		m.CurigaTuliTelingaRujuk = support.Nullable(d.CurigaTuliTelingaRujuk)
		m.PenurunanPendengaranTelingaKiri = support.Nullable(d.PenurunanPendengaranTelingaKiri)
		m.PenurunanPendengaranTelingaKanan = support.Nullable(d.PenurunanPendengaranTelingaKanan)
		m.MendengarBisikanTelingaKiri = support.Nullable(d.MendengarBisikanTelingaKiri)
		m.MendengarBisikanTelingaKanan = support.Nullable(d.MendengarBisikanTelingaKanan)
		m.CongekTelingaKiri = support.Nullable(d.CongekTelingaKiri)
		m.CongekTelingaKanan = support.Nullable(d.CongekTelingaKanan)
		m.CongekTelingaRujuk = support.Nullable(d.CongekTelingaRujuk)
		m.SumbatanSerumenTelingaKiri = support.Nullable(d.SumbatanSerumenTelingaKiri)
		m.SumbatanSerumenTelingaKanan = support.Nullable(d.SumbatanSerumenTelingaKanan)
		m.SumbatanSerumenTelingaRujuk = support.Nullable(d.SumbatanSerumenTelingaRujuk)
		m.HasilSkrining = strings.TrimSpace(d.HasilSkrining)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningIndraPendengaran]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningIndraPendengaran) any { return Str(m.Nip) }},
	},
}
