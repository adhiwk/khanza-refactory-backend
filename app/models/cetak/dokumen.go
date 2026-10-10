// Package cetak struktur dokumen cetak (pengganti laporan Jasper Khanza) yang dirender view resources/views/cetak.
package cetak

import "html/template"

type Dokumen struct {
	Judul       string
	Nomor       string
	Instansi    Instansi
	Identitas   []Baris
	Bagian      []Bagian
	TandaTangan []TandaTangan
	Dicetak     string
}

type Instansi struct {
	Nama      string       `gorm:"column:nama_instansi"`
	Alamat    string       `gorm:"column:alamat_instansi"`
	Kabupaten string       `gorm:"column:kabupaten"`
	Propinsi  string       `gorm:"column:propinsi"`
	Kontak    string       `gorm:"column:kontak"`
	Email     string       `gorm:"column:email"`
	Logo      template.URL `gorm:"-"` // data URI logo dari setting.logo
}

type Baris struct {
	Label string
	Nilai string
}

// Bagian satu kelompok isi: daftar label-nilai dan/atau tabel.
type Bagian struct {
	Judul string
	Baris []Baris
	Tabel *Tabel
}

type Tabel struct {
	Kolom []string
	Isi   [][]string
}

type TandaTangan struct {
	Peran string
	Nama  string
	Kode  string
}
