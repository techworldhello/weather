package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"weather/pkg"
)

func main() {
	w := weather.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/weather", w.GetWeatherResponse)

	log.Println("Starting server on :4000...")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
