package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKesiapanAnestesi checklist kesiapan anestesi (RMChecklistKesiapanAnestesi).
var FormChecklistKesiapanAnestesi = &Form[model.ChecklistKesiapanAnestesi, request.ChecklistKesiapanAnestesiData]{
	Slug:  "checklist-kesiapan-anestesi",
	Label: "checklist kesiapan anestesi",
	Spec: repo.Spec{
		Table:  "checklist_kesiapan_anestesi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip", "kd_dokter"},
	},
	SetKey: func(m *model.ChecklistKesiapanAnestesi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKesiapanAnestesi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu: func(m *model.ChecklistKesiapanAnestesi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKesiapanAnestesi) []string {
		return []string{derefStr(m.Nip), derefStr(m.KdDokter)}
	},
	Fill: func(m *model.ChecklistKesiapanAnestesi, d request.ChecklistKesiapanAnestesiData) error {
		m.Nip = support.Nullable(d.Nip)
		m.KdDokter = support.Nullable(d.KdDokter)
		m.Tindakan = strings.TrimSpace(d.Tindakan)
		m.TeknikAnestesi = strings.TrimSpace(d.TeknikAnestesi)
		m.Listrik1 = support.Nullable(d.Listrik1)
		m.Listrik2 = support.Nullable(d.Listrik2)
		m.Listrik3 = support.Nullable(d.Listrik3)
		m.Listrik4 = support.Nullable(d.Listrik4)
		m.Gasmedis1 = support.Nullable(d.Gasmedis1)
		m.Gasmedis2 = support.Nullable(d.Gasmedis2)
		m.Gasmedis3 = support.Nullable(d.Gasmedis3)
		m.Gasmedis4 = support.Nullable(d.Gasmedis4)
		m.Gasmedis5 = support.Nullable(d.Gasmedis5)
		m.Gasmedis6 = support.Nullable(d.Gasmedis6)
		m.Mesinanes1 = support.Nullable(d.Mesinanes1)
		m.Mesinanes2 = support.Nullable(d.Mesinanes2)
		m.Mesinanes3 = support.Nullable(d.Mesinanes3)
		m.Mesinanes4 = support.Nullable(d.Mesinanes4)
		m.Mesinanes5 = support.Nullable(d.Mesinanes5)
		m.Jalannapas1 = support.Nullable(d.Jalannapas1)
		m.Jalannapas2 = support.Nullable(d.Jalannapas2)
		m.Jalannapas3 = support.Nullable(d.Jalannapas3)
		m.Jalannapas4 = support.Nullable(d.Jalannapas4)
		m.Jalannapas5 = support.Nullable(d.Jalannapas5)
		m.Jalannapas6 = support.Nullable(d.Jalannapas6)
		m.Jalannapas7 = support.Nullable(d.Jalannapas7)
		m.Jalannapas8 = support.Nullable(d.Jalannapas8)
		m.Jalannapas9 = support.Nullable(d.Jalannapas9)
		m.Lainlain1 = support.Nullable(d.Lainlain1)
		m.Lainlain2 = support.Nullable(d.Lainlain2)
		m.Lainlain3 = support.Nullable(d.Lainlain3)
		m.Lainlain4 = support.Nullable(d.Lainlain4)
		m.Lainlain5 = support.Nullable(d.Lainlain5)
		m.Lainlain6 = support.Nullable(d.Lainlain6)
		m.Lainlain7 = support.Nullable(d.Lainlain7)
		m.Lainlain8 = support.Nullable(d.Lainlain8)
		m.Obatobat1 = support.Nullable(d.Obatobat1)
		m.Obatobat2 = support.Nullable(d.Obatobat2)
		m.Obatobat3 = support.Nullable(d.Obatobat3)
		m.Obatobat4 = support.Nullable(d.Obatobat4)
		m.Obatobat5 = support.Nullable(d.Obatobat5)
		m.Obatobat6 = support.Nullable(d.Obatobat6)
		m.KeteranganLainnya = support.Nullable(d.KeteranganLainnya)
		return nil
	},
	Refs: []Ref[model.ChecklistKesiapanAnestesi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.ChecklistKesiapanAnestesi) any { return StrPtr(m.Nip) }},
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.ChecklistKesiapanAnestesi) any { return StrPtr(m.KdDokter) }},
	},
}
