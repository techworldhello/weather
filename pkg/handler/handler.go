package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"weather/pkg/cache"
	"weather/pkg/response"
	"weather/pkg/providers"
)

type API struct {
	services []providers.WeatherService
	log      *logrus.Entry
	cache    *cache.Cache
}

func New(logger *logrus.Entry, client providers.HttpClient) *API {
	return &API{
		services: []providers.WeatherService{
			providers.NewWeatherStack(logger, client),
			providers.NewOpenWeather(logger, client),
		},
		log:   logger,
		cache: cache.New(3 * time.Second),
	}
}

func (s API) GetWeatherResponse(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["city"]
	if len(param) == 0 {
		s.log.Error("No city query in request")
		response.InvalidRequest(w)
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
		return
	}

	resBytes, err := json.Marshal(customResp)
	if err != nil {
		s.log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}
