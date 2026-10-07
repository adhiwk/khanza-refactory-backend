package seeders

import (
	"fmt"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"

	"goravel/app/modules/rbac"
)

type seedUser struct {
	orm.Model
	orm.SoftDeletes
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (seedUser) TableName() string { return "users" }

type RbacSeeder struct{}

func (s *RbacSeeder) Signature() string { return "RbacSeeder" }

func (s *RbacSeeder) Run() error {
	svc := rbac.Default()

	matrix := map[string][]string{
		rbac.SuperAdminRole: {},
		"admin": {
			"users.view", "users.create", "users.update", "users.delete",
			"roles.view", "roles.manage",
			"patients.view", "patients.create", "patients.update", "patients.delete",
		},
		"dokter":  {"patients.view", "patients.update"},
		"perawat": {"patients.view"},
	}

	for role, perms := range matrix {
		if _, err := svc.CreateRole(role); err != nil {
			return err
		}
		if len(perms) > 0 {
			if err := svc.SyncRolePermissions(role, perms...); err != nil {
				return err
			}
		}
	}

	name := fmt.Sprint(facades.Config().Env("SUPERADMIN_NAME", "Super Admin"))
	email := fmt.Sprint(facades.Config().Env("SUPERADMIN_EMAIL", "admin@goravel.local"))
	password := fmt.Sprint(facades.Config().Env("SUPERADMIN_PASSWORD", "masuk123!"))

	var u seedUser
	if err := facades.Orm().Query().Where("email", email).First(&u); err != nil {
		return err
	}
	if u.ID == 0 {
		hashed, err := facades.Hash().Make(password)
		if err != nil {
			return err
		}
		u = seedUser{Name: name, Email: email, Password: hashed}
		if err := facades.Orm().Query().Create(&u); err != nil {
			return err
		}
	}

	return svc.AssignRole(u.ID, rbac.SuperAdminRole)
}
