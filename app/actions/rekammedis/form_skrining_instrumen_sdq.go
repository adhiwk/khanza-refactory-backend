package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningInstrumenSdq skrining instrumen SDQ (RMSkriningInstrumenSDQ).
var FormSkriningInstrumenSdq = &Form[model.SkriningInstrumenSdq, request.SkriningInstrumenSdqData]{
	Slug:  "skrining-instrumen-sdq",
	Label: "skrining instrumen SDQ",
	Spec: repo.Spec{
		Table:  "skrining_instrumen_sdq",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningInstrumenSdq, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningInstrumenSdq) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningInstrumenSdq) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningInstrumenSdq) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningInstrumenSdq, d request.SkriningInstrumenSdqData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Pernyataansdq1 = support.Nullable(d.Pernyataansdq1)
		m.NilaiSdq1 = d.NilaiSdq1
		m.Pernyataansdq2 = support.Nullable(d.Pernyataansdq2)
		m.NilaiSdq2 = d.NilaiSdq2
		m.Pernyataansdq3 = support.Nullable(d.Pernyataansdq3)
		m.NilaiSdq3 = d.NilaiSdq3
		m.Pernyataansdq4 = support.Nullable(d.Pernyataansdq4)
		m.NilaiSdq4 = d.NilaiSdq4
		m.Pernyataansdq5 = support.Nullable(d.Pernyataansdq5)
		m.NilaiSdq5 = d.NilaiSdq5
		m.Pernyataansdq6 = support.Nullable(d.Pernyataansdq6)
		m.NilaiSdq6 = d.NilaiSdq6
		m.Pernyataansdq7 = support.Nullable(d.Pernyataansdq7)
		m.NilaiSdq7 = d.NilaiSdq7
		m.Pernyataansdq8 = support.Nullable(d.Pernyataansdq8)
		m.NilaiSdq8 = d.NilaiSdq8
		m.Pernyataansdq9 = support.Nullable(d.Pernyataansdq9)
		m.NilaiSdq9 = d.NilaiSdq9
		m.Pernyataansdq10 = support.Nullable(d.Pernyataansdq10)
		m.NilaiSdq10 = d.NilaiSdq10
		m.Pernyataansdq11 = support.Nullable(d.Pernyataansdq11)
		m.NilaiSdq11 = d.NilaiSdq11
		m.Pernyataansdq12 = support.Nullable(d.Pernyataansdq12)
		m.NilaiSdq12 = d.NilaiSdq12
		m.Pernyataansdq13 = support.Nullable(d.Pernyataansdq13)
		m.NilaiSdq13 = d.NilaiSdq13
		m.Pernyataansdq14 = support.Nullable(d.Pernyataansdq14)
		m.NilaiSdq14 = d.NilaiSdq14
		m.Pernyataansdq15 = support.Nullable(d.Pernyataansdq15)
		m.NilaiSdq15 = d.NilaiSdq15
		m.Pernyataansdq16 = support.Nullable(d.Pernyataansdq16)
		m.NilaiSdq16 = d.NilaiSdq16
		m.Pernyataansdq17 = support.Nullable(d.Pernyataansdq17)
		m.NilaiSdq17 = d.NilaiSdq17
		m.Pernyataansdq18 = support.Nullable(d.Pernyataansdq18)
		m.NilaiSdq18 = d.NilaiSdq18
		m.Pernyataansdq19 = support.Nullable(d.Pernyataansdq19)
		m.NilaiSdq19 = d.NilaiSdq19
		m.Pernyataansdq20 = support.Nullable(d.Pernyataansdq20)
		m.NilaiSdq20 = d.NilaiSdq20
		m.Pernyataansdq21 = support.Nullable(d.Pernyataansdq21)
		m.NilaiSdq21 = d.NilaiSdq21
		m.Pernyataansdq22 = support.Nullable(d.Pernyataansdq22)
		m.NilaiSdq22 = d.NilaiSdq22
		m.Pernyataansdq23 = support.Nullable(d.Pernyataansdq23)
		m.NilaiSdq23 = d.NilaiSdq23
		m.Pernyataansdq24 = support.Nullable(d.Pernyataansdq24)
		m.NilaiSdq24 = d.NilaiSdq24
		m.Pernyataansdq25 = support.Nullable(d.Pernyataansdq25)
		m.NilaiSdq25 = d.NilaiSdq25
		m.NilaiTotalSdq = d.NilaiTotalSdq
		m.GejalaEmosional = support.Nullable(d.GejalaEmosional)
		m.NilaiGejalaEmosional = d.NilaiGejalaEmosional
		m.MasalahPerilaku = support.Nullable(d.MasalahPerilaku)
		m.NilaiMasalahPerilaku = d.NilaiMasalahPerilaku
		m.Hiperaktivitas = support.Nullable(d.Hiperaktivitas)
		m.NilaiHiperaktivitas = d.NilaiHiperaktivitas
		m.TemanSebaya = support.Nullable(d.TemanSebaya)
		m.NilaiTemanSebaya = d.NilaiTemanSebaya
		m.Kekuatan = support.Nullable(d.Kekuatan)
		m.NilaiKekuatan = d.NilaiKekuatan
		m.Kesulitan = support.Nullable(d.Kesulitan)
		m.NilaiKesulitan = d.NilaiKesulitan
		m.Keterangan = support.Nullable(d.Keterangan)
		return nil
	},
	Refs: []Ref[model.SkriningInstrumenSdq]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningInstrumenSdq) any { return Str(m.Nip) }},
	},
}
