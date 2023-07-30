package models

import "gorm.io/gorm"

type Waqaf_Sedeqah struct {
	gorm.Model
	Id               int    `gorm:"primary_key;type:int" json:"id"`
	Id_Admin         int    `gorm:"type:int" json:"id_admin"`
	Admin            Admin  `gorm:"foreignKey:id_admin"`
	Judul_Wakaf      string `gorm:"type:varchar(100)" json:"judul_wakaf"`
	Tipe             string `gorm:"type:varchar(10)" json:"tipe"`
	Tanggal_Mulai    string `gorm:"type:datetime" json:"tanggal_mulai"`
	Tanggal_Berakhir string `gorm:"type:datetime" json:"tanggal_berakhir"`
	Nama_Bank        string `gorm:"type:varchar(100)" json:"nama_bank"`
	Rekening_Bank    string `gorm:"type:varchar(100)" json:"rekening_bank"`
}
