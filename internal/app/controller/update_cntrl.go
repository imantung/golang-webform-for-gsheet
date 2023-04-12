package controller

import (
	"bytes"
	"fmt"
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

// Show update form
func (u *UpdateCntrlImpl) Form(ec echo.Context) error {
	var data struct {
		Gsheet string `param:"gsheet"`
		Row    int    `param:"row"`
		Error  string `query:"err"`
	}

	if err := ec.Bind(&data); err != nil {
		return err
	}

	if err := validate.Struct(&data); err != nil {
		return ec.HTML(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("%+v\n", data) // debug

	var buf bytes.Buffer
	tmpl.ExecuteTemplate(&buf, UpdateFormTmplFile, UpdateFormTmplData{
		Row:   data.Row,
		Error: data.Error,
		Opts:  DefaultUpdateFormOpts,
	})
	return ec.HTML(http.StatusServiceUnavailable, buf.String())
}

// Accept form submit
func (u *UpdateCntrlImpl) Submit(ec echo.Context) error {
	var data struct {
		Gsheet   string `param:"gsheet" validate:"required"`
		Row      int    `param:"row" validate:"required"`
		Name     string `form:"name" validate:"required"`
		Gender   string `form:"gender" validate:"required"`
		Level    string `form:"level" validate:"required"`
		State    string `form:"state" validate:"required"`
		Major    string `form:"major" validate:"required"`
		Activity string `form:"activity" validate:"required"`
	}

	if err := ec.Bind(&data); err != nil {
		return err
	}

	if err := validate.Struct(&data); err != nil {
		path := fmt.Sprintf("/update/%s/r/%d?err=%s",
			data.Gsheet, data.Row, validationErrorMessage(err))
		return ec.Redirect(http.StatusSeeOther, path)
	}

	fmt.Printf("%+v\n", data) // debug

	return ec.HTML(http.StatusServiceUnavailable, "Unavailable")
}
