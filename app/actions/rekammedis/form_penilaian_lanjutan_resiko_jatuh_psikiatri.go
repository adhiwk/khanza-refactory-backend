package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianLanjutanResikoJatuhPsikiatri penilaian lanjutan risiko jatuh psikiatri (RMPenilaianLanjutanRisikoJatuhPsikiatri).
var FormPenilaianLanjutanResikoJatuhPsikiatri = &Form[model.PenilaianLanjutanResikoJatuhPsikiatri, request.PenilaianLanjutanResikoJatuhPsikiatriData]{
	Slug:  "penilaian-lanjutan-resiko-jatuh-psikiatri",
	Label: "penilaian lanjutan risiko jatuh psikiatri",
	Spec: repo.Spec{
		Table:  "penilaian_lanjutan_resiko_jatuh_psikiatri",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianLanjutanResikoJatuhPsikiatri, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianLanjutanResikoJatuhPsikiatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianLanjutanResikoJatuhPsikiatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianLanjutanResikoJatuhPsikiatri) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianLanjutanResikoJatuhPsikiatri, d request.PenilaianLanjutanResikoJatuhPsikiatriData) error {
		if err := Enum("penilaian_jatuhedmonson_skala5", d.PenilaianJatuhedmonsonSkala5, "-", "Bipolar/Gangguan Schizoaffective", "Penggunaan Obat-obatan Terlarang, Ketergantungan Alkohol", "Gangguan Depresi Mayor", "Demensia/Delirium"); err != nil {
			return err
		}
		if err := Enum("penilaian_jatuhedmonson_skala6", d.PenilaianJatuhedmonsonSkala6, "-", "Mandiri/Keseimbangan Baik/Imobilisasi", "Dengan Alat Bantu (Kursi Roda, Walker, Dll)", "Vertigo/Kelemahan", "Goyah/Membutuhkan Bantuan & Menyadari Kemampuan", "Goyah Tapi Lupa Keterbatasan"); err != nil {
			return err
		}
		m.PenilaianJatuhedmonsonSkala1 = support.Nullable(d.PenilaianJatuhedmonsonSkala1)
		m.PenilaianJatuhedmonsonNilai1 = d.PenilaianJatuhedmonsonNilai1
		m.PenilaianJatuhedmonsonSkala2 = support.Nullable(d.PenilaianJatuhedmonsonSkala2)
		m.PenilaianJatuhedmonsonNilai2 = d.PenilaianJatuhedmonsonNilai2
		m.PenilaianJatuhedmonsonSkala3 = support.Nullable(d.PenilaianJatuhedmonsonSkala3)
		m.PenilaianJatuhedmonsonNilai3 = d.PenilaianJatuhedmonsonNilai3
		m.PenilaianJatuhedmonsonSkala4 = support.Nullable(d.PenilaianJatuhedmonsonSkala4)
		m.PenilaianJatuhedmonsonNilai4 = d.PenilaianJatuhedmonsonNilai4
		m.PenilaianJatuhedmonsonSkala5 = support.Nullable(d.PenilaianJatuhedmonsonSkala5)
		m.PenilaianJatuhedmonsonNilai5 = d.PenilaianJatuhedmonsonNilai5
		m.PenilaianJatuhedmonsonSkala6 = support.Nullable(d.PenilaianJatuhedmonsonSkala6)
		m.PenilaianJatuhedmonsonNilai6 = d.PenilaianJatuhedmonsonNilai6
		m.PenilaianJatuhedmonsonTotalnilai = d.PenilaianJatuhedmonsonTotalnilai
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Saran = support.Nullable(d.Saran)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianLanjutanResikoJatuhPsikiatri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianLanjutanResikoJatuhPsikiatri) any { return Str(m.Nip) }},
	},
}
