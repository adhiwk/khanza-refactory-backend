package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningKankerKolorektal skrining kanker kolorektal (RMSkriningKankerKolorektal).
var FormSkriningKankerKolorektal = &Form[model.SkriningKankerKolorektal, request.SkriningKankerKolorektalData]{
	Slug:  "skrining-kanker-kolorektal",
	Label: "skrining kanker kolorektal",
	Spec: repo.Spec{
		Table:  "skrining_kanker_kolorektal",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningKankerKolorektal, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningKankerKolorektal) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningKankerKolorektal) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningKankerKolorektal) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningKankerKolorektal, d request.SkriningKankerKolorektalData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.RiwayatPolipAdenomatosa = support.Nullable(d.RiwayatPolipAdenomatosa)
		m.RiwayatBabBerdarah = support.Nullable(d.RiwayatBabBerdarah)
		m.RiwayatReseksiKuratif = support.Nullable(d.RiwayatReseksiKuratif)
		m.ColokDubur = support.Nullable(d.ColokDubur)
		m.RiwayatKolorektalKeluarga = support.Nullable(d.RiwayatKolorektalKeluarga)
		m.DarahSamarFeses = support.Nullable(d.DarahSamarFeses)
		m.RujukFaskesLanjut = support.Nullable(d.RujukFaskesLanjut)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.KeteranganKesimpulan = support.Nullable(d.KeteranganKesimpulan)
		return nil
	},
	Refs: []Ref[model.SkriningKankerKolorektal]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningKankerKolorektal) any { return Str(m.Nip) }},
	},
}
