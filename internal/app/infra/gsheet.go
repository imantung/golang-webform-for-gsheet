package infra

import (
	"context"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func init() {
	di.Provide(NewGSheetService)
}

func NewGSheetService(cfg *Config) (*sheets.Service, error) {
	ctx := context.Background()
	return sheets.NewService(ctx,
		option.WithCredentialsFile(cfg.GSheet.CredPath), // or option.WithCredentialsFile() to read credential from file
	)
}
