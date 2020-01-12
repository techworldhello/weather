package cache

import (
	"log"
	"sync"
	"time"
	"weather/pkg/response"
)

type Cache struct {
	data        map[string]response.CustomResponse
	LastUpdated time.Time
	TTL         time.Duration
	mu          sync.RWMutex
}

func New(expiration time.Duration) *Cache {
	//cached := make(map[string]response.CustomResponse)
	//cached["weather"] = response.CustomResponse{
	//	Temperature: 20,
	//	WindSpeed: 10,
	//}
	return &Cache{
		TTL: expiration,
		//data: cached,
	}
}

func (c *Cache) Get(key string) (response.CustomResponse, bool) { // TODO: add mutex so 2 request cant write to cache at the same time
	log.Printf("data to get: %v", c.data["weather"])
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isExpired() && len(c.data) != 0 {
		return c.data[key], true
	}
	log.Print("cache empty")
	return c.data[key], false
}

func (c *Cache) isExpired() bool {
	elapsed := time.Since(c.LastUpdated)
	log.Printf("elapsed %v", elapsed*time.Second)
	if c.TTL > 0 {
		log.Printf("returned %v", c.TTL <= elapsed*time.Second)
		return c.TTL <= elapsed*time.Second // 3 < 5
	}
	return false
}

func (c *Cache) Add(key string, resp response.CustomResponse) {
	c.mu.Lock()
	if c.data == nil {
		c.data = make(map[string]response.CustomResponse)
	}
	c.data[key] = resp

	log.Printf("data added %v", c.data[key])
	c.LastUpdated = time.Now()
	c.mu.Unlock()
}

func (c *Cache) Clear() {
	for k := range c.data {
		delete(c.data, k)
	}
}
