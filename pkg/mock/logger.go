package mock

import (
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"weather/pkg/response"
)

var testLog, _ = test.NewNullLogger()
var TestLogger = logrus.NewEntry(testLog)

type MockWeatherStack struct {}

func (m MockWeatherStack) GetWeatherData(city string) (w response.CustomResponse, err error) {
	return w, nil
}

type MockOpenWeather struct {}

func (m MockOpenWeather) GetWeatherData(city string) (w response.CustomResponse, err error) {
	return w, nil
}
