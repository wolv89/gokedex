package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	timeout time.Duration
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeout time.Duration) Cache {
	cache := Cache{
		make(map[string]cacheEntry),
		timeout,
		&sync.RWMutex{},
	}

	go func() {
		ticker := time.NewTicker(timeout)
		defer ticker.Stop()
		for range ticker.C {
			cache.Reap()
		}
	}()

	return cache
}

func (cache *Cache) Add(key string, val []byte) {

	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		time.Now(),
		val,
	}

}

func (cache *Cache) Get(key string) ([]byte, bool) {

	cache.mu.Lock()
	defer cache.mu.Unlock()

	if _, ok := cache.entries[key]; !ok {
		return nil, false
	}

	return cache.entries[key].val, true

}

func (cache *Cache) Reap() {

	cache.mu.Lock()
	defer cache.mu.Unlock()

	now := time.Now()

	for key, entry := range cache.entries {
		if now.After(entry.createdAt.Add(cache.timeout)) {
			delete(cache.entries, key)
		}
	}

}
