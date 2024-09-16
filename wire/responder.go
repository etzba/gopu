package wire

import (
	"net/http"

	"github.com/etzba/gopu/pkg/logger"
)

type Responder interface {
	SendOK(w http.ResponseWriter)
	SendNothing(w http.ResponseWriter)
	SendError(w http.ResponseWriter, err error)
	SendBadRequest(w http.ResponseWriter)
}

type Respond struct {
	Logger *logger.Log
}

func (r Respond) SendOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK!"))
}

func (r Respond) SendNothing(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("No content"))
}

func (r Respond) SendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (r Respond) SendBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request"))
}
