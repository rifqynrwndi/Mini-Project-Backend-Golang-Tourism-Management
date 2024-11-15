package trash_report

import (
	"gorm.io/gorm"
	"tourism-monitoring/entities"
)

type TrashReportRepo struct {
	db *gorm.DB
}

func NewTrashReportRepo(db *gorm.DB) *TrashReportRepo {
	return &TrashReportRepo{db: db}
}

func (repo TrashReportRepo) GetTrashReportByPlaceID(id int) ([]entities.TrashReport, error) {
	var trashReports []entities.TrashReport
	if err := repo.db.Preload("ObjekWisata").Where("objek_wisata_id = ?", id).Find(&trashReports).Error; err != nil {
		return nil, err
	}
	return trashReports, nil
}

func (repo TrashReportRepo) GetTrashReportByID(id int) (entities.TrashReport, error) {
	var trashReports []entities.TrashReport
	if err := repo.db.Preload("ObjekWisata").Where("id = ?", id).Find(&trashReports).Error; err != nil {
		return entities.TrashReport{}, err
	}

	if len(trashReports) == 0 {
		return entities.TrashReport{}, gorm.ErrRecordNotFound
	}
	return trashReports[0], nil
}

func (repo TrashReportRepo) InsertTrashReport(trashReport entities.TrashReport) (entities.TrashReport, error) {
	if err := repo.db.Create(&trashReport).Error; err != nil {
		return entities.TrashReport{}, err
	}
	
	var reportWithRelations entities.TrashReport
    if err := repo.db.Preload("ObjekWisata").First(&reportWithRelations, trashReport.ID).Error; err != nil {
        return entities.TrashReport{}, err
    }
    return reportWithRelations, nil
}

func (repo TrashReportRepo) UpdateTrashReport(id int, trashReport entities.TrashReport) (entities.TrashReport, error) {
	if err := repo.db.Model(&trashReport).Where("id = ?", id).Updates(trashReport).Error; err != nil {
		return entities.TrashReport{}, err
	}
	var reportWithRelations entities.TrashReport
    if err := repo.db.Preload("ObjekWisata").First(&reportWithRelations, trashReport.ID).Error; err != nil {
        return entities.TrashReport{}, err
    }
    return reportWithRelations, nil
}

func (repo TrashReportRepo) DeleteTrashReport(id int) error {
	var trashReport entities.TrashReport
	if err := repo.db.Where("id = ?", id).Delete(&trashReport).Error; err != nil {
		return err
	}
	return nil
}

func (repo TrashReportRepo) GetTotalTrashReportsCount() (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.TrashReport{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}