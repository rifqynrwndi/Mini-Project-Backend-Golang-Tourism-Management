package auth

import "tourism-monitoring/entities"

type User struct {
	ID       int `gorm:"primaryKey"`
	Nama     string
	Email    string
	Password string
}

func FromEntities(user entities.User) User {
	return User{
		ID:       user.ID,
		Nama:     user.Nama,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:       user.ID,
		Nama:     user.Nama,
		Email:    user.Email,
		Password: user.Password,
	}
}
