package config

import (
	"time"
)

type Config struct {
	Name		 string         `json:"name" yaml:"name" env:"APP_NAME" envDefault:"stupid-news-backend"`
	Version      string         `json:"version" yaml:"version" env:"APP_VERSION" envDefault:"1.0.0"`
	Environment  string         `json:"environment" yaml:"environment" env:"APP_ENVIRONMENT" envDefault:"development"`
	Server      ServerConfig   `json:"server" yaml:"server" env:"SERVER"`
	Database    DatabaseConfig `json:"database" yaml:"database" env:"DATABASE"`
}

type ServerConfig struct {
	Host         string        `json:"host" yaml:"host" env:"HOST" envDefault:"localhost"`
	Port         int           `json:"port" yaml:"port" env:"PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout" env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout" env:"WRITE_TIMEOUT" envDefault:"10s"`
}

type DatabaseConfig struct {
	Host         string `json:"host" yaml:"host" env:"DB_HOST" envDefault:"localhost"`
	Port         int    `json:"port" yaml:"port" env:"DB_PORT" envDefault:"5432"`
	User         string `json:"user" yaml:"user" env:"DB_USER" envDefault:"postgres"`
	Password     string `json:"password" yaml:"password" env:"DB_PASSWORD"`
	Name         string `json:"name" yaml:"name" env:"DB_NAME" envDefault:"myapp"`
	SSLMode      string `json:"ssl_mode" yaml:"ssl_mode" env:"DB_SSL_MODE" envDefault:"disable"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns" env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns" env:"DB_MAX_IDLE_CONNS" envDefault:"25"`
}
