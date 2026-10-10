// Package rekammedis use case form asesmen rekam medis (source/src/rekammedis RM*.java).
// Aturan yang sama di semua form Khanza:
//   - no_rawat harus terdaftar; waktu asesmen tidak boleh sebelum registrasi (kecuali Admin Utama),
//   - petugas pengisi = akun login (kecuali Admin Utama),
//   - ubah/hapus hanya oleh petugas yang bersangkutan dan maksimal 2 x 24 jam (kecuali Admin Utama).
package rekammedis

import (
	"fmt"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

const batasUbah = 48 * time.Hour

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrSebelumRegistrasi  = support.Invalid("waktu asesmen tidak boleh sebelum waktu registrasi")
	ErrTanpaPegawai       = support.Forbidden("akun belum dihubungkan ke pegawai (users.kd_pegawai)")
	ErrBukanPetugas       = support.Forbidden("hanya bisa diisi/diubah/dihapus oleh petugas yang bersangkutan")
	ErrKadaluarsa         = support.Forbidden("perubahan data / penghapusan data tidak boleh lebih dari 2 x 24 jam")
)

// Actor pengguna yang menjalankan aksi.
type Actor struct {
	KodePegawai string
	SuperAdmin  bool
}

// Ref referensi kolom ke master (foreign key) dengan pesan yang jelas.
type Ref[M any] struct {
	Column string
	Table  string
	RefCol string
	Label  string
	Value  func(m *M) any // nil = kosong / tidak dicek
}

// Form deskripsi satu form asesmen (dibangkitkan per tabel).
type Form[M any, D any] struct {
	Slug    string
	Label   string
	Spec    repo.Spec
	SetKey  func(m *M, key repo.Key) error // isi kolom kunci dari input; waktu kosong = sekarang
	KeyOf   func(m *M) repo.Key
	Waktu   func(m *M) time.Time // waktu asesmen; zero bila form tidak punya kolom waktu
	Petugas func(m *M) []string  // kolom petugas pengisi
	Fill    func(m *M, d D) error
	Refs    []Ref[M]
}

// Record satu asesmen beserta daftar kode detail (masalah/rencana keperawatan).
type Record[M any] struct {
	Asesmen *M                  `json:"asesmen"`
	Detail  map[string][]string `json:"detail,omitempty"`
}

type Action[M any, D any] struct {
	form *Form[M, D]
	repo repo.Store[M]
	now  func() time.Time
}

func NewAction[M any, D any](form *Form[M, D], store repo.Store[M]) *Action[M, D] {
	return &Action[M, D]{form: form, repo: store, now: time.Now}
}

func (a *Action[M, D]) Form() *Form[M, D] {
	return a.form
}

func (a *Action[M, D]) List(f repo.Filter, page, limit int) ([]M, int64, error) {
	return a.repo.Paginate(f, page, limit)
}

func (a *Action[M, D]) Detail(key repo.Key) (*Record[M], error) {
	m, err := a.find(key)
	if err != nil {
		return nil, err
	}
	return a.record(m)
}

func (a *Action[M, D]) Create(key repo.Key, d D, details map[string][]string, actor Actor) (*Record[M], error) {
	m := new(M)
	if err := a.form.SetKey(m, key); err != nil {
		return nil, err
	}
	if err := a.form.Fill(m, d); err != nil {
		return nil, err
	}
	k, err := a.repo.Konteks(a.form.KeyOf(m)["no_rawat"])
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	if !actor.SuperAdmin {
		if w := a.form.Waktu(m); !w.IsZero() && k.SebelumRegistrasi(w) {
			return nil, ErrSebelumRegistrasi
		}
		if err := a.pemilik(m, actor); err != nil {
			return nil, err
		}
	}
	exists, err := a.repo.Exists(a.form.KeyOf(m))
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, support.Conflict(a.form.Label + " untuk rawat & waktu tersebut sudah ada")
	}
	if err := a.cekReferensi(m, details); err != nil {
		return nil, err
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		if err := a.repo.Create(tx, m); err != nil {
			return err
		}
		return a.simpanDetail(tx, m, details)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(a.form.KeyOf(m))
}

// Update mengganti isi asesmen (PUT); kunci tidak berubah, detail diganti seluruhnya.
func (a *Action[M, D]) Update(key repo.Key, d D, details map[string][]string, actor Actor) (*Record[M], error) {
	m, err := a.editable(key, actor)
	if err != nil {
		return nil, err
	}
	if err := a.form.Fill(m, d); err != nil {
		return nil, err
	}
	if !actor.SuperAdmin {
		if err := a.pemilik(m, actor); err != nil {
			return nil, err
		}
	}
	if err := a.cekReferensi(m, details); err != nil {
		return nil, err
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		if err := a.repo.Save(tx, m); err != nil {
			return err
		}
		return a.simpanDetail(tx, m, details)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(a.form.KeyOf(m))
}

func (a *Action[M, D]) Delete(key repo.Key, actor Actor) error {
	m, err := a.editable(key, actor)
	if err != nil {
		return err
	}
	return support.Transaction(func(tx contractsorm.Query) error {
		return a.repo.Delete(tx, a.form.KeyOf(m))
	})
}

func (a *Action[M, D]) find(key repo.Key) (*M, error) {
	for k, v := range key {
		key[k] = strings.TrimSpace(v)
	}
	m, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, support.NotFound("data " + a.form.Label + " tidak ditemukan")
	}
	return m, nil
}

func (a *Action[M, D]) record(m *M) (*Record[M], error) {
	r := &Record[M]{Asesmen: m}
	if len(a.form.Spec.Details) > 0 {
		d, err := a.repo.Details(a.form.KeyOf(m)["no_rawat"])
		if err != nil {
			return nil, err
		}
		r.Detail = d
	}
	return r, nil
}

// editable data lama boleh diubah/dihapus oleh actor (pemilik & 2 x 24 jam, kecuali Admin Utama).
func (a *Action[M, D]) editable(key repo.Key, actor Actor) (*M, error) {
	m, err := a.find(key)
	if err != nil {
		return nil, err
	}
	if actor.SuperAdmin {
		return m, nil
	}
	if err := a.pemilik(m, actor); err != nil {
		return nil, err
	}
	if w := a.form.Waktu(m); !w.IsZero() && a.now().Sub(w) > batasUbah {
		return nil, ErrKadaluarsa
	}
	return m, nil
}

// pemilik akun login harus salah satu petugas pengisi form (akses.getkode() di Khanza).
func (a *Action[M, D]) pemilik(m *M, actor Actor) error {
	petugas := a.form.Petugas(m)
	if len(petugas) == 0 {
		return nil
	}
	if actor.KodePegawai == "" {
		return ErrTanpaPegawai
	}
	for _, p := range petugas {
		if p == actor.KodePegawai {
			return nil
		}
	}
	return ErrBukanPetugas
}

func (a *Action[M, D]) cekReferensi(m *M, details map[string][]string) error {
	for _, r := range a.form.Refs {
		v := r.Value(m)
		if v == nil {
			continue
		}
		ok, err := a.repo.RefExists(r.Table, r.RefCol, v)
		if err != nil {
			return err
		}
		if !ok {
			return support.NotFound(fmt.Sprintf("data %s (%v) tidak ditemukan", r.Label, v))
		}
	}
	for _, d := range a.form.Spec.Details {
		for _, kode := range details[d.Name] {
			ok, err := a.repo.RefExists(d.RefTable, d.RefColumn, kode)
			if err != nil {
				return err
			}
			if !ok {
				return support.NotFound(fmt.Sprintf("kode %s %s tidak ditemukan", strings.ReplaceAll(strings.TrimPrefix(d.Name, "kode_"), "_", " "), kode))
			}
		}
	}
	return nil
}

func (a *Action[M, D]) simpanDetail(tx contractsorm.Query, m *M, details map[string][]string) error {
	if len(a.form.Spec.Details) == 0 {
		return nil
	}
	clean := map[string][]string{}
	for name, codes := range details {
		seen := map[string]bool{}
		for _, c := range codes {
			if c = strings.TrimSpace(c); c != "" && !seen[c] {
				seen[c] = true
				clean[name] = append(clean[name], c)
			}
		}
	}
	return a.repo.ReplaceDetails(tx, a.form.KeyOf(m)["no_rawat"], clean)
}

// ---- helper untuk kode form yang dibangkitkan ----

// Str nilai referensi string; kosong = tidak dicek.
func Str(s string) any {
	if s = strings.TrimSpace(s); s == "" || s == "-" {
		return nil
	}
	return s
}

// StrPtr nilai referensi *string; nil/kosong = tidak dicek.
func StrPtr(s *string) any {
	if s == nil {
		return nil
	}
	return Str(*s)
}

// Tm waktu dari *time.Time (zero bila nil).
func Tm(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

// TglJam menggabungkan tanggal (*time.Time) dan jam "HH:MM:SS".
func TglJam(tgl *time.Time, jam string) time.Time {
	if tgl == nil {
		return time.Time{}
	}
	w, err := time.ParseInLocation(support.DateTimeLayout, tgl.Format(support.DateLayout)+" "+jam, time.Local)
	if err != nil {
		return *tgl
	}
	return w
}

// KeyDateTime kunci datetime dari input (kosong = sekarang).
func KeyDateTime(s string) (*time.Time, error) {
	t, err := support.ParseDateTime(s)
	if err != nil || t != nil {
		return t, err
	}
	now := time.Now().Truncate(time.Second)
	return &now, nil
}

// KeyDate kunci tanggal dari input (kosong = hari ini).
func KeyDate(s string) (*time.Time, error) {
	t, err := support.DateOrToday(s)
	return &t, err
}

// KeyTime kunci jam dari input (kosong = sekarang).
func KeyTime(s string) (string, error) {
	return support.TimeOrNow(s)
}

// FmtDateTime / FmtDate teks kunci dari waktu model.
func FmtDateTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(support.DateTimeLayout)
}

func FmtDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(support.DateLayout)
}

// OptTime jam opsional "HH:MM:SS" (kosong = NULL / "").
func OptTime(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", nil
	}
	return s, support.ValidTime(s)
}

func derefStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefAny[T any](p *T) any {
	if p == nil {
		return nil
	}
	return *p
}

// Enum validasi nilai enum yang tidak bisa dinyatakan dengan rule "in" (nilai mengandung koma / kosong).
func Enum(field, value string, allowed ...string) error {
	value = strings.TrimSpace(value)
	for _, a := range allowed {
		if value == a {
			return nil
		}
	}
	if value == "" {
		return nil
	}
	return support.Invalid(fmt.Sprintf("nilai %s tidak valid", field))
}
