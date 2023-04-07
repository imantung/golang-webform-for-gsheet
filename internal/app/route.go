package app

import (
	"github.com/imantung/golang_webform_for_gsheet/internal/app/controller"
	"github.com/labstack/echo/v4"
)

func setRoute(
	e *echo.Echo,
	cntrl *controller.UpdateCntrl,
) {

	e.GET("update", cntrl.Form)
	e.POST("update", cntrl.Form)

}
