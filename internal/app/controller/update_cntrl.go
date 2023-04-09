package controller

import (
	"bytes"
	"net/http"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	UpdateCntrl interface {
		Form(ec echo.Context) error
		Submit(ec echo.Context) error
	}
	UpdateCntrlImpl struct{}
)

func init() {
	di.Provide(func() UpdateCntrl { return &UpdateCntrlImpl{} })
}

func (*UpdateCntrlImpl) Form(ec echo.Context) error {
	var buf bytes.Buffer

	tmpl.ExecuteTemplate(&buf, UpdateFormTmplFile, UpdateFormTmplData{
		Opts: UpdateFormOpts{
			Genders:    []string{"Male", "Female"},
			Levels:     []string{"1. Freshman", "2. Sophomore", "4. Senior", "3. Junior"},
			States:     []string{"CA", "SD", "NC", "WI", "MD", "NE", "MA", "FL", "SC", "AK", "NY", "NH", "RI"},
			Majors:     []string{"English", "Math", "Art", "Physics"},
			Activities: []string{"Drama Club", "Lacrosse", "Basketball", "Baseball", "Debate", "Track & Field"},
		},
	})
	return ec.HTML(http.StatusServiceUnavailable, buf.String())
}

func (*UpdateCntrlImpl) Submit(ec echo.Context) error {
	return ec.HTML(http.StatusServiceUnavailable, "Unavailable")
}
