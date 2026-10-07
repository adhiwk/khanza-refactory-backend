package user

import (
	"github.com/goravel/framework/contracts/http"
)

type StoreUserRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *StoreUserRequest) Authorize(ctx http.Context) error {
	return nil
}

// Diubah dari map[string]string menjadi map[string]any sesuai Goravel terbaru
func (r *StoreUserRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"name":     "required|string|max_len:100",
		"email":    "required|email|max_len:100",
		"password": "required|string|min_len:6",
	}
}

func (r *StoreUserRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

type UpdateUserRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *UpdateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateUserRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"name":     "string|max_len:100",
		"email":    "email|max_len:100",
		"password": "string|min_len:6",
	}
}

func (r *UpdateUserRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}
