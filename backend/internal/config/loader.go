package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

func Load() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("config.toml", &cfg)
	if err != nil {
		// Try to read from environment only
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
