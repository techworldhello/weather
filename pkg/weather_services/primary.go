package weather_services

import (
	"encoding/json"
	"log"
	"net/http"
	"weather/pkg/response"
)

type WeatherStack struct {
	Http http.Client
	Key  string
	City string
}

func NewWeatherStack(city string) *WeatherStack {
	return &WeatherStack{
		Http: http.Client{},
		Key:  "ea506498a37292032b031f4837caeca",
		City: city,
	}
}

func (ws WeatherStack) GetWeatherData() (w response.CustomResponse, err error) {
	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	if err != nil {
		return w, err
	}

	q := req.URL.Query()
	q.Add("access_key", ws.Key)
	q.Add("query", ws.City)
	req.URL.RawQuery = q.Encode()

	res, err := ws.Http.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return w, err
	}

	schema := response.WeatherStackResponse{}

	if err := json.NewDecoder(res.Body).Decode(&schema); err != nil {
		return w, err
	}

	return response.CustomResponse{
		Temperature: schema.Current.Temperature,
		WindSpeed:   schema.Current.WindSpeed,
	}, nil
}
