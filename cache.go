package gocache

import (
	"sync"
	"time"
)

type Cache struct {
	d time.Duration
	m sync.Map // we dont want to expose native method of sync.Map
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

func (c *Cache) Store(key, value interface{}) {
	c.m.Store(key, value)
	newObject(c, c.d, key)
}

type F = func() (interface{}, error)

func (c *Cache) StoreF(key interface{}, f F) error {
	value, err := f()
	if err != nil {
		return err
	}
	c.Store(key, value)
	return nil
}

func (c *Cache) Load(key interface{}) (interface{}, bool) {
	return c.m.Load(key)
}

func (c *Cache) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	actual, loaded = c.m.LoadOrStore(key, value)
	if !loaded {
		newObject(c, c.d, key)
	}
	return
}

func (c *Cache) LoadOrStoreF(key interface{}, f F) (actual interface{}, loaded bool, err error) {
	actual, loaded = c.Load(key)
	if !loaded {
		actual, err = f()
		if err != nil {
			return
		} else {
			c.Store(key, actual)
		}
	}
	return
}

func (c *Cache) Delete(key interface{}) {
	c.m.Delete(key)
}
