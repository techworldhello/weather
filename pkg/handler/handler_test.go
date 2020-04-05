package handler

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weather/pkg/cache"
	"weather/pkg/mock"
	"weather/pkg/response"
	"weather/pkg/providers"
)

func _TestGetWeatherResponseWorksREAL(t *testing.T) {
	mockLogrus := logrus.NewEntry(logrus.New())
	api := New(mockLogrus, &http.Client{})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://dummy.com/v1/weather?city=melbourne", nil)

	for i := 1; i <= 5; i++ {
		api.GetWeatherResponse(w, r)
		log.Print("response: ", i, w.Body)
		log.Print("result: ", i, w.Result())
	}
	time.Sleep(5 * time.Second)
	api.GetWeatherResponse(w, r)
	log.Print("final response: ", w.Body)
	log.Print("final result: ", w.Result())
}

func TestGetWeatherResponse(t *testing.T) {
	var (
		successReqFn = func(melbourne string) (w response.CustomResponse, err error) {
			return w, nil
		}
		serverErrorReqFn = func(melbourne string) (w response.CustomResponse, err error) {
			return w, errors.New("server error")
		}
	)
	expectations := []struct {
		name         string
		resp         string
		url          string
		statusCode   int
		mockServices []providers.WeatherService
	}{
		{
			name:       "200 OK",
			url:        "http://dummy.com/v1/weather?city=melbourne",
			statusCode: 200,
			resp:       `{"Temperature":0,"WindSpeed":0}`,
			mockServices: []providers.WeatherService{
				mock.MockWeatherStack{successReqFn}, mock.MockOpenWeather{successReqFn},
			},
		},
		{
			name:       "Invalid request - no query param",
			url:        "http://dummy.com/v1/weather",
			statusCode: 400,
			resp:       `{"StatusCode":400,"Message":"Invalid Request"}`,
		},
		{
			name:       "Internal server error",
			url:        "http://dummy.com/v1/weather?city=melbourne",
			statusCode: 500,
			resp:       `{"StatusCode":500,"Message":"All weather services are down, please try again later"}`,
			mockServices: []providers.WeatherService{
				mock.MockWeatherStack{serverErrorReqFn}, mock.MockOpenWeather{serverErrorReqFn},
			},
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			time.Sleep(1 * time.Second)
			api := API{services: expect.mockServices, log: mock.TestLogger, cache: &cache.Cache{}}
			r := httptest.NewRequest("GET", expect.url, nil)
			w := httptest.NewRecorder()
			api.GetWeatherResponse(w, r)

			assert.Equal(t, w.Body.String(), expect.resp)
			assert.Equal(t, w.Result().StatusCode, expect.statusCode)
		})
	}
}
