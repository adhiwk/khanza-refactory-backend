package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningTolac skrining TOLAC (RMSkriningTOLAC).
var FormSkriningTolac = &Form[model.SkriningTolac, request.SkriningTolacData]{
	Slug:  "skrining-tolac",
	Label: "skrining TOLAC",
	Spec: repo.Spec{
		Table:  "skrining_tolac",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.SkriningTolac, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningTolac) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningTolac) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningTolac) []string { return []string{derefStr(m.KdDokter)} },
	Fill: func(m *model.SkriningTolac, d request.SkriningTolacData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Gpa = support.Nullable(d.Gpa)
		m.Diagnosa = support.Nullable(d.Diagnosa)
		m.JumlahSc = support.Nullable(d.JumlahSc)
		m.TahunSc = support.Nullable(d.TahunSc)
		m.IndikasiSc = support.Nullable(d.IndikasiSc)
		m.JenisInsisi = support.Nullable(d.JenisInsisi)
		m.RiwayatPervaginam = support.Nullable(d.RiwayatPervaginam)
		m.TbjGram = support.Nullable(d.TbjGram)
		m.PresentasiJanin = support.Nullable(d.PresentasiJanin)
		m.InklusiRiwayatSc = support.Nullable(d.InklusiRiwayatSc)
		m.InklusiPanggulAdekuat = support.Nullable(d.InklusiPanggulAdekuat)
		m.InklusiJaninTunggalKepala = support.Nullable(d.InklusiJaninTunggalKepala)
		m.InklusiTbjSesuai = support.Nullable(d.InklusiTbjSesuai)
		m.EksklusiScKlasikRuptur = support.Nullable(d.EksklusiScKlasikRuptur)
		m.EksklusiSc2x = support.Nullable(d.EksklusiSc2x)
		m.EksklusiPlasentaPrevia = support.Nullable(d.EksklusiPlasentaPrevia)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.EdukasiDiberikan = support.Nullable(d.EdukasiDiberikan)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.KdDokter = support.Nullable(d.KdDokter)
		return nil
	},
	Refs: []Ref[model.SkriningTolac]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SkriningTolac) any { return StrPtr(m.KdDokter) }},
	},
}
