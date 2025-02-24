package models

import (
	"errors"
	
	"gorm.io/gorm"
)


func AddUser(login, password string) error {
	user := User{Login: login, Password: password}
	if err := DB.Create(&user).Error; err != nil {
		return errors.New("ошибка при добавлении пользователя в базу данных")
	}
	return nil
}

func ValidateUser(login, password string) (bool, error) {
	var user User
	// Поиск пользователя по логину
	if err := DB.Where("login = ? AND password = ?", login, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UserExists(login string) (bool, error) {
	var user User
	// Поиск пользователя по логину
	if err := DB.Where("login = ?", login).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
