package weather

import (
	"log"
	"time"
	"weather/pkg/response"
)

type Cache struct {
	response.CustomResponse
	lastUpdated time.Time
}

func (c Cache) processCache() (response.CustomResponse, bool) {
	if cacheExists(c.lastUpdated) && c.Temperature != 0 { // TODO: best way to check if customResponse struct is not empty?
		log.Print("using cache")
		return c.CustomResponse, true
	}
	return response.CustomResponse{}, false
}

func cacheExists(lastUpdatedCache time.Time) bool {
	elapsed := time.Since(lastUpdatedCache)
	log.Printf("elapsed %v", elapsed)
	return elapsed <= time.Second*3
}

func (c Cache) saveResponse(resp response.CustomResponse) {
	c.CustomResponse = resp
	c.lastUpdated = time.Now()
}
