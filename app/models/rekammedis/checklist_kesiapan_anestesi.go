package rekammedis

import "time"

// ChecklistKesiapanAnestesi tabel `checklist_kesiapan_anestesi` (checklist kesiapan anestesi, RMChecklistKesiapanAnestesi).
type ChecklistKesiapanAnestesi struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal           *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Nip               *string    `gorm:"column:nip" json:"nip"`
	KdDokter          *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
	Tindakan          string     `gorm:"column:tindakan" json:"tindakan"`
	TeknikAnestesi    string     `gorm:"column:teknik_anestesi" json:"teknik_anestesi"`
	Listrik1          *string    `gorm:"column:listrik1" json:"listrik1"`
	Listrik2          *string    `gorm:"column:listrik2" json:"listrik2"`
	Listrik3          *string    `gorm:"column:listrik3" json:"listrik3"`
	Listrik4          *string    `gorm:"column:listrik4" json:"listrik4"`
	Gasmedis1         *string    `gorm:"column:gasmedis1" json:"gasmedis1"`
	Gasmedis2         *string    `gorm:"column:gasmedis2" json:"gasmedis2"`
	Gasmedis3         *string    `gorm:"column:gasmedis3" json:"gasmedis3"`
	Gasmedis4         *string    `gorm:"column:gasmedis4" json:"gasmedis4"`
	Gasmedis5         *string    `gorm:"column:gasmedis5" json:"gasmedis5"`
	Gasmedis6         *string    `gorm:"column:gasmedis6" json:"gasmedis6"`
	Mesinanes1        *string    `gorm:"column:mesinanes1" json:"mesinanes1"`
	Mesinanes2        *string    `gorm:"column:mesinanes2" json:"mesinanes2"`
	Mesinanes3        *string    `gorm:"column:mesinanes3" json:"mesinanes3"`
	Mesinanes4        *string    `gorm:"column:mesinanes4" json:"mesinanes4"`
	Mesinanes5        *string    `gorm:"column:mesinanes5" json:"mesinanes5"`
	Jalannapas1       *string    `gorm:"column:jalannapas1" json:"jalannapas1"`
	Jalannapas2       *string    `gorm:"column:jalannapas2" json:"jalannapas2"`
	Jalannapas3       *string    `gorm:"column:jalannapas3" json:"jalannapas3"`
	Jalannapas4       *string    `gorm:"column:jalannapas4" json:"jalannapas4"`
	Jalannapas5       *string    `gorm:"column:jalannapas5" json:"jalannapas5"`
	Jalannapas6       *string    `gorm:"column:jalannapas6" json:"jalannapas6"`
	Jalannapas7       *string    `gorm:"column:jalannapas7" json:"jalannapas7"`
	Jalannapas8       *string    `gorm:"column:jalannapas8" json:"jalannapas8"`
	Jalannapas9       *string    `gorm:"column:jalannapas9" json:"jalannapas9"`
	Lainlain1         *string    `gorm:"column:lainlain1" json:"lainlain1"`
	Lainlain2         *string    `gorm:"column:lainlain2" json:"lainlain2"`
	Lainlain3         *string    `gorm:"column:lainlain3" json:"lainlain3"`
	Lainlain4         *string    `gorm:"column:lainlain4" json:"lainlain4"`
	Lainlain5         *string    `gorm:"column:lainlain5" json:"lainlain5"`
	Lainlain6         *string    `gorm:"column:lainlain6" json:"lainlain6"`
	Lainlain7         *string    `gorm:"column:lainlain7" json:"lainlain7"`
	Lainlain8         *string    `gorm:"column:lainlain8" json:"lainlain8"`
	Obatobat1         *string    `gorm:"column:obatobat1" json:"obatobat1"`
	Obatobat2         *string    `gorm:"column:obatobat2" json:"obatobat2"`
	Obatobat3         *string    `gorm:"column:obatobat3" json:"obatobat3"`
	Obatobat4         *string    `gorm:"column:obatobat4" json:"obatobat4"`
	Obatobat5         *string    `gorm:"column:obatobat5" json:"obatobat5"`
	Obatobat6         *string    `gorm:"column:obatobat6" json:"obatobat6"`
	KeteranganLainnya *string    `gorm:"column:keterangan_lainnya" json:"keterangan_lainnya"`
}

func (ChecklistKesiapanAnestesi) TableName() string {
	return "checklist_kesiapan_anestesi"
}
