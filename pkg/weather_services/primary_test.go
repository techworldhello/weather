package weather_services_test

import (
	"log"
	"testing"
	"weather/pkg/weather_services"
)

func TestGetWeatherResponse(t *testing.T) {
	w := weather_services.NewWeatherStack("Melbourne")
	resp, err := w.GetWeatherData()
	if err != nil {
		t.Errorf("error getting data: %v", err)

	}
	log.Print("response: ", resp)
}
