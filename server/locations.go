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
		s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		s.Respoder.SendOK(w, locations)
		s.shipper.SetCurrentUsersEnd(r)
	}
}

func (s *Server) postLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		s.shipper.Collect(now, r)
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

		locations = append(locations, location)
		s.Logger.Info("Add a new location to memory " + loc.Name)
		s.Respoder.SendOK(w)
		s.shipper.SetCurrentUsersEnd(r)
	}
}

func (s *Server) getLocationById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		idStr, _ := strings.CutPrefix(r.URL.Path, "/locations/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error("Failed to convert string to integer", err)
			s.Respoder.SendError(w, err)
			return
		}

		loc := locations[id]
		s.Logger.Info("Location is " + loc.Name)
		s.Respoder.SendOK(w)
		s.shipper.SetCurrentUsersEnd(r)
	}
}
