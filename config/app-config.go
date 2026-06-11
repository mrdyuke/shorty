package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App
		Server
		Pool
		DataBase
	}

	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
		Debug   bool   `env:"APP_DEBUG" envDefault:"false"`
	}

	Server struct {
		ServerPort            string        `env:"SERVER_PORT" envDefault:"8080"`
		ServerShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"25s"`
		ReadHeaderTimeout     time.Duration `env:"SERVER_READ_HEADER_TIMEOUT" envDefault:"2s"`
		ReadTimeout           time.Duration `env:"SERVER_READ_TIMEOUT"       envDefault:"5s"`
		WriteTimeout          time.Duration `env:"SERVER_WRITE_TIMEOUT"      envDefault:"10s"`
		IdleTimeout           time.Duration `env:"SERVER_IDLE_TIMEOUT"       envDefault:"30s"`
		MaxHeaderBytes        int           `env:"SERVER_MAX_HEADER_BYTES"   envDefault:"1048576"`
	}

	Pool struct {
		MaxConnections int32         `env:"POOL_MAX_CONNECTIONS" envDefault:"10"`
		MaxLifetime    time.Duration `env:"POOL_MAX_LIFETIME"    envDefault:"5m"`
	}

	DataBase struct {
		DBHost     string `env:"DATABASE_HOST"     envDefault:"localhost"`
		DBPort     uint16 `env:"DATABASE_PORT"     envDefault:"5432"`
		DBUser     string `env:"DATABASE_USER"     envDefault:"postgres"`
		DBName     string `env:"DATABASE_NAME"     envDefault:"postgres"`
		DBPassword string `env:"DATABASE_PASSWORD" envDefault:"root"`
	}
)

func NewConfig() (*Config, error) {

	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
