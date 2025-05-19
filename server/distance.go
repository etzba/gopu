package server

import (
	"math"
	"net/http"

	"github.com/etzba/gopu/wire"
)

func (s *Server) getDistance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		s.Respoder.SendOK(w)
	}
}

//
// --------------------------------------------------------------------- helpers ----------------------------------------------------------
//

func getDistanceByPythagoras(loc1, loc2 wire.Location) float64 {
	x := (loc1.Latitude - loc2.Latitude) * (loc1.Latitude - loc2.Latitude)
	y := (loc1.Longtitude - loc2.Longtitude) * (loc1.Longtitude - loc2.Longtitude)
	return math.Sqrt(x + y)
}

// https://en.wikipedia.org/wiki/Haversine_formula
func getDistanceByHarvestine() {

}
