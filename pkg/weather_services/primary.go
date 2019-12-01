package weather_services

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"weather/pkg/response"
)

type WeatherStack struct {
	http   http.Client
	apiKey string
	schema response.WeatherStackResponse
	log   *logrus.Entry
}

func NewWeatherStack(logger *logrus.Entry) *WeatherStack {
	return &WeatherStack{
		http:   http.Client{}, // TODO: how to mock in test without interfacing do??
		apiKey: "ea506498a37292032b031f4837caeca4",
		log: logger,
	}
}

func (w WeatherStack) GetWeatherData(city string) (r response.CustomResponse, err error) {
	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	if err != nil {
		w.log.Fatalf("error creating new request: %v", err)
		return r, err
	}

	q := req.URL.Query()
	q.Add("access_key", w.apiKey)
	q.Add("query", city)
	req.URL.RawQuery = q.Encode()
	w.log.Debugf("query params: %v", q)

	res, err := w.http.Do(req)
	if err != nil {
		w.log.Fatalf("error making GET request: %v", err)
		return r, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		w.log.Debugf("resp status code: %d", res.StatusCode)
		return r, err
	}

	// TODO: why doesnt it error out with the schema is different (error response)
	if err := json.NewDecoder(res.Body).Decode(&w.schema); err != nil {
		w.log.Fatalf("error decoding json: %v", err)
		return r, err
	}

	return response.CustomResponse{
		Temperature: w.schema.Current.Temperature,
		WindSpeed:   w.schema.Current.WindSpeed,
	}, nil
}
