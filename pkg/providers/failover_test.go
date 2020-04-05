package providers

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"weather/pkg/mock"
	"weather/pkg/response"
)

func _TestGetOpenWeatherResponseREAL(t *testing.T) {
	w := NewOpenWeather(mock.TestLogger, &http.Client{})
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}

func TestFailoverProviderReturns200(t *testing.T) {
	w := OpenWeather{
		http: mock.Client{200, mock.OpenWeatherResp},
		log: mock.TestLogger,
	}
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting weather data: %v", err)
	}
	expectedResp := response.CustomResponse{
		Temperature: 21,
		WindSpeed: 13,
	}
	assert.Equal(t, resp, expectedResp)
}
