package weather_services

import (
	"log"
	"testing"
)

func TestGetOpenWeatherResponse(t *testing.T) {
	w := NewOpenWeather()
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}
