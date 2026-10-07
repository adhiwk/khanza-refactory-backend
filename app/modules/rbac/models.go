package rbac

import "github.com/goravel/framework/database/orm"

// ModelTypeUser dipakai di kolom model_type (padanan morph type Spatie).
const ModelTypeUser = "users"

type Role struct {
	orm.Model
	Name        string        `gorm:"column:name" json:"name"`
	Permissions []*Permission `gorm:"many2many:role_has_permissions;" json:"permissions,omitempty"`
}

func (Role) TableName() string { return "roles" }

type Permission struct {
	orm.Model
	Name string `gorm:"column:name" json:"name"`
}

func (Permission) TableName() string { return "permissions" }

// Pivot: role <-> permission
type RoleHasPermission struct {
	RoleID       uint `gorm:"column:role_id;primaryKey"`
	PermissionID uint `gorm:"column:permission_id;primaryKey"`
}

func (RoleHasPermission) TableName() string { return "role_has_permissions" }

// Pivot: user <-> role
type ModelHasRole struct {
	RoleID    uint   `gorm:"column:role_id;primaryKey"`
	ModelType string `gorm:"column:model_type;primaryKey"`
	ModelID   uint   `gorm:"column:model_id;primaryKey"`
}

func (ModelHasRole) TableName() string { return "model_has_roles" }

// Pivot: user <-> permission (direct permission)
type ModelHasPermission struct {
	PermissionID uint   `gorm:"column:permission_id;primaryKey"`
	ModelType    string `gorm:"column:model_type;primaryKey"`
	ModelID      uint   `gorm:"column:model_id;primaryKey"`
}

func (ModelHasPermission) TableName() string { return "model_has_permissions" }
