package pemberianobat

import (
	"testing"

	model "goravel/app/models/pemberianobat"
)

func TestHargaJual(t *testing.T) {
	b := model.Barang{HBeli: 1000, Ralan: 1234, Kelas1: 1500, Karyawan: 1100}
	cases := []struct {
		name       string
		kenaikan   float64
		kolom      string
		pembulatan bool
		want       float64
	}{
		{"kolom ralan tanpa pembulatan", 0, "ralan", false, 1234},
		{"kolom ralan dibulatkan ke atas 100", 0, "ralan", true, 1300},
		{"kelipatan 100 tetap", 0, "kelas1", true, 1500},
		{"kenaikan 25% dari harga beli", 0.25, "ralan", false, 1250},
		{"kenaikan dibulatkan", 0.25, "karyawan", true, 1300},
		{"kolom tidak dikenal = ralan", 0, "x", false, 1234},
	}
	for _, c := range cases {
		if got := HargaJual(b, c.kenaikan, c.kolom, c.pembulatan); got != c.want {
			t.Errorf("%s: HargaJual = %v, want %v", c.name, got, c.want)
		}
	}
}
