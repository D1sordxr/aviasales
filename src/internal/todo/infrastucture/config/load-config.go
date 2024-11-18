package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"src/internal/infrastucture/config/config"
)

const BasicConfigPath = "./src/configs/app/local.yaml"

func NewConfig() (*config.Config, error) {
	var cfg config.Config

	if err := cleanenv.ReadConfig(BasicConfigPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err.Error())
	}

	return &cfg, nil
}
