package model

import (
	"errors"
	"github.com/blog/v1/utils"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"not null " json:"username"`
	Password string `gorm:"not null " json:"-"`
	Avatar   string `gorm:"null " json:"avatar"`
	Phone    string `gorm:"null " json:"phone"`
	Name     string `gorm:"null " json:"name"`
	Address  string `gorm:"null " json:"address"`
	Birthday string `gorm:"null " json:"birthday"`
}

func (u *Users) BeforeCreate(db *gorm.DB) error {
	if len(u.Password) < 6 {
		return errors.New("密码太简单了")
	}
	//对密码进行加密存储
	u.Password = utils.Password(u.Password)
	return nil
}
