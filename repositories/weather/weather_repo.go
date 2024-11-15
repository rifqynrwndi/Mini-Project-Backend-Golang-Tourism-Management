package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity int     `json:"humidity"`
		WindKph  float64 `json:"wind_kph"`
	} `json:"current"`
}

type WeatherRepo struct {
	apiKey string
	apiURL string
}

func NewWeatherRepo() *WeatherRepo {
	return &WeatherRepo{
		apiKey: os.Getenv("WEATHER_API_KEY"),
		apiURL: os.Getenv("WEATHER_API_URL"),
	}
}

func (repo *WeatherRepo) FetchWeatherData(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", repo.apiURL, repo.apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}

	return &weatherResp, nil
}
