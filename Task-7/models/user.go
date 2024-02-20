package models

import "gorm.io/gorm"

// GORM Model for User
type User struct {
	gorm.Model
	Name string `gorm:"name" primaryKey`
	Pass string `gorm:"pass" `
}
