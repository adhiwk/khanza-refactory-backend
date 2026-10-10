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
			"registrations.view", "registrations.create", "registrations.update", "registrations.delete",
			"masterdata.view", "masterdata.create", "masterdata.update", "masterdata.delete",
			"igd.view", "igd.create",
			"rawat_jalan.view", "rawat_jalan.create", "rawat_jalan.update", "rawat_jalan.delete",
			"rawat_inap.view", "rawat_inap.create", "rawat_inap.update", "rawat_inap.delete",
			"pelayanan.view", "pelayanan.create", "pelayanan.update", "pelayanan.delete",
			"farmasi_master.view", "farmasi_master.create", "farmasi_master.update", "farmasi_master.delete",
			"resep.view", "resep.create", "resep.update", "resep.delete",
			"billing.view", "billing.create", "billing.delete",
		},
		"dokter": {"patients.view", "patients.update", "masterdata.view",
			"rawat_jalan.view", "rawat_jalan.create", "rawat_jalan.update", "rawat_jalan.delete",
			"rawat_inap.view", "rawat_inap.create", "rawat_inap.update", "rawat_inap.delete",
			"pelayanan.view", "pelayanan.create", "pelayanan.update",
			"resep.view", "resep.create", "resep.update", "resep.delete"},
		"apoteker": {"patients.view", "masterdata.view", "rawat_jalan.view", "rawat_jalan.create", "rawat_jalan.delete",
			"rawat_inap.view", "rawat_inap.create", "rawat_inap.delete",
			"farmasi_master.view", "farmasi_master.create", "farmasi_master.update", "farmasi_master.delete", "resep.view"},
		"kasir": {"patients.view", "registrations.view", "billing.view", "billing.create", "billing.delete"},
		"perawat": {"patients.view", "masterdata.view", "igd.view", "igd.create",
			"rawat_jalan.view", "rawat_jalan.create", "rawat_jalan.update",
			"rawat_inap.view", "rawat_inap.create", "rawat_inap.update", "pelayanan.view"},
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
