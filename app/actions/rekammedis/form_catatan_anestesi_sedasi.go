package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanAnestesiSedasi catatan anastesi sedasi (RMCatatanAnastesiSedasi).
var FormCatatanAnestesiSedasi = &Form[model.CatatanAnestesiSedasi, request.CatatanAnestesiSedasiData]{
	Slug:  "catatan-anestesi-sedasi",
	Label: "catatan anastesi sedasi",
	Spec: repo.Spec{
		Table:  "catatan_anestesi_sedasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter_bedah", "kd_dokter_anestesi", "nip_perawat_ok", "nip_perawat_anestesi"},
	},
	SetKey: func(m *model.CatatanAnestesiSedasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.CatatanAnestesiSedasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.CatatanAnestesiSedasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.CatatanAnestesiSedasi) []string {
		return []string{m.KdDokterBedah, m.KdDokterAnestesi, derefStr(m.NipPerawatOk), derefStr(m.NipPerawatAnestesi)}
	},
	Fill: func(m *model.CatatanAnestesiSedasi, d request.CatatanAnestesiSedasiData) error {
		m.KdDokterBedah = strings.TrimSpace(d.KdDokterBedah)
		m.KdDokterAnestesi = strings.TrimSpace(d.KdDokterAnestesi)
		m.DiagnosaPreBedah = strings.TrimSpace(d.DiagnosaPreBedah)
		m.TindakanJenisPembedahan = strings.TrimSpace(d.TindakanJenisPembedahan)
		m.DiagnosaPascaBedah = strings.TrimSpace(d.DiagnosaPascaBedah)
		m.PreInduksiJam = support.Nullable(d.PreInduksiJam)
		m.PreInduksiKesadaran = support.Nullable(d.PreInduksiKesadaran)
		m.PreInduksiTd = support.Nullable(d.PreInduksiTd)
		m.PreInduksiNadi = support.Nullable(d.PreInduksiNadi)
		m.PreInduksiRr = support.Nullable(d.PreInduksiRr)
		m.PreInduksiSuhu = support.Nullable(d.PreInduksiSuhu)
		m.PreInduksiO2 = support.Nullable(d.PreInduksiO2)
		m.PreInduksiTb = support.Nullable(d.PreInduksiTb)
		m.PreInduksiBb = support.Nullable(d.PreInduksiBb)
		m.PreInduksiRhesus = support.Nullable(d.PreInduksiRhesus)
		m.PreInduksiHb = support.Nullable(d.PreInduksiHb)
		m.PreInduksiHt = support.Nullable(d.PreInduksiHt)
		m.PreInduksiLeko = support.Nullable(d.PreInduksiLeko)
		m.PreInduksiTrombo = support.Nullable(d.PreInduksiTrombo)
		m.PreInduksiBtct = support.Nullable(d.PreInduksiBtct)
		m.PreInduksiGds = support.Nullable(d.PreInduksiGds)
		m.PreInduksiLainlain = support.Nullable(d.PreInduksiLainlain)
		m.TeknikAlatHiopotensi = support.Nullable(d.TeknikAlatHiopotensi)
		m.TeknikAlatTci = support.Nullable(d.TeknikAlatTci)
		m.TeknikAlatCpb = support.Nullable(d.TeknikAlatCpb)
		m.TeknikAlatVentilasi = support.Nullable(d.TeknikAlatVentilasi)
		m.TeknikAlatBroncoskopy = support.Nullable(d.TeknikAlatBroncoskopy)
		m.TeknikAlatGlidescopi = support.Nullable(d.TeknikAlatGlidescopi)
		m.TeknikAlatUsg = support.Nullable(d.TeknikAlatUsg)
		m.TeknikAlatStimulatorSaraf = support.Nullable(d.TeknikAlatStimulatorSaraf)
		m.TeknikAlatLainlain = support.Nullable(d.TeknikAlatLainlain)
		m.MonitoringEkg = support.Nullable(d.MonitoringEkg)
		m.MonitoringEkgKeterangan = support.Nullable(d.MonitoringEkgKeterangan)
		m.MonitoringArteri = support.Nullable(d.MonitoringArteri)
		m.MonitoringArteriKeterangan = support.Nullable(d.MonitoringArteriKeterangan)
		m.MonitoringCvp = support.Nullable(d.MonitoringCvp)
		m.MonitoringCvpKeterangan = support.Nullable(d.MonitoringCvpKeterangan)
		m.MonitoringEtco = support.Nullable(d.MonitoringEtco)
		m.MonitoringStetoskop = support.Nullable(d.MonitoringStetoskop)
		m.MonitoringNibp = support.Nullable(d.MonitoringNibp)
		m.MonitoringNgt = support.Nullable(d.MonitoringNgt)
		m.MonitoringBis = support.Nullable(d.MonitoringBis)
		m.MonitoringCathAPulmo = support.Nullable(d.MonitoringCathAPulmo)
		m.MonitoringSpo2 = support.Nullable(d.MonitoringSpo2)
		m.MonitoringKateter = support.Nullable(d.MonitoringKateter)
		m.MonitoringTemp = support.Nullable(d.MonitoringTemp)
		m.MonitoringLainlain = support.Nullable(d.MonitoringLainlain)
		m.StatusFisikAsa = support.Nullable(d.StatusFisikAsa)
		m.StatusFisikAlergi = strings.TrimSpace(d.StatusFisikAlergi)
		m.StatusFisikAlergiKeterangan = support.Nullable(d.StatusFisikAlergiKeterangan)
		m.StatusFisikPenyulitSedasi = support.Nullable(d.StatusFisikPenyulitSedasi)
		m.PerencanaanLanjut = support.Nullable(d.PerencanaanLanjut)
		m.PerencanaanLanjutSedasi = support.Nullable(d.PerencanaanLanjutSedasi)
		m.PerencanaanLanjutSedasiKeterangan = support.Nullable(d.PerencanaanLanjutSedasiKeterangan)
		m.PerencanaanLanjutSpinal = support.Nullable(d.PerencanaanLanjutSpinal)
		m.PerencanaanLanjutAnestesiUmum = support.Nullable(d.PerencanaanLanjutAnestesiUmum)
		m.PerencanaanLanjutAnestesiUmumKeterangan = support.Nullable(d.PerencanaanLanjutAnestesiUmumKeterangan)
		m.PerencanaanLanjutBlokPerifer = support.Nullable(d.PerencanaanLanjutBlokPerifer)
		m.PerencanaanLanjutBlokPeriferKeterangan = support.Nullable(d.PerencanaanLanjutBlokPeriferKeterangan)
		m.PerencanaanLanjutEpidural = support.Nullable(d.PerencanaanLanjutEpidural)
		m.PerencanaanBatal = support.Nullable(d.PerencanaanBatal)
		m.PerencanaanBatalAlasan = support.Nullable(d.PerencanaanBatalAlasan)
		m.NipPerawatOk = support.Nullable(d.NipPerawatOk)
		m.NipPerawatAnestesi = support.Nullable(d.NipPerawatAnestesi)
		return nil
	},
	Refs: []Ref[model.CatatanAnestesiSedasi]{
		{Column: "kd_dokter_bedah", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.CatatanAnestesiSedasi) any { return Str(m.KdDokterBedah) }},
		{Column: "kd_dokter_anestesi", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.CatatanAnestesiSedasi) any { return Str(m.KdDokterAnestesi) }},
		{Column: "nip_perawat_ok", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanAnestesiSedasi) any { return StrPtr(m.NipPerawatOk) }},
		{Column: "nip_perawat_anestesi", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanAnestesiSedasi) any { return StrPtr(m.NipPerawatAnestesi) }},
	},
}
