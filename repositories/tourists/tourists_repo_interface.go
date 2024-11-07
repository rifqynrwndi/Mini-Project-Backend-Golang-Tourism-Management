package tourists

import "tourism-monitoring/entities"

type TouristsRepoInterface interface {
	GetAllTourists() ([]entities.User, error)
	GetTouristByID(id int) (entities.User, error)
	InsertTourist(user entities.User) (entities.User, error)
	UpdateTourist(id int, user entities.User) (entities.User, error)
	DeleteTourist(id int) error
}
