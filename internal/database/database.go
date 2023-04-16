package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wheelerjl/godoit/internal/variables"
)

type Client struct {
	DB *pgxpool.Pool
}

func NewDatabaseClient(variables variables.Variables) (client Client, err error) {
	connectionURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?search_path=%s",
		variables.DatabaseUser,
		variables.DatabasePass,
		variables.DatabaseHost,
		5432,
		"godoit",
		"godoit",
	)

	client.DB, err = pgxpool.New(context.Background(), connectionURL)
	if err != nil {
		return client, err
	}

	return client, nil
}
