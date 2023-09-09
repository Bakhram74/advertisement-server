package repository

import (
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	db.Querier
}
type SQLStore struct {
	*db.Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  db.New(connPool),
	}
}
