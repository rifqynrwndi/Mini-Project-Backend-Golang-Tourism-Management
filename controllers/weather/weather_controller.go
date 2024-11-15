package weather

import (
	"net/http"
	"tourism-monitoring/services/weather"

	"github.com/labstack/echo/v4"
)

type WeatherController struct {
	service *weather.WeatherService
}

func NewWeatherController(service *weather.WeatherService) *WeatherController {
	return &WeatherController{service: service}
}

func (controller *WeatherController) GetWeather(c echo.Context) error {
	city := c.QueryParam("city")
	if city == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "City is required"})
	}

	weatherData, err := controller.service.GetWeather(city)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, weatherData)
}
