package user

import (
	usermodel "goravel/app/models/user"

	"github.com/goravel/framework/facades"
)

type Repository interface {
	FindByID(id uint) (*usermodel.User, error)
	FindByEmail(email string) (*usermodel.User, error)
	GetAll() ([]usermodel.User, error)
	GetPaginated(page, limit int) ([]usermodel.User, int64, error)
	Create(user *usermodel.User) error
	Update(user *usermodel.User) error
	Delete(id uint) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindByID(id uint) (*usermodel.User, error) {
	var u usermodel.User
	if err := facades.Orm().Query().Where("id = ?", id).First(&u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, nil
	}
	return &u, nil
}

func (r *repository) FindByEmail(email string) (*usermodel.User, error) {
	var u usermodel.User
	if err := facades.Orm().Query().Where("email = ?", email).First(&u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, nil
	}
	return &u, nil
}

func (r *repository) GetAll() ([]usermodel.User, error) {
	var users []usermodel.User
	if err := facades.Orm().Query().Find(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetPaginated(page, limit int) ([]usermodel.User, int64, error) {
	var users []usermodel.User
	var total int64
	if err := facades.Orm().Query().Paginate(page, limit, &users, &total); err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *repository) Create(user *usermodel.User) error {
	return facades.Orm().Query().Create(user)
}

func (r *repository) Update(user *usermodel.User) error {
	return facades.Orm().Query().Save(user)
}

func (r *repository) Delete(id uint) error {
	_, err := facades.Orm().Query().Where("id = ?", id).Delete(&usermodel.User{})
	return err
}
