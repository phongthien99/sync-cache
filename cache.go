package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	data    *sync.Map
	janitor *janitor
}

func NewCache(cleanupInterval time.Duration) *Cache {
	var cache = &Cache{
		data: new(sync.Map),
		janitor: &janitor{
			interval: cleanupInterval,
			stop:     make(chan bool),
		},
	}
	go cache.janitor.Run(cache)
	return cache
}

func (c *Cache) Get(key string) (any, error) {
	value, ok := c.data.Load(key)
	if !ok {
		return nil, ErrNotFound // Key not found
	}
	return (value).(*CacheItem).Value, nil
}

func (c *Cache) Set(key string, value any, d time.Duration) {

	var e int64

	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	item := &CacheItem{
		Value:      value,
		Expiration: e,
	}

	c.data.Store(key, item)

}

func (c *Cache) Delete(key string) {
	c.data.Delete(key)
}

func (c *Cache) Clear() {
	c.data.Range(func(key, _ any) bool {
		c.data.Delete(key)
		return true
	})
}

func (c *Cache) Size() int {
	size := 0
	c.data.Range(func(_, _ any) bool {
		size++
		return true
	})
	return size
}

func (c *Cache) work() {
	now := time.Now().UnixNano()
	c.data.Range(func(key, value any) bool {
		var expiration = value.(int64)
		if now > expiration && expiration != 0 {
			c.Delete(key.(string))
		}
		return true
	})
}
