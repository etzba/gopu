package main

import (
	"github.com/etzba/gopu/pkg/logger"
	"github.com/etzba/gopu/server"
)

func main() {
	logger := logger.New()
	server := server.New(logger, ":8080")
	if err := server.Run(); err != nil {
		panic(err)
	}
}
