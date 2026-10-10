package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianTambahanBunuhDiri penilaian tambahan bunuh diri (RMPenilaianTambahanBunuhDiri).
var FormPenilaianTambahanBunuhDiri = &Form[model.PenilaianTambahanBunuhDiri, request.PenilaianTambahanBunuhDiriData]{
	Slug:  "penilaian-tambahan-bunuh-diri",
	Label: "penilaian tambahan bunuh diri",
	Spec: repo.Spec{
		Table:  "penilaian_tambahan_bunuh_diri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianTambahanBunuhDiri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianTambahanBunuhDiri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianTambahanBunuhDiri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianTambahanBunuhDiri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianTambahanBunuhDiri, d request.PenilaianTambahanBunuhDiriData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.StatikHidupSendiri = support.Nullable(d.StatikHidupSendiri)
		m.StatikSkorhidupSendiri = d.StatikSkorhidupSendiri
		m.StatikUpayaSuicide = support.Nullable(d.StatikUpayaSuicide)
		m.StatikSkorupayaSuicide = d.StatikSkorupayaSuicide
		m.StatikKeluargaSuicide = support.Nullable(d.StatikKeluargaSuicide)
		m.StatikSkorkeluargaSuicide = d.StatikSkorkeluargaSuicide
		m.StatikDiagnosaGangguanJiwa = support.Nullable(d.StatikDiagnosaGangguanJiwa)
		m.StatikSkordiagnosaGangguanJiwa = d.StatikSkordiagnosaGangguanJiwa
		m.StatikDisabilitasBerat = support.Nullable(d.StatikDisabilitasBerat)
		m.StatikSkordisabilitasBerat = d.StatikSkordisabilitasBerat
		m.StatikBerpisah = support.Nullable(d.StatikBerpisah)
		m.StatikSkorberpisah = d.StatikSkorberpisah
		m.StatikKehilanganKerja = support.Nullable(d.StatikKehilanganKerja)
		m.StatikSkorkehilanganKerja = d.StatikSkorkehilanganKerja
		m.StatikSkortotal = d.StatikSkortotal
		m.DinamisIdeBunuhDiri = support.Nullable(d.DinamisIdeBunuhDiri)
		m.DinamisSkorideBunuhDiri = d.DinamisSkorideBunuhDiri
		m.DinamisMaksudSuicide = support.Nullable(d.DinamisMaksudSuicide)
		m.DinamisSkormaksudSuicide = d.DinamisSkormaksudSuicide
		m.DinamisStressBerat = support.Nullable(d.DinamisStressBerat)
		m.DinamisSkorstressBerat = d.DinamisSkorstressBerat
		m.DinamisKeputusasaan = support.Nullable(d.DinamisKeputusasaan)
		m.DinamisSkorkeputusasaan = d.DinamisSkorkeputusasaan
		m.DinamisKejadianSignifikan = support.Nullable(d.DinamisKejadianSignifikan)
		m.DinamisSkorkejadianSignifikan = d.DinamisSkorkejadianSignifikan
		m.DinamisKehilanganKontrol = support.Nullable(d.DinamisKehilanganKontrol)
		m.DinamisSkorkehilanganKontrol = d.DinamisSkorkehilanganKontrol
		m.DinamisPenggunaanNapza = support.Nullable(d.DinamisPenggunaanNapza)
		m.DinamisSkorpenggunaanNapza = d.DinamisSkorpenggunaanNapza
		m.DinamisSkortotal = d.DinamisSkortotal
		m.FaktorFaktorPencegahan = support.Nullable(d.FaktorFaktorPencegahan)
		m.TotalSkor = d.TotalSkor
		m.LevelSkor = support.Nullable(d.LevelSkor)
		return nil
	},
	Refs: []Ref[model.PenilaianTambahanBunuhDiri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianTambahanBunuhDiri) any { return Str(m.Nip) }},
	},
}
