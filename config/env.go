package config

import "github.com/kelseyhightower/envconfig"

type Env struct {
	Port string `required:"true"`

	MongoUrl            string `required:"true" split_words:"true"`
	MongoDatabaseName   string `required:"true" split_words:"true"`
	MongoCollectionName string `required:"true" split_words:"true"`
	MongoTimeout        int    `required:"true" split_words:"true"`

	Blacklists string `required:"true" split_words:"true"`
	AdminToken string `required:"true" split_words:"true"`
}

func GetEnv() *Env {
	env := &Env{}

	if err := envconfig.Process("SHORT_URL_SERVICE", env); err != nil {
		panic(err)
	}

	return env
}
