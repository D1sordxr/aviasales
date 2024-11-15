package config

import (
	db "aviasales/src/internal/db/config"
	api "aviasales/src/internal/http/config"
)

type Config struct {
	AppConfig     `yaml:"app"`
	db.DBConfig   `yaml:"db"`
	api.APIConfig `yaml:"http_server"`
}

type AppConfig struct {
	Mode string `yaml:"mode" env-default:"local"`
}
