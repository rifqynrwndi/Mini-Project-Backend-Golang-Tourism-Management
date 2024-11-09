package places

import (
	"strconv"
	"tourism-monitoring/controllers/base"
	"tourism-monitoring/controllers/places/response"
	"tourism-monitoring/entities"
	"tourism-monitoring/services/places"

	"github.com/labstack/echo/v4"
)

type PlacesController struct {
	placesService *places.PlacesService
}

func NewPlacesController(service *places.PlacesService) *PlacesController {
	return &PlacesController{placesService: service}
}

func (controller PlacesController) GetAllPlaces(c echo.Context) error {
	places, err := controller.placesService.GetAllPlaces()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromPlaceEntities(places))
}

func (controller PlacesController) GetPlaceById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	places, err := controller.placesService.GetPlaceByID(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromPlaceEntity(places))
}

func (controller PlacesController) InsertPlace(c echo.Context) error {
	place := new(entities.Place)
	if err := c.Bind(place); err != nil {
		return base.ErrorResponse(c, err)
	}
	createdPlace, err := controller.placesService.InsertPlace(*place)
    if err != nil {
        return base.ErrorResponse(c, err)
    }
    return base.SuccesResponse(c, createdPlace)
}

func (controller PlacesController) UpdatePlace(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	place := new(entities.Place)
	if err := c.Bind(place); err != nil {
		return base.ErrorResponse(c, err)
	}

	place.ID = id

	updatedPlace, err := controller.placesService.UpdatePlace(id, *place)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.FromPlaceEntity(updatedPlace))
}

func (controller PlacesController) DeletePlace(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.placesService.DeletePlace(id); err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, "Place deleted successfully")
}