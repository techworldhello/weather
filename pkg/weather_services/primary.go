package weather_services

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"weather/pkg/response"
)

type WeatherStack struct {
	http   HttpClient
	apiKey string
	log    *logrus.Entry
}

func NewWeatherStack(logger *logrus.Entry, client HttpClient) *WeatherStack {
	return &WeatherStack{
		http:   client,
		apiKey: "ea506498a37292032b031f4837caeca4",
		log:    logger,
	}
}

func (w WeatherStack) GetWeatherData(city string) (r response.CustomResponse, err error) {
	req, err := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	if err != nil {
		w.log.Fatalf("error creating new request: %v", err)
		return r, err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("access_key", w.apiKey)
	q.Add("query", city)
	req.URL.RawQuery = q.Encode()
	w.log.Debugf("query params: %v", q)

	res, err := w.http.Do(req)
	logrus.Printf("err: %v", err)
	if err != nil {
		w.log.Fatalf("error making GET request: %v", err)
		return r, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		w.log.Debugf("resp status code: %d", res.StatusCode)
		return r, err
	}

	var schema *response.WeatherStackResponse

	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&schema); err != nil {
		w.log.Errorf("1st error decoding json: %v", err)
		return r, err
	}

	return response.CustomResponse{
		Temperature: schema.Current.Temperature,
		WindSpeed:   schema.Current.WindSpeed,
	}, nil
}
