package main

import (
	"context"
	advertisement_server "github.com/Bakhram74/advertisement-server.git"
	"github.com/Bakhram74/advertisement-server.git/internal/config"
	"github.com/Bakhram74/advertisement-server.git/internal/handler"
	"github.com/Bakhram74/advertisement-server.git/internal/repository"
	"github.com/Bakhram74/advertisement-server.git/internal/service"
	"github.com/Bakhram74/advertisement-server.git/pkg/client/postgresql"
	"github.com/Bakhram74/advertisement-server.git/pkg/logging"
	_ "github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	storage := cfg.Storage

	connPool, err := postgresql.NewClient(context.TODO(), 3, postgresql.Config{
		Host:     storage.Host,
		Port:     storage.Port,
		Username: storage.Username,
		Password: storage.Password,
		DBName:   storage.Database,
		SSLMode:  storage.SSLMode,
	})
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	store := repository.NewStore(connPool)
	repos := repository.NewRepository(store)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	svr := new(advertisement_server.Server)
	logger.Infof("server is listening port %s:%s", cfg.Host, cfg.Port)
	if err = svr.Run(cfg.Host, cfg.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}
