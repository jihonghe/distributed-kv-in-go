package main

import (
	"distributed-kv-in-go/server/cache"
	"distributed-kv-in-go/server/http"
	"distributed-kv-in-go/server/tcp"
)

func main() {
	c := cache.New(cache.InMemoryCache)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
