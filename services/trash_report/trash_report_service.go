package trash_report

import (
	"tourism-monitoring/entities"
	"tourism-monitoring/repositories/trash_report"
)

type TrashReportService struct {
	repo *trash_report.TrashReportRepo
}

func NewTrashReportService(repo *trash_report.TrashReportRepo) *TrashReportService {
	return &TrashReportService{repo: repo}
}

func (service TrashReportService) GetTrashReportByPlaceID(id int) ([]entities.TrashReport, error) {
	return service.repo.GetTrashReportByPlaceID(id)
}

func (service TrashReportService) GetTrashReportByID(id int) (entities.TrashReport, error) {
	return service.repo.GetTrashReportByID(id)
}

func (service TrashReportService) InsertTrashReport(trashReport entities.TrashReport) (entities.TrashReport, error) {
	return service.repo.InsertTrashReport(trashReport)
}

func (service TrashReportService) UpdateTrashReport(id int, trashReport entities.TrashReport) (entities.TrashReport, error) {
	return service.repo.UpdateTrashReport(id, trashReport)
}

func (service TrashReportService) DeleteTrashReport(id int) error {
	return service.repo.DeleteTrashReport(id)
}
