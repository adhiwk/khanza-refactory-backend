package support

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

// KhanzaConnection koneksi database SIMRS Khanza (skema sik.sql).
const KhanzaConnection = "mysql_kedua"

// DB query baru pada database Khanza.
func DB() contractsorm.Query {
	return facades.Orm().Connection(KhanzaConnection).Query()
}

// Transaction menjalankan fn dalam satu transaksi database Khanza; error dari fn membatalkan seluruh perubahan.
func Transaction(fn func(tx contractsorm.Query) error) error {
	return facades.Orm().Connection(KhanzaConnection).Transaction(fn)
}

// IsDuplicate true bila err adalah pelanggaran unique/primary key MySQL.
func IsDuplicate(err error) bool {
	var myErr *mysql.MySQLError
	return errors.As(err, &myErr) && myErr.Number == 1062
}

// IsForeignKey true bila err adalah pelanggaran foreign key MySQL (data masih dipakai / referensi tidak ada).
func IsForeignKey(err error) bool {
	var myErr *mysql.MySQLError
	return errors.As(err, &myErr) && (myErr.Number == 1451 || myErr.Number == 1452)
}
