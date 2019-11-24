package response

type WeatherStackResponse struct {
	Current struct {
		ObservationTime     string   `json:"observation_time"`
		Temperature         int      `json:"temperature"`
		WeatherCode         int      `json:"weather_code"`
		WeatherIcons        []string `json:"weather_icons"`
		WeatherDescriptions []string `json:"weather_descriptions"`
		WindSpeed           int      `json:"wind_speed"`
		WindDegree          int      `json:"wind_degree"`
		WindDir             string   `json:"wind_dir"`
		Pressure            int      `json:"pressure"`
		Precip              float64  `json:"precip"`
		Humidity            int      `json:"humidity"`
		Cloudcover          int      `json:"cloudcover"`
		Feelslike           int      `json:"feelslike"`
		UvIndex             int      `json:"uv_index"`
		Visibility          int      `json:"visibility"`
		IsDay               string   `json:"is_day"`
	} `json:"current"`
}

type OpenWeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
}
