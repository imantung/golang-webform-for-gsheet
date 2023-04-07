package infra

import (
	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

func init() {
	di.Provide(NewServer)
}

// NewServer return new instance of server
func NewServer(cfg *Config) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.Debug = cfg.Debug
	return e
}
