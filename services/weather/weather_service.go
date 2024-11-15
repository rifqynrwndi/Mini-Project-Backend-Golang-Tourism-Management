package weather

import (
	"fmt"
	"tourism-monitoring/repositories/weather"
)

type WeatherService struct {
	repo *weather.WeatherRepo
}

func NewWeatherService(repo *weather.WeatherRepo) *WeatherService {
	return &WeatherService{repo: repo}
}

func (service *WeatherService) GetWeather(city string) (*weather.WeatherResponse, error) {
	if city == "" {
		return nil, fmt.Errorf("city is required")
	}

	data, err := service.repo.FetchWeatherData(city)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}

	return data, nil
}
