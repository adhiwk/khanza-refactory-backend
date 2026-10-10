package support

import (
	"strconv"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/http"
)

const (
	DateLayout     = "2006-01-02"
	TimeLayout     = "15:04:05"
	DateTimeLayout = "2006-01-02 15:04:05"
	maxPageLimit   = 500
)

var (
	ErrInvalidDate     = Invalid("format tanggal harus YYYY-MM-DD")
	ErrInvalidTime     = Invalid("format jam harus HH:MM:SS")
	ErrInvalidDateTime = Invalid("format waktu harus YYYY-MM-DD HH:MM:SS")
)

// PageParams membaca ?page= & ?limit= dengan default 1 / 10 dan batas atas limit.
func PageParams(ctx http.Context) (int, int) {
	page, _ := strconv.Atoi(ctx.Request().Input("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Input("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > maxPageLimit {
		limit = maxPageLimit
	}
	return page, limit
}

// Nullable string kosong menjadi NULL.
func Nullable(s string) *string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return &s
}

func StrPtr(s string) *string {
	return &s
}

func Deref[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}

func OrDefault(s, def string) string {
	if s = strings.TrimSpace(s); s != "" {
		return s
	}
	return def
}

// ParseDate "YYYY-MM-DD"; string kosong menghasilkan nil.
func ParseDate(s string) (*time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	t, err := time.ParseInLocation(DateLayout, s, time.Local)
	if err != nil {
		return nil, ErrInvalidDate
	}
	return &t, nil
}

// ParseDateTime "YYYY-MM-DD HH:MM:SS"; string kosong menghasilkan nil.
func ParseDateTime(s string) (*time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	t, err := time.ParseInLocation(DateTimeLayout, s, time.Local)
	if err != nil {
		return nil, ErrInvalidDateTime
	}
	return &t, nil
}

// ValidTime memastikan format "HH:MM:SS".
func ValidTime(s string) error {
	if _, err := time.Parse(TimeLayout, s); err != nil {
		return ErrInvalidTime
	}
	return nil
}

// Today tanggal hari ini pukul 00:00 waktu lokal.
func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

// DateOrToday tanggal dari string atau hari ini bila kosong.
func DateOrToday(s string) (time.Time, error) {
	t, err := ParseDate(s)
	if err != nil {
		return time.Time{}, err
	}
	if t == nil {
		return Today(), nil
	}
	return *t, nil
}

// TimeOrNow jam dari string atau jam sekarang bila kosong.
func TimeOrNow(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Now().Format(TimeLayout), nil
	}
	if err := ValidTime(s); err != nil {
		return "", err
	}
	return s, nil
}
