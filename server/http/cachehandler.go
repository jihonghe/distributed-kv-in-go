package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// cacheHandler
// 实现方法ServeHTTP(w http.ResponseWriter, r *http.Request)，就是一个handler
type cacheHandler struct {
	*Server
}

// ServeHTTP
// @description support method: put, get, del
// localhost:12345/cache/key-name
func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.RequestURI(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodGet:
		h.getHandler(w, r)
	case http.MethodPut:
		h.setHandler(w, r)
	case http.MethodDelete:
		h.delHandler(w, r)
	}
}

func (h *cacheHandler) delHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.RequestURI(), "/")[3]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Del(key); err != nil {
		log.Printf("failed to delete from cache by key=%s, error: %s", key, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *cacheHandler) setHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.RequestURI(), "/")[3]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body content, error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(value) == 0 {
		log.Printf("failed to set kv-pair(key=%s, value=%s), null value is not allowed", key, value)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.Set(key, value); err != nil {
		log.Printf("failed to set kv pair(key=%s, value=%s) to cache, error: %s", key, value, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *cacheHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.RequestURI(), "/")[3]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := h.Get(key)
	if err != nil {
		log.Printf("failed to get value by key=%s, error: %s", key, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(value) == 0 {
		log.Printf("key='%s' is not exist in cache", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(value); err != nil {
		log.Printf("failed to write value data[value=%s] to response, error: %s", string(value), err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}
