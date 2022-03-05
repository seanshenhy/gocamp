package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(32);not null" json:"password"`
}

func (User) TableName() string {
	return "user"
}

func GetUserName(id int) (*User, error) {
	var user User
	gdb := db.Select("username").Where("id = ?", id).First(&user)
	if gdb.Error != nil {
		return &user, errors.Wrap(gdb.Error, "model：GetUserName 用户查询异常")
	}
	return &user, nil
}
