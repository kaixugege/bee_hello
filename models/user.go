package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int `gorm:"default:1"` // 0 管理员 1正常用户
}

func (db *DB) QueryUserByEmailAndPassword(email, password string) (User, error) {
	var user User
	if err := db.db.Model(
		&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func QueryUserByName(name string) (user User, err error) {
	return user, db.Where("name  = ?", name).Take(&user).Error
}
