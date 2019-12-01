package response

type WeatherStackResponse struct {
	Current struct {
		Temperature int `json:"temperature"`
		WindSpeed   int `json:"wind_speed"`
	}
}

type OpenWeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	}
	Wind struct {
		Speed float64 `json:"speed"`
	}
}
