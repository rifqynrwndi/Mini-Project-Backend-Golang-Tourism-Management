package visit_report

import "tourism-monitoring/entities"

type VisitReportRepoInterface interface {
	GetAllVisitReports() ([]entities.VisitReport, error)
	GetVisitReportByID(id int) (entities.VisitReport, error)
	InsertVisitReport(visitReport entities.VisitReport) (entities.VisitReport, error)
	UpdateVisitReport(id int, visitReport entities.VisitReport) (entities.VisitReport, error)
	DeleteVisitReport(id int) error
}