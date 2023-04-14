package infra

import (
	"context"
	"html/template"
	"net/http"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func init() {
	di.Provide(LoadTemplate)
	di.Provide(NewServer)
	di.Provide(NewGSheetService)
}

func LoadTemplate(cfg *Config) *template.Template {
	return template.Must(template.ParseGlob(cfg.Pages))
}

func NewServer(cfg *Config, tmpl *template.Template) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, ec echo.Context) {
		code := http.StatusInternalServerError

		message := "Internal Server Error"
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}
		ec.Logger().Error(err)

		w := ec.Response().Writer

		tmpl.ExecuteTemplate(w, "error.gohtml", struct {
			Code    int
			Message string
			Detail  string
		}{
			Code:    code,
			Message: message,
			Detail:  err.Error(),
		})

	}
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
