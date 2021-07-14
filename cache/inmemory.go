package cache

import (
	"errors"
	"fmt"
	"sync"
)

type inMemoryCache struct {
	// data sync.Map 支持并发安全读写的kv
	data sync.Map
	// 记录缓存状态
	Stat
}

func newInMemoryCache() *inMemoryCache {
	return new(inMemoryCache)
}

func (c *inMemoryCache) Set(key string, value []byte) error {
	c.data.Store(key, value)
	c.add(key, value)

	return nil
}

func (c *inMemoryCache) Get(key string) ([]byte, error) {
	if v, ok := c.data.Load(key); ok {
		value, yes := v.([]byte)
		if !yes {
			return nil, errors.New(fmt.Sprintf("cannot convert data(%+v) to type []byte", v))
		}
		return value, nil
	}

	return nil, nil
}

func (c *inMemoryCache) Del(key string) error {
	c.data.Delete(key)

	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}
