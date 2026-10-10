package rekammedis

import (
	"strings"

	icdmodel "goravel/app/models/icd"
	kamarmodel "goravel/app/models/kamarinap"
	obatmodel "goravel/app/models/pemberianobat"
	"goravel/app/models/perawatan"
	icdrepo "goravel/app/repository/icd"
	kamarrepo "goravel/app/repository/kamarinap"
	obatrepo "goravel/app/repository/pemberianobat"
	pemeriksaanrepo "goravel/app/repository/pemeriksaan"
	repo "goravel/app/repository/rekammedis"
	tindakanrepo "goravel/app/repository/tindakan"
	"goravel/app/support"
)

var ErrPasienNotFound = support.NotFound("data pasien tidak ditemukan")

// formLain form rekam medis di luar engine asesmen yang ikut ditandai pada riwayat.
var formLain = []FormInfo{{Slug: "triase-igd", Label: "triase IGD", Table: "data_triase_igd"}}

// RiwayatKunjungan isi rekam medis satu kunjungan (RMRiwayatPerawatan).
type RiwayatKunjungan struct {
	repo.Kunjungan
	Diagnosa    []icdmodel.Kode         `json:"diagnosa"`
	Prosedur    []icdmodel.Kode         `json:"prosedur"`
	Pemeriksaan []perawatan.Pemeriksaan `json:"pemeriksaan"`
	Tindakan    []perawatan.Tindakan    `json:"tindakan"`
	Obat        []obatmodel.Pemberian   `json:"obat"`
	Kamar       []kamarmodel.KamarInap  `json:"kamar"`
	Asesmen     []FormInfo              `json:"asesmen"`
}

type Riwayat struct {
	Pasien    *repo.Pasien       `json:"pasien"`
	Kunjungan []RiwayatKunjungan `json:"kunjungan"`
}

// RiwayatAction menyusun riwayat rekam medis dari modul pelayanan yang sudah ada.
type RiwayatAction struct {
	riwayat     repo.RiwayatStore
	icd         icdrepo.Repository
	pemeriksaan pemeriksaanrepo.Repository
	tindakan    tindakanrepo.Repository
	obat        obatrepo.Repository
	kamar       kamarrepo.Repository
}

func NewRiwayatAction(riwayat repo.RiwayatStore, icd icdrepo.Repository, pemeriksaan pemeriksaanrepo.Repository,
	tindakan tindakanrepo.Repository, obat obatrepo.Repository, kamar kamarrepo.Repository) *RiwayatAction {
	return &RiwayatAction{riwayat: riwayat, icd: icd, pemeriksaan: pemeriksaan, tindakan: tindakan, obat: obat, kamar: kamar}
}

// Riwayat kunjungan pasien terbaru lebih dulu (paginasi per kunjungan).
func (a *RiwayatAction) Riwayat(noRkmMedis, tglAwal, tglAkhir string, page, limit int) (*Riwayat, int64, error) {
	p, err := a.riwayat.Pasien(strings.TrimSpace(noRkmMedis))
	if err != nil {
		return nil, 0, err
	}
	if p == nil {
		return nil, 0, ErrPasienNotFound
	}
	kunjungan, total, err := a.riwayat.Kunjungan(p.NoRkmMedis, tglAwal, tglAkhir, page, limit)
	if err != nil {
		return nil, 0, err
	}
	forms := append(append([]FormInfo{}, DaftarForm...), formLain...)
	tables := make([][2]string, len(forms))
	label := map[string]FormInfo{}
	for i, f := range forms {
		tables[i] = [2]string{f.Slug, f.Table}
		label[f.Slug] = f
	}

	out := &Riwayat{Pasien: p, Kunjungan: make([]RiwayatKunjungan, 0, len(kunjungan))}
	for _, k := range kunjungan {
		rk := RiwayatKunjungan{Kunjungan: k, Asesmen: []FormInfo{}}
		if rk.Diagnosa, err = a.icd.List(icdmodel.Diagnosa, k.NoRawat, ""); err != nil {
			return nil, 0, err
		}
		if rk.Prosedur, err = a.icd.List(icdmodel.Prosedur, k.NoRawat, ""); err != nil {
			return nil, 0, err
		}
		for _, rawat := range []perawatan.Rawat{perawatan.Ralan, perawatan.Ranap} {
			pem, err := a.pemeriksaan.List(rawat, k.NoRawat)
			if err != nil {
				return nil, 0, err
			}
			tin, err := a.tindakan.List(rawat, k.NoRawat)
			if err != nil {
				return nil, 0, err
			}
			obt, err := a.obat.List(k.NoRawat, rawat)
			if err != nil {
				return nil, 0, err
			}
			rk.Pemeriksaan = append(rk.Pemeriksaan, pem...)
			rk.Tindakan = append(rk.Tindakan, tin...)
			rk.Obat = append(rk.Obat, obt...)
		}
		if rk.Kamar, err = a.kamar.Riwayat(k.NoRawat); err != nil {
			return nil, 0, err
		}
		slugs, err := a.riwayat.FormTerisi(k.NoRawat, tables)
		if err != nil {
			return nil, 0, err
		}
		for _, s := range slugs {
			rk.Asesmen = append(rk.Asesmen, label[s])
		}
		out.Kunjungan = append(out.Kunjungan, rk)
	}
	return out, total, nil
}
