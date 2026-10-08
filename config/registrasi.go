package config

import (
	"goravel/app/facades"
)

func init() {
	config := facades.Config()
	config.Add("registrasi", map[string]any{
		// Dasar urutan no_reg (antrian) per tanggal, sama dengan URUTNOREG di SIMRS Khanza:
		// "dokter", "poli", atau "dokter + poli".
		"urut_no_reg": config.Env("URUT_NO_REG", "dokter"),
	})
}
