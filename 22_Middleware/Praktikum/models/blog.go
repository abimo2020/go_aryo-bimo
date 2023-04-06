package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Judul  string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
	UserID uint   `json:"user_id" form:"user_id"`
}
