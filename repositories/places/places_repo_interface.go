package places

import (
	"tourism-monitoring/entities"
)

type PlacesRepoInterface interface {
	GetAllPlaces() ([]entities.Place, error)
	GetPlaceByID(id int) (entities.Place, error)
	InsertPlace(place entities.Place) (entities.Place, error)
	UpdatePlace(id int, place entities.Place) (entities.Place, error)
	DeletePlace(id int) error
	GetTotalPlacesCount() (int64, error)
}

