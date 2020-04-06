package response

import (
	"fmt"
	"log"
	"net/http"
)

type StatusCode int

const (
	notFound StatusCode = iota
	badRequest
	serverError
)

func (s StatusCode) Int() int {
	var statusCodes = [...]int {
		http.StatusNotFound,
		http.StatusBadRequest,
		http.StatusInternalServerError,
	}
	if s < notFound || s > serverError {
		log.Print("unknown status code")
	}
	return statusCodes[s]
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(notFound.Int())
	writeToClient(notFound, w)
}

func InvalidRequest(w http.ResponseWriter) {
	w.WriteHeader(badRequest.Int())
	writeToClient(badRequest, w)
}

func ServerError(w http.ResponseWriter) {
	w.WriteHeader(serverError.Int())
	writeToClient(serverError, w)
}

func writeToClient(s StatusCode, w http.ResponseWriter) {
	var err error
	switch s {
	case notFound:
		_, err = fmt.Fprint(w, `{"StatusCode":404,"Message":"Not found. Pls try /v1/weather endpoint"}`)
	case badRequest:
		_, err = fmt.Fprint(w, `{"StatusCode":400,"Message":"Invalid Request"}`)
	case serverError:
		_, err = fmt.Fprint(w, `{"StatusCode":500,"Message":"All weather services are down, please try again later"}`)
	default:
		log.Printf("this status code hasn't been handled yet: %d", s.Int())
	}
	if err != nil {
		log.Printf("error writing to stream: %v", err)
	}
}
