package weather_services

import (
	"log"
	"net/http"
	"testing"
	"weather/pkg/mock"
)

func _TestGetWeatherResponseWorks(t *testing.T) {
	w := WeatherStack{

	}
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}

func TestGetWeatherResponseREAL(t *testing.T) {
	w := NewWeatherStack(mock.TestLogger, &http.Client{})
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}
