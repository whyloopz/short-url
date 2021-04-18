package config

import "github.com/kelseyhightower/envconfig"

type Env struct {
	Port string `required:"true"`
}

func GetEnv() *Env {
	env := &Env{}

	if err := envconfig.Process("SHORT_URL_SERVICE", env); err != nil {
		panic(err)
	}

	return env
}
