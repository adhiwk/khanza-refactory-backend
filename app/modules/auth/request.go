package auth

import (
	"github.com/goravel/framework/contracts/http"
)

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *LoginRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *LoginRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"email":    "required|email",
		"password": "required|string",
	}
}

type RegisterRequest struct {
	Name                 string `form:"name" json:"name"`
	Email                string `form:"email" json:"email"`
	Password             string `form:"password" json:"password"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation"`
}

func (r *RegisterRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RegisterRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"name":                  "required|string|max_len:100",
		"email":                 "required|email|max_len:100",
		"password":              "required|string|min_len:6",
		"password_confirmation": "required|eq_field:password",
	}
}

func (r *RegisterRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"password_confirmation.required": "Konfirmasi password wajib diisi",
		"password_confirmation.eq_field": "Konfirmasi password tidak sama dengan password",
	}
}

func (r *RegisterRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"password_confirmation": "konfirmasi password",
	}
}
