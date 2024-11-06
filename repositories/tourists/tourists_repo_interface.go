package tourists

import "tourism-monitoring/entities"

type TouristsRepoInterface interface {
	GetAllTourists() ([]entities.User, error)
}
