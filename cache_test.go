package gocache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second)
	_, ok := cache.Load("test")
	assert.False(t, ok)
	cache.Store("test", "value")
	val, ok := cache.Load("test")
	assert.True(t, ok)
	assert.Equal(t, "value", val.(string))
	time.Sleep(2 * time.Second)
	_, ok = cache.Load("test")
	assert.False(t, ok)
}

func TestDeleteBeforeInvalidation(t *testing.T) {
	cache := NewCache(time.Second)
	_, ok := cache.Load("test")
	assert.False(t, ok)
	cache.Store("test", "value")
	val, ok := cache.Load("test")
	assert.True(t, ok)
	assert.Equal(t, "value", val.(string))
	cache.Delete("test")
	time.Sleep(2 * time.Second)
	_, ok = cache.Load("test")
	assert.False(t, ok)
}
