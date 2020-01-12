package cache

import (
	"testing"
	"time"
	"weather/pkg/response"
)

func TestCacheDoesNotExist(t *testing.T) {
	c := Cache{LastUpdated: time.Now().Local().Add(5 * time.Second)}
	resp, cacheExists := c.Get("weather")
	if cacheExists {
		t.Error("cache should not exist (lastUpdated exceeds 3 seconds)")
	}
	cr := response.CustomResponse{}
	if resp != cr {
		t.Error("response should not exist in cache")
	}
}

func TestCacheExists(t *testing.T) {
	cached := make(map[string]response.CustomResponse)
	cached["weather"] = response.CustomResponse{
		Temperature: 20,
		WindSpeed: 10,
	}
	c := Cache{
		LastUpdated: time.Now(),
		data: cached,
	}
	resp, cacheExists := c.Get("weather")
	if !cacheExists {
		t.Error("cache should exist but doesn't")
	}
	if resp != cached["weather"] {
		t.Error("cache should hold an existing response")
	}
}
