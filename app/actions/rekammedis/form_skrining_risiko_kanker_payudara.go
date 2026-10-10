package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningRisikoKankerPayudara skrining risiko kanker payudara (RMSkriningRisikoKankerPayudara).
var FormSkriningRisikoKankerPayudara = &Form[model.SkriningRisikoKankerPayudara, request.SkriningRisikoKankerPayudaraData]{
	Slug:  "skrining-risiko-kanker-payudara",
	Label: "skrining risiko kanker payudara",
	Spec: repo.Spec{
		Table:  "skrining_risiko_kanker_payudara",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningRisikoKankerPayudara, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningRisikoKankerPayudara) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningRisikoKankerPayudara) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningRisikoKankerPayudara) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningRisikoKankerPayudara, d request.SkriningRisikoKankerPayudaraData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.FaktorRisikoAwal1 = support.Nullable(d.FaktorRisikoAwal1)
		m.NilaiRisikoAwal1 = support.Nullable(d.NilaiRisikoAwal1)
		m.FaktorRisikoAwal2 = support.Nullable(d.FaktorRisikoAwal2)
		m.NilaiRisikoAwal2 = support.Nullable(d.NilaiRisikoAwal2)
		m.FaktorRisikoAwal3 = support.Nullable(d.FaktorRisikoAwal3)
		m.NilaiRisikoAwal3 = support.Nullable(d.NilaiRisikoAwal3)
		m.FaktorRisikoAwal4 = support.Nullable(d.FaktorRisikoAwal4)
		m.NilaiRisikoAwal4 = support.Nullable(d.NilaiRisikoAwal4)
		m.FaktorRisikoAwal5 = support.Nullable(d.FaktorRisikoAwal5)
		m.NilaiRisikoAwal5 = support.Nullable(d.NilaiRisikoAwal5)
		m.FaktorRisikoAwal6 = support.Nullable(d.FaktorRisikoAwal6)
		m.NilaiRisikoAwal6 = support.Nullable(d.NilaiRisikoAwal6)
		m.FaktorRisikoAwal7 = support.Nullable(d.FaktorRisikoAwal7)
		m.NilaiRisikoAwal7 = support.Nullable(d.NilaiRisikoAwal7)
		m.FaktorRisikoAwal8 = support.Nullable(d.FaktorRisikoAwal8)
		m.NilaiRisikoAwal8 = support.Nullable(d.NilaiRisikoAwal8)
		m.FaktorRisikoAwal9 = support.Nullable(d.FaktorRisikoAwal9)
		m.NilaiRisikoAwal9 = support.Nullable(d.NilaiRisikoAwal9)
		m.FaktorRisikoAwal10 = support.Nullable(d.FaktorRisikoAwal10)
		m.NilaiRisikoAwal10 = support.Nullable(d.NilaiRisikoAwal10)
		m.FaktorRisikoAwal11 = support.Nullable(d.FaktorRisikoAwal11)
		m.NilaiRisikoAwal11 = support.Nullable(d.NilaiRisikoAwal11)
		m.FaktorRisikoAwal12 = support.Nullable(d.FaktorRisikoAwal12)
		m.NilaiRisikoAwal12 = support.Nullable(d.NilaiRisikoAwal12)
		m.FaktorRisikoAwal13 = support.Nullable(d.FaktorRisikoAwal13)
		m.NilaiRisikoAwal13 = support.Nullable(d.NilaiRisikoAwal13)
		m.FaktorRisikoAwal14 = support.Nullable(d.FaktorRisikoAwal14)
		m.NilaiRisikoAwal14 = support.Nullable(d.NilaiRisikoAwal14)
		m.FaktorRisikoTinggi1 = support.Nullable(d.FaktorRisikoTinggi1)
		m.NilaiRisikoTinggi1 = support.Nullable(d.NilaiRisikoTinggi1)
		m.FaktorRisikoTinggi2 = support.Nullable(d.FaktorRisikoTinggi2)
		m.NilaiRisikoTinggi2 = support.Nullable(d.NilaiRisikoTinggi2)
		m.FaktorRisikoTinggi3 = support.Nullable(d.FaktorRisikoTinggi3)
		m.NilaiRisikoTinggi3 = support.Nullable(d.NilaiRisikoTinggi3)
		m.FaktorRisikoTinggi4 = support.Nullable(d.FaktorRisikoTinggi4)
		m.NilaiRisikoTinggi4 = support.Nullable(d.NilaiRisikoTinggi4)
		m.FaktorRisikoTinggi5 = support.Nullable(d.FaktorRisikoTinggi5)
		m.NilaiRisikoTinggi5 = support.Nullable(d.NilaiRisikoTinggi5)
		m.FaktorRisikoTinggi6 = support.Nullable(d.FaktorRisikoTinggi6)
		m.NilaiRisikoTinggi6 = support.Nullable(d.NilaiRisikoTinggi6)
		m.FaktorRisikoTinggi7 = support.Nullable(d.FaktorRisikoTinggi7)
		m.NilaiRisikoTinggi7 = support.Nullable(d.NilaiRisikoTinggi7)
		m.FaktorRisikoTinggi8 = support.Nullable(d.FaktorRisikoTinggi8)
		m.NilaiRisikoTinggi8 = support.Nullable(d.NilaiRisikoTinggi8)
		m.FaktorRisikoTinggi9 = support.Nullable(d.FaktorRisikoTinggi9)
		m.NilaiRisikoTinggi9 = support.Nullable(d.NilaiRisikoTinggi9)
		m.FaktorRisikoTinggi10 = support.Nullable(d.FaktorRisikoTinggi10)
		m.NilaiRisikoTinggi10 = support.Nullable(d.NilaiRisikoTinggi10)
		m.FaktorRisikoTinggi11 = support.Nullable(d.FaktorRisikoTinggi11)
		m.NilaiRisikoTinggi11 = support.Nullable(d.NilaiRisikoTinggi11)
		m.FaktorRisikoTinggi12 = support.Nullable(d.FaktorRisikoTinggi12)
		m.NilaiRisikoTinggi12 = support.Nullable(d.NilaiRisikoTinggi12)
		m.FaktorRisikoTinggi13 = support.Nullable(d.FaktorRisikoTinggi13)
		m.NilaiRisikoTinggi13 = support.Nullable(d.NilaiRisikoTinggi13)
		m.FaktorKecurigaanGanas1 = support.Nullable(d.FaktorKecurigaanGanas1)
		m.NilaiKecurigaanGanas1 = support.Nullable(d.NilaiKecurigaanGanas1)
		m.FaktorKecurigaanGanas2 = support.Nullable(d.FaktorKecurigaanGanas2)
		m.NilaiKecurigaanGanas2 = support.Nullable(d.NilaiKecurigaanGanas2)
		m.FaktorKecurigaanGanas3 = support.Nullable(d.FaktorKecurigaanGanas3)
		m.NilaiKecurigaanGanas3 = support.Nullable(d.NilaiKecurigaanGanas3)
		m.FaktorKecurigaanGanas4 = support.Nullable(d.FaktorKecurigaanGanas4)
		m.NilaiKecurigaanGanas4 = support.Nullable(d.NilaiKecurigaanGanas4)
		m.FaktorKecurigaanGanas5 = support.Nullable(d.FaktorKecurigaanGanas5)
		m.NilaiKecurigaanGanas5 = support.Nullable(d.NilaiKecurigaanGanas5)
		m.FaktorKecurigaanGanas6 = support.Nullable(d.FaktorKecurigaanGanas6)
		m.NilaiKecurigaanGanas6 = support.Nullable(d.NilaiKecurigaanGanas6)
		m.FaktorKecurigaanGanas7 = support.Nullable(d.FaktorKecurigaanGanas7)
		m.NilaiKecurigaanGanas7 = support.Nullable(d.NilaiKecurigaanGanas7)
		m.FaktorKecurigaanGanas8 = support.Nullable(d.FaktorKecurigaanGanas8)
		m.NilaiKecurigaanGanas8 = support.Nullable(d.NilaiKecurigaanGanas8)
		m.TotalSkor = support.Nullable(d.TotalSkor)
		m.HasilSadanis = support.Nullable(d.HasilSadanis)
		m.TindakLanjutSadanis = support.Nullable(d.TindakLanjutSadanis)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningRisikoKankerPayudara]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningRisikoKankerPayudara) any { return Str(m.Nip) }},
	},
}
