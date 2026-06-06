package config

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App
		Server
		Pool
		DataBase
	}

	App struct {
		Name      string `env:"APP_NAME"`
		Version   string `env:"APP_VERSION"`
		IsRelease bool
	}

	Server struct {
		ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	}

	Pool struct {
		MaxIdle        int           `env:"POOL_MAX_IDLE"        envDefault:"5"`
		MaxConnections int           `env:"POOL_MAX_CONNECTIONS" envDefault:"10"`
		MaxLifetime    time.Duration `env:"POOL_MAX_LIFETIME"    envDefault:"5m"`
	}

	DataBase struct {
		Host     string `env:"DATABASE_HOST"     envDefault:"localhost"`
		DBPort   string `env:"DATABASE_PORT"     envDefault:"5432"`
		User     string `env:"DATABASE_USER"     envDefault:"postgres"`
		Name     string `env:"DATABASE_NAME"     envDefault:"postgres"`
		Password string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	}
)

func NewConfig() (*Config, error) {

	cfg := new(Config)

	isRelease := flag.Bool("release", false, "")
	flag.Parse()
	cfg.IsRelease = *isRelease

	if !cfg.IsRelease {
		if err := godotenv.Load(".env"); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
