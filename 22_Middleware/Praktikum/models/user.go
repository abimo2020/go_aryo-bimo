package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `gorm:"-"`
	Blogs    []Blog `json: "blogs"`
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where("user_id = ?", u.ID).Delete(&Blog{})
	return
}
