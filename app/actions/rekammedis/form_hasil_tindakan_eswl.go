package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormHasilTindakanEswl hasil tindakan ESWL (RMHasilTindakanESWL).
var FormHasilTindakanEswl = &Form[model.HasilTindakanEswl, request.HasilTindakanEswlData]{
	Slug:  "hasil-tindakan-eswl",
	Label: "hasil tindakan ESWL",
	Spec: repo.Spec{
		Table:  "hasil_tindakan_eswl",
		Keys:   []string{"no_rawat", "mulai"},
		Waktu:  "mulai",
		Search: []string{"no_rawat", "kd_dokter", "nip"},
	},
	SetKey: func(m *model.HasilTindakanEswl, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Mulai, err = KeyDateTime(key["mulai"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.HasilTindakanEswl) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "mulai": FmtDateTime(m.Mulai)}
	},
	Waktu:   func(m *model.HasilTindakanEswl) time.Time { return Tm(m.Mulai) },
	Petugas: func(m *model.HasilTindakanEswl) []string { return []string{m.KdDokter, m.Nip} },
	Fill: func(m *model.HasilTindakanEswl, d request.HasilTindakanEswlData) error {
		vSelesai, err := support.ParseDateTime(d.Selesai)
		if err != nil {
			return err
		}
		m.Selesai = vSelesai
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Nip = strings.TrimSpace(d.Nip)
		m.Diagnosa = strings.TrimSpace(d.Diagnosa)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.ObatAnalgesik = strings.TrimSpace(d.ObatAnalgesik)
		m.ObatLain = strings.TrimSpace(d.ObatLain)
		m.UraianTindakan = strings.TrimSpace(d.UraianTindakan)
		m.UraianTindakanFocus = strings.TrimSpace(d.UraianTindakanFocus)
		m.UraianTindakanRate = strings.TrimSpace(d.UraianTindakanRate)
		m.UraianTindakanPower = strings.TrimSpace(d.UraianTindakanPower)
		m.UraianTindakanShock = strings.TrimSpace(d.UraianTindakanShock)
		m.Diintegrasi = strings.TrimSpace(d.Diintegrasi)
		m.Kekurangan = strings.TrimSpace(d.Kekurangan)
		m.Anjungan = strings.TrimSpace(d.Anjungan)
		return nil
	},
	Refs: []Ref[model.HasilTindakanEswl]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.HasilTindakanEswl) any { return Str(m.KdDokter) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.HasilTindakanEswl) any { return Str(m.Nip) }},
	},
}
