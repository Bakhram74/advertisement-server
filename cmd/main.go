package main

import (
	"fmt"

	"github.com/Bakhram74/advertisement-server.git/internal/config"
	"github.com/Bakhram74/advertisement-server.git/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("test")
	cfg := config.GetConfig()
	fmt.Print(cfg)
}
