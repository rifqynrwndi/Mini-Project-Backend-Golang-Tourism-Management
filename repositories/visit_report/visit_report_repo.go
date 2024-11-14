package visit_report

import (
	"tourism-monitoring/entities"

	"gorm.io/gorm"
)

type VisitReportRepo struct {
	db *gorm.DB
}

func NewVisitReportRepo(db *gorm.DB) *VisitReportRepo {
	return &VisitReportRepo{db: db}
}

func (repo VisitReportRepo) GetAllVisitReports(limit int, offset int) ([]entities.VisitReport, error) {
	var visitReports []entities.VisitReport
	if err := repo.db.Preload("Wisatawan").Preload("ObjekWisata").Find(&visitReports).Error; err != nil {
		return nil, err
	}
	return visitReports, nil
}

func (repo VisitReportRepo) GetVisitReportByID(id int) (entities.VisitReport, error) {
	var visitReport entities.VisitReport
	if err := repo.db.Preload("Wisatawan").Preload("ObjekWisata").Where("id = ?", id).First(&visitReport).Error; err != nil {
		return entities.VisitReport{}, err
	}
	return visitReport, nil
}

func (repo VisitReportRepo) InsertVisitReport(visitReport entities.VisitReport) (entities.VisitReport, error) {
	if err := repo.db.Create(&visitReport).Error; err != nil {
		return entities.VisitReport{}, err
	}

	var reportWithRelations entities.VisitReport
	if err := repo.db.Preload("Wisatawan").Preload("ObjekWisata").First(&reportWithRelations, visitReport.ID).Error; err != nil {
		return entities.VisitReport{}, err
	}
	return reportWithRelations, nil
}

func (repo VisitReportRepo) UpdateVisitReport(id int, visitReport entities.VisitReport) (entities.VisitReport, error) {
	if err := repo.db.Model(&visitReport).Where("id = ?", id).Updates(visitReport).Error; err != nil {
		return entities.VisitReport{}, err
	}

	var reportWithRelations entities.VisitReport
	if err := repo.db.Preload("Wisatawan").Preload("ObjekWisata").First(&reportWithRelations, visitReport.ID).Error; err != nil {
		return entities.VisitReport{}, err
	}
	return reportWithRelations, nil
}

func (repo VisitReportRepo) DeleteVisitReport(id int) error {
	var visitReport entities.VisitReport
	if err := repo.db.Where("id = ?", id).Delete(&visitReport).Error; err != nil {
		return err
	}
	return nil
}

func (repo VisitReportRepo) GetTotalVisitReportsCount() (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.VisitReport{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo VisitReportRepo) GetAverageVisitsForPlace(placeID int) (float64, error) {
    var count int64
    if err := repo.db.Model(&entities.VisitReport{}).Where("objek_wisata_id = ?", placeID).Count(&count).Error; err != nil {
        return 0, err
    }

    return float64(count), nil
}
