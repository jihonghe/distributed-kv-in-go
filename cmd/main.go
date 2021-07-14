package main

import (
	"distributed-kv-in-go/cache"
	"distributed-kv-in-go/server/httpserver"
)

func main() {
	c := cache.New(cache.InMemoryCache)
	httpserver.New(c).Listen()
}
