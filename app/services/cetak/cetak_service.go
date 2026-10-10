// Package cetak penyusun dokumen cetak: kop instansi, identitas, dan konversi model ke baris label-nilai.
package cetak

import (
	"encoding/base64"
	"fmt"
	"html/template"
	nethttp "net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	model "goravel/app/models/cetak"
	repo "goravel/app/repository/cetak"
	"goravel/app/support"
)

// namaKolom kolom nama untuk tabel referensi yang kodenya diterjemahkan saat cetak.
var namaKolom = map[string]string{
	"dokter": "nm_dokter", "petugas": "nama", "pegawai": "nama", "penyakit": "nm_penyakit", "icd9": "deskripsi_panjang",
	"poliklinik": "nm_poli", "bangsal": "nm_bangsal", "databarang": "nama_brng", "jns_perawatan": "nm_perawatan",
	"jns_perawatan_lab": "nm_perawatan", "jns_perawatan_radiologi": "nm_perawatan", "master_imunisasi": "nama_imunisasi",
	"master_triase_macam_kasus": "macam_kasus", "master_masalah_mpp": "nama_masalah",
}

// NamaKolom kolom nama untuk tabel referensi (termasuk seluruh master masalah/rencana keperawatan); kosong bila tidak dikenal.
func NamaKolom(table string) string {
	switch {
	case strings.HasPrefix(table, "master_masalah_keperawatan"):
		return "nama_masalah"
	case strings.HasPrefix(table, "master_rencana_keperawatan"):
		return "rencana_keperawatan"
	}
	return namaKolom[table]
}

type Service struct {
	repo repo.Repository
	now  func() time.Time
}

func NewService(repo repo.Repository) *Service {
	return &Service{repo: repo, now: time.Now}
}

// Dokumen kerangka dokumen dengan kop instansi & waktu cetak.
func (s *Service) Dokumen(judul, nomor string) (*model.Dokumen, error) {
	inst, logo, err := s.repo.Instansi()
	if err != nil {
		return nil, err
	}
	if len(logo) > 0 {
		inst.Logo = template.URL("data:" + nethttp.DetectContentType(logo) + ";base64," + base64.StdEncoding.EncodeToString(logo))
	}
	return &model.Dokumen{Judul: strings.ToUpper(judul), Nomor: nomor, Instansi: *inst, Dicetak: s.now().Format("02-01-2006 15:04:05")}, nil
}

func (s *Service) IdentitasRawat(noRawat string) ([]model.Baris, error) {
	b, err := s.repo.IdentitasRawat(noRawat)
	if err == nil && b == nil {
		err = support.NotFound("data registrasi tidak ditemukan")
	}
	return b, err
}

func (s *Service) IdentitasPasien(noRkmMedis string) ([]model.Baris, error) {
	b, err := s.repo.IdentitasPasien(noRkmMedis)
	if err == nil && b == nil {
		err = support.NotFound("data pasien tidak ditemukan")
	}
	return b, err
}

// Nama nama dari kode referensi tabel (kode itu sendiri bila tidak ditemukan).
func (s *Service) Nama(table, keyCol string, kode any) string {
	col := NamaKolom(table)
	k := fmt.Sprint(kode)
	if col == "" || k == "" {
		return k
	}
	n, err := s.repo.Nama(table, keyCol, col, kode)
	if err != nil || n == "" {
		return k
	}
	return n
}

// Ref referensi kolom model ke tabel master untuk diterjemahkan.
type Ref struct {
	Column, Table, RefCol string
}

// Baris mengubah struct (tag json = nama kolom) menjadi baris label-nilai; kolom di skip dilewati,
// kolom ber-referensi ditampilkan "kode - nama".
func (s *Service) Baris(m any, skip map[string]bool, refs []Ref) []model.Baris {
	refOf := map[string]Ref{}
	for _, r := range refs {
		refOf[r.Column] = r
	}
	v := reflect.Indirect(reflect.ValueOf(m))
	t := v.Type()
	out := []model.Baris{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		col := strings.Split(f.Tag.Get("json"), ",")[0]
		if col == "" || col == "-" || skip[col] {
			continue
		}
		nilai := Format(v.Field(i).Interface())
		if r, ok := refOf[col]; ok && nilai != "-" {
			if n := s.Nama(r.Table, r.RefCol, nilai); n != nilai {
				nilai = nilai + " - " + n
			}
		}
		out = append(out, model.Baris{Label: Label(col), Nilai: nilai})
	}
	return out
}

// Label "keluhan_utama" -> "Keluhan Utama"; singkatan pendek ("td", "gcs") ditulis kapital.
func Label(col string) string {
	parts := strings.Split(col, "_")
	if len(parts) == 1 && len(col) <= 3 {
		return strings.ToUpper(col)
	}
	for i, p := range parts {
		if p != "" {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

// Format nilai untuk dicetak: nil/kosong = "-", waktu dd-mm-yyyy [HH:MM:SS], angka tanpa nol berlebih.
func Format(x any) string {
	v := reflect.ValueOf(x)
	if !v.IsValid() || (v.Kind() == reflect.Pointer && v.IsNil()) {
		return "-"
	}
	x = reflect.Indirect(v).Interface()
	var s string
	switch t := x.(type) {
	case time.Time:
		if t.IsZero() {
			return "-"
		}
		if t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0 {
			return t.Format("02-01-2006")
		}
		return t.Format("02-01-2006 15:04:05")
	case float64:
		s = strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		s = strings.TrimSpace(t)
	default:
		s = fmt.Sprint(t)
	}
	if s == "" {
		return "-"
	}
	return s
}
