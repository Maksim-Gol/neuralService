package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type (
	HTTPConfig struct {
		Port string `yaml:"port" env-default:":3000"`
	}

	PostgresConfig struct {
		DBHost string `yaml:"host" env-default:"localhost"`
		DBPort string `yaml:"port" env-default:"5432"`
		DBUser string `yaml:"user" env-default:"user"`
		DBName string `yaml:"name" env-default:"postgres"`
	}

	Config struct {
		Env string `yaml:"env" env-default:"local"`

		HTTP     HTTPConfig     `yaml:"http"`
		Postgres PostgresConfig `yaml:"postgres"`
	}
)

func MustLoad() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", err)
	}

	return &cfg

}
