package weather

import "weather/pkg/response"

type WeatherService interface {
	GetWeatherData(city string) (w response.CustomResponse, err error)
}
