package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningInstrumenSrq skrining SRQ (RMSkriningSRQ).
var FormSkriningInstrumenSrq = &Form[model.SkriningInstrumenSrq, request.SkriningInstrumenSrqData]{
	Slug:  "skrining-instrumen-srq",
	Label: "skrining SRQ",
	Spec: repo.Spec{
		Table:  "skrining_instrumen_srq",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningInstrumenSrq, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningInstrumenSrq) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningInstrumenSrq) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningInstrumenSrq) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningInstrumenSrq, d request.SkriningInstrumenSrqData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Pernyataansrq1 = support.Nullable(d.Pernyataansrq1)
		m.NilaiSrq1 = d.NilaiSrq1
		m.Pernyataansrq2 = support.Nullable(d.Pernyataansrq2)
		m.NilaiSrq2 = d.NilaiSrq2
		m.Pernyataansrq3 = support.Nullable(d.Pernyataansrq3)
		m.NilaiSrq3 = d.NilaiSrq3
		m.Pernyataansrq4 = support.Nullable(d.Pernyataansrq4)
		m.NilaiSrq4 = d.NilaiSrq4
		m.Pernyataansrq5 = support.Nullable(d.Pernyataansrq5)
		m.NilaiSrq5 = d.NilaiSrq5
		m.Pernyataansrq6 = support.Nullable(d.Pernyataansrq6)
		m.NilaiSrq6 = d.NilaiSrq6
		m.Pernyataansrq7 = support.Nullable(d.Pernyataansrq7)
		m.NilaiSrq7 = d.NilaiSrq7
		m.Pernyataansrq8 = support.Nullable(d.Pernyataansrq8)
		m.NilaiSrq8 = d.NilaiSrq8
		m.Pernyataansrq9 = support.Nullable(d.Pernyataansrq9)
		m.NilaiSrq9 = d.NilaiSrq9
		m.Pernyataansrq10 = support.Nullable(d.Pernyataansrq10)
		m.NilaiSrq10 = d.NilaiSrq10
		m.Pernyataansrq11 = support.Nullable(d.Pernyataansrq11)
		m.NilaiSrq11 = d.NilaiSrq11
		m.Pernyataansrq12 = support.Nullable(d.Pernyataansrq12)
		m.NilaiSrq12 = d.NilaiSrq12
		m.Pernyataansrq13 = support.Nullable(d.Pernyataansrq13)
		m.NilaiSrq13 = d.NilaiSrq13
		m.Pernyataansrq14 = support.Nullable(d.Pernyataansrq14)
		m.NilaiSrq14 = d.NilaiSrq14
		m.Pernyataansrq15 = support.Nullable(d.Pernyataansrq15)
		m.NilaiSrq15 = d.NilaiSrq15
		m.Pernyataansrq16 = support.Nullable(d.Pernyataansrq16)
		m.NilaiSrq16 = d.NilaiSrq16
		m.Pernyataansrq17 = support.Nullable(d.Pernyataansrq17)
		m.NilaiSrq17 = d.NilaiSrq17
		m.Pernyataansrq18 = support.Nullable(d.Pernyataansrq18)
		m.NilaiSrq18 = d.NilaiSrq18
		m.Pernyataansrq19 = support.Nullable(d.Pernyataansrq19)
		m.NilaiSrq19 = d.NilaiSrq19
		m.Pernyataansrq20 = support.Nullable(d.Pernyataansrq20)
		m.NilaiSrq20 = d.NilaiSrq20
		m.NilaiTotalSrq = d.NilaiTotalSrq
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.SkriningInstrumenSrq]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningInstrumenSrq) any { return Str(m.Nip) }},
	},
}
