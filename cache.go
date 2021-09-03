package gocache

import (
	"sync"
	"time"
)

type Cache struct {
	d time.Duration
	sync.Map
}

func NewCache(d time.Duration) *Cache {
	return &Cache{d: d}
}

type Object struct {
	c   *Cache
	key interface{}
}

func newObject(c *Cache, d time.Duration, key interface{}) *Object {
	o := Object{c: c, key: key}
	time.AfterFunc(d, func() {
		c.Delete(key)
	})
	return &o
}

func (c *Cache) Store(key, value string) {
	c.Map.Store(key, value)
	newObject(c, c.d, key)
}

func (c *Cache) Load(key string) (interface{}, bool) {
	return c.Map.Load(key)
}

func (c *Cache) Delete(key interface{}) {
	c.Map.Delete(key)
}
