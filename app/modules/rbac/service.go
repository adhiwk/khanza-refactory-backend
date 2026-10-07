package rbac

import (
	"errors"
	"fmt"
	"sync"
)

// SuperAdminRole melewati semua pengecekan (padanan Gate::before Spatie).
const SuperAdminRole = "super-admin"

var (
	ErrRoleNotFound       = errors.New("rbac: role tidak ditemukan")
	ErrPermissionNotFound = errors.New("rbac: permission tidak ditemukan")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service { return &Service{repo: repo} }

var (
	once    sync.Once
	service *Service
)

// Default mengembalikan service singleton, dipakai middleware & gate.
func Default() *Service {
	once.Do(func() { service = NewService(NewRepository()) })
	return service
}

// ---------- helper internal ----------

func (s *Service) roleIDs(names []string, create bool) ([]uint, error) {
	ids := make([]uint, 0, len(names))
	for _, n := range names {
		role, err := s.repo.FindRole(n)
		if err != nil {
			return nil, err
		}
		if role == nil {
			if !create {
				return nil, fmt.Errorf("%w: %s", ErrRoleNotFound, n)
			}
			if role, err = s.repo.CreateRole(n); err != nil {
				return nil, err
			}
		}
		ids = append(ids, role.ID)
	}
	return ids, nil
}

func (s *Service) permIDs(names []string, create bool) ([]uint, error) {
	ids := make([]uint, 0, len(names))
	for _, n := range names {
		p, err := s.repo.FindPermission(n)
		if err != nil {
			return nil, err
		}
		if p == nil {
			if !create {
				return nil, fmt.Errorf("%w: %s", ErrPermissionNotFound, n)
			}
			if p, err = s.repo.CreatePermission(n); err != nil {
				return nil, err
			}
		}
		ids = append(ids, p.ID)
	}
	return ids, nil
}

// ---------- master data ----------

// CreateRole membuat role bila belum ada (idempotent).
func (s *Service) CreateRole(name string) (*Role, error) {
	if r, err := s.repo.FindRole(name); err != nil || r != nil {
		return r, err
	}
	return s.repo.CreateRole(name)
}

// CreatePermission membuat permission bila belum ada (idempotent).
func (s *Service) CreatePermission(name string) (*Permission, error) {
	if p, err := s.repo.FindPermission(name); err != nil || p != nil {
		return p, err
	}
	return s.repo.CreatePermission(name)
}

func (s *Service) DeleteRole(name string) error       { return s.repo.DeleteRole(name) }
func (s *Service) DeletePermission(name string) error { return s.repo.DeletePermission(name) }
func (s *Service) Roles() ([]Role, error)             { return s.repo.ListRoles() }
func (s *Service) Permissions() ([]Permission, error) { return s.repo.ListPermissions() }

// ---------- role <-> permission ----------

func (s *Service) GivePermissionToRole(role string, perms ...string) error {
	r, err := s.CreateRole(role)
	if err != nil {
		return err
	}
	ids, err := s.permIDs(perms, true)
	if err != nil {
		return err
	}
	return s.repo.AttachRolePermissions(r.ID, ids...)
}

func (s *Service) RevokePermissionFromRole(role string, perms ...string) error {
	r, err := s.repo.FindRole(role)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("%w: %s", ErrRoleNotFound, role)
	}
	ids, err := s.permIDs(perms, false)
	if err != nil {
		return err
	}
	return s.repo.DetachRolePermissions(r.ID, ids...)
}

func (s *Service) SyncRolePermissions(role string, perms ...string) error {
	r, err := s.CreateRole(role)
	if err != nil {
		return err
	}
	ids, err := s.permIDs(perms, true)
	if err != nil {
		return err
	}
	return s.repo.ReplaceRolePermissions(r.ID, ids...)
}

// ---------- user <-> role ----------

func (s *Service) AssignRole(userID uint, roles ...string) error {
	ids, err := s.roleIDs(roles, false)
	if err != nil {
		return err
	}
	return s.repo.AttachUserRoles(userID, ids...)
}

func (s *Service) RemoveRole(userID uint, roles ...string) error {
	ids, err := s.roleIDs(roles, false)
	if err != nil {
		return err
	}
	return s.repo.DetachUserRoles(userID, ids...)
}

func (s *Service) SyncRoles(userID uint, roles ...string) error {
	ids, err := s.roleIDs(roles, false)
	if err != nil {
		return err
	}
	return s.repo.ReplaceUserRoles(userID, ids...)
}

// ---------- user <-> permission (langsung) ----------

func (s *Service) GivePermissionTo(userID uint, perms ...string) error {
	ids, err := s.permIDs(perms, false)
	if err != nil {
		return err
	}
	return s.repo.AttachUserPermissions(userID, ids...)
}

func (s *Service) RevokePermissionTo(userID uint, perms ...string) error {
	ids, err := s.permIDs(perms, false)
	if err != nil {
		return err
	}
	return s.repo.DetachUserPermissions(userID, ids...)
}

func (s *Service) SyncPermissions(userID uint, perms ...string) error {
	ids, err := s.permIDs(perms, false)
	if err != nil {
		return err
	}
	return s.repo.ReplaceUserPermissions(userID, ids...)
}

// ---------- pengecekan ----------

func (s *Service) GetRoleNames(userID uint) ([]string, error) {
	ids, err := s.repo.UserRoleIDs(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.RoleNames(ids)
}

// GetAllPermissions = permission langsung + permission via role.
func (s *Service) GetAllPermissions(userID uint) ([]string, error) {
	direct, err := s.repo.UserPermissionIDs(userID)
	if err != nil {
		return nil, err
	}
	roleIDs, err := s.repo.UserRoleIDs(userID)
	if err != nil {
		return nil, err
	}
	viaRoles, err := s.repo.RolePermissionIDs(roleIDs...)
	if err != nil {
		return nil, err
	}
	seen := map[uint]struct{}{}
	all := make([]uint, 0, len(direct)+len(viaRoles))
	for _, id := range append(direct, viaRoles...) {
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			all = append(all, id)
		}
	}
	return s.repo.PermissionNames(all)
}

func (s *Service) HasRole(userID uint, role string) (bool, error) {
	return s.HasAnyRole(userID, role)
}

func (s *Service) HasAnyRole(userID uint, roles ...string) (bool, error) {
	have, err := s.GetRoleNames(userID)
	if err != nil {
		return false, err
	}
	set := toSet(have)
	for _, r := range roles {
		if _, ok := set[r]; ok {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) HasAllRoles(userID uint, roles ...string) (bool, error) {
	have, err := s.GetRoleNames(userID)
	if err != nil {
		return false, err
	}
	set := toSet(have)
	for _, r := range roles {
		if _, ok := set[r]; !ok {
			return false, nil
		}
	}
	return true, nil
}

// HasPermissionTo: murni cek permission (langsung / via role), tanpa bypass super-admin.
func (s *Service) HasPermissionTo(userID uint, perm string) (bool, error) {
	all, err := s.GetAllPermissions(userID)
	if err != nil {
		return false, err
	}
	_, ok := toSet(all)[perm]
	return ok, nil
}

// Can = super-admin bypass + HasPermissionTo. Error DB dianggap "tidak boleh".
func (s *Service) Can(userID uint, perm string) bool {
	if ok, err := s.HasRole(userID, SuperAdminRole); err == nil && ok {
		return true
	}
	ok, err := s.HasPermissionTo(userID, perm)
	return err == nil && ok
}

func toSet(items []string) map[string]struct{} {
	m := make(map[string]struct{}, len(items))
	for _, i := range items {
		m[i] = struct{}{}
	}
	return m
}