package trash_report

import (
	"fmt"
	"strconv"
	"tourism-monitoring/controllers/base"
	"tourism-monitoring/controllers/pagination"
	"tourism-monitoring/controllers/trash_report/request"
	"tourism-monitoring/controllers/trash_report/response"
	"tourism-monitoring/services/trash_report"

	"github.com/labstack/echo/v4"
)

type TrashReportController struct {
	service *trash_report.TrashReportService
}

func NewTrashReportController(service *trash_report.TrashReportService) *TrashReportController {
	return &TrashReportController{service: service}
}

func (controller *TrashReportController) GetTrashReportByPlaceID(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalCount, err := controller.service.GetTotalTrashReportsCount()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	reports, err := controller.service.GetTrashReportByPlaceID(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return pagination.SuccessPaginatedResponse(c, reports, page, limit, totalCount)
}

func (controller *TrashReportController) GetTrashReportByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	report, err := controller.service.GetTrashReportByID(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromTrashReportEntity(report))
}

func (controller *TrashReportController) InsertTrashReport(c echo.Context) error {
	req := new(request.TrashReportRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}

	report, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	createdReport, err := controller.service.InsertTrashReport(report)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromTrashReportEntity(createdReport))
}

func (controller *TrashReportController) UpdateTrashReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	req := new(request.TrashReportRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err)
	}

	report, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	report.ID = id
	if report.ObjekWisataID == 0 {
		return base.ErrorResponse(c, fmt.Errorf("objek_wisata_id tidak valid"))
	}

	updatedReport, err := controller.service.UpdateTrashReport(id, report)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.FromTrashReportEntity(updatedReport))
}

func (controller *TrashReportController) DeleteTrashReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.service.DeleteTrashReport(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, "Trash report deleted successfully")
}
