package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianTambahanGeriatri penilaian tambahan geriatri (RMPenilaianTambahanGeriatri).
var FormPenilaianTambahanGeriatri = &Form[model.PenilaianTambahanGeriatri, request.PenilaianTambahanGeriatriData]{
	Slug:  "penilaian-tambahan-geriatri",
	Label: "penilaian tambahan geriatri",
	Spec: repo.Spec{
		Table:  "penilaian_tambahan_geriatri",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.PenilaianTambahanGeriatri, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianTambahanGeriatri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianTambahanGeriatri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianTambahanGeriatri) []string { return []string{m.Nik} },
	Fill: func(m *model.PenilaianTambahanGeriatri, d request.PenilaianTambahanGeriatriData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("kualitas_hidup_perawatan_diri", d.KualitasHidupPerawatanDiri, "Tidak Mempunyai Kesulitan Dalam Perawatan Diri Sendiri", "Mengalami Kesulitan Untuk Membasuh Badan, Mandi Atau Berpakaian", "Tidak Mampu Membasuh Badan, Mandi/Berpakaian Sendiri"); err != nil {
			return err
		}
		if err := Enum("skala_nyeri", d.SkalaNyeri, "0 - Tidak Nyeri", "2 - Dapat Ditoleransi(Aktifitas Tidak Tergangu)", "4 - Dapat Ditoleransi(Beberapa Aktifitas Sedikit Terganggu)", "5 - Tidak Dapat Ditoleransi(Masih Bisa Menggunakan Telp, Menonton TV/Membaca)", "6 - Tidak Dapat Ditoleransi(Tidak Bisa Menggunakan Telp, Menonton TV/Membaca)", "8 - Tidak Dapat Ditoleransi(Masih Bisa Berbicara Kerenya Nyeri)", "10 - Tidak Dapat Ditoleransi(Tidak Bisa Berbicara Kerenya Nyeri)"); err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nik = strings.TrimSpace(d.Nik)
		m.AsalMasuk = support.Nullable(d.AsalMasuk)
		m.KondisiMasuk = support.Nullable(d.KondisiMasuk)
		m.KeteranganKondisiMasuk = support.Nullable(d.KeteranganKondisiMasuk)
		m.Anamnesis = strings.TrimSpace(d.Anamnesis)
		m.DiagnosaMedis = support.Nullable(d.DiagnosaMedis)
		m.RiwayatImmunoTelinga = support.Nullable(d.RiwayatImmunoTelinga)
		m.RiwayatImmunoSinus = support.Nullable(d.RiwayatImmunoSinus)
		m.RiwayatImmunoAntibiotik = support.Nullable(d.RiwayatImmunoAntibiotik)
		m.RiwayatImmunoPneumonia = support.Nullable(d.RiwayatImmunoPneumonia)
		m.RiwayatImmunoAbses = support.Nullable(d.RiwayatImmunoAbses)
		m.RiwayatImmunoSariawan = support.Nullable(d.RiwayatImmunoSariawan)
		m.RiwayatImmunoMemerlukanAntibiotik = support.Nullable(d.RiwayatImmunoMemerlukanAntibiotik)
		m.RiwayatImmunoInfeksiDalam = support.Nullable(d.RiwayatImmunoInfeksiDalam)
		m.RiwayatImmunoImmunodefisiensiPrimer = support.Nullable(d.RiwayatImmunoImmunodefisiensiPrimer)
		m.RiwayatImmunoJenisKangker = support.Nullable(d.RiwayatImmunoJenisKangker)
		m.RiwayatImmunoInfeksiOportunistik = support.Nullable(d.RiwayatImmunoInfeksiOportunistik)
		m.PolaAktifitasTidur = support.Nullable(d.PolaAktifitasTidur)
		m.KeteranganPolaAktifitasTidur = support.Nullable(d.KeteranganPolaAktifitasTidur)
		m.PolaAktifitasObatTidur = support.Nullable(d.PolaAktifitasObatTidur)
		m.KeteranganPolaAktifitasObatTidur = support.Nullable(d.KeteranganPolaAktifitasObatTidur)
		m.PolaAktifitasOlahraga = support.Nullable(d.PolaAktifitasOlahraga)
		m.KeteranganPolaAktifitasOlahraga = support.Nullable(d.KeteranganPolaAktifitasOlahraga)
		m.KualitasHidupMobilitas = support.Nullable(d.KualitasHidupMobilitas)
		m.KualitasHidupPerawatanDiri = support.Nullable(d.KualitasHidupPerawatanDiri)
		m.KualitasHidupAktifitasSeharihari = support.Nullable(d.KualitasHidupAktifitasSeharihari)
		m.KualitasHidupRasaNyeri = support.Nullable(d.KualitasHidupRasaNyeri)
		m.SkalaNyeri = support.Nullable(d.SkalaNyeri)
		return nil
	},
	Refs: []Ref[model.PenilaianTambahanGeriatri]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.PenilaianTambahanGeriatri) any { return Str(m.Nik) }},
	},
}
