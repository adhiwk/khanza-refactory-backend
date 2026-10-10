package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningInstrumenAmt skrining instrumen AMT (RMSkriningInstrumenAMT).
var FormSkriningInstrumenAmt = &Form[model.SkriningInstrumenAmt, request.SkriningInstrumenAmtData]{
	Slug:  "skrining-instrumen-amt",
	Label: "skrining instrumen AMT",
	Spec: repo.Spec{
		Table:  "skrining_instrumen_amt",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningInstrumenAmt, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningInstrumenAmt) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningInstrumenAmt) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningInstrumenAmt) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningInstrumenAmt, d request.SkriningInstrumenAmtData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Pernyataanamt1 = support.Nullable(d.Pernyataanamt1)
		m.NilaiAmt1 = d.NilaiAmt1
		m.Pernyataanamt2 = support.Nullable(d.Pernyataanamt2)
		m.NilaiAmt2 = d.NilaiAmt2
		m.Pernyataanamt3 = support.Nullable(d.Pernyataanamt3)
		m.NilaiAmt3 = d.NilaiAmt3
		m.Pernyataanamt4 = support.Nullable(d.Pernyataanamt4)
		m.NilaiAmt4 = d.NilaiAmt4
		m.Pernyataanamt5 = support.Nullable(d.Pernyataanamt5)
		m.NilaiAmt5 = d.NilaiAmt5
		m.Pernyataanamt6 = support.Nullable(d.Pernyataanamt6)
		m.NilaiAmt6 = d.NilaiAmt6
		m.Pernyataanamt7 = support.Nullable(d.Pernyataanamt7)
		m.NilaiAmt7 = d.NilaiAmt7
		m.Pernyataanamt8 = support.Nullable(d.Pernyataanamt8)
		m.NilaiAmt8 = d.NilaiAmt8
		m.Pernyataanamt9 = support.Nullable(d.Pernyataanamt9)
		m.NilaiAmt9 = d.NilaiAmt9
		m.Pernyataanamt10 = support.Nullable(d.Pernyataanamt10)
		m.NilaiAmt10 = d.NilaiAmt10
		m.NilaiTotalAmt = d.NilaiTotalAmt
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.SkriningInstrumenAmt]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningInstrumenAmt) any { return Str(m.Nip) }},
	},
}
