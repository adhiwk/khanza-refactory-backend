package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilEndoskopiFaringLaring hasil endoskopi faring laring (RMHasilEndoskopiFaringLaring).
var FormHasilEndoskopiFaringLaring = &Form[model.HasilEndoskopiFaringLaring, request.HasilEndoskopiFaringLaringData]{
	Slug:  "hasil-endoskopi-faring-laring",
	Label: "hasil endoskopi faring laring",
	Spec: repo.Spec{
		Table:  "hasil_endoskopi_faring_laring",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.HasilEndoskopiFaringLaring, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.HasilEndoskopiFaringLaring) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.HasilEndoskopiFaringLaring) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.HasilEndoskopiFaringLaring) []string { return []string{m.KdDokter} },
	Fill: func(m *model.HasilEndoskopiFaringLaring, d request.HasilEndoskopiFaringLaringData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.DiagnosaKlinis = strings.TrimSpace(d.DiagnosaKlinis)
		m.KirimanDari = strings.TrimSpace(d.KirimanDari)
		m.FaringUvula = support.Nullable(d.FaringUvula)
		m.FaringArkusFaring = support.Nullable(d.FaringArkusFaring)
		m.FaringDindingPosterior = support.Nullable(d.FaringDindingPosterior)
		m.FaringTonsil = support.Nullable(d.FaringTonsil)
		m.LaringTonsilLingual = support.Nullable(d.LaringTonsilLingual)
		m.LaringValekula = support.Nullable(d.LaringValekula)
		m.LaringSinusPiriformis = support.Nullable(d.LaringSinusPiriformis)
		m.LaringEpiglotis = support.Nullable(d.LaringEpiglotis)
		m.LaringArytenoid = support.Nullable(d.LaringArytenoid)
		m.LaringPlikaVentrikularis = support.Nullable(d.LaringPlikaVentrikularis)
		m.LaringPitaSuara = support.Nullable(d.LaringPitaSuara)
		m.LaringRimaVocalis = support.Nullable(d.LaringRimaVocalis)
		m.LaringLainlain = support.Nullable(d.LaringLainlain)
		m.Kesan = support.Nullable(d.Kesan)
		m.Saran = support.Nullable(d.Saran)
		return nil
	},
	Refs: []Ref[model.HasilEndoskopiFaringLaring]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilEndoskopiFaringLaring) any { return Str(m.KdDokter) }},
	},
}
