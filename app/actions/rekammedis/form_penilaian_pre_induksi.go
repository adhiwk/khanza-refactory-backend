package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPreInduksi penilaian pre induksi (RMPenilaianPreInduksi).
var FormPenilaianPreInduksi = &Form[model.PenilaianPreInduksi, request.PenilaianPreInduksiData]{
	Slug:  "penilaian-pre-induksi",
	Label: "penilaian pre induksi",
	Spec: repo.Spec{
		Table:  "penilaian_pre_induksi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianPreInduksi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianPreInduksi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianPreInduksi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPreInduksi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianPreInduksi, d request.PenilaianPreInduksiData) error {
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Tensi = support.Nullable(d.Tensi)
		m.Nadi = support.Nullable(d.Nadi)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Ekg = support.Nullable(d.Ekg)
		m.LainLain = support.Nullable(d.LainLain)
		m.Asesmen = support.Nullable(d.Asesmen)
		m.Perencanaan = support.Nullable(d.Perencanaan)
		m.InfusPerifier = support.Nullable(d.InfusPerifier)
		m.Cvc = support.Nullable(d.Cvc)
		m.Posisi = support.Nullable(d.Posisi)
		m.Premedikasi = support.Nullable(d.Premedikasi)
		m.PremedikasiKeterangan = support.Nullable(d.PremedikasiKeterangan)
		m.Induksi = support.Nullable(d.Induksi)
		m.InduksiKeterangan = support.Nullable(d.InduksiKeterangan)
		m.FaceMaskNo = support.Nullable(d.FaceMaskNo)
		m.NasopharingNo = support.Nullable(d.NasopharingNo)
		m.EttNo = support.Nullable(d.EttNo)
		m.EttJenis = support.Nullable(d.EttJenis)
		m.EttViksasi = support.Nullable(d.EttViksasi)
		m.LmaNo = support.Nullable(d.LmaNo)
		m.LmaJenis = support.Nullable(d.LmaJenis)
		m.Tracheostomi = support.Nullable(d.Tracheostomi)
		m.BronchoscopiFiberoptik = support.Nullable(d.BronchoscopiFiberoptik)
		m.Glidescopi = support.Nullable(d.Glidescopi)
		m.LainLainTatalaksana = support.Nullable(d.LainLainTatalaksana)
		m.IntubasiSesudahTidur = support.Nullable(d.IntubasiSesudahTidur)
		m.IntubasiOral = support.Nullable(d.IntubasiOral)
		m.IntubasiTracheostomi = support.Nullable(d.IntubasiTracheostomi)
		m.IntubasiKeterangan = support.Nullable(d.IntubasiKeterangan)
		m.SulitVentilasi = support.Nullable(d.SulitVentilasi)
		m.SulitIntubasi = support.Nullable(d.SulitIntubasi)
		m.Ventilasi = support.Nullable(d.Ventilasi)
		m.TeknikRegionalJenis = support.Nullable(d.TeknikRegionalJenis)
		m.TeknikRegionalLokasi = support.Nullable(d.TeknikRegionalLokasi)
		m.TeknikRegionalJenisJarum = support.Nullable(d.TeknikRegionalJenisJarum)
		m.TeknikRegionalKateter = support.Nullable(d.TeknikRegionalKateter)
		m.TeknikRegionalKateterViksasi = support.Nullable(d.TeknikRegionalKateterViksasi)
		m.TeknikRegionalObatObatan = support.Nullable(d.TeknikRegionalObatObatan)
		m.TeknikRegionalKomplikasi = support.Nullable(d.TeknikRegionalKomplikasi)
		m.TeknikRegionalHasil = support.Nullable(d.TeknikRegionalHasil)
		return nil
	},
	Refs: []Ref[model.PenilaianPreInduksi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianPreInduksi) any { return Str(m.KdDokter) }},
	},
}
