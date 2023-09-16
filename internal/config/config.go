package config

import (
	"github.com/Bakhram74/advertisement-server.git/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v5"
	"sync"
	"time"
)

type Config struct {
	HttpAddress         string        `yaml:"http_address" env-default:"0.0.0.0:8080"`
	TokenSymmetricKey   string        `yaml:"token_symmetric_key" env-default:"12wccsd323da23cds323cdscdsc23qwe"`
	AccessTokenDuration time.Duration `yaml:"access_token_duration"`
	Storage             StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Database string `yaml:"database" env-default:"advertisement"`
	Username string `yaml:"username" env-default:"root"`
	Password string `yaml:"password" env-default:"secret"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

var instance Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		err := cleanenv.ReadConfig("/Users/user/go/src/advertisement-server/config.yml", &instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return &instance
}
