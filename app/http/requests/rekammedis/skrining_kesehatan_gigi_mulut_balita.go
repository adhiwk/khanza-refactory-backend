package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKesehatanGigiMulutBalitaData isian skrining kesehatan gigi mulut balita.
type SkriningKesehatanGigiMulutBalitaData struct {
	Tanggal                    string `form:"tanggal" json:"tanggal"`
	PernahPemeriksaanGigimulut string `form:"pernah_pemeriksaan_gigimulut" json:"pernah_pemeriksaan_gigimulut"`
	SudahTumbuhGigi            string `form:"sudah_tumbuh_gigi" json:"sudah_tumbuh_gigi"`
	JumlahGigiTumbuh           string `form:"jumlah_gigi_tumbuh" json:"jumlah_gigi_tumbuh"`
	KondisiKebersihanGigimulut string `form:"kondisi_kebersihan_gigimulut" json:"kondisi_kebersihan_gigimulut"`
	KebiasaanSusuBotol         string `form:"kebiasaan_susu_botol" json:"kebiasaan_susu_botol"`
	MengemilManis              string `form:"mengemil_manis" json:"mengemil_manis"`
	MenyikatGigiSebelumTidur   string `form:"menyikat_gigi_sebelum_tidur" json:"menyikat_gigi_sebelum_tidur"`
	MengemutMakanan            string `form:"mengemut_makanan" json:"mengemut_makanan"`
	LidahKotor                 string `form:"lidah_kotor" json:"lidah_kotor"`
	CelahBibir                 string `form:"celah_bibir" json:"celah_bibir"`
	HasilSkrining              string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan                 string `form:"keterangan" json:"keterangan"`
	Nip                        string `form:"nip" json:"nip"`
}

func skriningKesehatanGigiMulutBalitaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                      "required|date",
		"pernah_pemeriksaan_gigimulut": "in:Pernah,Tidak Pernah",
		"sudah_tumbuh_gigi":            "in:Sudah,Belum",
		"jumlah_gigi_tumbuh":           "in:<20,>20,0",
		"kondisi_kebersihan_gigimulut": "in:Bersih,Kotor",
		"kebiasaan_susu_botol":         "in:Ya,Tidak",
		"mengemil_manis":               "in:Ya,Tidak",
		"menyikat_gigi_sebelum_tidur":  "in:Ya,Tidak",
		"mengemut_makanan":             "in:Ya,Tidak",
		"lidah_kotor":                  "in:Ya,Tidak",
		"celah_bibir":                  "in:Tidak Ada,Ada",
		"hasil_skrining":               "string|max_len:50",
		"keterangan":                   "string|max_len:100",
		"nip":                          "required|string|max_len:20",
	}
	return rules
}

// SkriningKesehatanGigiMulutBalitaStore simpan skrining kesehatan gigi mulut balita; kolom waktu kunci kosong = sekarang.
type SkriningKesehatanGigiMulutBalitaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKesehatanGigiMulutBalitaData
}

func (r *SkriningKesehatanGigiMulutBalitaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutBalitaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKesehatanGigiMulutBalitaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKesehatanGigiMulutBalitaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningKesehatanGigiMulutBalitaStore) Payload() SkriningKesehatanGigiMulutBalitaData {
	return r.SkriningKesehatanGigiMulutBalitaData
}

func (r *SkriningKesehatanGigiMulutBalitaStore) DetailValues() map[string][]string { return nil }

// SkriningKesehatanGigiMulutBalitaUpdate ubah skrining kesehatan gigi mulut balita (PUT); kunci lewat query string.
type SkriningKesehatanGigiMulutBalitaUpdate struct {
	SkriningKesehatanGigiMulutBalitaData
}

func (r *SkriningKesehatanGigiMulutBalitaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutBalitaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKesehatanGigiMulutBalitaRules()
}

func (r *SkriningKesehatanGigiMulutBalitaUpdate) Payload() SkriningKesehatanGigiMulutBalitaData {
	return r.SkriningKesehatanGigiMulutBalitaData
}

func (r *SkriningKesehatanGigiMulutBalitaUpdate) DetailValues() map[string][]string { return nil }
