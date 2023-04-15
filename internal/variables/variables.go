package variables

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type Variables struct {
	Env    string `json:"env" envconfig:"environment"`
	Port   int    `json:"port" default:"8080"`
	Debug  bool   `json:"debug" default:"false"`
	DbHost string `json:"db_host" envconfig:"db_host" required:"true"`
	DbUser string `json:"db_user" envconfig:"db_user" required:"true"`
	DbPass string `json:"db_pass" envconfig:"db_pass" required:"true"`
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
