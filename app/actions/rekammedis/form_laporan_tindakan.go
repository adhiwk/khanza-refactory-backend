package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormLaporanTindakan laporan tindakan (RMLaporanTindakan).
var FormLaporanTindakan = &Form[model.LaporanTindakan, request.LaporanTindakanData]{
	Slug:  "laporan-tindakan",
	Label: "laporan tindakan",
	Spec: repo.Spec{
		Table:  "laporan_tindakan",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter", "nip"},
	},
	SetKey: func(m *model.LaporanTindakan, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.LaporanTindakan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.LaporanTindakan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.LaporanTindakan) []string { return []string{m.KdDokter, derefStr(m.Nip)} },
	Fill: func(m *model.LaporanTindakan, d request.LaporanTindakanData) error {
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Nip = support.Nullable(d.Nip)
		m.DiagnosaPraTindakan = strings.TrimSpace(d.DiagnosaPraTindakan)
		m.DiagnosaPascaTindakan = strings.TrimSpace(d.DiagnosaPascaTindakan)
		m.TindakanMedik = strings.TrimSpace(d.TindakanMedik)
		m.Uraian = strings.TrimSpace(d.Uraian)
		m.Hasil = strings.TrimSpace(d.Hasil)
		m.Kesimpulan = strings.TrimSpace(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.LaporanTindakan]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.LaporanTindakan) any { return Str(m.KdDokter) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.LaporanTindakan) any { return StrPtr(m.Nip) }},
	},
}
