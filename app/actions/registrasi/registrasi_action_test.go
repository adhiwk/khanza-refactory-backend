package registrasi

import (
	"testing"
	"time"
)

func TestHitungUmur(t *testing.T) {
	now := time.Date(2026, 10, 8, 9, 0, 0, 0, time.Local)
	date := func(y int, m time.Month, d int) *time.Time {
		v := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
		return &v
	}

	cases := []struct {
		lahir *time.Time
		umur  int
		stts  string
	}{
		{date(1990, 10, 8), 36, "Th"},
		{date(1990, 10, 9), 35, "Th"},
		{date(2025, 10, 8), 1, "Th"},
		{date(2026, 3, 9), 6, "Bl"},
		{date(2026, 9, 20), 18, "Hr"},
		{date(2026, 10, 8), 0, "Hr"},
		{nil, 0, "Th"},
	}
	for _, c := range cases {
		umur, stts := hitungUmur(c.lahir, now)
		if umur != c.umur || stts != c.stts {
			t.Errorf("hitungUmur(%v) = %d %s, want %d %s", c.lahir, umur, stts, c.umur, c.stts)
		}
	}
}
