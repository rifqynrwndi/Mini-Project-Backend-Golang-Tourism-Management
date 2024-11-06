package services

import "tourism-monitoring/entities"

type AuthInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
}
