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

func (repo VisitReportRepo) GetAllVisitReports() ([]entities.VisitReport, error) {
	var visitReports []entities.VisitReport
	if err := repo.db.Find(&visitReports).Error; err != nil {
		return nil, err
	}
	return visitReports, nil
}

func (repo VisitReportRepo) GetVisitReportByID(id int) (entities.VisitReport, error) {
	var visitReports []entities.VisitReport
	if err := repo.db.Where("id = ?", id).Find(&visitReports).Error; err != nil {
		return entities.VisitReport{}, err
	}

	if len(visitReports) == 0 {
        return entities.VisitReport{}, gorm.ErrRecordNotFound
    }
	return visitReports[0], nil
}

func (repo VisitReportRepo) InsertVisitReport(visitReport entities.VisitReport) (entities.VisitReport, error) {
	if err := repo.db.Create(&visitReport).Error; err != nil {
		return entities.VisitReport{}, err
	}
	return visitReport, nil
}

func (repo VisitReportRepo) UpdateVisitReport(id int, visitReport entities.VisitReport) (entities.VisitReport, error) {
    if err := repo.db.Model(&visitReport).Where("id = ?", id).Updates(visitReport).Error; err != nil {
        return entities.VisitReport{}, err
    }
    return visitReport, nil
}

func (repo VisitReportRepo) DeleteVisitReport(id int) error {
	var visitReport entities.VisitReport
	if err := repo.db.Where("id = ?", id).Delete(&visitReport).Error; err != nil {
		return err
	}
	return nil
}


