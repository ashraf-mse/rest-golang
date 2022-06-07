package main

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Des   string `json:"des" binding:"required"`
}
