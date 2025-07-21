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
		now := time.Now()
		defer s.shipper.Collect(now, r)
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			s.Logger.Error("Failed to parse multipart form", err)
			s.Respoder.SendError(w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			s.Respoder.SendBadRequest(w)
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				s.Logger.Error("could not close file", err)
			}
		}()

		buff := make([]byte, 0)
		_, err = file.Read(buff)
		if err != nil {
			s.Logger.Error("Cannot read buffer ", err)
			s.Respoder.SendError(w, err)
			return
		}

		if err = os.MkdirAll("./uploads", os.ModePerm); err != nil {
			s.Logger.Error("Cannot create dir ", err)
			s.Respoder.SendError(w, err)
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			s.Logger.Error("Cannot create file ", err)
			s.Respoder.SendError(w, err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				s.Logger.Error("could not close file", err)
			}
		}()

		_, err = f.Write(buff)
		if err != nil {
			s.Logger.Error("Cannot write file ", err)
			s.Respoder.SendError(w, err)
			return
		}

		_, err = io.Copy(f, file)
		if err != nil {
			s.Logger.Error("Cannot copy file to system ", err)
			s.Respoder.SendError(w, err)
			return
		}

		s.Logger.Info("Upload file successfully. filename: " + fileHeader.Filename)
		s.Respoder.SendNothing(w)
	}
}
