package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.EqualFold(r.Method, http.MethodGet) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cacheStat := h.GetStat()
	status, err := json.Marshal(cacheStat)
	if err != nil {
		log.Printf("failed to marshal cache.Stat[stat=%+v] to []byte, error: %s", cacheStat, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(status); err != nil {
		log.Printf("failed to write status data[status=%+v] to response, error: %s", cacheStat, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
