package kamarinap

import (
	"testing"
	"time"

	model "goravel/app/models/kamarinap"
)

func TestLamaInap(t *testing.T) {
	at := func(d, h int) time.Time { return time.Date(2026, 10, d, h, 0, 0, 0, time.Local) }
	cases := []struct {
		name          string
		masuk, keluar time.Time
		set           model.Pengaturan
		want          float64
	}{
		{"hari sama di bawah jam minimal", at(1, 8), at(1, 10), model.Pengaturan{JamMinimal: 6}, 0},
		{"hari sama melewati jam minimal", at(1, 8), at(1, 15), model.Pengaturan{JamMinimal: 6}, 1},
		{"beda hari dihitung kalender", at(1, 23), at(2, 1), model.Pengaturan{JamMinimal: 6}, 1},
		{"tiga hari", at(1, 8), at(4, 8), model.Pengaturan{}, 3},
		{"hitung hari awal", at(1, 8), at(4, 8), model.Pengaturan{HitungHariAwal: true}, 4},
		{"hari sama + hari awal", at(1, 8), at(1, 9), model.Pengaturan{JamMinimal: 6, HitungHariAwal: true}, 1},
	}
	for _, c := range cases {
		if got := LamaInap(c.masuk, c.keluar, c.set); got != c.want {
			t.Errorf("%s: LamaInap = %v, want %v", c.name, got, c.want)
		}
	}
}
