package infra

import (
	"context"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func init() {
	di.Provide(NewServer)
	di.Provide(NewGSheetService)
}

func NewServer(cfg *Config) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.Debug = cfg.Debug
	return e
}

func NewGSheetService(cfg *Config) (*sheets.Service, error) {
	ctx := context.Background()
	return sheets.NewService(ctx,
		option.WithCredentialsFile(cfg.GSheet.CredPath), // or option.WithCredentialsFile() to read credential from file
	)
}
