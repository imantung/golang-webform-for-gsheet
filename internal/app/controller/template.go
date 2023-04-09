package controller

import (
	"html/template"
)

var tmpl = template.Must(template.ParseGlob("web/template/*.gohtml"))

var UpdateFormTmplFile = "update_form.gohtml"

type (
	UpdateFormTmplData struct {
		Opts UpdateFormOpts
	}
	UpdateFormOpts struct {
		Genders    []string
		Levels     []string
		States     []string
		Majors     []string
		Activities []string
	}
)
