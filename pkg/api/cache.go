package api

import (
	"github.com/bluele/gcache"
	"github.com/sirupsen/logrus"
)

type Cache struct {
	OrderCache gcache.Cache
}

var Caching Cache

func (c *Cache) CreteCache() gcache.Cache {

	gc := gcache.New(50).
		LRU().
		Build()

	c.OrderCache = gc

	return gc
}

func (c *Cache) SetCache(key interface{}, value interface{}) {
	c.CreteCache()

	err := Caching.OrderCache.Set(key, value)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func (c *Cache) GetCache(key interface{}) (interface{}, error) {
	get, err := Caching.OrderCache.Get(key)
	if err != nil {
		logrus.Error("NO DATA IN CACHE", err)
		return nil, err
	}

	return get, nil
}
