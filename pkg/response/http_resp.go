package response

import (
	"fmt"
	"log"
	"net/http"
)

type ErrorJSON struct {
	StatusCode int
	Message    string
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(w, `{"StatusCode":404,"Message":"Not found. Pls try /v1/weather endpoint"}`)
	if err != nil {
		log.Printf("error writing to stream: %v", err)
	}
}

func InvalidRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := fmt.Fprint(w, `{"StatusCode":400,"Message":"Invalid Request"}`)
	if err != nil {
		log.Printf("error writing to stream: %v", err)
	}
}

func ServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := fmt.Fprint(w, `{"StatusCode":500,"Message":"All weather services are down, please try again later"}`)
	if err != nil {
		log.Printf("error writing to stream: %v", err)
	}
}
