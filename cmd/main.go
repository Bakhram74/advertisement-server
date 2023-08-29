package main

import (
	"fmt"
	"github.com/Bakhram74/advertisement-server.git/internal/config"
	"github.com/Bakhram74/advertisement-server.git/pkg/logging"
	_ "github.com/jackc/pgx/v5"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("test")
	cfg := config.GetConfig()
	fmt.Print(cfg)
}
