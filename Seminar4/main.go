package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
	Delete(k string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработаем конструктор и методы кеша, чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{
		data: make(map[string]string),
	}
}

type cacheImpl struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.data[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = v
}

func (c *cacheImpl) Delete(k string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, k)
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)

	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))

	c.Delete("test")

	fmt.Println(db.Get("test"))
}
