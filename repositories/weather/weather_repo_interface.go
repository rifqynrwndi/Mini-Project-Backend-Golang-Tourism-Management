package weather

type WeatherRepoInterface interface {
	FetchWeatherData(city string) (*WeatherResponse, error)
}