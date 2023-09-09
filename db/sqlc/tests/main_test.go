package tests

import (
	"context"
	"fmt"
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
	"github.com/Bakhram74/advertisement-server.git/internal/config"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	conn, err := pgx.Connect(context.Background(), postgresUrl(cfg.Storage))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	testQueries = db.New(conn)
	os.Exit(m.Run())
}

func postgresUrl(cfg config.StorageConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)
}
