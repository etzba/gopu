package server

import (
	"net/http"
	"time"
)

func (s *Server) getLocations() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		httpRequestCounter.Inc()
		numberOfConcurrentUsers.Add(1)
		now := time.Now()
		numberOfConcurrentUsers.Dec()
		httpRequestDuration.Observe(float64(time.Since(now)))
		s.Respoder.SendOK(w)
	}
}

func (s *Server) getLocationById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		httpRequestCounter.Inc()
		numberOfConcurrentUsers.Add(1)
		now := time.Now()
		numberOfConcurrentUsers.Dec()
		httpRequestDuration.Observe(float64(time.Since(now)))
		s.Respoder.SendOK(w)
	}
}
