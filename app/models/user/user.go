package user

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	orm.SoftDeletes
}

func (User) TableName() string {
	return "users"
}
