package Entities

import "gorm.io/gorm"

func Init(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
