package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"weather/pkg/cache"
	"weather/pkg/response"
	"weather/pkg/weather_services"
)

type API struct {
	services []weather_services.WeatherService
	log      *logrus.Entry
	cache    *cache.Cache
}

func New(logger *logrus.Entry, client weather_services.HttpClient) *API {
	return &API{
		services: []weather_services.WeatherService{
			weather_services.NewWeatherStack(logger, client),
			weather_services.NewOpenWeather(logger, client),
		},
		log:   logger,
		cache: cache.New(3 * time.Second),
	}
}

func (s API) GetWeatherResponse(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["city"]
	if len(param) == 0 {
		s.log.Error("No city query in request")
		response.InvalidRequest(w) // look into middleware
		return
	}

	var customResp *response.CustomResponse
	cityQueried := param[0]

	result, cacheExists := s.cache.Get(cityQueried)
	if cacheExists {
		customResp = &result
		s.log.Info("Using cache...")
	} else {
		for _, service := range s.services {
			resp, err := service.GetWeatherData(cityQueried)
			if err == nil {
				customResp = &resp
				s.cache.Add(cityQueried, resp)
				break
			}
		}
	}

	if customResp == nil {
		response.ServerError(w)
	}

	resBytes, err := json.Marshal(customResp)
	if err != nil {
		s.log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}
