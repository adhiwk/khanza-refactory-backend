package config

import (
	"goravel/app/facades"
)

func init() {
	config := facades.Config()
	config.Add("farmasi", map[string]any{
		// Stok obat per batch/faktur (AKTIFKANBATCHOBAT di setting Khanza).
		"aktifkan_batch": config.Env("AKTIFKAN_BATCH_OBAT", false),
		// Kolom HPP obat: "dasar" atau "h_beli" (HPPFARMASI di setting Khanza).
		"hpp": config.Env("HPP_FARMASI", "dasar"),
		// Harga jual obat dibulatkan ke atas kelipatan 100 (PEMBULATANHARGAOBAT di setting Khanza).
		"pembulatan_harga": config.Env("PEMBULATAN_HARGA_OBAT", false),
	})
}
