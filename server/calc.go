package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) getResultByUrlPath() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		s.Respoder.SendOK(w)
	}
}

// https://go.dev/blog/json
// client sent json payload as follow: '[4.543, 3, 1]'. floats and integers allowed, string will return bad request
func (s *Server) postNumbersAddition() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		var nums []interface{}
		if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
			s.Logger.Error("Failed to create new decoder", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		intResult := 0
		floatResult := 0.0000
		for _, num := range nums {
			switch v := num.(type) {
			case int:
				intResult = intResult + v
			case float64:
				floatResult = floatResult + v
			default:
				s.Respoder.SendBadRequest(w)
				return
			}
		}

		switch {
		case intResult != 0 && floatResult == 0:
			var result interface{} = []int{intResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		case floatResult != 0 && floatResult == 0:
			var result interface{} = []float64{floatResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		default:
			var result interface{} = []float64{floatResult + float64(intResult)}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		}
	}
}

func (s *Server) postNumbersSubtraction() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		var nums []interface{}
		if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
			s.Logger.Error("Failed to create new decoder", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		intResult := 0
		floatResult := 0.0000
		for _, num := range nums {
			switch v := num.(type) {
			case int:
				if intResult == 0 {
					intResult = intResult + v
				} else {
					intResult = intResult - v
				}
			case float64:
				if floatResult == 0 {
					floatResult = floatResult + v
				} else {
					floatResult = floatResult - v
				}
			default:
				s.Respoder.SendBadRequest(w)
				return
			}
		}

		switch {
		case intResult != 0 && floatResult == 0:
			var result interface{} = []int{intResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		case floatResult != 0 && floatResult == 0:
			var result interface{} = []float64{floatResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		default:
			var result interface{} = []float64{floatResult + float64(intResult)}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		}
	}
}
func (s *Server) postNumbersMultiply() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		var nums []interface{}
		if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
			s.Logger.Error("Failed to create new decoder", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		intResult := 0
		floatResult := 0.0000
		for _, num := range nums {
			switch v := num.(type) {
			case int:
				if intResult == 0 {
					intResult = intResult + v
				} else {
					intResult = intResult * v
				}
			case float64:
				if floatResult == 0 {
					floatResult = floatResult + v
				} else {
					floatResult = floatResult * v
				}
			default:
				s.Respoder.SendBadRequest(w)
				return
			}
		}

		switch {
		case intResult != 0 && floatResult == 0:
			var result interface{} = []int{intResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
			return
		case floatResult != 0 && floatResult == 0:
			var result interface{} = []float64{floatResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		default:
			var result interface{} = []float64{floatResult + float64(intResult)}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		}
	}
}

func (s *Server) postNumbersDivide() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer s.shipper.Collect(now, r)
		s.Logger.Info("Server go request" + " method: " + r.Method + " uri: " + r.RequestURI)
		var nums []interface{}
		if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
			s.Logger.Error("Failed to create new decoder", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		intResult := 0
		floatResult := 0.0000
		for _, num := range nums {
			switch v := num.(type) {
			case int:
				if intResult == 0 {
					intResult = intResult + v
				} else {
					intResult = intResult / v
				}
			case float64:
				if floatResult == 0 {
					floatResult = floatResult + v
				} else {
					floatResult = floatResult / v
				}
			default:
				s.Respoder.SendBadRequest(w)
				return
			}
		}

		switch {
		case intResult != 0 && floatResult == 0:
			var result interface{} = []int{intResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		case floatResult != 0 && floatResult == 0:
			var result interface{} = []float64{floatResult}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		default:
			var result interface{} = []float64{floatResult + float64(intResult)}
			s.Logger.Info(fmt.Sprintf("Given numbers were %v and result was %v", nums, result))
			s.Respoder.SendOK(w, result)
		}
	}
}
