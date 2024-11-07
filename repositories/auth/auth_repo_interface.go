package auth

import "tourism-monitoring/entities"

type AuthRepoInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetLastUserID() (int, error)
}
