package variables

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type Variables struct {
	Env          string `json:"env" envconfig:"environment" required:"true"`
	Port         int    `json:"port" default:"8080"`
	Debug        bool   `json:"debug" default:"false"`
	DatabaseHost string `json:"database_host" envconfig:"database_host" required:"true"`
	DatabaseUser string `json:"database_user" envconfig:"database_user" required:"true"`
	DatabasePass string `json:"database_pass" envconfig:"database_pass" required:"true"`
	DiscordToken string `json:"discord_token" envconfig:"discord_token" required:"true"`
}

func ProcessVariables() (Variables, error) {
	variables := Variables{}
	if err := envconfig.Process("godoit", &variables); err != nil {
		return Variables{}, err
	}
	if err := validator.New().Struct(variables); err != nil {
		return Variables{}, err
	}
	return variables, nil
}
