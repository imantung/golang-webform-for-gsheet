package infra

import (
	"time"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/kelseyhightower/envconfig"
)

// NOTE: Generate .env or .envrc file using `go generate ./internal/app/infra/config.go > .envrc `
//go:generate echo export APP_ADDRESS=:8089
//go:generate echo export APP_READ_TIMEOUT=5s
//go:generate echo export APP_WRITE_TIMEOUT=10s
//go:generate echo export APP_DEBUG=true

const (
	ConfigPrefix = "APP"
)

type (
	Config struct {
		Address      string        `envconfig:"ADDRESS" default:":8089" required:"true"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
		WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
		Debug        bool          `envconfig:"DEBUG" default:"true"`
	}
)

func init() {
	di.Provide(LoadAppConfig)
}

func LoadAppConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ConfigPrefix, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
