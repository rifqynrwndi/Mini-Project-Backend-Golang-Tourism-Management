package tourists

import (
	"tourism-monitoring/entities"

	"gorm.io/gorm"
)

type TouristsRepo struct {
	db *gorm.DB
}

func NewTouristsRepo(db *gorm.DB) *TouristsRepo {
	return &TouristsRepo{db: db}
}

func (repo TouristsRepo) GetAllTourists() ([]entities.User, error) {
	var users []entities.User
	if err := repo.db.Select("nama", "usia", "asal", "jenis_kelamin", "tipe_wisatawan").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
