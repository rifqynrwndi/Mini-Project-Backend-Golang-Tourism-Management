package tourists

import (
	"tourism-monitoring/controllers/tourists/response"
	"tourism-monitoring/controllers/base"
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
	tourists, err := controller.touristsService.GetAllTourists()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromTouristEntities(tourists))
}
