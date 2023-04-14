package repo

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"go.uber.org/dig"
	"google.golang.org/api/sheets/v4"
)

//go:generate mockgen -source=$GOFILE -destination=$PROJ/internal/generated/mock_$GOPACKAGE/$GOFILE

type (
	Student struct {
		Row      int
		Name     string
		Gender   string
		Level    string
		State    string
		Major    string
		Activity string
	}
	StudentRepo interface {
		FindOne(gsheet string, row int) (*Student, error)
		Update(gsheet string, row int, student *Student) error
	}
	StudentRepoImpl struct {
		dig.In
		Gs *sheets.Service
	}
)

const (
	SheetName = "Database"
)

func init() {
	di.Provide(func(impl StudentRepoImpl) StudentRepo {
		return &impl
	})
}

func (s *StudentRepoImpl) FindOne(gsheet string, row int) (*Student, error) {
	resp, err := s.Gs.Spreadsheets.Values.Get(gsheet, s.getRange(row)).Do()
	if err != nil {
		return nil, err
	}
	if len(resp.Values) < 1 {
		return nil, fmt.Errorf("missing row: %d", row)
	}

	return s.convertToStudent(resp.Values[0])
}

func (s *StudentRepoImpl) Update(gsheet string, row int, student *Student) error {
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, s.convertToRowValue(student))
	_, err := s.Gs.Spreadsheets.Values.Update(gsheet, s.getRange(row), &vr).ValueInputOption("RAW").Do()
	return err
}

func (s *StudentRepoImpl) convertToStudent(val []interface{}) (*Student, error) {
	if len(val) < 7 {
		return nil, errors.New("expecting 7 column")
	}
	row, _ := strconv.Atoi(val[0].(string))
	return &Student{
		Row:      row,
		Name:     val[1].(string),
		Gender:   val[2].(string),
		Level:    val[3].(string),
		State:    val[4].(string),
		Major:    val[5].(string),
		Activity: val[6].(string),
	}, nil
}

func (s *StudentRepoImpl) convertToRowValue(student *Student) []interface{} {
	return []interface{}{
		student.Row,
		student.Name,
		student.Gender,
		student.Level,
		student.State,
		student.Major,
		student.Activity,
	}
}

func (s *StudentRepoImpl) getRange(row int) string {
	return fmt.Sprintf("%s!%d:%d", SheetName, row, row)
}
