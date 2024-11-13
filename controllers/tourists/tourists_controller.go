package tourists

import (
	"strconv"
	"tourism-monitoring/controllers/base"
	"tourism-monitoring/controllers/pagination"
	"tourism-monitoring/controllers/tourists/response"
	"tourism-monitoring/entities"
	"tourism-monitoring/services/tourists"

	"github.com/labstack/echo/v4"
)

type TouristsController struct {
	touristsService *tourists.TouristsService
}

func NewTouristsController(service *tourists.TouristsService) *TouristsController {
	return &TouristsController{touristsService: service}
}

func (controller TouristsController) GetAllTourists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalCount, err := controller.touristsService.GetTotalTouristsCount()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	
	tourists, err := controller.touristsService.GetAllTourists()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return pagination.SuccessPaginatedResponse(c, tourists, page, limit, totalCount)
}

func (controller TouristsController) GetTouristByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tourists, err := controller.touristsService.GetTouristByID(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromTouristEntity(tourists))
}

func (controller TouristsController) InsertTourist(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return base.ErrorResponse(c, err)
	}

	createdUser, err := controller.touristsService.InsertTourist(*user)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, createdUser)
}

func (controller TouristsController) UpdateTourist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return base.ErrorResponse(c, err)
	}

	updatedUser, err := controller.touristsService.UpdateTourist(id, *user)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.FromTouristEntity(updatedUser))
}

func (controller TouristsController) DeleteTourist(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.touristsService.DeleteTourist(id); err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, "User deleted successfully")
}

