package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID       uint `gorm:"primarykey"`
	CreateAt string
	UserName string `gorm:"unique"`
	Password string
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
