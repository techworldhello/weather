package mock

import "weather/pkg/response"

type MockWeatherStack struct {
	GetWeatherMock func(city string) (w response.CustomResponse, err error)
}

func (m MockWeatherStack) GetWeatherData(city string) (w response.CustomResponse, err error) {
	return m.GetWeatherMock(city)
}

type MockOpenWeather struct {
	GetWeatherMock func(city string) (w response.CustomResponse, err error)
}

func (m MockOpenWeather) GetWeatherData(city string) (w response.CustomResponse, err error) {
	return m.GetWeatherMock(city)
}

