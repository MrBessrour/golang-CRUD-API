package main

import (
	"gorm.io/gorm"
)

//this model represent a database table
type Post struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);" json:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100);" json:"des" binding:"required"`
	Status string `gorm:"type:varchar(200);" json:"status"`
}
