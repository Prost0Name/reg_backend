package model

// Здесь можно добавить дополнительные методы для работы с пользователями, если это необходимо

type User struct {
	Login    string `json:"login" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

// CreateUser создает нового пользователя и сохраняет его в базе данных
func CreateUser(login string, password string) error {
	user := DBUser{
		Login:    login,
		Password: password,
	}
	return DB.Create(&user).Error
}
