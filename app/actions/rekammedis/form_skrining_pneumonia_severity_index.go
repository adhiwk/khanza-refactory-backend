package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningPneumoniaSeverityIndex skrining pneumonia severity index (RMSkriningPneumoniaSeverityIndex).
var FormSkriningPneumoniaSeverityIndex = &Form[model.SkriningPneumoniaSeverityIndex, request.SkriningPneumoniaSeverityIndexData]{
	Slug:  "skrining-pneumonia-severity-index",
	Label: "skrining pneumonia severity index",
	Spec: repo.Spec{
		Table:  "skrining_pneumonia_severity_index",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningPneumoniaSeverityIndex, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningPneumoniaSeverityIndex) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningPneumoniaSeverityIndex) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningPneumoniaSeverityIndex) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningPneumoniaSeverityIndex, d request.SkriningPneumoniaSeverityIndexData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("rekomendasi", d.Rekomendasi, "Rawat Jalan", "Observasi/Rawat Inap Singkat", "Rawat Inap", "Rawat Inap, Pertimbangkan ICU"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.NilaiUmur = d.NilaiUmur
		m.TinggalDiPantiJompo = support.Nullable(d.TinggalDiPantiJompo)
		m.NilaiTinggalDiPantiJompo = d.NilaiTinggalDiPantiJompo
		m.GagalJantung = support.Nullable(d.GagalJantung)
		m.NilaiGagalJantung = d.NilaiGagalJantung
		m.PenyakitHati = support.Nullable(d.PenyakitHati)
		m.NilaiPenyakitHati = d.NilaiPenyakitHati
		m.PenyakitGinjal = support.Nullable(d.PenyakitGinjal)
		m.NilaiPenyakitGinjal = d.NilaiPenyakitGinjal
		m.Kanker = support.Nullable(d.Kanker)
		m.NilaiKanker = d.NilaiKanker
		m.PenyakitSerebrovaskuler = support.Nullable(d.PenyakitSerebrovaskuler)
		m.NilaiPenyakitSerebrovaskuler = d.NilaiPenyakitSerebrovaskuler
		m.DisorentasiMental = support.Nullable(d.DisorentasiMental)
		m.NilaiDisorentasiMental = d.NilaiDisorentasiMental
		m.FrekuensiNapas = support.Nullable(d.FrekuensiNapas)
		m.NilaiFrekuensiNapas = d.NilaiFrekuensiNapas
		m.TdSistolik = support.Nullable(d.TdSistolik)
		m.NilaiTdSistolik = d.NilaiTdSistolik
		m.Suhu = support.Nullable(d.Suhu)
		m.NilaiSuhu = d.NilaiSuhu
		m.Nadi = support.Nullable(d.Nadi)
		m.NilaiNadi = d.NilaiNadi
		m.PhDarah = support.Nullable(d.PhDarah)
		m.NilaiPhDarah = d.NilaiPhDarah
		m.Natrium = support.Nullable(d.Natrium)
		m.NilaiNatrium = d.NilaiNatrium
		m.Bun = support.Nullable(d.Bun)
		m.NilaiBun = d.NilaiBun
		m.Pao = support.Nullable(d.Pao)
		m.NilaiPao = d.NilaiPao
		m.Glukosa = support.Nullable(d.Glukosa)
		m.NilaiGlukosa = d.NilaiGlukosa
		m.EfusiPleura = support.Nullable(d.EfusiPleura)
		m.NilaiEfusiPleura = d.NilaiEfusiPleura
		m.Hematokrit = support.Nullable(d.Hematokrit)
		m.NilaiHematokrit = d.NilaiHematokrit
		m.TotalSkor = d.TotalSkor
		m.Kelas = support.Nullable(d.Kelas)
		m.SkorInterpretasi = support.Nullable(d.SkorInterpretasi)
		m.Mortalitas = support.Nullable(d.Mortalitas)
		m.Rekomendasi = support.Nullable(d.Rekomendasi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningPneumoniaSeverityIndex]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningPneumoniaSeverityIndex) any { return Str(m.Nip) }},
	},
}
