package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianLanjutanResikoJatuhAnak penilaian lanjutan risiko jatuh anak (RMPenilaianLanjutanRisikoJatuhAnak).
var FormPenilaianLanjutanResikoJatuhAnak = &Form[model.PenilaianLanjutanResikoJatuhAnak, request.PenilaianLanjutanResikoJatuhAnakData]{
	Slug:  "penilaian-lanjutan-resiko-jatuh-anak",
	Label: "penilaian lanjutan risiko jatuh anak",
	Spec: repo.Spec{
		Table:  "penilaian_lanjutan_resiko_jatuh_anak",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianLanjutanResikoJatuhAnak, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianLanjutanResikoJatuhAnak) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianLanjutanResikoJatuhAnak) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianLanjutanResikoJatuhAnak) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianLanjutanResikoJatuhAnak, d request.PenilaianLanjutanResikoJatuhAnakData) error {
		if err := Enum("penilaian_humptydumpty_skala3", d.PenilaianHumptydumptySkala3, "Kelainan Neurologi", "Perubahan Dalam Oksigen(Masalah Saluran Nafas, Dehidrasi, Anemia, Anoreksia / Sakit Kepala, Dll)", "Kelainan Psikis / Perilaku", "Diagnosa Lain"); err != nil {
			return err
		}
		if err := Enum("penilaian_humptydumpty_skala7", d.PenilaianHumptydumptySkala7, "Bermacam-macam Obat Yang Digunakan : Obat Sedative (Kecuali Pasien ICU Yang Menggunakan sedasi dan paralisis), Hipnotik, Barbiturat, Fenoti-Azin, Antidepresan, Laksans/Diuretika,Narkotik", "Salah Satu Dari Pengobatan Di Atas", "Pengobatan Lain"); err != nil {
			return err
		}
		m.PenilaianHumptydumptySkala1 = support.Nullable(d.PenilaianHumptydumptySkala1)
		m.PenilaianHumptydumptyNilai1 = d.PenilaianHumptydumptyNilai1
		m.PenilaianHumptydumptySkala2 = support.Nullable(d.PenilaianHumptydumptySkala2)
		m.PenilaianHumptydumptyNilai2 = d.PenilaianHumptydumptyNilai2
		m.PenilaianHumptydumptySkala3 = support.Nullable(d.PenilaianHumptydumptySkala3)
		m.PenilaianHumptydumptyNilai3 = d.PenilaianHumptydumptyNilai3
		m.PenilaianHumptydumptySkala4 = support.Nullable(d.PenilaianHumptydumptySkala4)
		m.PenilaianHumptydumptyNilai4 = d.PenilaianHumptydumptyNilai4
		m.PenilaianHumptydumptySkala5 = support.Nullable(d.PenilaianHumptydumptySkala5)
		m.PenilaianHumptydumptyNilai5 = d.PenilaianHumptydumptyNilai5
		m.PenilaianHumptydumptySkala6 = support.Nullable(d.PenilaianHumptydumptySkala6)
		m.PenilaianHumptydumptyNilai6 = d.PenilaianHumptydumptyNilai6
		m.PenilaianHumptydumptySkala7 = support.Nullable(d.PenilaianHumptydumptySkala7)
		m.PenilaianHumptydumptyNilai7 = d.PenilaianHumptydumptyNilai7
		m.PenilaianHumptydumptyTotalnilai = d.PenilaianHumptydumptyTotalnilai
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Saran = support.Nullable(d.Saran)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianLanjutanResikoJatuhAnak]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianLanjutanResikoJatuhAnak) any { return Str(m.Nip) }},
	},
}
