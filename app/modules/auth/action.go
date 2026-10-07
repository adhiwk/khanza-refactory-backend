package auth

import (
	"errors"

	usermodel "goravel/app/models/user"
	userrepo "goravel/app/repository/user"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

var (
	ErrInvalidCredentials = errors.New("email atau password salah")
	ErrEmailTaken         = errors.New("email sudah terdaftar")
)

type Action struct {
	userRepo userrepo.Repository
}

func NewAction(userRepo userrepo.Repository) *Action {
	return &Action{userRepo: userRepo}
}

func (a *Action) Login(ctx http.Context, email, plainPassword string) (string, error) {
	u, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if u == nil {
		return "", ErrInvalidCredentials
	}

	if !facades.Hash().Check(plainPassword, u.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := facades.Auth(ctx).LoginUsingID(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Action) Register(name, email, plainPassword string) (*usermodel.User, error) {
	existing, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrEmailTaken
	}

	hashed, err := facades.Hash().Make(plainPassword)
	if err != nil {
		return nil, err
	}

	newUser := &usermodel.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}
	if err := a.userRepo.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (a *Action) Me(ctx http.Context) (*usermodel.User, error) {
	var u usermodel.User
	if err := facades.Auth(ctx).User(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
