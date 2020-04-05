package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"weather/pkg/handler"
	"weather/pkg/response"
)

func main() {
	var logger = initLogger("Weather App🌤")
	w := handler.New(logger, &http.Client{})

	mux := mux.NewRouter()
	mux.HandleFunc("/v1/weather", w.GetWeatherResponse).Methods("GET").Queries("city", "{city}")
	mux.HandleFunc("/", response.NotFound)

	logger.Info("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Fatal(err)
	}
}

func initLogger(name string) *logrus.Entry {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return logrus.WithFields(logrus.Fields{
		"app_name":    name,
		"environment": os.Getenv("DEPLOYMENT_ENVIRONMENT"),
	})
}
