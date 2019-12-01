package weather_services

import (
	"log"
	"testing"
)

func TestGetWeatherResponseWorks(t *testing.T) {
	w := WeatherStack{


	}
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}

func TestGetWeatherResponse(t *testing.T) {
	w := NewWeatherStack()
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}
