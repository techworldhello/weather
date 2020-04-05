package cache

import (
	"log"
	"testing"
	"time"
	"weather/pkg/response"
)

func TestCache(t *testing.T) {
	expectations := []struct {
		name        string
		citySaved   string
		lastUpdated time.Time
		response    response.CustomResponse
		cacheExists bool
	}{
		{
			name:        "Cache exits",
			citySaved:   "Melbourne",
			lastUpdated: time.Now(),
			response: response.CustomResponse{
				Temperature: 20,
				WindSpeed:   10,
			},
			cacheExists: true,
		},
		{
			name:        "Cache exits different city",
			citySaved:   "Sydney",
			lastUpdated: time.Now(),
			response: response.CustomResponse{
				Temperature: 25,
				WindSpeed:   15,
			},
			cacheExists: true,
		},
		{
			name:        "Cache does not exist",
			citySaved:   "Melbourne",
			lastUpdated: time.Now().Add(-5 * time.Second),
			response: response.CustomResponse{
				Temperature: 20,
				WindSpeed:   10,
			},
			cacheExists: false,
		},
	}

	cacheTTL := 3 * time.Second
	emptyResp := response.CustomResponse{}

	for i, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			cache := New(cacheTTL)
			cache.Add(expect.citySaved, expect.response)
			cache.LastUpdated = map[string]time.Time{expect.citySaved: expect.lastUpdated}

			resp, cacheExists := cache.Get(expect.citySaved)
			if cacheExists != expect.cacheExists {
				t.Errorf("Expected: %t\nGot: %t", expect.cacheExists, cacheExists)
			}
			if resp != expect.response {
				t.Errorf("Expected: %v\nGot: %v", expect.response, resp)
			}
			log.Printf("cached data <%d>: %v", i+1, cache.data[expect.citySaved])

			cache.Clear()
			if cache.data[expect.citySaved] != emptyResp {
				t.Error("Cache should have been cleared but is not")
			}
		})
	}
}
