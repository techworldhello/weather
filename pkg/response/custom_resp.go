package response

type CustomResponse struct {
	Temperature int `json: "temperature_degrees"`
	WindSpeed   int `json: "wind_speed"`
}
