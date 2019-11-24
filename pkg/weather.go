package weather

import "weather/pkg/response"

type WeatherService interface {
	GetWeatherData() (w response.CustomResponse, err error)
}
