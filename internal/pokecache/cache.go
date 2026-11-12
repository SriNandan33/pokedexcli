package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	data map[string]cacheEntry
	mux  *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		data: map[string]cacheEntry{},
		mux:  &sync.Mutex{},
	}
	go c.reapPeriodically(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.data[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, ok := c.data[key]
	return entry.val, ok
}

func (c *Cache) reapPeriodically(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > interval {
				delete(c.data, key)
			}
		}
		c.mux.Unlock()
	}
}
