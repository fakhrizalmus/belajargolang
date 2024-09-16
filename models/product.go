package models

type Product struct {
	Id         int64  `gorm:"PrimaryKey" json:"id"`
	NamaProduk string `gorm:"type:varchar(300)" json:"namaproduk" binding:"required"`
	Deskripsi  string `gorm:"type:text" json:"deskripsi" binding:"required"`
}
