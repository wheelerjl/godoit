package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/wheelerjl/godoit/internal/variables"
)

const (
	port    = 5432
	timeout = 3
	name    = "godoit"
)

type Client struct {
	DB *sql.DB
}

func NewDatabaseClient(variables variables.Variables) (client Client, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable connect_timeout=%d",
		variables.DbHost, port, variables.DbUser, variables.DbPass, name, timeout)

	client.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return client, err
	}

	return client, nil
}
