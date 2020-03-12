package weather_services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"weather/pkg/response"
)

type OpenWeather struct {
	http   HttpClient
	apiKey string
	log    *logrus.Entry
}

func NewOpenWeather(logger *logrus.Entry, client HttpClient) *OpenWeather {
	return &OpenWeather{
		http:   client,
		apiKey: "2326504fb9b100bee21400190e4dbe6d",
		log:    logger,
	}
}

func (o OpenWeather) GetWeatherData(city string) (w response.CustomResponse, err error) {
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		o.log.Errorf("error creating new request: %v", err)
		return w, err
	}

	q := req.URL.Query()
	q.Add("q", city+",AU")
	q.Add("appid", o.apiKey)
	req.URL.RawQuery = q.Encode()
	o.log.Debugf("query params: %v", q)

	res, err := o.http.Do(req)
	if err != nil {
		o.log.Errorf("error making GET request: %v", err)
		return w, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		o.log.Debugf("resp status code: %d", res.StatusCode)
		return w, errors.New(fmt.Sprintf("http status code: %v", res.StatusCode))
	}

	var schema response.OpenWeatherResponse

	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&schema); err != nil {
		o.log.Errorf("error decoding json: %v", err)
		return w, err
	}

	return response.CustomResponse{
		Temperature: int(schema.Main.Temp) - 273,
		WindSpeed:   int(math.Round(schema.Wind.Speed / 5 * 18)),
	}, nil
}
