package obat

import (
	"errors"
	"math"
	"strings"
	"time"

	obatModel "goravel/app/models/obat"
	obatRepo "goravel/app/repository/obat"
)

var (
	ErrNotFound      = errors.New("data obat tidak ditemukan")
	ErrAlreadyExists = errors.New("kode barang sudah digunakan")
	ErrInvalidStatus = errors.New("status harus '0' atau '1'")
	ErrInvalidExpire = errors.New("format expire harus YYYY-MM-DD")
)

// ===== Request / Response =====

type ListRequest struct {
	Search       string `form:"search" json:"search"`
	Status       string `form:"status" json:"status"`
	Kdjns        string `form:"kdjns" json:"kdjns"`
	KodeKategori string `form:"kode_kategori" json:"kode_kategori"`
	KodeGolongan string `form:"kode_golongan" json:"kode_golongan"`
	Page         int    `form:"page" json:"page"`
	Limit        int    `form:"limit" json:"limit"`
}

type ListResponse struct {
	Data []obatModel.Obat `json:"data"`
	Meta Meta             `json:"meta"`
}

type Meta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

// ObatRequest dipakai untuk create & update.
// Expire dikirim sebagai string "YYYY-MM-DD" (boleh kosong / null).
type ObatRequest struct {
	KodeBrng     string  `form:"kode_brng" json:"kode_brng"`
	NamaBrng     string  `form:"nama_brng" json:"nama_brng"`
	KodeSatbesar string  `form:"kode_satbesar" json:"kode_satbesar"`
	KodeSat      string  `form:"kode_sat" json:"kode_sat"`
	LetakBarang  string  `form:"letak_barang" json:"letak_barang"`
	Dasar        float64 `form:"dasar" json:"dasar"`
	HBeli        float64 `form:"h_beli" json:"h_beli"`
	Ralan        float64 `form:"ralan" json:"ralan"`
	Kelas1       float64 `form:"kelas1" json:"kelas1"`
	Kelas2       float64 `form:"kelas2" json:"kelas2"`
	Kelas3       float64 `form:"kelas3" json:"kelas3"`
	Utama        float64 `form:"utama" json:"utama"`
	Vip          float64 `form:"vip" json:"vip"`
	Vvip         float64 `form:"vvip" json:"vvip"`
	BeliLuar     float64 `form:"beliluar" json:"beliluar"`
	JualBebas    float64 `form:"jualbebas" json:"jualbebas"`
	Karyawan     float64 `form:"karyawan" json:"karyawan"`
	StokMinimal  float64 `form:"stokminimal" json:"stokminimal"`
	Kdjns        string  `form:"kdjns" json:"kdjns"`
	Isi          float64 `form:"isi" json:"isi"`
	Kapasitas    float64 `form:"kapasitas" json:"kapasitas"`
	Expire       *string `form:"expire" json:"expire"`
	Status       string  `form:"status" json:"status"`
	KodeIndustri string  `form:"kode_industri" json:"kode_industri"`
	KodeKategori string  `form:"kode_kategori" json:"kode_kategori"`
	KodeGolongan string  `form:"kode_golongan" json:"kode_golongan"`
}

// ===== Action =====

type Action struct {
	obatRepo obatRepo.Repository
}

func NewAction(obatRepo obatRepo.Repository) *Action {
	return &Action{obatRepo: obatRepo}
}

func (a *Action) List(req ListRequest) (*ListResponse, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	list, total, err := a.obatRepo.Paginate(obatRepo.Filter{
		Search:       strings.TrimSpace(req.Search),
		Status:       req.Status,
		Kdjns:        req.Kdjns,
		KodeKategori: req.KodeKategori,
		KodeGolongan: req.KodeGolongan,
		Page:         req.Page,
		Limit:        req.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &ListResponse{
		Data: list,
		Meta: Meta{
			Page:      req.Page,
			Limit:     req.Limit,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(req.Limit))),
		},
	}, nil
}

func (a *Action) Detail(kode string) (*obatModel.Obat, error) {
	data, err := a.obatRepo.FindByKode(kode)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req ObatRequest) (*obatModel.Obat, error) {
	req.KodeBrng = strings.TrimSpace(req.KodeBrng)

	exists, err := a.obatRepo.ExistsByKode(req.KodeBrng)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &obatModel.Obat{KodeBrng: req.KodeBrng}
	if err := fill(data, req); err != nil {
		return nil, err
	}

	if err := a.obatRepo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Update: kode_brng (primary key) tidak bisa diubah, diambil dari URL.
func (a *Action) Update(kode string, req ObatRequest) (*obatModel.Obat, error) {
	data, err := a.Detail(kode)
	if err != nil {
		return nil, err
	}

	if err := fill(data, req); err != nil {
		return nil, err
	}

	if err := a.obatRepo.Save(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Delete(kode string) error {
	affected, err := a.obatRepo.Delete(kode)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *Action) UpdateStatus(kode string, status string) (*obatModel.Obat, error) {
	if status != "0" && status != "1" {
		return nil, ErrInvalidStatus
	}

	if _, err := a.Detail(kode); err != nil {
		return nil, err
	}

	if _, err := a.obatRepo.UpdateStatus(kode, status); err != nil {
		return nil, err
	}
	return a.Detail(kode)
}

// ===== helper =====

func fill(data *obatModel.Obat, req ObatRequest) error {
	if req.Status != "0" && req.Status != "1" {
		return ErrInvalidStatus
	}

	expire, err := parseDate(req.Expire)
	if err != nil {
		return err
	}

	data.NamaBrng = req.NamaBrng
	data.KodeSatbesar = req.KodeSatbesar
	data.KodeSat = req.KodeSat
	data.LetakBarang = req.LetakBarang
	data.Dasar = req.Dasar
	data.HBeli = req.HBeli
	data.Ralan = req.Ralan
	data.Kelas1 = req.Kelas1
	data.Kelas2 = req.Kelas2
	data.Kelas3 = req.Kelas3
	data.Utama = req.Utama
	data.Vip = req.Vip
	data.Vvip = req.Vvip
	data.BeliLuar = req.BeliLuar
	data.JualBebas = req.JualBebas
	data.Karyawan = req.Karyawan
	data.StokMinimal = req.StokMinimal
	data.Kdjns = req.Kdjns
	data.Isi = req.Isi
	data.Kapasitas = req.Kapasitas
	data.Expire = expire
	data.Status = req.Status
	data.KodeIndustri = req.KodeIndustri
	data.KodeKategori = req.KodeKategori
	data.KodeGolongan = req.KodeGolongan

	return nil
}

func parseDate(s *string) (*time.Time, error) {
	if s == nil || strings.TrimSpace(*s) == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", strings.TrimSpace(*s))
	if err != nil {
		return nil, ErrInvalidExpire
	}
	return &t, nil
}
