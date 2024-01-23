package repository

import (
	todo "github.com/aWatLove/togo-list/pkg/model"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) { // TODO: сделать нормальную обработку ошибок, что пользователя такого нет
	var user todo.User
	result := r.db.Where("username = ? and password_hash = ?", username, password).First(&user)
	return user, result.Error
}
