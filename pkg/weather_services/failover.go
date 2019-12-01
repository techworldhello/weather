package weather_services

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"weather/pkg/response"
)

type OpenWeather struct {
	http   http.Client
	apiKey string
	schema response.OpenWeatherResponse
	log    *logrus.Entry
}

func NewOpenWeather(logger *logrus.Entry) *OpenWeather {
	return &OpenWeather{
		http:   http.Client{},
		apiKey: os.Getenv("OPENWEATHER_KEY"),
		log:    logger,
	}
}

func (o OpenWeather) GetWeatherData(city string) (w response.CustomResponse, err error) {
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		o.log.Fatalf("error creating new request: %v", err)
		return w, err
	}

	q := req.URL.Query()
	q.Add("q", city+",AU")
	q.Add("appid", o.apiKey)
	req.URL.RawQuery = q.Encode()
	o.log.Debugf("query params: %v", q)

	res, err := o.http.Do(req)
	if err != nil {
		o.log.Fatalf("error making GET request: %v", err)
		return w, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		o.log.Debugf("resp status code: %d", res.StatusCode)
		return w, err
	}

	if err := json.NewDecoder(res.Body).Decode(&o.schema); err != nil {
		o.log.Fatalf("error decoding json: %v", err)
		return w, err
	}

	return response.CustomResponse{
		Temperature: int(o.schema.Main.Temp) - 273,
		WindSpeed:   int(math.Round(o.schema.Wind.Speed / 5 * 18)),
	}, nil
}
