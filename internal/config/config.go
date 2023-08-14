package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

const (
	configPathEnv = "CONFIG_PATH"
)

// TODO: add ENVs
type Config struct {
	Env         string           `yaml:"environment" env-default:"prod"`
	StoragePath string           `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServerConfig `yaml:"http_server" env-default:"local"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv(configPathEnv)
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
