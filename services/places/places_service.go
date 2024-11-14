package places

import (
	"tourism-monitoring/entities"
	"tourism-monitoring/repositories/places"
)

type PlacesService struct {
	placesRepo *places.PlacesRepo
}

func NewPlacesService(repo *places.PlacesRepo) *PlacesService {
	return &PlacesService{placesRepo: repo}
}

func (service PlacesService) GetAllPlaces() ([]entities.Place, error) {
	return service.placesRepo.GetAllPlaces()
}

func (service PlacesService) GetPlaceByID(id int) (entities.Place, error) {
	return service.placesRepo.GetPlaceByID(id)
}

func (service PlacesService) InsertPlace(place entities.Place) (entities.Place, error) {
	return service.placesRepo.InsertPlace(place)
}

func (service PlacesService) UpdatePlace(id int, place entities.Place) (entities.Place, error) {
	return service.placesRepo.UpdatePlace(id, place)
}

func (service PlacesService) DeletePlace(id int) error {
	return service.placesRepo.DeletePlace(id)
}

func (service PlacesService) GetTotalPlacesCount() (int64, error) {
	return service.placesRepo.GetTotalPlacesCount()
}