package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.Mutex
}

func New(interval time.Duration) Cache {
	cache := Cache{entries: map[string]cacheEntry{}, mutex: &sync.Mutex{}}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	e, ok := c.entries[key]
	return e.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) >= interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}
