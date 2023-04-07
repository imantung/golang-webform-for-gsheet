package controller

import (
	"net/http"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	UpdateCntrl struct {
	}
)

func init() {
	di.Provide(func() *UpdateCntrl { return &UpdateCntrl{} })
}

func (*UpdateCntrl) Form(ec echo.Context) error {
	return ec.HTML(http.StatusServiceUnavailable, "Unavailable")
}

func (*UpdateCntrl) Submit(ec echo.Context) error {
	return ec.HTML(http.StatusServiceUnavailable, "Unavailable")
}
