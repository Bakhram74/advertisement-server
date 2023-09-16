package config

import (
	"github.com/Bakhram74/advertisement-server.git/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v5"
	"time"
)

type Config struct {
	HttpAddress         string        `yaml:"http_address"`
	TokenSymmetricKey   string        `yaml:"token_symmetric_key"`
	AccessTokenDuration time.Duration `yaml:"access_token_duration"`
	Storage             StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port" `
	Database string `yaml:"database"`
	Username string `yaml:"username" `
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
}

var instance Config

func GetConfig() Config {

	logger := logging.GetLogger()
	logger.Info("read application configuration")
	err := cleanenv.ReadConfig("/Users/user/go/src/advertisement-server/config.yaml", &instance)
	if err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		logger.Info(help)
		logger.Fatal(err)
	}

	return instance
}
