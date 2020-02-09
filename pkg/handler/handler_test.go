package handler

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weather/pkg/response"
)

func TestGetWeatherResponseWorksREAL(t *testing.T) {
	mockLogrus := logrus.NewEntry(logrus.New())
	api := New(mockLogrus, &http.Client{})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://dummy.com/v1/weather?city=melbourne", nil)

	for i := 1; i<=5; i++ {
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
	expectations := []struct{
		name string
		resp response.CustomResponse
		url  string
	}{
		{
			name: "200 OK",
			resp: response.CustomResponse{
				Temperature: 20,
				WindSpeed:   10,
			},
			url: "http://dummy.com/v1/weather?city=melbourne",
		},
		{
			name: "no query param",
			resp: response.CustomResponse{
				Temperature: 0,
				WindSpeed:   0,
			},
			url: "http://dummy.com/v1/weather",
		},
		{
			name: "no endpoint",
			resp: response.CustomResponse{
				Temperature: 0,
				WindSpeed:   0,
			},
			url: "http://dummy.com",
		},
		{
			name: "services are down", // will need to mock weather service to return 500
			resp: response.CustomResponse{
				Temperature: 0,
				WindSpeed:   0,
			},
			url: "http://dummy.com/v1/weather?city=melbourne",
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			api := API{}
			r := httptest.NewRequest("GET", expect.url, nil)
			w := httptest.NewRecorder()
			api.GetWeatherResponse(w, r)

			log.Print("response: ", w.Body)
			log.Print("result: ", w.Result())
			//if w.Body != expect.resp {
			//
			//}
			if w.Result().StatusCode != 200 {
				t.Error("expected a 200 OK response")
			}
		})
	}
}
