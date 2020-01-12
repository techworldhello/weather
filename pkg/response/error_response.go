package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorJSON struct {
	StatusCode int
	Message    string
}

func (e ErrorJSON) MarshallError() (b byte, err error) {
	return b, nil
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	e := &ErrorJSON{
		StatusCode: 404,
		Message: "Not found. Pls try /v1/weather endpoint",
	}
	b, err := json.Marshal(&e)
	if err != nil {
		log.Fatalf("error marshalling errorJSON: %v", err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatalf("error writing data to conn: %v", err)
		return
	}
}

func InvalidRequest(w http.ResponseWriter) {
	e := &ErrorJSON{
		StatusCode: 400,
		Message: "Invalid Request",
	}
	b, err := json.Marshal(&e)
	if err != nil {
		log.Printf("error marshalling errorJSON: %v", err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Printf("error writing data to conn: %v", err)
		return
	}
}

func ServerError(w http.ResponseWriter) {
	e := &ErrorJSON{
		StatusCode: 500,
		Message: "All weather services are down, please try again later",
	}
	b, err := json.Marshal(&e)
	if err != nil {
		log.Printf("error marshalling errorJSON: %v", err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Printf("error writing data to conn: %v", err)
		return
	}
}
