package providers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"weather/pkg/response"
)

type WeatherStack struct {
	http   HttpClient
	log    *logrus.Entry
}

func NewWeatherStack(logger *logrus.Entry, client HttpClient) *WeatherStack {
	return &WeatherStack{
		http:   client,
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
	q.Add("access_key", os.Getenv("WEATHERSTACK_KEY"))
	q.Add("query", city)
	req.URL.RawQuery = q.Encode()
	w.log.Debugf("query params: %v", q)

	res, err := w.http.Do(req)
	defer res.Body.Close()
	if err != nil {
		w.log.Fatalf("error making GET request: %v", err)
		return r, err
	}

	if res.StatusCode != http.StatusOK {
		w.log.Debugf("resp status code: %d", res.StatusCode)
		return r, err
	}

	var schema *response.WeatherStackResponse

	if err := json.NewDecoder(res.Body).Decode(&schema); err != nil {
		w.log.Errorf("error decoding response: %v", err)
		return r, err
	}

	return response.CustomResponse{
		Temperature: schema.Current.Temperature,
		WindSpeed:   schema.Current.WindSpeed,
	}, nil
}
