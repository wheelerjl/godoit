package main

import (
	"log"

	"github.com/wheelerjl/godoit/internal/config"
	"github.com/wheelerjl/godoit/internal/server"
)

func main() {
	conf, err := config.ProcessConfig()
	if err != nil {
		log.Fatalf("unable to process config: %v", err)
	}

	s := server.NewServer(conf)
	if err := s.Start(); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}
