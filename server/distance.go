package server

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/etzba/gopu/wire"
)

func (s *Server) getDistance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		now := time.Now()
		defer s.shipper.Collect(now, r)
		dis := wire.Distance{}
		if err := json.NewDecoder(r.Body).Decode(&dis); err != nil {
			s.Logger.Error("Failed to create new decoder", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		wireDistance := wire.Distance{
			Formula:    dis.Formula,
			FromPlace:  dis.FromPlace,
			Longtitude: dis.Longtitude,
			Latitude:   dis.Latitude,
		}
		place := wire.Location{}
		for _, p := range locations {
			if dis.FromPlace == p.Name {
				place.Name = p.Name
				place.Latitude = p.Latitude
				place.Longtitude = p.Longtitude
			}
		}
		if place.Name == "" {
			s.Respoder.SendBadRequest(w)
			return
		}

		distance := getDistanceByPythagoras(place, wireDistance) * 100
		obj := make(map[string]float64, 1)
		obj["distance"] = distance
		s.Logger.Info(fmt.Sprintf("Client located about %fkm from %s", distance, place.Name))
		s.Respoder.SendOK(w, obj)
	}
}

//
// --------------------------------------------------------------------- helpers ----------------------------------------------------------
//

func getDistanceByPythagoras(loc1 wire.Location, loc2 wire.Distance) float64 {
	x := (loc1.Latitude - loc2.Latitude) * (loc1.Latitude - loc2.Latitude)
	y := (loc1.Longtitude - loc2.Longtitude) * (loc1.Longtitude - loc2.Longtitude)
	return math.Sqrt(x + y)
}

// https://en.wikipedia.org/wiki/Haversine_formula
func getDistanceByHarvestine() {

}
