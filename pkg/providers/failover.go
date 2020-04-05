package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"weather/pkg/response"
)

type OpenWeather struct {
	http   HttpClient
	log    *logrus.Entry
}

func NewOpenWeather(logger *logrus.Entry, client HttpClient) *OpenWeather {
	return &OpenWeather{
		http:   client,
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
	q.Add("appid", os.Getenv("OPENWEATHER_KEY"))
	req.URL.RawQuery = q.Encode()
	o.log.Debugf("query params: %v", q)

	res, err := o.http.Do(req)
	defer res.Body.Close()
	if err != nil {
		o.log.Errorf("error making GET request: %v", err)
		return w, err
	}

	if res.StatusCode != http.StatusOK {
		o.log.Debugf("resp status code: %d", res.StatusCode)
		return w, errors.New(fmt.Sprintf("http status code: %v", res.StatusCode))
	}

	var schema response.OpenWeatherResponse

	if err := json.NewDecoder(res.Body).Decode(&schema); err != nil {
		o.log.Errorf("error decoding response: %v", err)
		return w, err
	}

	return response.CustomResponse{
		Temperature: int(schema.Main.Temp) - 273,
		WindSpeed:   int(math.Round(schema.Wind.Speed / 5 * 18)),
	}, nil
}
