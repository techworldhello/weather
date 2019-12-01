package weather

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http/httptest"
	"testing"
	"weather/pkg/response"
)

func TestGetWeatherResponseWorksREAL(t *testing.T) {
	mockLogrus := logrus.NewEntry(logrus.New())
	api := New(mockLogrus)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://dummy.com/v1/weather?city=melbourne", nil)
	api.GetWeatherResponse(w, r)

	log.Print("response: ", w.Body)
	log.Print("result: ", w.Result())
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
			api := API{response: expect.resp}
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
