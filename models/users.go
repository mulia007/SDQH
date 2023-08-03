package models

type Users struct {
	Id       int    `gorm:"primary_key;type:int" json:"id"`
	Username string `gorm:"type:nvarchar(50)" json:"username"`
	Password string `gorm:"type:nvarchar(200)" json:"password"`
}
