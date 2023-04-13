package infra

import (
	"net/http"
	"text/template"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

var tmpl = template.Must(template.ParseFiles("web/error/error.gohtml"))

func init() {
	di.Provide(NewServer)
}

func NewServer(cfg *Config) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.HideBanner = true
	e.Debug = cfg.Debug
	return e
}

func customHTTPErrorHandler(err error, ec echo.Context) {
	code := http.StatusInternalServerError

	message := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}
	ec.Logger().Error(err)

	w := ec.Response().Writer

	tmpl.Execute(w, struct {
		Code    int
		Message string
		Detail  string
	}{
		Code:    code,
		Message: message,
		Detail:  err.Error(),
	})

}
