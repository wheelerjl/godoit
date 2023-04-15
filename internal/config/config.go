package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env   string `json:"env" envconfig:"environment"`
	Port  int    `json:"port" default:"8080"`
	Debug bool   `json:"debug" default:"false"`
}

func ProcessConfig() (Config, error) {
	config := Config{}
	if err := envconfig.Process("godoit", &config); err != nil {
		return Config{}, err
	}
	if err := validator.New().Struct(config); err != nil {
		return Config{}, err
	}
	return config, nil
}
