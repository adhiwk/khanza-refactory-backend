package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianFisioterapi penilaian fisioterapi (RMPenilaianFisioterapi).
var FormPenilaianFisioterapi = &Form[model.PenilaianFisioterapi, request.PenilaianFisioterapiData]{
	Slug:  "penilaian-fisioterapi",
	Label: "penilaian fisioterapi",
	Spec: repo.Spec{
		Table:  "penilaian_fisioterapi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianFisioterapi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianFisioterapi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianFisioterapi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianFisioterapi) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianFisioterapi, d request.PenilaianFisioterapiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.KeluhanUtama = strings.TrimSpace(d.KeluhanUtama)
		m.Rps = strings.TrimSpace(d.Rps)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = strings.TrimSpace(d.Hr)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.NyeriTekan = strings.TrimSpace(d.NyeriTekan)
		m.NyeriGerak = strings.TrimSpace(d.NyeriGerak)
		m.NyeriDiam = strings.TrimSpace(d.NyeriDiam)
		m.Palpasi = strings.TrimSpace(d.Palpasi)
		m.LuasGerakSendi = strings.TrimSpace(d.LuasGerakSendi)
		m.KekuatanOtot = strings.TrimSpace(d.KekuatanOtot)
		m.Statis = strings.TrimSpace(d.Statis)
		m.Dinamis = strings.TrimSpace(d.Dinamis)
		m.Kognitif = strings.TrimSpace(d.Kognitif)
		m.Auskultasi = strings.TrimSpace(d.Auskultasi)
		m.AlatBantu = strings.TrimSpace(d.AlatBantu)
		m.KetBantu = strings.TrimSpace(d.KetBantu)
		m.Prothesa = strings.TrimSpace(d.Prothesa)
		m.KetPro = strings.TrimSpace(d.KetPro)
		m.Deformitas = strings.TrimSpace(d.Deformitas)
		m.KetDeformitas = strings.TrimSpace(d.KetDeformitas)
		m.Resikojatuh = strings.TrimSpace(d.Resikojatuh)
		m.KetResikojatuh = strings.TrimSpace(d.KetResikojatuh)
		m.Adl = strings.TrimSpace(d.Adl)
		m.LainlainFungsional = strings.TrimSpace(d.LainlainFungsional)
		m.KetFisik = strings.TrimSpace(d.KetFisik)
		m.PemeriksaanMusculoskeletal = strings.TrimSpace(d.PemeriksaanMusculoskeletal)
		m.PemeriksaanNeuromuscular = strings.TrimSpace(d.PemeriksaanNeuromuscular)
		m.PemeriksaanCardiopulmonal = strings.TrimSpace(d.PemeriksaanCardiopulmonal)
		m.PemeriksaanIntegument = strings.TrimSpace(d.PemeriksaanIntegument)
		m.PengukuranMusculoskeletal = strings.TrimSpace(d.PengukuranMusculoskeletal)
		m.PengukuranNeuromuscular = strings.TrimSpace(d.PengukuranNeuromuscular)
		m.PengukuranCardiopulmonal = strings.TrimSpace(d.PengukuranCardiopulmonal)
		m.PengukuranIntegument = strings.TrimSpace(d.PengukuranIntegument)
		m.Penunjang = strings.TrimSpace(d.Penunjang)
		m.DiagnosisFisio = strings.TrimSpace(d.DiagnosisFisio)
		m.RencanaTerapi = strings.TrimSpace(d.RencanaTerapi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianFisioterapi]{
		{Column: "nip", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.PenilaianFisioterapi) any { return Str(m.Nip) }},
	},
}
