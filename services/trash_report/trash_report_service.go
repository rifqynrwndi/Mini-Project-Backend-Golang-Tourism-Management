package trash_report

import (
	"tourism-monitoring/entities"
	"tourism-monitoring/repositories/trash_report"
)

type TrashReportService struct {
	trashReportRepo trash_report.TrashReportRepoInterface
}

func NewTrashReportService(repo trash_report.TrashReportRepoInterface) *TrashReportService {
	return &TrashReportService{trashReportRepo: repo}
}

func (service TrashReportService) GetTrashReportByPlaceID(id int) ([]entities.TrashReport, error) {
	return service.trashReportRepo.GetTrashReportByPlaceID(id)
}

func (service TrashReportService) GetTrashReportByID(id int) (entities.TrashReport, error) {
	return service.trashReportRepo.GetTrashReportByID(id)
}

func (service TrashReportService) InsertTrashReport(trashReport entities.TrashReport) (entities.TrashReport, error) {
	return service.trashReportRepo.InsertTrashReport(trashReport)
}

func (service TrashReportService) UpdateTrashReport(id int, trashReport entities.TrashReport) (entities.TrashReport, error) {
	return service.trashReportRepo.UpdateTrashReport(id, trashReport)
}

func (service TrashReportService) DeleteTrashReport(id int) error {
	return service.trashReportRepo.DeleteTrashReport(id)
}

func (service TrashReportService) GetTotalTrashReportsCount() (int64, error) {
	return service.trashReportRepo.GetTotalTrashReportsCount()
}
