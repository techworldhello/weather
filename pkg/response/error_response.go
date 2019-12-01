package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorJSON struct {
	StatusCode int
	Message    string
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	e := &errorJSON{
		StatusCode: 404,
		Message: "Not found. Pls try /v1/weather endpoint",
	}
	b, err := json.Marshal(&e)
	if err != nil {
		log.Fatalf("error marshalling errorJSON: %v", err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatalf("error writing data to conn: %v", err)
		return
	}
}
