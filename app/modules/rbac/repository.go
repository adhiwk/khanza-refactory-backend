package rbac

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type Repository interface {
	// Master data (nil, nil bila tidak ditemukan)
	FindRole(name string) (*Role, error)
	FindPermission(name string) (*Permission, error)
	CreateRole(name string) (*Role, error)
	CreatePermission(name string) (*Permission, error)
	DeleteRole(name string) error
	DeletePermission(name string) error
	ListRoles() ([]Role, error)
	ListPermissions() ([]Permission, error)
	RoleNames(ids []uint) ([]string, error)
	PermissionNames(ids []uint) ([]string, error)

	// User <-> Role
	UserRoleIDs(userID uint) ([]uint, error)
	AttachUserRoles(userID uint, roleIDs ...uint) error
	DetachUserRoles(userID uint, roleIDs ...uint) error
	ReplaceUserRoles(userID uint, roleIDs ...uint) error

	// User <-> Permission (langsung)
	UserPermissionIDs(userID uint) ([]uint, error)
	AttachUserPermissions(userID uint, permIDs ...uint) error
	DetachUserPermissions(userID uint, permIDs ...uint) error
	ReplaceUserPermissions(userID uint, permIDs ...uint) error

	// Role <-> Permission
	RolePermissionIDs(roleIDs ...uint) ([]uint, error)
	AttachRolePermissions(roleID uint, permIDs ...uint) error
	DetachRolePermissions(roleID uint, permIDs ...uint) error
	ReplaceRolePermissions(roleID uint, permIDs ...uint) error
}

type repository struct{}

func NewRepository() Repository { return &repository{} }

// ---------- master data ----------

func (r *repository) FindRole(name string) (*Role, error) {
	var role Role
	if err := facades.Orm().Query().Where("name", name).First(&role); err != nil {
		return nil, err
	}
	if role.ID == 0 {
		return nil, nil
	}
	return &role, nil
}

func (r *repository) FindPermission(name string) (*Permission, error) {
	var p Permission
	if err := facades.Orm().Query().Where("name", name).First(&p); err != nil {
		return nil, err
	}
	if p.ID == 0 {
		return nil, nil
	}
	return &p, nil
}

func (r *repository) CreateRole(name string) (*Role, error) {
	role := &Role{Name: name}
	if err := facades.Orm().Query().Create(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (r *repository) CreatePermission(name string) (*Permission, error) {
	p := &Permission{Name: name}
	if err := facades.Orm().Query().Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *repository) DeleteRole(name string) error {
	_, err := facades.Orm().Query().Where("name", name).Delete(&Role{})
	return err
}

func (r *repository) DeletePermission(name string) error {
	_, err := facades.Orm().Query().Where("name", name).Delete(&Permission{})
	return err
}

func (r *repository) ListRoles() ([]Role, error) {
	var roles []Role
	err := facades.Orm().Query().With("Permissions").Order("name asc").Find(&roles)
	return roles, err
}

func (r *repository) ListPermissions() ([]Permission, error) {
	var perms []Permission
	err := facades.Orm().Query().Order("name asc").Find(&perms)
	return perms, err
}

func (r *repository) RoleNames(ids []uint) ([]string, error) {
	var names []string
	if len(ids) == 0 {
		return names, nil
	}
	err := facades.Orm().Query().Model(&Role{}).Where("id in ?", ids).Pluck("name", &names)
	return names, err
}

func (r *repository) PermissionNames(ids []uint) ([]string, error) {
	var names []string
	if len(ids) == 0 {
		return names, nil
	}
	err := facades.Orm().Query().Model(&Permission{}).Where("id in ?", ids).Pluck("name", &names)
	return names, err
}

// ---------- user <-> role ----------

func (r *repository) UserRoleIDs(userID uint) ([]uint, error) {
	var ids []uint
	err := facades.Orm().Query().Model(&ModelHasRole{}).
		Where("model_type", ModelTypeUser).Where("model_id", userID).
		Pluck("role_id", &ids)
	return ids, err
}

func attachUserRoles(q orm.Query, userID uint, roleIDs []uint) error {
	for _, id := range roleIDs {
		n, err := q.Model(&ModelHasRole{}).
			Where("role_id", id).Where("model_type", ModelTypeUser).Where("model_id", userID).Count()
		if err != nil {
			return err
		}
		if n > 0 {
			continue
		}
		if err := q.Create(&ModelHasRole{RoleID: id, ModelType: ModelTypeUser, ModelID: userID}); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) AttachUserRoles(userID uint, roleIDs ...uint) error {
	return attachUserRoles(facades.Orm().Query(), userID, roleIDs)
}

func (r *repository) DetachUserRoles(userID uint, roleIDs ...uint) error {
	if len(roleIDs) == 0 {
		return nil
	}
	_, err := facades.Orm().Query().
		Where("model_type", ModelTypeUser).Where("model_id", userID).Where("role_id in ?", roleIDs).
		Delete(&ModelHasRole{})
	return err
}

func (r *repository) ReplaceUserRoles(userID uint, roleIDs ...uint) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		if _, err := tx.Where("model_type", ModelTypeUser).Where("model_id", userID).Delete(&ModelHasRole{}); err != nil {
			return err
		}
		return attachUserRoles(tx, userID, roleIDs)
	})
}

// ---------- user <-> permission ----------

func (r *repository) UserPermissionIDs(userID uint) ([]uint, error) {
	var ids []uint
	err := facades.Orm().Query().Model(&ModelHasPermission{}).
		Where("model_type", ModelTypeUser).Where("model_id", userID).
		Pluck("permission_id", &ids)
	return ids, err
}

func attachUserPermissions(q orm.Query, userID uint, permIDs []uint) error {
	for _, id := range permIDs {
		n, err := q.Model(&ModelHasPermission{}).
			Where("permission_id", id).Where("model_type", ModelTypeUser).Where("model_id", userID).Count()
		if err != nil {
			return err
		}
		if n > 0 {
			continue
		}
		if err := q.Create(&ModelHasPermission{PermissionID: id, ModelType: ModelTypeUser, ModelID: userID}); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) AttachUserPermissions(userID uint, permIDs ...uint) error {
	return attachUserPermissions(facades.Orm().Query(), userID, permIDs)
}

func (r *repository) DetachUserPermissions(userID uint, permIDs ...uint) error {
	if len(permIDs) == 0 {
		return nil
	}
	_, err := facades.Orm().Query().
		Where("model_type", ModelTypeUser).Where("model_id", userID).Where("permission_id in ?", permIDs).
		Delete(&ModelHasPermission{})
	return err
}

func (r *repository) ReplaceUserPermissions(userID uint, permIDs ...uint) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		if _, err := tx.Where("model_type", ModelTypeUser).Where("model_id", userID).Delete(&ModelHasPermission{}); err != nil {
			return err
		}
		return attachUserPermissions(tx, userID, permIDs)
	})
}

// ---------- role <-> permission ----------

func (r *repository) RolePermissionIDs(roleIDs ...uint) ([]uint, error) {
	var ids []uint
	if len(roleIDs) == 0 {
		return ids, nil
	}
	err := facades.Orm().Query().Model(&RoleHasPermission{}).
		Where("role_id in ?", roleIDs).Pluck("permission_id", &ids)
	return ids, err
}

func attachRolePermissions(q orm.Query, roleID uint, permIDs []uint) error {
	for _, id := range permIDs {
		n, err := q.Model(&RoleHasPermission{}).
			Where("role_id", roleID).Where("permission_id", id).Count()
		if err != nil {
			return err
		}
		if n > 0 {
			continue
		}
		if err := q.Create(&RoleHasPermission{RoleID: roleID, PermissionID: id}); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) AttachRolePermissions(roleID uint, permIDs ...uint) error {
	return attachRolePermissions(facades.Orm().Query(), roleID, permIDs)
}

func (r *repository) DetachRolePermissions(roleID uint, permIDs ...uint) error {
	if len(permIDs) == 0 {
		return nil
	}
	_, err := facades.Orm().Query().
		Where("role_id", roleID).Where("permission_id in ?", permIDs).
		Delete(&RoleHasPermission{})
	return err
}

func (r *repository) ReplaceRolePermissions(roleID uint, permIDs ...uint) error {
	return facades.Orm().Transaction(func(tx orm.Query) error {
		if _, err := tx.Where("role_id", roleID).Delete(&RoleHasPermission{}); err != nil {
			return err
		}
		return attachRolePermissions(tx, roleID, permIDs)
	})
}
