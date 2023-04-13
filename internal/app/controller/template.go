package controller

import (
	"html/template"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/repo"
)

var tmpl = template.Must(template.ParseGlob("web/template/*.gohtml"))

// NOTE: put template related variable and type here

// >>>>>>>>>>>>>>>>>>>> update_form.gohtml >>>>>>>>>>>>>>>>>>>>>

var (
	UpdateFormTmplFile    = "update_form.gohtml"
	DefaultUpdateFormOpts = UpdateFormOpts{
		Genders:    []string{"Male", "Female"},
		Levels:     []string{"1. Freshman", "2. Sophomore", "4. Senior", "3. Junior"},
		States:     []string{"CA", "SD", "NC", "WI", "MD", "NE", "MA", "FL", "SC", "AK", "NY", "NH", "RI"},
		Majors:     []string{"English", "Math", "Art", "Physics"},
		Activities: []string{"Drama Club", "Lacrosse", "Basketball", "Baseball", "Debate", "Track & Field"},
	}
)

type (
	UpdateFormTmplData struct {
		Row     int
		Error   string
		Success string
		Student *repo.Student
		Opts    UpdateFormOpts
	}
	UpdateFormOpts struct {
		Genders    []string
		Levels     []string
		States     []string
		Majors     []string
		Activities []string
	}
)

// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
