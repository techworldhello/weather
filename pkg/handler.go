package weather

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"weather/pkg/response"
	"weather/pkg/weather_services"
)

type API struct {
	services []WeatherService
	memo     Cache
	response response.CustomResponse
	log      *logrus.Entry
}

func New(logger *logrus.Entry) *API {
	return &API{
		services: []WeatherService{
			weather_services.NewWeatherStack(logger),
			weather_services.NewOpenWeather(logger),
		},
		memo: Cache{
			lastUpdated: time.Now(),
		},
		log: logger,
	}
}

func (s API) GetWeatherResponse(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["city"]
	if len(param) == 0 {
		s.log.Fatal("You must provide a city query")
		return
	}

	result, cacheExists := s.memo.processCache()
	if cacheExists {
		s.response = result
		s.log.Info("Using cache..")
	} else {
		for _, service := range s.services {
			resp, err := service.GetWeatherData(param[0])
			if err == nil {
				s.response = resp
				s.memo.saveResponse(resp)
				break
			}
		}
	}

	resBytes, err := json.Marshal(&s.response)
	if err != nil {
		s.log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}
