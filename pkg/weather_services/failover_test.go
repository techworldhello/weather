package weather_services

import (
	"log"
	"net/http"
	"testing"
	"weather/pkg/mock"
)

func _TestGetOpenWeatherResponseREAL(t *testing.T) {
	w := NewOpenWeather(mock.TestLogger, &http.Client{})
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}
