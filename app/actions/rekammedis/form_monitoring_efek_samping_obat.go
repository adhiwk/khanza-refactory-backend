package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormMonitoringEfekSampingObat pelaporan efek samping obat (RMPelaporanEfekSampingObat).
var FormMonitoringEfekSampingObat = &Form[model.MonitoringEfekSampingObat, request.MonitoringEfekSampingObatData]{
	Slug:  "monitoring-efek-samping-obat",
	Label: "pelaporan efek samping obat",
	Spec: repo.Spec{
		Table:  "monitoring_efek_samping_obat",
		Keys:   []string{"no_rawat"},
		Waktu:  "",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.MonitoringEfekSampingObat, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.MonitoringEfekSampingObat) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.MonitoringEfekSampingObat) time.Time { return time.Time{} },
	Petugas: func(m *model.MonitoringEfekSampingObat) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.MonitoringEfekSampingObat, d request.MonitoringEfekSampingObatData) error {
		vTanggal, err := support.ParseDate(d.Tanggal)
		if err != nil {
			return err
		}
		vTanggalKejadian, err := support.ParseDate(d.TanggalKejadian)
		if err != nil {
			return err
		}
		vTanggalKesudahan, err := support.ParseDate(d.TanggalKesudahan)
		if err != nil {
			return err
		}
		vTanggalMulai1, err := support.ParseDate(d.TanggalMulai1)
		if err != nil {
			return err
		}
		vTanggalAkhir1, err := support.ParseDate(d.TanggalAkhir1)
		if err != nil {
			return err
		}
		vTanggalMulai2, err := support.ParseDate(d.TanggalMulai2)
		if err != nil {
			return err
		}
		vTanggalAkhir2, err := support.ParseDate(d.TanggalAkhir2)
		if err != nil {
			return err
		}
		vTanggalMulai3, err := support.ParseDate(d.TanggalMulai3)
		if err != nil {
			return err
		}
		vTanggalAkhir3, err := support.ParseDate(d.TanggalAkhir3)
		if err != nil {
			return err
		}
		vTanggalMulai4, err := support.ParseDate(d.TanggalMulai4)
		if err != nil {
			return err
		}
		vTanggalAkhir4, err := support.ParseDate(d.TanggalAkhir4)
		if err != nil {
			return err
		}
		vTanggalMulai5, err := support.ParseDate(d.TanggalMulai5)
		if err != nil {
			return err
		}
		vTanggalAkhir5, err := support.ParseDate(d.TanggalAkhir5)
		if err != nil {
			return err
		}
		vTanggalMulai6, err := support.ParseDate(d.TanggalMulai6)
		if err != nil {
			return err
		}
		vTanggalAkhir6, err := support.ParseDate(d.TanggalAkhir6)
		if err != nil {
			return err
		}
		vTanggalMulai7, err := support.ParseDate(d.TanggalMulai7)
		if err != nil {
			return err
		}
		vTanggalAkhir7, err := support.ParseDate(d.TanggalAkhir7)
		if err != nil {
			return err
		}
		vTanggalMulai8, err := support.ParseDate(d.TanggalMulai8)
		if err != nil {
			return err
		}
		vTanggalAkhir8, err := support.ParseDate(d.TanggalAkhir8)
		if err != nil {
			return err
		}
		vTanggalMulai9, err := support.ParseDate(d.TanggalMulai9)
		if err != nil {
			return err
		}
		vTanggalAkhir9, err := support.ParseDate(d.TanggalAkhir9)
		if err != nil {
			return err
		}
		vTanggalMulai10, err := support.ParseDate(d.TanggalMulai10)
		if err != nil {
			return err
		}
		vTanggalAkhir10, err := support.ParseDate(d.TanggalAkhir10)
		if err != nil {
			return err
		}
		m.NoLaporan = strings.TrimSpace(d.NoLaporan)
		m.Tanggal = vTanggal
		m.Profesi = strings.TrimSpace(d.Profesi)
		m.Nik = support.Nullable(d.Nik)
		m.KdBangsal = support.Nullable(d.KdBangsal)
		m.BeratBadan = support.Nullable(d.BeratBadan)
		m.PasienHamil = strings.TrimSpace(d.PasienHamil)
		m.Kesudahan = strings.TrimSpace(d.Kesudahan)
		m.PenyakitLain = strings.TrimSpace(d.PenyakitLain)
		m.PenyakitUtama = strings.TrimSpace(d.PenyakitUtama)
		m.TanggalKejadian = vTanggalKejadian
		m.Manifestasi = strings.TrimSpace(d.Manifestasi)
		m.MasalahKualitas = strings.TrimSpace(d.MasalahKualitas)
		m.RiwayatEso = strings.TrimSpace(d.RiwayatEso)
		m.TanggalKesudahan = vTanggalKesudahan
		m.HasilKesudahan = strings.TrimSpace(d.HasilKesudahan)
		m.Obat1 = strings.TrimSpace(d.Obat1)
		m.Sedian1 = strings.TrimSpace(d.Sedian1)
		m.ObatJkn1 = strings.TrimSpace(d.ObatJkn1)
		m.Batch1 = strings.TrimSpace(d.Batch1)
		m.Cara1 = strings.TrimSpace(d.Cara1)
		m.Dosis1 = strings.TrimSpace(d.Dosis1)
		m.TanggalMulai1 = vTanggalMulai1
		m.TanggalAkhir1 = vTanggalAkhir1
		m.Indikasi1 = strings.TrimSpace(d.Indikasi1)
		m.Obat2 = strings.TrimSpace(d.Obat2)
		m.Sedian2 = strings.TrimSpace(d.Sedian2)
		m.ObatJkn2 = strings.TrimSpace(d.ObatJkn2)
		m.Batch2 = strings.TrimSpace(d.Batch2)
		m.Cara2 = strings.TrimSpace(d.Cara2)
		m.Dosis2 = strings.TrimSpace(d.Dosis2)
		m.TanggalMulai2 = vTanggalMulai2
		m.TanggalAkhir2 = vTanggalAkhir2
		m.Indikasi2 = strings.TrimSpace(d.Indikasi2)
		m.Obat3 = strings.TrimSpace(d.Obat3)
		m.Sedian3 = strings.TrimSpace(d.Sedian3)
		m.ObatJkn3 = strings.TrimSpace(d.ObatJkn3)
		m.Batch3 = strings.TrimSpace(d.Batch3)
		m.Cara3 = strings.TrimSpace(d.Cara3)
		m.Dosis3 = strings.TrimSpace(d.Dosis3)
		m.TanggalMulai3 = vTanggalMulai3
		m.TanggalAkhir3 = vTanggalAkhir3
		m.Indikasi3 = strings.TrimSpace(d.Indikasi3)
		m.Obat4 = strings.TrimSpace(d.Obat4)
		m.Sedian4 = strings.TrimSpace(d.Sedian4)
		m.ObatJkn4 = strings.TrimSpace(d.ObatJkn4)
		m.Batch4 = strings.TrimSpace(d.Batch4)
		m.Cara4 = strings.TrimSpace(d.Cara4)
		m.Dosis4 = strings.TrimSpace(d.Dosis4)
		m.TanggalMulai4 = vTanggalMulai4
		m.TanggalAkhir4 = vTanggalAkhir4
		m.Indikasi4 = strings.TrimSpace(d.Indikasi4)
		m.Obat5 = strings.TrimSpace(d.Obat5)
		m.Sedian5 = strings.TrimSpace(d.Sedian5)
		m.ObatJkn5 = strings.TrimSpace(d.ObatJkn5)
		m.Batch5 = strings.TrimSpace(d.Batch5)
		m.Cara5 = strings.TrimSpace(d.Cara5)
		m.Dosis5 = strings.TrimSpace(d.Dosis5)
		m.TanggalMulai5 = vTanggalMulai5
		m.TanggalAkhir5 = vTanggalAkhir5
		m.Indikasi5 = strings.TrimSpace(d.Indikasi5)
		m.Obat6 = strings.TrimSpace(d.Obat6)
		m.Sedian6 = strings.TrimSpace(d.Sedian6)
		m.ObatJkn6 = strings.TrimSpace(d.ObatJkn6)
		m.Batch6 = strings.TrimSpace(d.Batch6)
		m.Cara6 = strings.TrimSpace(d.Cara6)
		m.Dosis6 = strings.TrimSpace(d.Dosis6)
		m.TanggalMulai6 = vTanggalMulai6
		m.TanggalAkhir6 = vTanggalAkhir6
		m.Indikasi6 = strings.TrimSpace(d.Indikasi6)
		m.Obat7 = strings.TrimSpace(d.Obat7)
		m.Sedian7 = strings.TrimSpace(d.Sedian7)
		m.ObatJkn7 = strings.TrimSpace(d.ObatJkn7)
		m.Batch7 = strings.TrimSpace(d.Batch7)
		m.Cara7 = strings.TrimSpace(d.Cara7)
		m.Dosis7 = strings.TrimSpace(d.Dosis7)
		m.TanggalMulai7 = vTanggalMulai7
		m.TanggalAkhir7 = vTanggalAkhir7
		m.Indikasi7 = strings.TrimSpace(d.Indikasi7)
		m.Obat8 = strings.TrimSpace(d.Obat8)
		m.Sedian8 = strings.TrimSpace(d.Sedian8)
		m.ObatJkn8 = strings.TrimSpace(d.ObatJkn8)
		m.Batch8 = strings.TrimSpace(d.Batch8)
		m.Cara8 = strings.TrimSpace(d.Cara8)
		m.Dosis8 = strings.TrimSpace(d.Dosis8)
		m.TanggalMulai8 = vTanggalMulai8
		m.TanggalAkhir8 = vTanggalAkhir8
		m.Indikasi8 = strings.TrimSpace(d.Indikasi8)
		m.Obat9 = strings.TrimSpace(d.Obat9)
		m.Sedian9 = strings.TrimSpace(d.Sedian9)
		m.ObatJkn9 = strings.TrimSpace(d.ObatJkn9)
		m.Batch9 = strings.TrimSpace(d.Batch9)
		m.Cara9 = strings.TrimSpace(d.Cara9)
		m.Dosis9 = strings.TrimSpace(d.Dosis9)
		m.TanggalMulai9 = vTanggalMulai9
		m.TanggalAkhir9 = vTanggalAkhir9
		m.Indikasi9 = strings.TrimSpace(d.Indikasi9)
		m.Obat10 = strings.TrimSpace(d.Obat10)
		m.Sedian10 = strings.TrimSpace(d.Sedian10)
		m.ObatJkn10 = strings.TrimSpace(d.ObatJkn10)
		m.Batch10 = strings.TrimSpace(d.Batch10)
		m.Cara10 = strings.TrimSpace(d.Cara10)
		m.Dosis10 = strings.TrimSpace(d.Dosis10)
		m.TanggalMulai10 = vTanggalMulai10
		m.TanggalAkhir10 = vTanggalAkhir10
		m.Indikasi10 = strings.TrimSpace(d.Indikasi10)
		m.ReaksiSetelah = strings.TrimSpace(d.ReaksiSetelah)
		m.ReaksiSama = strings.TrimSpace(d.ReaksiSama)
		m.Naranjo1 = strings.TrimSpace(d.Naranjo1)
		m.NilaiNaranjo1 = strings.TrimSpace(d.NilaiNaranjo1)
		m.Naranjo2 = strings.TrimSpace(d.Naranjo2)
		m.NilaiNaranjo2 = strings.TrimSpace(d.NilaiNaranjo2)
		m.Naranjo3 = strings.TrimSpace(d.Naranjo3)
		m.NilaiNaranjo3 = strings.TrimSpace(d.NilaiNaranjo3)
		m.Naranjo4 = strings.TrimSpace(d.Naranjo4)
		m.NilaiNaranjo4 = strings.TrimSpace(d.NilaiNaranjo4)
		m.Naranjo5 = strings.TrimSpace(d.Naranjo5)
		m.NilaiNaranjo5 = strings.TrimSpace(d.NilaiNaranjo5)
		m.Naranjo6 = strings.TrimSpace(d.Naranjo6)
		m.NilaiNaranjo6 = strings.TrimSpace(d.NilaiNaranjo6)
		m.Naranjo7 = strings.TrimSpace(d.Naranjo7)
		m.NilaiNaranjo7 = strings.TrimSpace(d.NilaiNaranjo7)
		m.Naranjo8 = strings.TrimSpace(d.Naranjo8)
		m.NilaiNaranjo8 = strings.TrimSpace(d.NilaiNaranjo8)
		m.Naranjo9 = strings.TrimSpace(d.Naranjo9)
		m.NilaiNaranjo9 = strings.TrimSpace(d.NilaiNaranjo9)
		m.Naranjo10 = strings.TrimSpace(d.Naranjo10)
		m.NilaiNaranjo10 = strings.TrimSpace(d.NilaiNaranjo10)
		m.TotalNilaiNaranjo = strings.TrimSpace(d.TotalNilaiNaranjo)
		m.KategoriNaranjo = strings.TrimSpace(d.KategoriNaranjo)
		return nil
	},
	Refs: []Ref[model.MonitoringEfekSampingObat]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.MonitoringEfekSampingObat) any { return StrPtr(m.Nik) }},
		{Column: "kd_bangsal", Table: "bangsal", RefCol: "kd_bangsal", Label: "bangsal", Value: func(m *model.MonitoringEfekSampingObat) any { return StrPtr(m.KdBangsal) }},
	},
}
