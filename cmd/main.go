package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"weather/pkg"
	"weather/pkg/response"
)

var logger = initLogger("Real-time weather 🌤")

func main() {
	w := weather.New(logger)

	mux := mux.NewRouter()
	mux.HandleFunc("/v1/weather", w.GetWeatherResponse).Methods("GET").Queries("city", "{city}")
	mux.HandleFunc("/", response.NotFound)

	logger.Info("Starting server on :3000...")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		logger.Fatal(err)
	}
}

func initLogger(name string) *logrus.Entry {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return logrus.WithFields(logrus.Fields{
		"app_name": name,
		"environment": os.Getenv("DEPLOYMENT_ENVIRONMENT"),
	})
}
