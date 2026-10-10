package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPreOperasi penilaian pre operasi (RMPenilaianPreOperasi).
var FormPenilaianPreOperasi = &Form[model.PenilaianPreOperasi, request.PenilaianPreOperasiData]{
	Slug:  "penilaian-pre-operasi",
	Label: "penilaian pre operasi",
	Spec: repo.Spec{
		Table:  "penilaian_pre_operasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianPreOperasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianPreOperasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianPreOperasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPreOperasi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianPreOperasi, d request.PenilaianPreOperasiData) error {
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.RingkasanKlinik = support.Nullable(d.RingkasanKlinik)
		m.PemeriksaanFisik = support.Nullable(d.PemeriksaanFisik)
		m.PemeriksaanDiagnostik = support.Nullable(d.PemeriksaanDiagnostik)
		m.DiagnosaPreOperasi = support.Nullable(d.DiagnosaPreOperasi)
		m.RencanaTindakanBedah = support.Nullable(d.RencanaTindakanBedah)
		m.HalHalYangPerludiPersiapkan = support.Nullable(d.HalHalYangPerludiPersiapkan)
		m.TerapiPreOperasi = support.Nullable(d.TerapiPreOperasi)
		return nil
	},
	Refs: []Ref[model.PenilaianPreOperasi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianPreOperasi) any { return Str(m.KdDokter) }},
	},
}
