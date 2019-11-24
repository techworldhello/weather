package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"weather/pkg/response"
	"weather/pkg/weather_services"
)

type API struct {
	Services []WeatherService
}

func New() *API { // variadic args in case of more services in future?
	return &API{
		Services: []WeatherService{
			weather_services.NewWeatherStack("Melbourne"),
			weather_services.NewOpenWeather("Melbourne"),
		},
	}
}

func (s API) GetWeatherResponse(w http.ResponseWriter, r *http.Request) {
	// get results from cache, if failure, get primary service

	// check city requirement

	var response response.CustomResponse

	for _, service := range s.Services {
		resp, err := service.GetWeatherData()
		if err == nil {
			response = resp
		} else {
			log.Fatal(err)
		}
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
