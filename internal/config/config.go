package config

import (
	"sync"

	"github.com/Bakhram74/advertisement-server.git/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host string `yaml:"host" env-default:"127.0.0.1"`
	Port string `yaml:"port" env-default:"8080"`

	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host" env-default:"localhost"`
	Port     string `json:"port" env-default:"5432"`
	Database string `json:"database" env-default:"advertisement"`
	Username string `json:"username" env-default:"localhost"`
	Password string `json:"password" env-default:"1234"`
	SSLMode  string `json:"ssl_mode" env-default:"disable"`
}

var instance Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		err := cleanenv.ReadEnv(&instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}

	})
	return &instance
}
