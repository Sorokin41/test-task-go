package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string `yaml:"address" env-default:":8080"`
	Timeout     int    `yaml:"timeout" env-default:"5s"`
	IdleTimeout int    `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return &config
}
