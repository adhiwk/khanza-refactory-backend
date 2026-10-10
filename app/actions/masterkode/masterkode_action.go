// Package masterkode use case master kode rekam medis (MasterMasalahKeperawatan*, MasterRencanaKeperawatan*,
// MasterMasalahMPP, MasterTriase*, MasterImunisasi): kode 3 digit otomatis bila kosong, hapus permanen.
package masterkode

import (
	"strconv"
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/masterkode"
	repo "goravel/app/repository/masterkode"
	"goravel/app/support"
)

func masalah(slug, label, table string) *repo.Spec {
	return &repo.Spec{Slug: slug, Label: label, Table: table, KodeCol: "kode_masalah", NamaCol: "nama_masalah", NamaMax: 100, KodeWidth: 3}
}

func rencana(slug, label, table string, induk *repo.Spec) *repo.Spec {
	return &repo.Spec{Slug: slug, Label: label, Table: table, KodeCol: "kode_rencana", NamaCol: "rencana_keperawatan", NamaMax: 1000,
		IndukCol: "kode_masalah", Induk: induk, KodeWidth: 3}
}

var (
	masalahUmum       = masalah("masalah-keperawatan", "masalah keperawatan", "master_masalah_keperawatan")
	masalahAnak       = masalah("masalah-keperawatan-anak", "masalah keperawatan anak", "master_masalah_keperawatan_anak")
	masalahGeriatri   = masalah("masalah-keperawatan-geriatri", "masalah keperawatan geriatri", "master_masalah_keperawatan_geriatri")
	masalahGigi       = masalah("masalah-keperawatan-gigi", "masalah keperawatan gigi", "master_masalah_keperawatan_gigi")
	masalahIGD        = masalah("masalah-keperawatan-igd", "masalah keperawatan IGD", "master_masalah_keperawatan_igd")
	masalahMata       = masalah("masalah-keperawatan-mata", "masalah keperawatan mata", "master_masalah_keperawatan_mata")
	masalahNeonatus   = masalah("masalah-keperawatan-neonatus", "masalah keperawatan neonatus", "master_masalah_keperawatan_neonatus")
	masalahPsikiatri  = masalah("masalah-keperawatan-psikiatri", "masalah keperawatan psikiatri", "master_masalah_keperawatan_psikiatri")
	triasePemeriksaan = &repo.Spec{Slug: "triase-pemeriksaan", Label: "pemeriksaan triase", Table: "master_triase_pemeriksaan",
		KodeCol: "kode_pemeriksaan", NamaCol: "nama_pemeriksaan", NamaMax: 150, KodeWidth: 3}
)

func skala(n string) *repo.Spec {
	return &repo.Spec{Slug: "triase-skala" + n, Label: "skala " + n + " triase", Table: "master_triase_skala" + n,
		KodeCol: "kode_skala" + n, NamaCol: "pengkajian_skala" + n, NamaMax: 150, IndukCol: "kode_pemeriksaan", Induk: triasePemeriksaan, KodeWidth: 3}
}

// Specs seluruh master kode (slug unik dipakai di route).
var Specs = []*repo.Spec{
	masalahUmum, masalahAnak, masalahGeriatri, masalahGigi, masalahIGD, masalahMata, masalahNeonatus, masalahPsikiatri,
	masalah("masalah-mpp", "masalah MPP", "master_masalah_mpp"),
	rencana("rencana-keperawatan", "rencana keperawatan", "master_rencana_keperawatan", masalahUmum),
	rencana("rencana-keperawatan-anak", "rencana keperawatan anak", "master_rencana_keperawatan_anak", masalahAnak),
	rencana("rencana-keperawatan-geriatri", "rencana keperawatan geriatri", "master_rencana_keperawatan_geriatri", masalahGeriatri),
	rencana("rencana-keperawatan-gigi", "rencana keperawatan gigi", "master_rencana_keperawatan_gigi", masalahGigi),
	rencana("rencana-keperawatan-igd", "rencana keperawatan IGD", "master_rencana_keperawatan_igd", masalahIGD),
	rencana("rencana-keperawatan-mata", "rencana keperawatan mata", "master_rencana_keperawatan_mata", masalahMata),
	rencana("rencana-keperawatan-neonatus", "rencana keperawatan neonatus", "master_rencana_keperawatan_neonatus", masalahNeonatus),
	rencana("rencana-keperawatan-psikiatri", "rencana keperawatan psikiatri", "master_rencana_keperawatan_psikiatri", masalahPsikiatri),
	{Slug: "triase-macam-kasus", Label: "macam kasus triase", Table: "master_triase_macam_kasus", KodeCol: "kode_kasus",
		NamaCol: "macam_kasus", NamaMax: 150, KodeWidth: 3},
	triasePemeriksaan,
	skala("1"), skala("2"), skala("3"), skala("4"), skala("5"),
	{Slug: "imunisasi", Label: "imunisasi", Table: "master_imunisasi", KodeCol: "kode_imunisasi", NamaCol: "nama_imunisasi", NamaMax: 100, KodeWidth: 3},
}

var (
	ErrNamaKosong  = support.Invalid("nama wajib diisi")
	ErrIndukKosong = support.Invalid("kode_induk wajib diisi")
)

type Action struct {
	spec *repo.Spec
	repo repo.Repository
}

func NewAction(spec *repo.Spec, repo repo.Repository) *Action {
	return &Action{spec: spec, repo: repo}
}

func (a *Action) Spec() *repo.Spec {
	return a.spec
}

func (a *Action) List(search, kodeInduk string, page, limit int) ([]model.Kode, int64, error) {
	return a.repo.Paginate(a.spec, strings.TrimSpace(search), strings.TrimSpace(kodeInduk), page, limit)
}

func (a *Action) Detail(kode string) (*model.Kode, error) {
	k, err := a.repo.Find(a.spec, strings.TrimSpace(kode))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, support.NotFound("data " + a.spec.Label + " tidak ditemukan")
	}
	return k, nil
}

// Create kode kosong dibuat otomatis (angka terbesar + 1, 3 digit).
func (a *Action) Create(in model.Kode) (*model.Kode, error) {
	k, err := a.clean(in)
	if err != nil {
		return nil, err
	}
	if k.Kode != "" {
		exists, err := a.repo.Exists(a.spec, k.Kode)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, support.Conflict("kode " + a.spec.Label + " sudah digunakan")
		}
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		if k.Kode == "" {
			if k.Kode, err = a.repo.NextKode(tx, a.spec); err != nil {
				return err
			}
		}
		return a.repo.Create(tx, a.spec, k)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(k.Kode)
}

func (a *Action) Update(kode string, in model.Kode) (*model.Kode, error) {
	old, err := a.Detail(kode)
	if err != nil {
		return nil, err
	}
	in.Kode = old.Kode
	k, err := a.clean(in)
	if err != nil {
		return nil, err
	}
	if err := a.repo.Update(a.spec, k); err != nil {
		return nil, err
	}
	return a.Detail(k.Kode)
}

// Delete hapus permanen hanya bila kode tidak dipakai. Foreign key Khanza bersifat ON DELETE CASCADE, sehingga
// hapus langsung (seperti form Java) akan ikut menghapus data asesmen pasien dan master turunannya.
func (a *Action) Delete(kode string) error {
	k, err := a.Detail(kode)
	if err != nil {
		return err
	}
	pemakai, err := a.repo.Pemakai(a.spec, k.Kode)
	if err != nil {
		return err
	}
	if len(pemakai) > 0 {
		return support.Conflict("kode " + a.spec.Label + " masih dipakai pada: " + strings.Join(pemakai, ", "))
	}
	return a.repo.Delete(a.spec, k.Kode)
}

func (a *Action) clean(in model.Kode) (model.Kode, error) {
	k := model.Kode{Kode: strings.TrimSpace(in.Kode), Nama: strings.TrimSpace(in.Nama), KodeInduk: strings.TrimSpace(in.KodeInduk)}
	if k.Nama == "" {
		return k, ErrNamaKosong
	}
	if len(k.Nama) > a.spec.NamaMax {
		return k, support.Invalid("nama maksimal " + strconv.Itoa(a.spec.NamaMax) + " karakter")
	}
	if a.spec.IndukCol == "" {
		k.KodeInduk = ""
		return k, nil
	}
	if k.KodeInduk == "" {
		return k, ErrIndukKosong
	}
	ok, err := a.repo.Exists(a.spec.Induk, k.KodeInduk)
	if err != nil {
		return k, err
	}
	if !ok {
		return k, support.NotFound("data " + a.spec.Induk.Label + " tidak ditemukan")
	}
	return k, nil
}
