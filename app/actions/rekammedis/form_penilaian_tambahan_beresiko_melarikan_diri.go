package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianTambahanBeresikoMelarikanDiri penilaian tambahan melarikan diri (RMPenilaianTambahanMelarikanDiri).
var FormPenilaianTambahanBeresikoMelarikanDiri = &Form[model.PenilaianTambahanBeresikoMelarikanDiri, request.PenilaianTambahanBeresikoMelarikanDiriData]{
	Slug:  "penilaian-tambahan-beresiko-melarikan-diri",
	Label: "penilaian tambahan melarikan diri",
	Spec: repo.Spec{
		Table:  "penilaian_tambahan_beresiko_melarikan_diri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianTambahanBeresikoMelarikanDiri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianTambahanBeresikoMelarikanDiri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianTambahanBeresikoMelarikanDiri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianTambahanBeresikoMelarikanDiri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianTambahanBeresikoMelarikanDiri, d request.PenilaianTambahanBeresikoMelarikanDiriData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.StatikRiwayatMelarikanDiri = support.Nullable(d.StatikRiwayatMelarikanDiri)
		m.StatikSkorriwayatMelarikanDiri = d.StatikSkorriwayatMelarikanDiri
		m.StatikRiwayatPenolakanPengobatan = support.Nullable(d.StatikRiwayatPenolakanPengobatan)
		m.StatikSkorriwayatPenolakanPengobatan = d.StatikSkorriwayatPenolakanPengobatan
		m.StatikUsiaDibawah35 = support.Nullable(d.StatikUsiaDibawah35)
		m.StatikSkorusiaDibawah35 = d.StatikSkorusiaDibawah35
		m.StatikLakiLaki = support.Nullable(d.StatikLakiLaki)
		m.StatikSkorlakiLaki = d.StatikSkorlakiLaki
		m.StatikDiagnosisSkizofrenia = support.Nullable(d.StatikDiagnosisSkizofrenia)
		m.StatikSkordiagnosisSkizofrenia = d.StatikSkordiagnosisSkizofrenia
		m.StatikBelumMenikah = support.Nullable(d.StatikBelumMenikah)
		m.StatikSkorbelumMenikah = d.StatikSkorbelumMenikah
		m.StatikRiwayatPenggunaanNapza = support.Nullable(d.StatikRiwayatPenggunaanNapza)
		m.StatikSkoriwayatPenggunaanNapza = d.StatikSkoriwayatPenggunaanNapza
		m.StatikDiagnosisGangguanKepribadian = support.Nullable(d.StatikDiagnosisGangguanKepribadian)
		m.StatikSkordiagnosisGangguanKepribadian = d.StatikSkordiagnosisGangguanKepribadian
		m.StatikRiwayatKriminal = support.Nullable(d.StatikRiwayatKriminal)
		m.StatikSkorriwayatKriminal = d.StatikSkorriwayatKriminal
		m.StatikSkortotal = d.StatikSkortotal
		m.DinamisAntiTreatment = support.Nullable(d.DinamisAntiTreatment)
		m.DinamisSkorantiTreatment = d.DinamisSkorantiTreatment
		m.DinamisPenggunaanNapza = support.Nullable(d.DinamisPenggunaanNapza)
		m.DinamisSkorpenggunaanNapza = d.DinamisSkorpenggunaanNapza
		m.DinamisKebosanan = support.Nullable(d.DinamisKebosanan)
		m.DinamisSkorkebosanan = d.DinamisSkorkebosanan
		m.DinamisPerintahHalusinasi = support.Nullable(d.DinamisPerintahHalusinasi)
		m.DinamisSkorperintahHalusinasi = d.DinamisSkorperintahHalusinasi
		m.DinamisHilangnyaKontrolDiri = support.Nullable(d.DinamisHilangnyaKontrolDiri)
		m.DinamisSkorhilangnyaKontrolDiri = d.DinamisSkorhilangnyaKontrolDiri
		m.DinamisSeksualTidakWajar = support.Nullable(d.DinamisSeksualTidakWajar)
		m.DinamisSkorseksualTidakWajar = d.DinamisSkorseksualTidakWajar
		m.DinamisKemarahanFrustasi = support.Nullable(d.DinamisKemarahanFrustasi)
		m.DinamisSkorkemarahanFrustasi = d.DinamisSkorkemarahanFrustasi
		m.DinamisKetakutanPerawatan = support.Nullable(d.DinamisKetakutanPerawatan)
		m.DinamisSkorketakutanPerawatan = d.DinamisSkorketakutanPerawatan
		m.DinamisSkortotal = d.DinamisSkortotal
		m.FaktorFaktorPencegahan = support.Nullable(d.FaktorFaktorPencegahan)
		m.TotalSkor = d.TotalSkor
		m.LevelSkor = support.Nullable(d.LevelSkor)
		return nil
	},
	Refs: []Ref[model.PenilaianTambahanBeresikoMelarikanDiri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianTambahanBeresikoMelarikanDiri) any { return Str(m.Nip) }},
	},
}
