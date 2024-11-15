package tourists

import (
	"tourism-monitoring/entities"
	"tourism-monitoring/repositories/tourists"
)

type TouristsService struct {
	touristsRepo tourists.TouristsRepoInterface
}

func NewTouristsService(repo tourists.TouristsRepoInterface) *TouristsService {
	return &TouristsService{touristsRepo: repo}
}

func (service TouristsService) GetAllTourists() ([]entities.User, error) {
	return service.touristsRepo.GetAllTourists()
}

func (service TouristsService) GetTouristByID(id int) (entities.User, error) {
	return service.touristsRepo.GetTouristByID(id)
}

func (service TouristsService) InsertTourist(user entities.User) (entities.User, error) {
	return service.touristsRepo.InsertTourist(user)
}

func (service TouristsService) UpdateTourist(id int, user entities.User) (entities.User, error) {
	return service.touristsRepo.UpdateTourist(id, user)
}

func (service TouristsService) DeleteTourist(id int) error {
	return service.touristsRepo.DeleteTourist(id)
}

func (service TouristsService) GetTotalTouristsCount() (int64, error) {
	return service.touristsRepo.GetTotalTouristsCount()
}
