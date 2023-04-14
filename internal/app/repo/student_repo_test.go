package repo_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/repo"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func TestStudentRepo_FindOne(t *testing.T) {
	testcases := []struct {
		TestName        string
		GSheet          string
		Row             int
		RespError       error
		Values          [][]interface{}
		ExpectedStudent *repo.Student
		ExpectedError   string
	}{
		{
			TestName: "success",
			GSheet:   "gsheet",
			Row:      99,
			Values: [][]interface{}{
				{"99", "some-name", "some-gender", "some-level", "some-state", "some-major", "some-activity"},
			},
			ExpectedStudent: &repo.Student{
				Row:      99,
				Name:     "some-name",
				Gender:   "some-gender",
				Level:    "some-level",
				State:    "some-state",
				Major:    "some-major",
				Activity: "some-activity",
			},
		},
		{
			TestName:      "missing row",
			GSheet:        "gsheet",
			Row:           99,
			RespError:     errors.New("some-error"),
			ExpectedError: "Get \"https://sheets.googleapis.com/v4/spreadsheets/gsheet/values/Database%2199%3A99?alt=json&prettyPrint=false\": some-error",
		},
		{
			TestName:      "missing row",
			GSheet:        "gsheet",
			Row:           99,
			Values:        [][]interface{}{},
			ExpectedError: "missing row: 99",
		},
		{
			TestName: "missing column",
			GSheet:   "gsheet",
			Row:      99,
			Values: [][]interface{}{
				{"99", "some-name"},
			},
			ExpectedError: "expecting 7 column",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			path := fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%s/values/%s%%21%d%%3A%d?alt=json&prettyPrint=false", tt.GSheet, repo.SheetName, tt.Row, tt.Row)

			httpmock.RegisterResponder("GET", path,
				func(req *http.Request) (*http.Response, error) {
					if tt.RespError != nil {
						return nil, tt.RespError
					}
					return httpmock.NewJsonResponse(200, map[string]interface{}{
						"majorDimension": "ROWS",
						"values":         tt.Values,
					})
				},
			)

			gs, err := sheets.NewService(context.Background(), option.WithHTTPClient(http.DefaultClient))
			assert.NoError(t, err)

			impl := repo.StudentRepoImpl{
				Gs: gs,
			}

			student, err := impl.FindOne(tt.GSheet, tt.Row)
			if tt.ExpectedError != "" {
				assert.EqualError(t, err, tt.ExpectedError)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.ExpectedStudent, student)

		})
	}

}
