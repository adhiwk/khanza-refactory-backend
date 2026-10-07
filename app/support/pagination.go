package support

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PaginationLink struct {
	Url    *string `json:"url"`
	Label  string  `json:"label"`
	Page   *int    `json:"page"`
	Active bool    `json:"active"`
}

type LaravelPaginationResponse struct {
	CurrentPage  int              `json:"current_page"`
	Data         interface{}      `json:"data"`
	FirstPageUrl string           `json:"first_page_url"`
	From         *int             `json:"from"`
	LastPage     int              `json:"last_page"`
	LastPageUrl  string           `json:"last_page_url"`
	Links        []PaginationLink `json:"links"`
	NextPageUrl  *string          `json:"next_page_url"`
	Path         string           `json:"path"`
	PerPage      int              `json:"per_page"`
	PrevPageUrl  *string          `json:"prev_page_url"`
	To           *int             `json:"to"`
	Total        int64            `json:"total"`
}

func FormatLaravelPagination(ctx http.Context, data interface{}, page, limit int, total int64) LaravelPaginationResponse {
	// 1. Hitung Last Page
	lastPage := int(total / int64(limit))
	if total%int64(limit) > 0 {
		lastPage++
	}
	if lastPage == 0 {
		lastPage = 1
	}

	// 2. Ambil base URL secara aman dari konfigurasi app.url dan path request saat ini
	appUrl := facades.Config().GetString("app.url", "http://localhost:3000")
	currentPath := ctx.Request().Url() // contoh: /api/v1/proyek

	baseUrl := fmt.Sprintf("%s%s", appUrl, currentPath)

	firstPageUrl := fmt.Sprintf("%s?page=1", baseUrl)
	lastPageUrl := fmt.Sprintf("%s?page=%d", baseUrl, lastPage)

	// 3. Hitung From dan To (indeks item yang ditampilkan)
	var from, to *int
	if total > 0 {
		f := ((page - 1) * limit) + 1
		t := f + limit - 1
		if int64(t) > total {
			t = int(total)
		}
		from = &f
		to = &t
	}

	// 4. Buat Prev & Next URL
	var prevUrl, nextUrl *string
	if page > 1 {
		u := fmt.Sprintf("%s?page=%d", baseUrl, page-1)
		prevUrl = &u
	}
	if page < lastPage {
		u := fmt.Sprintf("%s?page=%d", baseUrl, page+1)
		nextUrl = &u
	}

	// 5. Bangun Array Links
	var links []PaginationLink

	var prevPageNum *int
	if page > 1 {
		p := page - 1
		prevPageNum = &p
	}
	links = append(links, PaginationLink{
		Url:    prevUrl,
		Label:  "&laquo; Previous",
		Page:   prevPageNum,
		Active: false,
	})

	for i := 1; i <= lastPage; i++ {
		u := fmt.Sprintf("%s?page=%d", baseUrl, i)
		pageNum := i
		links = append(links, PaginationLink{
			Url:    &u,
			Label:  fmt.Sprintf("%d", i),
			Page:   &pageNum,
			Active: i == page,
		})
	}

	var nextPageNum *int
	if page < lastPage {
		p := page + 1
		nextPageNum = &p
	}
	links = append(links, PaginationLink{
		Url:    nextUrl,
		Label:  "Next &raquo;",
		Page:   nextPageNum,
		Active: false,
	})

	return LaravelPaginationResponse{
		CurrentPage:  page,
		Data:         data,
		FirstPageUrl: firstPageUrl,
		From:         from,
		LastPage:     lastPage,
		LastPageUrl:  lastPageUrl,
		Links:        links,
		NextPageUrl:  nextUrl,
		Path:         baseUrl,
		PerPage:      limit,
		PrevPageUrl:  prevUrl,
		To:           to,
		Total:        total,
	}
}
