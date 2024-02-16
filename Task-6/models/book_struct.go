package models

import "gorm.io/gorm"

// Book struct as GORM Model
type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn" gorm:"primaryKey"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
}
