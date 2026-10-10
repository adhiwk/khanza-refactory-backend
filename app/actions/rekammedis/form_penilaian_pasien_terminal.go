package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianPasienTerminal penilaian pasien terminal (RMPenilaianPasienTerminal).
var FormPenilaianPasienTerminal = &Form[model.PenilaianPasienTerminal, request.PenilaianPasienTerminalData]{
	Slug:  "penilaian-pasien-terminal",
	Label: "penilaian pasien terminal",
	Spec: repo.Spec{
		Table:  "penilaian_pasien_terminal",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianPasienTerminal, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianPasienTerminal) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianPasienTerminal) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianPasienTerminal) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianPasienTerminal, d request.PenilaianPasienTerminalData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Diagnosa = strings.TrimSpace(d.Diagnosa)
		m.Rps = strings.TrimSpace(d.Rps)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.KeadaanUmum = support.Nullable(d.KeadaanUmum)
		m.Kesadaran = support.Nullable(d.Kesadaran)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Spo2 = strings.TrimSpace(d.Spo2)
		m.SkalaNyeri = strings.TrimSpace(d.SkalaNyeri)
		m.TahapPasienMenjelangAjal = support.Nullable(d.TahapPasienMenjelangAjal)
		m.TandaKlinisMenjelangKematian = support.Nullable(d.TandaKlinisMenjelangKematian)
		m.KebutuhanSpiritualPasien = support.Nullable(d.KebutuhanSpiritualPasien)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianPasienTerminal]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianPasienTerminal) any { return Str(m.Nip) }},
	},
}
