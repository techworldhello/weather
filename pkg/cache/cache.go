package cache

import (
	"log"
	"sync"
	"time"
	"weather/pkg/response"
)

type Cache struct {
	data        map[string]response.CustomResponse
	LastUpdated map[string]time.Time
	TTL         time.Duration
	mu          sync.RWMutex
}

func New(expiration time.Duration) *Cache {
	return &Cache{TTL: expiration}
}

func (c *Cache) Get(key string) (response.CustomResponse, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isExpired(key) && len(c.data) != 0 {
		return c.data[key], true
	}
	log.Print("cache empty")
	return c.data[key], false
}

func (c *Cache) isExpired(key string) bool {
	elapsed := time.Since(c.LastUpdated[key])
	log.Printf("elapsed %v seconds", elapsed / time.Second)
	if c.TTL > 0 {
		return c.TTL <= elapsed
	}
	return false
}

func (c *Cache) Add(key string, resp response.CustomResponse) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.data == nil || c.LastUpdated == nil {
		c.data = make(map[string]response.CustomResponse)
		c.LastUpdated = make(map[string]time.Time)
	}
	c.data[key] = resp
	c.LastUpdated[key] = time.Now()

	log.Printf("data added %v", c.data[key])
}

func (c *Cache) Clear() {
	for k := range c.data {
		delete(c.data, k)
	}
}
