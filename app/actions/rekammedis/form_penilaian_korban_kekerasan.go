package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianKorbanKekerasan penilaian korban kekerasan (RMPenilaianKorbanKekerasan).
var FormPenilaianKorbanKekerasan = &Form[model.PenilaianKorbanKekerasan, request.PenilaianKorbanKekerasanData]{
	Slug:  "penilaian-korban-kekerasan",
	Label: "penilaian korban kekerasan",
	Spec: repo.Spec{
		Table:  "penilaian_korban_kekerasan",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianKorbanKekerasan, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianKorbanKekerasan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianKorbanKekerasan) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianKorbanKekerasan) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianKorbanKekerasan, d request.PenilaianKorbanKekerasanData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Informasi = support.Nullable(d.Informasi)
		m.HubunganDenganPasien = support.Nullable(d.HubunganDenganPasien)
		m.JumlahSaudara = support.Nullable(d.JumlahSaudara)
		m.KondisiKeluaga = support.Nullable(d.KondisiKeluaga)
		m.HubunganOrangTerdekat = support.Nullable(d.HubunganOrangTerdekat)
		m.KekerasanYangDialami = support.Nullable(d.KekerasanYangDialami)
		m.TempatKejadian = support.Nullable(d.TempatKejadian)
		m.LamaKekerasan = d.LamaKekerasan
		m.PeriodeKekerasan = support.Nullable(d.PeriodeKekerasan)
		m.SeberapaSeringMengalami = support.Nullable(d.SeberapaSeringMengalami)
		m.PemicuKekerasan = support.Nullable(d.PemicuKekerasan)
		m.YangMelakukanKekerasan = support.Nullable(d.YangMelakukanKekerasan)
		m.DampakKekerasan = support.Nullable(d.DampakKekerasan)
		m.TandaTandaDidapatkan = support.Nullable(d.TandaTandaDidapatkan)
		m.MemerlukanPendampingan = support.Nullable(d.MemerlukanPendampingan)
		m.RiwayatKelainan = support.Nullable(d.RiwayatKelainan)
		m.PemeriksaanKepala = support.Nullable(d.PemeriksaanKepala)
		m.PemeriksaanThoraks = support.Nullable(d.PemeriksaanThoraks)
		m.PemeriksaanLeher = support.Nullable(d.PemeriksaanLeher)
		m.PemeriksaanAbdomen = support.Nullable(d.PemeriksaanAbdomen)
		m.PemeriksaanGenitalia = support.Nullable(d.PemeriksaanGenitalia)
		m.PemeriksaanEkstrimitasAtas = support.Nullable(d.PemeriksaanEkstrimitasAtas)
		m.PemeriksaanEkstrimitasBawah = strings.TrimSpace(d.PemeriksaanEkstrimitasBawah)
		m.PemeriksaanAnus = support.Nullable(d.PemeriksaanAnus)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianKorbanKekerasan]{
		{Column: "nip", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.PenilaianKorbanKekerasan) any { return Str(m.Nip) }},
	},
}
