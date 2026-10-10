package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMedisRalanRehabMedik penilaian awal medis ralan rehab medik (RMPenilaianAwalMedisRalanRehabMedik).
var FormPenilaianMedisRalanRehabMedik = &Form[model.PenilaianMedisRalanRehabMedik, request.PenilaianMedisRalanRehabMedikData]{
	Slug:  "penilaian-medis-ralan-rehab-medik",
	Label: "penilaian awal medis ralan rehab medik",
	Spec: repo.Spec{
		Table:  "penilaian_medis_ralan_rehab_medik",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMedisRalanRehabMedik, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMedisRalanRehabMedik) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMedisRalanRehabMedik) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMedisRalanRehabMedik) []string { return []string{derefStr(m.KdDokter)} },
	Fill: func(m *model.PenilaianMedisRalanRehabMedik, d request.PenilaianMedisRalanRehabMedikData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		vFisioterapi, err := support.ParseDate(d.Fisioterapi)
		if err != nil {
			return err
		}
		vTerapiOkupasi, err := support.ParseDate(d.TerapiOkupasi)
		if err != nil {
			return err
		}
		vTerapiWicara, err := support.ParseDate(d.TerapiWicara)
		if err != nil {
			return err
		}
		vTerapiAkupuntur, err := support.ParseDate(d.TerapiAkupuntur)
		if err != nil {
			return err
		}
		vTerapiLainnya, err := support.ParseDate(d.TerapiLainnya)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = support.Nullable(d.KdDokter)
		m.Anamnesis = support.Nullable(d.Anamnesis)
		m.Hubungan = support.Nullable(d.Hubungan)
		m.KeluhanUtama = support.Nullable(d.KeluhanUtama)
		m.Rps = support.Nullable(d.Rps)
		m.Rpd = support.Nullable(d.Rpd)
		m.Alergi = support.Nullable(d.Alergi)
		m.Kesadaran = support.Nullable(d.Kesadaran)
		m.Nyeri = support.Nullable(d.Nyeri)
		m.SkalaNyeri = support.Nullable(d.SkalaNyeri)
		m.Td = support.Nullable(d.Td)
		m.Nadi = support.Nullable(d.Nadi)
		m.Suhu = support.Nullable(d.Suhu)
		m.Rr = support.Nullable(d.Rr)
		m.Bb = support.Nullable(d.Bb)
		m.Kepala = support.Nullable(d.Kepala)
		m.KeteranganKepala = support.Nullable(d.KeteranganKepala)
		m.Thoraks = support.Nullable(d.Thoraks)
		m.KeteranganThoraks = support.Nullable(d.KeteranganThoraks)
		m.Abdomen = support.Nullable(d.Abdomen)
		m.KeteranganAbdomen = support.Nullable(d.KeteranganAbdomen)
		m.Ekstremitas = support.Nullable(d.Ekstremitas)
		m.KeteranganEkstremitas = support.Nullable(d.KeteranganEkstremitas)
		m.Columna = support.Nullable(d.Columna)
		m.KeteranganColumna = support.Nullable(d.KeteranganColumna)
		m.Muskulos = support.Nullable(d.Muskulos)
		m.KeteranganMuskulos = support.Nullable(d.KeteranganMuskulos)
		m.Lainnya = support.Nullable(d.Lainnya)
		m.ResikoJatuh = support.Nullable(d.ResikoJatuh)
		m.ResikoNutrisional = support.Nullable(d.ResikoNutrisional)
		m.KebutuhanFungsional = support.Nullable(d.KebutuhanFungsional)
		m.DiagnosaMedis = support.Nullable(d.DiagnosaMedis)
		m.DiagnosaFungsi = support.Nullable(d.DiagnosaFungsi)
		m.PenunjangLain = support.Nullable(d.PenunjangLain)
		m.Fisio = support.Nullable(d.Fisio)
		m.Okupasi = support.Nullable(d.Okupasi)
		m.Wicara = support.Nullable(d.Wicara)
		m.Akupuntur = support.Nullable(d.Akupuntur)
		m.Tatalain = support.Nullable(d.Tatalain)
		m.FrekuensiTerapi = support.Nullable(d.FrekuensiTerapi)
		m.Fisioterapi = vFisioterapi
		m.TerapiOkupasi = vTerapiOkupasi
		m.TerapiWicara = vTerapiWicara
		m.TerapiAkupuntur = vTerapiAkupuntur
		m.TerapiLainnya = vTerapiLainnya
		m.Edukasi = support.Nullable(d.Edukasi)
		return nil
	},
	Refs: []Ref[model.PenilaianMedisRalanRehabMedik]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMedisRalanRehabMedik) any { return StrPtr(m.KdDokter) }},
	},
}
