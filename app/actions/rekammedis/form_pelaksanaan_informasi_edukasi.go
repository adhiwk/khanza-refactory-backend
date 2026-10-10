package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPelaksanaanInformasiEdukasi pelaksanaan informasi edukasi (RMPelaksanaanInformasiEdukasi).
var FormPelaksanaanInformasiEdukasi = &Form[model.PelaksanaanInformasiEdukasi, request.PelaksanaanInformasiEdukasiData]{
	Slug:  "pelaksanaan-informasi-edukasi",
	Label: "pelaksanaan informasi edukasi",
	Spec: repo.Spec{
		Table:  "pelaksanaan_informasi_edukasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.PelaksanaanInformasiEdukasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PelaksanaanInformasiEdukasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PelaksanaanInformasiEdukasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PelaksanaanInformasiEdukasi) []string { return []string{m.Nik} },
	Fill: func(m *model.PelaksanaanInformasiEdukasi, d request.PelaksanaanInformasiEdukasiData) error {
		m.Nik = strings.TrimSpace(d.Nik)
		m.MateriEdukasi = support.Nullable(d.MateriEdukasi)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.DiberikanPada = strings.TrimSpace(d.DiberikanPada)
		m.KeteranganDiberikanPada = strings.TrimSpace(d.KeteranganDiberikanPada)
		m.LamaEdukasi = support.Nullable(d.LamaEdukasi)
		m.MetodeEdukasi = support.Nullable(d.MetodeEdukasi)
		m.HasilVerifikasi = support.Nullable(d.HasilVerifikasi)
		m.Status = support.Nullable(d.Status)
		return nil
	},
	Refs: []Ref[model.PelaksanaanInformasiEdukasi]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.PelaksanaanInformasiEdukasi) any { return Str(m.Nik) }},
	},
}
