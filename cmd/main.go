package main

import "github.com/Bakhram74/advertisement-server.git/pkg/logging"

func main() {
	logger := logging.GetLogger()
	logger.Info("test")
}
