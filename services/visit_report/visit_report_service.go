package visit_report

import (
	"tourism-monitoring/entities"
	"tourism-monitoring/repositories/visit_report"
)

type VisitReportService struct {
	repo visit_report.VisitReportRepoInterface
}

func NewVisitReportService(repo visit_report.VisitReportRepoInterface) *VisitReportService {
	return &VisitReportService{repo: repo}
}

func (service *VisitReportService) GetAllVisitReports(page int, limit int) ([]entities.VisitReport, error) {
	offset := (page - 1) * limit
	return service.repo.GetAllVisitReports(limit, offset)
}

func (service *VisitReportService) GetVisitReportByID(id int) (entities.VisitReport, error) {
	return service.repo.GetVisitReportByID(id)
}

func (service *VisitReportService) InsertVisitReport(visitReport entities.VisitReport, transportMode string, distanceKM float64) (entities.VisitReport, error) {
	visitReport.EstimasiEmisiKarbon = calculateCarbonEmission(transportMode, distanceKM)
	return service.repo.InsertVisitReport(visitReport)
}

func (service *VisitReportService) UpdateVisitReport(id int, visitReport entities.VisitReport, transportMode string, distanceKM float64) (entities.VisitReport, error) {
	visitReport.EstimasiEmisiKarbon = calculateCarbonEmission(transportMode, distanceKM)
	visitReport.ID = id
	return service.repo.UpdateVisitReport(id, visitReport)
}

func (service *VisitReportService) DeleteVisitReport(id int) error {
	return service.repo.DeleteVisitReport(id)
}

func (service *VisitReportService) GetTotalVisitReportsCount() (int64, error) {
	return service.repo.GetTotalVisitReportsCount()
}

func calculateCarbonEmission(transportMode string, distanceKM float64) float64 {
	emissionFactor := 0.21
	switch transportMode {
	case "train":
		emissionFactor = 0.05
	case "plane":
		emissionFactor = 0.25
	case "bus":
		emissionFactor = 0.1
	case "motorcycle":
		emissionFactor = 0.15
	case "car":
		emissionFactor = 0.2
	case "bicycle":
		emissionFactor = 0.01
	}
	return emissionFactor * distanceKM
}
