package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianTambahanPerilakuKekerasanData isian penilaian tambahan perilaku kekerasan.
type PenilaianTambahanPerilakuKekerasanData struct {
	Tanggal                            string `form:"tanggal" json:"tanggal"`
	Nip                                string `form:"nip" json:"nip"`
	StatikInsidenKekerasanBaruIni      string `form:"statik_insiden_kekerasan_baru_ini" json:"statik_insiden_kekerasan_baru_ini"`
	StatikSkorinsidenKekerasanBaruIni  *int   `form:"statik_skorinsiden_kekerasan_baru_ini" json:"statik_skorinsiden_kekerasan_baru_ini"`
	StatikRiwayatPenggunaanSenjata     string `form:"statik_riwayat_penggunaan_senjata" json:"statik_riwayat_penggunaan_senjata"`
	StatikSkorriwayatPenggunaanSenjata *int   `form:"statik_skorriwayat_penggunaan_senjata" json:"statik_skorriwayat_penggunaan_senjata"`
	StatikLakiLaki                     string `form:"statik_laki_laki" json:"statik_laki_laki"`
	StatikSkorlakiLaki                 *int   `form:"statik_skorlaki_laki" json:"statik_skorlaki_laki"`
	StatikUsiaDibawah35                string `form:"statik_usia_dibawah_35" json:"statik_usia_dibawah_35"`
	StatikSkorusiaDibawah35            *int   `form:"statik_skorusia_dibawah_35" json:"statik_skorusia_dibawah_35"`
	StatikRiwayatKriminal              string `form:"statik_riwayat_kriminal" json:"statik_riwayat_kriminal"`
	StatikSkorriwayatKriminal          *int   `form:"statik_skorriwayat_kriminal" json:"statik_skorriwayat_kriminal"`
	StatikIdeKekerasan                 string `form:"statik_ide_kekerasan" json:"statik_ide_kekerasan"`
	StatikSkorideKekerasan             *int   `form:"statik_skoride_kekerasan" json:"statik_skoride_kekerasan"`
	StatikKekerasanAnakAnak            string `form:"statik_kekerasan_anak_anak" json:"statik_kekerasan_anak_anak"`
	StatikSkorkekerasanAnakAnak        *int   `form:"statik_skorkekerasan_anak_anak" json:"statik_skorkekerasan_anak_anak"`
	StatikPeranDalamHidup              string `form:"statik_peran_dalam_hidup" json:"statik_peran_dalam_hidup"`
	StatikSkorperanDalamHidup          *int   `form:"statik_skorperan_dalam_hidup" json:"statik_skorperan_dalam_hidup"`
	StatikPenggunaanNapza              string `form:"statik_penggunaan_napza" json:"statik_penggunaan_napza"`
	StatikSkorpenggunaanNapza          *int   `form:"statik_skorpenggunaan_napza" json:"statik_skorpenggunaan_napza"`
	StatikSkortotal                    *int   `form:"statik_skortotal" json:"statik_skortotal"`
	DinamisIdeMelukaiOrangLain         string `form:"dinamis_ide_melukai_orang_lain" json:"dinamis_ide_melukai_orang_lain"`
	DinamisSkorideMelukaiOrangLain     *int   `form:"dinamis_skoride_melukai_orang_lain" json:"dinamis_skoride_melukai_orang_lain"`
	DinamisAksesKekerasan              string `form:"dinamis_akses_kekerasan" json:"dinamis_akses_kekerasan"`
	DinamisSkoraksesKekerasan          *int   `form:"dinamis_skorakses_kekerasan" json:"dinamis_skorakses_kekerasan"`
	DinamisIdeParanoid                 string `form:"dinamis_ide_paranoid" json:"dinamis_ide_paranoid"`
	DinamisSkorideParanoid             *int   `form:"dinamis_skoride_paranoid" json:"dinamis_skoride_paranoid"`
	DinamisPerintahHalusinasi          string `form:"dinamis_perintah_halusinasi" json:"dinamis_perintah_halusinasi"`
	DinamisSkorperintahHalusinasi      *int   `form:"dinamis_skorperintah_halusinasi" json:"dinamis_skorperintah_halusinasi"`
	DinamisFrustasiAgitasi             string `form:"dinamis_frustasi_agitasi" json:"dinamis_frustasi_agitasi"`
	DinamisSkorfrustasiAgitasi         *int   `form:"dinamis_skorfrustasi_agitasi" json:"dinamis_skorfrustasi_agitasi"`
	DinamisKesenanganKekerasan         string `form:"dinamis_kesenangan_kekerasan" json:"dinamis_kesenangan_kekerasan"`
	DinamisSkorkesenanganKekerasan     *int   `form:"dinamis_skorkesenangan_kekerasan" json:"dinamis_skorkesenangan_kekerasan"`
	DinamisSeksualTidakWajar           string `form:"dinamis_seksual_tidak_wajar" json:"dinamis_seksual_tidak_wajar"`
	DinamisSkorseksualTidakWajar       *int   `form:"dinamis_skorseksual_tidak_wajar" json:"dinamis_skorseksual_tidak_wajar"`
	DinamisHilangnyaKontrolDiri        string `form:"dinamis_hilangnya_kontrol_diri" json:"dinamis_hilangnya_kontrol_diri"`
	DinamisSkorhilangnyaKontrolDiri    *int   `form:"dinamis_skorhilangnya_kontrol_diri" json:"dinamis_skorhilangnya_kontrol_diri"`
	DinamisPengguaanNapza              string `form:"dinamis_pengguaan_napza" json:"dinamis_pengguaan_napza"`
	DinamisSkorpengguaanNapza          *int   `form:"dinamis_skorpengguaan_napza" json:"dinamis_skorpengguaan_napza"`
	DinamisSkortotal                   *int   `form:"dinamis_skortotal" json:"dinamis_skortotal"`
	FaktorFaktorPencegahan             string `form:"faktor_faktor_pencegahan" json:"faktor_faktor_pencegahan"`
	TotalSkor                          *int   `form:"total_skor" json:"total_skor"`
	LevelSkor                          string `form:"level_skor" json:"level_skor"`
}

func penilaianTambahanPerilakuKekerasanRules() map[string]any {
	rules := map[string]any{
		"tanggal":                               "required|date",
		"nip":                                   "required|string|max_len:20",
		"statik_insiden_kekerasan_baru_ini":     "in:Ya,Tidak,Tidak Tahu",
		"statik_skorinsiden_kekerasan_baru_ini": "int",
		"statik_riwayat_penggunaan_senjata":     "in:Ya,Tidak,Tidak Tahu",
		"statik_skorriwayat_penggunaan_senjata": "int",
		"statik_laki_laki":                      "in:Ya,Tidak,Tidak Tahu",
		"statik_skorlaki_laki":                  "int",
		"statik_usia_dibawah_35":                "in:Ya,Tidak,Tidak Tahu",
		"statik_skorusia_dibawah_35":            "int",
		"statik_riwayat_kriminal":               "in:Ya,Tidak,Tidak Tahu",
		"statik_skorriwayat_kriminal":           "int",
		"statik_ide_kekerasan":                  "in:Ya,Tidak,Tidak Tahu",
		"statik_skoride_kekerasan":              "int",
		"statik_kekerasan_anak_anak":            "in:Ya,Tidak,Tidak Tahu",
		"statik_skorkekerasan_anak_anak":        "int",
		"statik_peran_dalam_hidup":              "in:Ya,Tidak,Tidak Tahu",
		"statik_skorperan_dalam_hidup":          "int",
		"statik_penggunaan_napza":               "in:Ya,Tidak,Tidak Tahu",
		"statik_skorpenggunaan_napza":           "int",
		"statik_skortotal":                      "int",
		"dinamis_ide_melukai_orang_lain":        "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skoride_melukai_orang_lain":    "int",
		"dinamis_akses_kekerasan":               "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorakses_kekerasan":           "int",
		"dinamis_ide_paranoid":                  "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skoride_paranoid":              "int",
		"dinamis_perintah_halusinasi":           "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorperintah_halusinasi":       "int",
		"dinamis_frustasi_agitasi":              "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorfrustasi_agitasi":          "int",
		"dinamis_kesenangan_kekerasan":          "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorkesenangan_kekerasan":      "int",
		"dinamis_seksual_tidak_wajar":           "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorseksual_tidak_wajar":       "int",
		"dinamis_hilangnya_kontrol_diri":        "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorhilangnya_kontrol_diri":    "int",
		"dinamis_pengguaan_napza":               "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorpengguaan_napza":           "int",
		"dinamis_skortotal":                     "int",
		"faktor_faktor_pencegahan":              "string|max_len:500",
		"total_skor":                            "int",
		"level_skor":                            "in:Rendah(<7),Sedang(7-14),Tinggi(>14)",
	}
	return rules
}

// PenilaianTambahanPerilakuKekerasanStore simpan penilaian tambahan perilaku kekerasan; kolom waktu kunci kosong = sekarang.
type PenilaianTambahanPerilakuKekerasanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianTambahanPerilakuKekerasanData
}

func (r *PenilaianTambahanPerilakuKekerasanStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanPerilakuKekerasanStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianTambahanPerilakuKekerasanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianTambahanPerilakuKekerasanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianTambahanPerilakuKekerasanStore) Payload() PenilaianTambahanPerilakuKekerasanData {
	return r.PenilaianTambahanPerilakuKekerasanData
}

func (r *PenilaianTambahanPerilakuKekerasanStore) DetailValues() map[string][]string { return nil }

// PenilaianTambahanPerilakuKekerasanUpdate ubah penilaian tambahan perilaku kekerasan (PUT); kunci lewat query string.
type PenilaianTambahanPerilakuKekerasanUpdate struct {
	PenilaianTambahanPerilakuKekerasanData
}

func (r *PenilaianTambahanPerilakuKekerasanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanPerilakuKekerasanUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianTambahanPerilakuKekerasanRules()
}

func (r *PenilaianTambahanPerilakuKekerasanUpdate) Payload() PenilaianTambahanPerilakuKekerasanData {
	return r.PenilaianTambahanPerilakuKekerasanData
}

func (r *PenilaianTambahanPerilakuKekerasanUpdate) DetailValues() map[string][]string { return nil }
