package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanPersalinan catatan persalinan (RMCatatanPersalinan).
var FormCatatanPersalinan = &Form[model.CatatanPersalinan, request.CatatanPersalinanData]{
	Slug:  "catatan-persalinan",
	Label: "catatan persalinan",
	Spec: repo.Spec{
		Table:  "catatan_persalinan",
		Keys:   []string{"no_rawat"},
		Waktu:  "",
		Search: []string{"no_rawat", "kd_dokter", "nip"},
	},
	SetKey: func(m *model.CatatanPersalinan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.CatatanPersalinan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.CatatanPersalinan) time.Time { return time.Time{} },
	Petugas: func(m *model.CatatanPersalinan) []string { return []string{m.KdDokter, m.Nip} },
	Fill: func(m *model.CatatanPersalinan, d request.CatatanPersalinanData) error {
		vMulai, err := support.ParseDateTime(d.Mulai)
		if err != nil {
			return err
		}
		vSelesai, err := support.ParseDateTime(d.Selesai)
		if err != nil {
			return err
		}
		m.Mulai = vMulai
		m.Selesai = vSelesai
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Nip = strings.TrimSpace(d.Nip)
		m.Catatan = support.Nullable(d.Catatan)
		m.WaktuPersalinanKala1 = support.Nullable(d.WaktuPersalinanKala1)
		m.WaktuPersalinanKala2 = support.Nullable(d.WaktuPersalinanKala2)
		m.WaktuPersalinanKala3 = support.Nullable(d.WaktuPersalinanKala3)
		m.WaktuPersalinanJumlah = support.Nullable(d.WaktuPersalinanJumlah)
		m.Perineum = support.Nullable(d.Perineum)
		m.JahitanLuar1 = support.Nullable(d.JahitanLuar1)
		m.JahitanLuar2 = support.Nullable(d.JahitanLuar2)
		m.JahitanDalam1 = support.Nullable(d.JahitanDalam1)
		m.JahitanDalam2 = support.Nullable(d.JahitanDalam2)
		m.Anak = support.Nullable(d.Anak)
		m.StatusLahir = support.Nullable(d.StatusLahir)
		m.ApgarScore = support.Nullable(d.ApgarScore)
		m.Bb = support.Nullable(d.Bb)
		m.Pb = support.Nullable(d.Pb)
		m.Kelainan = support.Nullable(d.Kelainan)
		m.Ketuban = support.Nullable(d.Ketuban)
		m.Placenta = support.Nullable(d.Placenta)
		m.Ukuran = support.Nullable(d.Ukuran)
		m.TaliPusat = support.Nullable(d.TaliPusat)
		m.Insertio = support.Nullable(d.Insertio)
		m.DarahKeluarKala1 = support.Nullable(d.DarahKeluarKala1)
		m.DarahKeluarKala2 = support.Nullable(d.DarahKeluarKala2)
		m.DarahKeluarKala3 = support.Nullable(d.DarahKeluarKala3)
		m.DarahKeluarKala4 = support.Nullable(d.DarahKeluarKala4)
		m.DarahKeluarJumlah = support.Nullable(d.DarahKeluarJumlah)
		m.KondisiUmum = support.Nullable(d.KondisiUmum)
		m.Td = support.Nullable(d.Td)
		m.Nadi = support.Nullable(d.Nadi)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.KontraksiUterus = support.Nullable(d.KontraksiUterus)
		m.Ppv = support.Nullable(d.Ppv)
		m.Pengobatan = support.Nullable(d.Pengobatan)
		return nil
	},
	Refs: []Ref[model.CatatanPersalinan]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.CatatanPersalinan) any { return Str(m.KdDokter) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanPersalinan) any { return Str(m.Nip) }},
	},
}
