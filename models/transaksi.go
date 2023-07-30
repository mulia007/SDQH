package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Id             int           `gorm:"primary_key;type:int" json:"id"`
	Id_Waqaf       int           `gorm:"type:int" json:"id_waqaf"`
	Waqaf_Sedeqah  Waqaf_Sedeqah `gorm:"foreignKey:id_waqaf"`
	Id_Users       int           `gorm:"type:int" json:"id_users"`
	Users          Users         `gorm:"foreignKey:id_users"`
	Nominal        float32       `gorm:"type:decimal(18,2)" json:"nominal"`
	Bukti_Transfer []byte        `gorm:"type:blob" json:"bukti_transfer"`
	Tanggal        string        `gorm:"type:datetime" json:"tanggal"`
}
