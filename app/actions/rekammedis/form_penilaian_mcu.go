package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianMcu MCU (RMMCU).
var FormPenilaianMcu = &Form[model.PenilaianMcu, request.PenilaianMcuData]{
	Slug:  "penilaian-mcu",
	Label: "MCU",
	Spec: repo.Spec{
		Table:  "penilaian_mcu",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter"},
	},
	SetKey: func(m *model.PenilaianMcu, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.PenilaianMcu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.PenilaianMcu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianMcu) []string { return []string{m.KdDokter} },
	Fill: func(m *model.PenilaianMcu, d request.PenilaianMcuData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Informasi = strings.TrimSpace(d.Informasi)
		m.Rps = strings.TrimSpace(d.Rps)
		m.Rpk = strings.TrimSpace(d.Rpk)
		m.Rpd = strings.TrimSpace(d.Rpd)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Keadaan = strings.TrimSpace(d.Keadaan)
		m.Kesadaran = strings.TrimSpace(d.Kesadaran)
		m.Td = strings.TrimSpace(d.Td)
		m.Nadi = strings.TrimSpace(d.Nadi)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Tb = strings.TrimSpace(d.Tb)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Bmi = strings.TrimSpace(d.Bmi)
		m.KasifikasiBmi = strings.TrimSpace(d.KasifikasiBmi)
		m.LingkarPinggang = strings.TrimSpace(d.LingkarPinggang)
		m.RisikoLingkarPinggang = strings.TrimSpace(d.RisikoLingkarPinggang)
		m.Submandibula = strings.TrimSpace(d.Submandibula)
		m.Axilla = strings.TrimSpace(d.Axilla)
		m.Supraklavikula = strings.TrimSpace(d.Supraklavikula)
		m.Leher = strings.TrimSpace(d.Leher)
		m.Inguinal = strings.TrimSpace(d.Inguinal)
		m.Oedema = strings.TrimSpace(d.Oedema)
		m.SinusFrontalis = strings.TrimSpace(d.SinusFrontalis)
		m.SinusMaxilaris = strings.TrimSpace(d.SinusMaxilaris)
		m.Rambut = strings.TrimSpace(d.Rambut)
		m.Palpebra = strings.TrimSpace(d.Palpebra)
		m.Sklera = strings.TrimSpace(d.Sklera)
		m.Cornea = strings.TrimSpace(d.Cornea)
		m.ButaWarna = strings.TrimSpace(d.ButaWarna)
		m.Konjungtiva = strings.TrimSpace(d.Konjungtiva)
		m.Lensa = strings.TrimSpace(d.Lensa)
		m.Pupil = strings.TrimSpace(d.Pupil)
		m.MenggunakanKacamata = strings.TrimSpace(d.MenggunakanKacamata)
		m.Visus = strings.TrimSpace(d.Visus)
		m.LuasLapangPandang = strings.TrimSpace(d.LuasLapangPandang)
		m.KeteranganLuasLapangPandang = strings.TrimSpace(d.KeteranganLuasLapangPandang)
		m.LubangTelinga = strings.TrimSpace(d.LubangTelinga)
		m.DaunTelinga = strings.TrimSpace(d.DaunTelinga)
		m.SelaputPendengaran = strings.TrimSpace(d.SelaputPendengaran)
		m.ProcMastoideus = strings.TrimSpace(d.ProcMastoideus)
		m.SeptumNasi = strings.TrimSpace(d.SeptumNasi)
		m.LubangHidung = strings.TrimSpace(d.LubangHidung)
		m.Sinus = strings.TrimSpace(d.Sinus)
		m.Bibir = strings.TrimSpace(d.Bibir)
		m.Gusi = strings.TrimSpace(d.Gusi)
		m.Gigi = strings.TrimSpace(d.Gigi)
		m.Caries = strings.TrimSpace(d.Caries)
		m.Lidah = strings.TrimSpace(d.Lidah)
		m.Faring = strings.TrimSpace(d.Faring)
		m.Tonsil = strings.TrimSpace(d.Tonsil)
		m.KelenjarLimfe = strings.TrimSpace(d.KelenjarLimfe)
		m.KelenjarGondok = strings.TrimSpace(d.KelenjarGondok)
		m.GerakanDada = strings.TrimSpace(d.GerakanDada)
		m.VocalFemitus = strings.TrimSpace(d.VocalFemitus)
		m.PerkusiDada = strings.TrimSpace(d.PerkusiDada)
		m.BunyiNapas = strings.TrimSpace(d.BunyiNapas)
		m.BunyiTambahan = strings.TrimSpace(d.BunyiTambahan)
		m.IctusCordis = strings.TrimSpace(d.IctusCordis)
		m.BunyiJantung = strings.TrimSpace(d.BunyiJantung)
		m.Batas = strings.TrimSpace(d.Batas)
		m.Mamae = strings.TrimSpace(d.Mamae)
		m.KeteranganMamae = strings.TrimSpace(d.KeteranganMamae)
		m.Inspeksi = strings.TrimSpace(d.Inspeksi)
		m.Palpasi = strings.TrimSpace(d.Palpasi)
		m.Hepar = strings.TrimSpace(d.Hepar)
		m.PerkusiAbdomen = strings.TrimSpace(d.PerkusiAbdomen)
		m.Auskultasi = strings.TrimSpace(d.Auskultasi)
		m.Limpa = strings.TrimSpace(d.Limpa)
		m.Costovertebral = strings.TrimSpace(d.Costovertebral)
		m.Scoliosis = strings.TrimSpace(d.Scoliosis)
		m.KondisiKulit = strings.TrimSpace(d.KondisiKulit)
		m.PenyakitKulit = strings.TrimSpace(d.PenyakitKulit)
		m.EkstrimitasAtas = strings.TrimSpace(d.EkstrimitasAtas)
		m.EkstrimitasAtasKet = strings.TrimSpace(d.EkstrimitasAtasKet)
		m.EkstrimitasBawah = strings.TrimSpace(d.EkstrimitasBawah)
		m.EkstrimitasBawahKet = strings.TrimSpace(d.EkstrimitasBawahKet)
		m.AreaGenitalia = strings.TrimSpace(d.AreaGenitalia)
		m.KeteranganAreaGenitalia = strings.TrimSpace(d.KeteranganAreaGenitalia)
		m.AnusPerianal = strings.TrimSpace(d.AnusPerianal)
		m.KeteranganAnusPerianal = strings.TrimSpace(d.KeteranganAnusPerianal)
		m.Laborat = strings.TrimSpace(d.Laborat)
		m.Radiologi = strings.TrimSpace(d.Radiologi)
		m.Ekg = strings.TrimSpace(d.Ekg)
		m.Spirometri = strings.TrimSpace(d.Spirometri)
		m.Audiometri = strings.TrimSpace(d.Audiometri)
		m.Treadmill = strings.TrimSpace(d.Treadmill)
		m.RombergTest = strings.TrimSpace(d.RombergTest)
		m.BackStrength = strings.TrimSpace(d.BackStrength)
		m.AbiTanganKanan = strings.TrimSpace(d.AbiTanganKanan)
		m.AbiTanganKiri = strings.TrimSpace(d.AbiTanganKiri)
		m.AbiKakiKanan = strings.TrimSpace(d.AbiKakiKanan)
		m.AbiKakiKiri = strings.TrimSpace(d.AbiKakiKiri)
		m.Lainlain = strings.TrimSpace(d.Lainlain)
		m.Merokok = strings.TrimSpace(d.Merokok)
		m.Alkohol = strings.TrimSpace(d.Alkohol)
		m.Kesimpulan = strings.TrimSpace(d.Kesimpulan)
		m.Anjuran = strings.TrimSpace(d.Anjuran)
		return nil
	},
	Refs: []Ref[model.PenilaianMcu]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.PenilaianMcu) any { return Str(m.KdDokter) }},
	},
}
