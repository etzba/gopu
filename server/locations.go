package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/etzba/gopu/wire"
)

func (s *Server) getLocations() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		s.Respoder.SendOK(w, locations)
	}
}

func (s *Server) postLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		loc := wire.Location{}
		if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
			s.Logger.Error("Failed to create new decoder", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		location := wire.Location{
			Name:       loc.Name,
			Address:    loc.Address,
			Longtitude: loc.Longtitude,
			Latitude:   loc.Latitude,
		}

		s.locations = append(s.locations, location)
		s.Logger.Info("Add a new location to memory " + loc.Name)
		s.Respoder.SendOK(w)
	}
}

func (s *Server) getLocationById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		now := time.Now()
		defer s.shipper.Collect(now, r)
		idStr, _ := strings.CutPrefix(r.URL.Path, "/locations/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error("Failed to convert string to integer", err)
			s.Respoder.SendError(w, err)
			return
		}

		loc := s.locations[id]
		s.Logger.Info("Location is " + loc.Name)
		s.Respoder.SendOK(w)
	}
}
