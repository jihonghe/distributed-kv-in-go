package cache

import (
	"log"
	"strings"
)

const (
	InMemoryCache = "inMemoryCache"
)

type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Del(key string) error
	// GetStat 获取缓存系统状态
	GetStat() Stat
}

func New(typ string) Cache {
	var c Cache
	if strings.EqualFold(typ, InMemoryCache) {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("unknown cache type " + typ)
	}

	log.Println(typ, "ready to serve")
	return c
}
