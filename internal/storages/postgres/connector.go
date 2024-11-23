package postgres

import (
	"database/sql"
	"gw-exchanger/internal/storages"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) storages.Storage {
	return &PostgresStorage{db: db}
}
