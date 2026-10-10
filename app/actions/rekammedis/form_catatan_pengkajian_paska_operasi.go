package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanPengkajianPaskaOperasi catatan pengkajian paska operasi (RMCatatanPengkajianPaskaOperasi).
var FormCatatanPengkajianPaskaOperasi = &Form[model.CatatanPengkajianPaskaOperasi, request.CatatanPengkajianPaskaOperasiData]{
	Slug:  "catatan-pengkajian-paska-operasi",
	Label: "catatan pengkajian paska operasi",
	Spec: repo.Spec{
		Table:  "catatan_pengkajian_paska_operasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.CatatanPengkajianPaskaOperasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.CatatanPengkajianPaskaOperasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.CatatanPengkajianPaskaOperasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.CatatanPengkajianPaskaOperasi) []string { return []string{m.KdDokter} },
	Fill: func(m *model.CatatanPengkajianPaskaOperasi, d request.CatatanPengkajianPaskaOperasiData) error {
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.RawatPaskaOperasi = support.Nullable(d.RawatPaskaOperasi)
		m.Cairan = support.Nullable(d.Cairan)
		m.Antibiotika = support.Nullable(d.Antibiotika)
		m.Analgetika = support.Nullable(d.Analgetika)
		m.MedikamentosaLain = support.Nullable(d.MedikamentosaLain)
		m.Diet = support.Nullable(d.Diet)
		m.PemeriksaanLaborat = support.Nullable(d.PemeriksaanLaborat)
		m.Tranfusi = support.Nullable(d.Tranfusi)
		m.Lainlain = support.Nullable(d.Lainlain)
		return nil
	},
	Refs: []Ref[model.CatatanPengkajianPaskaOperasi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.CatatanPengkajianPaskaOperasi) any { return Str(m.KdDokter) }},
	},
}
