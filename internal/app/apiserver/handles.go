package apiserver

import (
	"fmt"
	"net/http"
)

func loggingHandle(s *APIServer, r *http.Request) {
	msg := fmt.Sprintf("from: %s, method: %s, query: %s, user-agent: %s, remote addr: %s",
		r.Host,
		r.Method,
		r.URL.Query(),
		r.RemoteAddr,
		r.UserAgent())
	s.logInfo(msg)
}

func (s *APIServer) handleMakeOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusCreated)
		loggingHandle(s, r)
	}
}

func (s *APIServer) handleGetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusCreated)
		loggingHandle(s, r)
	}
}
