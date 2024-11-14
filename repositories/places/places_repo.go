package places

import (
	"tourism-monitoring/entities"

	"gorm.io/gorm"
)

type PlacesRepo struct {
	db *gorm.DB
}

func NewPlacesRepo(db *gorm.DB) *PlacesRepo {
	return &PlacesRepo{db: db}
}

func (repo PlacesRepo) GetAllPlaces() ([]entities.Place, error) {
	var places []entities.Place
	if err := repo.db.Find(&places).Error; err != nil {
		return nil, err
	}
	return places, nil
}

func (repo PlacesRepo) GetPlaceByID(id int) (entities.Place, error) {
	var places []entities.Place
	if err := repo.db.Where("id = ?", id).Find(&places).Error; err != nil {
		return entities.Place{}, err
	}

	if len(places) == 0 {
		return entities.Place{}, gorm.ErrRecordNotFound
	}
	return places[0], nil
}

func (repo PlacesRepo) InsertPlace(place entities.Place) (entities.Place, error) {
	if err := repo.db.Create(&place).Error; err != nil {
		return entities.Place{}, err
	}
	return place, nil
}

func (repo PlacesRepo) UpdatePlace(id int, place entities.Place) (entities.Place, error) {
	if err := repo.db.Model(&place).Where("id = ?", id).Updates(place).Error; err != nil {
		return entities.Place{}, err
	}
	return place, nil
}

func (repo PlacesRepo) DeletePlace(id int) error {
	if err := repo.db.Where("id = ?", id).Delete(&entities.Place{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo PlacesRepo) GetTotalPlacesCount() (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.Place{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}


