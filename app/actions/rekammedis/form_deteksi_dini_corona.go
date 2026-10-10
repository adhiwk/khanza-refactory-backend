package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormDeteksiDiniCorona deteksi dini corona (RMDeteksiDiniCorona).
var FormDeteksiDiniCorona = &Form[model.DeteksiDiniCorona, request.DeteksiDiniCoronaData]{
	Slug:  "deteksi-dini-corona",
	Label: "deteksi dini corona",
	Spec: repo.Spec{
		Table:  "deteksi_dini_corona",
		Keys:   []string{"no_rawat"},
		Waktu:  "",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.DeteksiDiniCorona, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.DeteksiDiniCorona) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.DeteksiDiniCorona) time.Time { return time.Time{} },
	Petugas: func(m *model.DeteksiDiniCorona) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.DeteksiDiniCorona, d request.DeteksiDiniCoronaData) error {
		vTanggal, err := support.ParseDate(d.Tanggal)
		if err != nil {
			return err
		}
		vGejalaTanggalPertama, err := support.ParseDate(d.GejalaTanggalPertama)
		if err != nil {
			return err
		}
		vFaktorTanggalKedatangan, err := support.ParseDate(d.FaktorTanggalKedatangan)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = support.Nullable(d.Nip)
		m.GejalaDemam = support.Nullable(d.GejalaDemam)
		m.GejalaBatuk = support.Nullable(d.GejalaBatuk)
		m.GejalaSesak = support.Nullable(d.GejalaSesak)
		m.GejalaTanggalPertama = vGejalaTanggalPertama
		m.GejalaRiwayatSakit = support.Nullable(d.GejalaRiwayatSakit)
		m.GejalaRiwayatPeriksa = support.Nullable(d.GejalaRiwayatPeriksa)
		m.FaktorRiwayatPerjalanan = strings.TrimSpace(d.FaktorRiwayatPerjalanan)
		m.FaktorAsalDaerah = strings.TrimSpace(d.FaktorAsalDaerah)
		m.FaktorTanggalKedatangan = vFaktorTanggalKedatangan
		m.FaktorPaparanKontakpositif = strings.TrimSpace(d.FaktorPaparanKontakpositif)
		m.FaktorPaparanKontakpdp = strings.TrimSpace(d.FaktorPaparanKontakpdp)
		m.FaktorPaparanFaskespositif = strings.TrimSpace(d.FaktorPaparanFaskespositif)
		m.FaktorPaparanPerjalananln = strings.TrimSpace(d.FaktorPaparanPerjalananln)
		m.FaktorPaparanPasarhewan = strings.TrimSpace(d.FaktorPaparanPasarhewan)
		m.Kesimpulan = strings.TrimSpace(d.Kesimpulan)
		m.TindakLanjut = strings.TrimSpace(d.TindakLanjut)
		return nil
	},
	Refs: []Ref[model.DeteksiDiniCorona]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.DeteksiDiniCorona) any { return StrPtr(m.Nip) }},
	},
}
