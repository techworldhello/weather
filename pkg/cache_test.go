package weather

import (
	"testing"
	"time"
	"weather/pkg/response"
)

func TestCacheDoesNotExist(t *testing.T) {
	c := Cache{lastUpdated: time.Now().Local().Add(5 * time.Second)}
	resp, cacheExists := c.processCache()
	if cacheExists {
		t.Error("cache should not exist (lastUpdated exceeds 3 seconds)")
	}
	cr := response.CustomResponse{}
	if resp != cr {
		t.Error("response should not exist in cache")
	}
}

func TestCacheExists(t *testing.T) {
	cr := response.CustomResponse{
		Temperature: 20,
		WindSpeed: 10,
	}
	c := Cache{
		lastUpdated: time.Now(),
		CustomResponse: cr,
	}
	resp, cacheExists := c.processCache()
	if !cacheExists {
		t.Error("cache should exist but doesn't")
	}
	if resp != cr {
		t.Error("cache should hold an existing response")
	}
}
