package trash_report

import "tourism-monitoring/entities"

type TrashReportRepoInterface interface {
	GetTrashReportByPlaceID(id int) ([]entities.TrashReport, error)
	GetTrashReportByID(id int) (entities.TrashReport, error)
	InsertTrashReport(trashReport entities.TrashReport) (entities.TrashReport, error)
	UpdateTrashReport(id int, trashReport entities.TrashReport) (entities.TrashReport, error)
	DeleteTrashReport(id int) error
	GetTotalTrashReportCount() (int64, error)
}