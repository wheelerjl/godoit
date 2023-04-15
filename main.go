package main

import (
	"log"

	"github.com/wheelerjl/godoit/internal/database"
	"github.com/wheelerjl/godoit/internal/server"
	"github.com/wheelerjl/godoit/internal/variables"
)

func main() {
	var (
		config = server.Config{}
		err    error
	)

	config.Variables, err = variables.ProcessVariables()
	if err != nil {
		log.Fatalf("unable to process config: %v", err)
	}

	config.Database, err = database.NewDatabaseClient(config.Variables)
	if err != nil {
		log.Fatalf("unable to create database client: %v", err)
	}

	if err := server.NewServer(config).Start(); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}
