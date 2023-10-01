package core

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	
}

func LoadConfig() *Config {
	// Loading the environment variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		// sentry.CaptureException(err)
		log.Info("unable to load .env file: %e", err)
	}

	cfg := Config{} // 👈 new instance of `Config`

	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		// sentry.CaptureException(err)
		log.Info("unable to parse ennvironment variables: %e", err)
	}
	return &cfg
}
