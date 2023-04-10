package app

import (
	"github.com/imantung/golang_webform_for_gsheet/internal/app/controller"
	"github.com/labstack/echo/v4"
)

func setRoute(
	e *echo.Echo,
	updateCntrl controller.UpdateCntrl,
) {

	e.GET("/update/:gsheet/r/:row", updateCntrl.Form)
	e.POST("/update/:gsheet/r/:row", updateCntrl.Submit)
}
