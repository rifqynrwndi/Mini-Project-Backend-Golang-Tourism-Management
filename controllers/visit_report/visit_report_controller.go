package visit_report

import (
	"strconv"
	"tourism-monitoring/controllers/base"
	"tourism-monitoring/controllers/visit_report/response"
	"tourism-monitoring/controllers/visit_report/request"
	"tourism-monitoring/services/visit_report"

	"github.com/labstack/echo/v4"
)

type VisitReportController struct {
	service *visit_report.VisitReportService
}

func NewVisitReportController(service *visit_report.VisitReportService) *VisitReportController {
	return &VisitReportController{service: service}
}

func (controller *VisitReportController) GetAllVisitReports(c echo.Context) error {
	reports, err := controller.service.GetAllVisitReports()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromVisitReportEntities(reports))
}

func (controller *VisitReportController) GetVisitReportByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	visitReport, err := controller.service.GetVisitReportByID(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.FromVisitReportEntity(visitReport))
}

func (controller *VisitReportController) InsertVisitReport(c echo.Context) error {
	req := new(request.VisitReportRequest)

	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}

	report, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	createdReport, err := controller.service.InsertVisitReport(report, req.TransportMode, req.DistanceKM)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, createdReport)
}

func (controller *VisitReportController) UpdateVisitReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.VisitReportRequest)

	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}

	report, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	updatedReport, err := controller.service.UpdateVisitReport(id, report, req.TransportMode, req.DistanceKM)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, updatedReport)
}

func (controller *VisitReportController) DeleteVisitReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.service.DeleteVisitReport(id); err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, "Visit report deleted successfully")
}
