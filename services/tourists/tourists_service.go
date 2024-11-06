package tourists

import (
	"tourism-monitoring/entities"
	"tourism-monitoring/repositories/tourists"
)

type TouristsService struct {
	touristsRepo *tourists.TouristsRepo
}

func NewTouristsService(repo *tourists.TouristsRepo) *TouristsService {
	return &TouristsService{touristsRepo: repo}
}

func (service TouristsService) GetAllTourists() ([]entities.User, error) {
	return service.touristsRepo.GetAllTourists()
}
