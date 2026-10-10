package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"goravel/app/facades"
)

type M20261011000001AddKdPegawaiToUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20261011000001AddKdPegawaiToUsersTable) Signature() string {
	return "20261011000001_add_kd_pegawai_to_users_table"
}

// Up kd_pegawai menghubungkan akun API ke pegawai.nik Khanza (pengganti akses.getkode()).
func (r *M20261011000001AddKdPegawaiToUsersTable) Up() error {
	if facades.Schema().HasColumn("users", "kd_pegawai") {
		return nil
	}
	return facades.Schema().Table("users", func(table schema.Blueprint) {
		table.String("kd_pegawai", 20).Nullable()
		table.Index("kd_pegawai")
	})
}

// Down Reverse the migrations.
func (r *M20261011000001AddKdPegawaiToUsersTable) Down() error {
	return facades.Schema().Table("users", func(table schema.Blueprint) {
		table.DropIndex("kd_pegawai")
		table.DropColumn("kd_pegawai")
	})
}
