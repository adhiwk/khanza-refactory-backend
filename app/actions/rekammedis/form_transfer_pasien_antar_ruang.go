package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormTransferPasienAntarRuang transfer pasien antar ruang (RMTransferPasienAntarRuang).
var FormTransferPasienAntarRuang = &Form[model.TransferPasienAntarRuang, request.TransferPasienAntarRuangData]{
	Slug:  "transfer-pasien-antar-ruang",
	Label: "transfer pasien antar ruang",
	Spec: repo.Spec{
		Table:  "transfer_pasien_antar_ruang",
		Keys:   []string{"no_rawat", "tanggal_masuk"},
		Waktu:  "tanggal_masuk",
		Search: []string{"no_rawat", "nip_menyerahkan", "nip_menerima"},
	},
	SetKey: func(m *model.TransferPasienAntarRuang, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.TanggalMasuk, err = KeyDateTime(key["tanggal_masuk"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.TransferPasienAntarRuang) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal_masuk": FmtDateTime(m.TanggalMasuk)}
	},
	Waktu:   func(m *model.TransferPasienAntarRuang) time.Time { return Tm(m.TanggalMasuk) },
	Petugas: func(m *model.TransferPasienAntarRuang) []string { return []string{m.NipMenyerahkan, m.NipMenerima} },
	Fill: func(m *model.TransferPasienAntarRuang, d request.TransferPasienAntarRuangData) error {
		vTanggalPindah, err := support.ParseDateTime(d.TanggalPindah)
		if err != nil {
			return err
		}
		m.TanggalPindah = vTanggalPindah
		m.AsalRuang = support.Nullable(d.AsalRuang)
		m.RuangSelanjutnya = support.Nullable(d.RuangSelanjutnya)
		m.DiagnosaUtama = support.Nullable(d.DiagnosaUtama)
		m.DiagnosaSekunder = support.Nullable(d.DiagnosaSekunder)
		m.IndikasiPindahRuang = support.Nullable(d.IndikasiPindahRuang)
		m.KeteranganIndikasiPindahRuang = support.Nullable(d.KeteranganIndikasiPindahRuang)
		m.ProsedurYangSudahDilakukan = support.Nullable(d.ProsedurYangSudahDilakukan)
		m.ObatYangTelahDiberikan = support.Nullable(d.ObatYangTelahDiberikan)
		m.MetodePemindahanPasien = support.Nullable(d.MetodePemindahanPasien)
		m.PeralatanYangMenyertai = support.Nullable(d.PeralatanYangMenyertai)
		m.KeteranganPeralatanYangMenyertai = support.Nullable(d.KeteranganPeralatanYangMenyertai)
		m.PemeriksaanPenunjangYangDilakukan = support.Nullable(d.PemeriksaanPenunjangYangDilakukan)
		m.PasienKeluargaMenyetujui = support.Nullable(d.PasienKeluargaMenyetujui)
		m.NamaMenyetujui = support.Nullable(d.NamaMenyetujui)
		m.HubunganMenyetujui = support.Nullable(d.HubunganMenyetujui)
		m.KeluhanUtamaSebelumTransfer = support.Nullable(d.KeluhanUtamaSebelumTransfer)
		m.KeadaanUmumSebelumTransfer = support.Nullable(d.KeadaanUmumSebelumTransfer)
		m.TdSebelumTransfer = support.Nullable(d.TdSebelumTransfer)
		m.NadiSebelumTransfer = support.Nullable(d.NadiSebelumTransfer)
		m.RrSebelumTransfer = support.Nullable(d.RrSebelumTransfer)
		m.SuhuSebelumTransfer = support.Nullable(d.SuhuSebelumTransfer)
		m.KeluhanUtamaSesudahTransfer = support.Nullable(d.KeluhanUtamaSesudahTransfer)
		m.KeadaanUmumSesudahTransfer = support.Nullable(d.KeadaanUmumSesudahTransfer)
		m.TdSesudahTransfer = support.Nullable(d.TdSesudahTransfer)
		m.NadiSesudahTransfer = support.Nullable(d.NadiSesudahTransfer)
		m.RrSesudahTransfer = support.Nullable(d.RrSesudahTransfer)
		m.SuhuSesudahTransfer = support.Nullable(d.SuhuSesudahTransfer)
		m.NipMenyerahkan = strings.TrimSpace(d.NipMenyerahkan)
		m.NipMenerima = strings.TrimSpace(d.NipMenerima)
		return nil
	},
	Refs: []Ref[model.TransferPasienAntarRuang]{
		{Column: "nip_menyerahkan", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.TransferPasienAntarRuang) any { return Str(m.NipMenyerahkan) }},
		{Column: "nip_menerima", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.TransferPasienAntarRuang) any { return Str(m.NipMenerima) }},
	},
}
