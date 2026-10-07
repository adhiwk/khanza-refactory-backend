package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20261003000001CreateRbacTables struct{}

func (r *M20261003000001CreateRbacTables) Signature() string {
	return "20261003000001_create_rbac_tables"
}

func (r *M20261003000001CreateRbacTables) Up() error {
	if err := facades.Schema().Create("roles", func(table schema.Blueprint) {
		table.ID()
		table.String("name", 125)
		table.Timestamps()
		table.Unique("name")
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("permissions", func(table schema.Blueprint) {
		table.ID()
		table.String("name", 125)
		table.Timestamps()
		table.Unique("name")
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("role_has_permissions", func(table schema.Blueprint) {
		table.UnsignedBigInteger("role_id")
		table.UnsignedBigInteger("permission_id")
		table.Primary("role_id", "permission_id")
		table.Foreign("role_id").References("id").On("roles").CascadeOnDelete()
		table.Foreign("permission_id").References("id").On("permissions").CascadeOnDelete()
	}); err != nil {
		return err
	}

	if err := facades.Schema().Create("model_has_roles", func(table schema.Blueprint) {
		table.UnsignedBigInteger("role_id")
		table.String("model_type", 50)
		table.UnsignedBigInteger("model_id")
		table.Primary("role_id", "model_type", "model_id")
		table.Index("model_type", "model_id")
		table.Foreign("role_id").References("id").On("roles").CascadeOnDelete()
	}); err != nil {
		return err
	}

	return facades.Schema().Create("model_has_permissions", func(table schema.Blueprint) {
		table.UnsignedBigInteger("permission_id")
		table.String("model_type", 50)
		table.UnsignedBigInteger("model_id")
		table.Primary("permission_id", "model_type", "model_id")
		table.Index("model_type", "model_id")
		table.Foreign("permission_id").References("id").On("permissions").CascadeOnDelete()
	})
}

func (r *M20261003000001CreateRbacTables) Down() error {
	for _, t := range []string{"model_has_permissions", "model_has_roles", "role_has_permissions", "permissions", "roles"} {
		if err := facades.Schema().DropIfExists(t); err != nil {
			return err
		}
	}
	return nil
}
