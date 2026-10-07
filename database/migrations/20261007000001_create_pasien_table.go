package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"goravel/app/facades"
)

type M20261007000001CreatePasienTable struct{}

// Signature The unique signature for the migration.
func (r *M20261007000001CreatePasienTable) Signature() string {
	return "20261007000001_create_pasien_table"
}

// Up Run the migrations.
// Struktur mengikuti tabel `pasien` di docs/sik_khanza_1026.sql.
// Foreign key ke tabel master (penjab, kelurahan, dst) belum dibuat karena tabelnya belum ada di migrasi.
func (r *M20261007000001CreatePasienTable) Up() error {
	if !facades.Schema().HasTable("pasien") {
		return facades.Schema().Create("pasien", func(table schema.Blueprint) {
			table.String("no_rkm_medis", 15)
			table.String("nm_pasien", 40).Nullable()
			table.String("no_ktp", 20).Nullable()
			table.Enum("jk", []any{"L", "P"}).Nullable()
			table.String("tmp_lahir", 15).Nullable()
			table.Date("tgl_lahir").Nullable()
			table.String("nm_ibu", 40)
			table.String("alamat", 200).Nullable()
			table.Enum("gol_darah", []any{"A", "B", "O", "AB", "-"}).Nullable()
			table.String("pekerjaan", 60).Nullable()
			table.Enum("stts_nikah", []any{"BELUM MENIKAH", "MENIKAH", "JANDA", "DUDHA", "JOMBLO"}).Nullable()
			table.String("agama", 12).Nullable()
			table.Date("tgl_daftar").Nullable()
			table.String("no_tlp", 40).Nullable()
			table.String("umur", 30)
			table.Enum("pnd", []any{"TS", "TK", "SD", "SMP", "SMA", "SLTA/SEDERAJAT", "D1", "D2", "D3", "D4", "S1", "S2", "S3", "-"})
			table.Enum("keluarga", []any{"AYAH", "IBU", "ISTRI", "SUAMI", "SAUDARA", "ANAK", "DIRI SENDIRI", "LAIN-LAIN"}).Nullable()
			table.String("namakeluarga", 50)
			table.Char("kd_pj", 3)
			table.String("no_peserta", 25).Nullable()
			table.Integer("kd_kel")
			table.Integer("kd_kec")
			table.Integer("kd_kab")
			table.String("pekerjaanpj", 35)
			table.String("alamatpj", 100)
			table.String("kelurahanpj", 60)
			table.String("kecamatanpj", 60)
			table.String("kabupatenpj", 60)
			table.String("perusahaan_pasien", 8)
			table.Integer("suku_bangsa")
			table.Integer("bahasa_pasien")
			table.Integer("cacat_fisik")
			table.String("email", 50)
			table.String("nip", 30)
			table.Integer("kd_prop")
			table.String("propinsipj", 30)

			table.Primary("no_rkm_medis")
			table.Index("kd_pj")
			table.Index("kd_kec")
			table.Index("kd_kab")
			table.Index("kd_kel")
			table.Index("kd_prop")
			table.Index("nm_pasien")
			table.Index("alamat")
			table.Index("no_ktp")
			table.Index("no_peserta")
			table.Index("perusahaan_pasien")
			table.Index("suku_bangsa")
			table.Index("bahasa_pasien")
			table.Index("cacat_fisik")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261007000001CreatePasienTable) Down() error {
	return facades.Schema().DropIfExists("pasien")
}
