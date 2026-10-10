package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningAdiksiNikotin skrining adiksi nikotin (RMSkriningAdiksiNikotin).
var FormSkriningAdiksiNikotin = &Form[model.SkriningAdiksiNikotin, request.SkriningAdiksiNikotinData]{
	Slug:  "skrining-adiksi-nikotin",
	Label: "skrining adiksi nikotin",
	Spec: repo.Spec{
		Table:  "skrining_adiksi_nikotin",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningAdiksiNikotin, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningAdiksiNikotin) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningAdiksiNikotin) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningAdiksiNikotin) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningAdiksiNikotin, d request.SkriningAdiksiNikotinData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("skala_motivasi", d.SkalaMotivasi, "1. Saya SUDAH memutuskan TIDAK akan berhenti merokok seumur hidup Saya", "2. Saya TIDAK PERNAH berpikir untuk berhenti merokok. Saya TIDAK PUNYA rencana untuk berhenti", "3. Saya PERNAH berpikir untuk berhentl merokok, tetapi Saya TIDAK PUNYA rencana", "4. TERKADANG Saya berpikir untuk berhenti merokok, tetapi Saya tidak punya rencana", "5. Saya SERING berpikir untuk berhentl merokok, tetapi Saya tidak punya rencana", "6. Saya BERENCANA untuk berhenti merokok dalam 6 bulan ke depan", "7. Saya berencana untuk berhenti merokok dalam 30 hari ke depan", "8. Saya masih merokok, tetapi Saya mau berubah. Saya siap untuk berhenti merokok", "9. Saya sudah berhenti merokok, tetapi Saya khawatir akan merokok kembali, Saya butuh lingkungan tanpa asap rokok", "10. Saya sudah berhenti merokok"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.RokokDihisap = support.Nullable(d.RokokDihisap)
		m.NilaiRokokDihisap = d.NilaiRokokDihisap
		m.MenyalakanRokok = support.Nullable(d.MenyalakanRokok)
		m.NilaiMenyalakanRokok = d.NilaiMenyalakanRokok
		m.TidakRela = support.Nullable(d.TidakRela)
		m.NilaiTidakRela = d.NilaiTidakRela
		m.JamPertama = support.Nullable(d.JamPertama)
		m.NilaiJamPertama = d.NilaiJamPertama
		m.RasaIngin = support.Nullable(d.RasaIngin)
		m.NilaiRasaIngin = d.NilaiRasaIngin
		m.SakitBerat = support.Nullable(d.SakitBerat)
		m.NilaiSakitBerat = d.NilaiSakitBerat
		m.NilaiTotal = d.NilaiTotal
		m.KeteranganHasilSkrining = support.Nullable(d.KeteranganHasilSkrining)
		m.SkalaMotivasi = support.Nullable(d.SkalaMotivasi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningAdiksiNikotin]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningAdiksiNikotin) any { return Str(m.Nip) }},
	},
}
