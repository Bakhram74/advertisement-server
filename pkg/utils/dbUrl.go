package utils

import (
	"fmt"
	"github.com/Bakhram74/advertisement-server.git/internal/config"
)

func GetPostgresUrl(cfg config.StorageConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)
}
