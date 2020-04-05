package providers

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"weather/pkg/mock"
	"weather/pkg/response"
)

func _TestGetWeatherResponseREAL(t *testing.T) {
	w := NewWeatherStack(mock.TestLogger, &http.Client{})
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting data: %v", err)
	}
	log.Print("response: ", resp)
}

func TestPrimaryProviderReturns200(t *testing.T) {
	w := WeatherStack{
		http: mock.Client{200, mock.WeatherStackResp},
		log: mock.TestLogger,
	}
	resp, err := w.GetWeatherData("Melbourne")
	if err != nil {
		t.Errorf("error getting weather data: %v", err)
	}
	expectedResp := response.CustomResponse{
		Temperature: 11,
		WindSpeed: 13,
	}
	assert.Equal(t, resp, expectedResp)
}
