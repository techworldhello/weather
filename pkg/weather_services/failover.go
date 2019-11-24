package weather_services

import (
	"encoding/json"
	"math"
	"net/http"
	"weather/pkg/response"
)

type OpenWeather struct {
	Http http.Client
	Key  string
	City string
}

func NewOpenWeather(city string) *WeatherStack {
	return &WeatherStack{
		Http: http.Client{},
		Key:  "2326504fb9b100bee21400190e4dbe6d",
		City: city,
	}
}

func (ow OpenWeather) GetWeatherData() (w response.CustomResponse, err error) {
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		return w, err
	}

	q := req.URL.Query()
	q.Add("q", ow.City+",AU")
	q.Add("appid", ow.Key)
	req.URL.RawQuery = q.Encode()

	res, err := ow.Http.Do(req)
	if err != nil {
		return w, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return w, err
	}

	schema := response.OpenWeatherResponse{}

	if err := json.NewDecoder(res.Body).Decode(&schema); err != nil {
		return w, err
	}

	return response.CustomResponse{
		Temperature: int(schema.Main.Temp) - 273,
		WindSpeed:   int(math.Round(schema.Wind.Speed / 5 * 18)),
	}, nil
}
