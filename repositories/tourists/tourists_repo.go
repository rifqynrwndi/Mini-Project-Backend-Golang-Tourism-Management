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
	if err := repo.db.Select("id","nama", "usia", "asal", "jenis_kelamin", "tipe_wisatawan").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo TouristsRepo) GetTouristByID(id int) (entities.User, error) {
	var users []entities.User
	if err := repo.db.Select("id", "nama", "usia", "asal", "jenis_kelamin", "tipe_wisatawan").Where("id = ?", id).Find(&users).Error; err != nil {
		return entities.User{}, err
	}

	if len(users) == 0 {
		return entities.User{}, gorm.ErrRecordNotFound
	}
	return users[0], nil
}

func (repo TouristsRepo) InsertTourist(user entities.User) (entities.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (repo TouristsRepo) UpdateTourist(id int, user entities.User) (entities.User, error) {
	if err := repo.db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (repo TouristsRepo) DeleteTourist(id int) error {
	if err := repo.db.Where("id = ?", id).Delete(&entities.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo TouristsRepo) GetTotalTouristsCount() (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

