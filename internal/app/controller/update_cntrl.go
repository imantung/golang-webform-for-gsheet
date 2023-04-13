package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/imantung/golang_webform_for_gsheet/internal/app/repo"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	UpdateCntrl interface {
		Form(ec echo.Context) error
		Submit(ec echo.Context) error
	}
	UpdateCntrlImpl struct {
		dig.In
		Repo repo.StudentRepo
	}
)

func init() {
	di.Provide(func(impl UpdateCntrlImpl) UpdateCntrl { return &impl })
}

// Show update form
func (u *UpdateCntrlImpl) Form(ec echo.Context) error {
	var data struct {
		Gsheet  string `param:"gsheet"`
		Row     int    `param:"row"`
		Error   string `query:"err"`
		Success string `query:"success"`
	}

	if err := ec.Bind(&data); err != nil {
		return err
	}

	student, err := u.Repo.FindOne(data.Gsheet, data.Row)
	if err != nil {
		return err
	}

	if err := validate.Struct(&data); err != nil {
		return ec.HTML(http.StatusBadRequest, err.Error())
	}

	var buf bytes.Buffer
	tmpl.ExecuteTemplate(&buf, UpdateFormTmplFile, UpdateFormTmplData{
		Row:     data.Row,
		Error:   data.Error,
		Success: data.Success,
		Student: student,
		Opts:    DefaultUpdateFormOpts,
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

	path := fmt.Sprintf("/update/%s/r/%d", data.Gsheet, data.Row)
	if err := validate.Struct(&data); err != nil {
		return ec.Redirect(http.StatusSeeOther, path+"?err="+valdnErrMsg(err))
	}

	if err := u.Repo.Update(data.Gsheet, data.Row, &repo.Student{
		Row:      data.Row,
		Name:     data.Name,
		Gender:   data.Gender,
		Level:    data.Level,
		State:    data.State,
		Major:    data.Major,
		Activity: data.Activity,
	}); err != nil {
		return ec.Redirect(http.StatusSeeOther, path+"?err="+err.Error())
	}

	return ec.Redirect(http.StatusSeeOther, path+"?success=Data updated")
}
