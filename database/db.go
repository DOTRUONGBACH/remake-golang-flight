package database

import (
	"context"
	"jet/ent"
	"log"
	"os"
)

type Repository interface {
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (Repository, error) {
	client, err := ent.Open("postgres", connection_str)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating db schema from the automation migration tool")
		os.Exit(1)
	}

	return &PostgresDB{Client: client}, nil
}
