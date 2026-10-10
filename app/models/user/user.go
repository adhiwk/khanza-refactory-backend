package user

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	// KdPegawai pegawai.nik Khanza milik akun ini (kode petugas/dokter pada data pelayanan).
	KdPegawai *string `json:"kd_pegawai" gorm:"column:kd_pegawai;type:varchar(20)"`
	orm.SoftDeletes
}

func (User) TableName() string {
	return "users"
}
