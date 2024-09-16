package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const maxUploadSize = 5 * 1024 * 1024 * 1024

func (s *Server) uploadFileHandlerfunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			s.Logger.Error("failed to parse multipart form", err)
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			s.Respoder.SendBadRequest(w)
			return
		}
		defer file.Close()

		buff := make([]byte, 0)
		_, err = file.Read(buff)
		if err != nil {
			s.Respoder.SendError(w, err)
			return
		}

		if err = os.MkdirAll("./uploads", os.ModePerm); err != nil {
			s.Respoder.SendError(w, err)
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			s.Respoder.SendError(w, err)
			return
		}
		defer f.Close()

		_, err = f.Write(buff)
		if err != nil {
			s.Respoder.SendError(w, err)
			return
		}

		_, err = io.Copy(f, file)
		if err != nil {
			s.Respoder.SendError(w, err)
			return
		}

		s.Logger.Info("Upload file successfully" + "file: " + fileHeader.Filename)
		s.Respoder.SendNothing(w)
	}
}
