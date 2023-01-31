package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Host string `env:"DB_HOST" envDefault:"localhost"`
	User string `env:"DB_USER" envDefault:"root"`
}

func LoadConfig() Database {

	dbConf := Database{}

	err := env.Parse(&dbConf)
	if err != nil {
		log.Error().Err(err).Msg("Could not load config")
	}

	return dbConf

}
