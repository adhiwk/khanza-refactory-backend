package user

import (
	"strings"

	usermodel "goravel/app/models/user"
	userrepo "goravel/app/repository/user"
	"goravel/app/support"

	"github.com/goravel/framework/facades"
)

var (
	ErrNotFound        = support.NotFound("data user tidak ditemukan")
	ErrPegawaiNotFound = support.NotFound("data pegawai tidak ditemukan")
)

type Action struct {
	repo userrepo.Repository
}

func NewAction(repo userrepo.Repository) *Action {
	return &Action{
		repo: repo,
	}
}

func (a *Action) List(page, limit int) ([]usermodel.User, int64, error) {
	return a.repo.GetPaginated(page, limit)
}

func (a *Action) Detail(id uint) (*usermodel.User, error) {
	return a.repo.FindByID(id)
}

func (a *Action) Create(name, email, plainPassword string) (*usermodel.User, error) {
	// Hash password menggunakan Goravel Hash facade
	hashedPassword, err := facades.Hash().Make(plainPassword)
	if err != nil {
		return nil, err
	}

	user := &usermodel.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = a.repo.Create(user)
	return user, err
}

func (a *Action) Update(id uint, name, email, plainPassword string) (*usermodel.User, error) {
	user, err := a.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if plainPassword != "" {
		hashedPassword, err := facades.Hash().Make(plainPassword)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	err = a.repo.Update(user)
	return user, err
}

// SetPegawai menghubungkan akun ke pegawai Khanza; kd kosong melepas hubungan.
func (a *Action) SetPegawai(id uint, kd string) (*usermodel.User, error) {
	user, err := a.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID == 0 {
		return nil, ErrNotFound
	}
	kd = strings.TrimSpace(kd)
	if kd == "" {
		user.KdPegawai = nil
	} else {
		ok, err := a.repo.PegawaiExists(kd)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, ErrPegawaiNotFound
		}
		user.KdPegawai = &kd
	}
	err = a.repo.Update(user)
	return user, err
}

func (a *Action) Delete(id uint) error {
	return a.repo.Delete(id)
}
