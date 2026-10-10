package jabatan

// Jabatan tabel `jabatan` (jabatan).
type Jabatan struct {
	KdJbtn string  `gorm:"column:kd_jbtn;primaryKey;type:char(4);not null" json:"kd_jbtn"`
	NmJbtn *string `gorm:"column:nm_jbtn;type:varchar(25)" json:"nm_jbtn"`
}

func (Jabatan) TableName() string {
	return "jabatan"
}
