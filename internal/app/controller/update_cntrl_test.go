package controller_test

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/imantung/golang_webform_for_gsheet/internal/app/controller"
	"github.com/imantung/golang_webform_for_gsheet/internal/app/repo"
	"github.com/imantung/golang_webform_for_gsheet/internal/generated/mock_repo"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestUpdateController_Form(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../web/template/*.gohtml"))

	testcases := []struct {
		TestName           string
		Gsheet             string
		Row                string
		OnMockStudentRepo  func(*mock_repo.MockStudentRepo)
		ExpectedError      string
		ExpectedOutputFile string
	}{
		{
			TestName: "error",
			Gsheet:   "gsheet",
			Row:      "99",
			OnMockStudentRepo: func(m *mock_repo.MockStudentRepo) {
				m.EXPECT().FindOne("gsheet", 99).Return(nil, errors.New("some-error"))
			},
			ExpectedError: "some-error",
		},
		{
			TestName: "success",
			Gsheet:   "gsheet",
			Row:      "99",
			OnMockStudentRepo: func(m *mock_repo.MockStudentRepo) {
				m.EXPECT().FindOne("gsheet", 99).
					Return(&repo.Student{
						Row:      99,
						Name:     "some-name",
						Gender:   "Male",
						Level:    "3. Junior",
						State:    "SC",
						Major:    "English",
						Activity: "Lacrosse",
					}, nil)
			},
			ExpectedOutputFile: "test_case/update_form_success_1.html",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			ec := e.NewContext(req, rec)
			ec.SetPath("update/:gsheet/r/:row")
			ec.SetParamNames("gsheet", "row")
			ec.SetParamValues(tt.Gsheet, tt.Row)

			ctrl := gomock.NewController(t)
			repo := mock_repo.NewMockStudentRepo(ctrl)
			if tt.OnMockStudentRepo != nil {
				tt.OnMockStudentRepo(repo)
			}

			cntrl := controller.NewUpdateCntrl(repo, tmpl)

			err := cntrl.Form(ec)
			if tt.ExpectedError != "" {
				require.EqualError(t, err, tt.ExpectedError)
			} else {
				// fmt.Println(rec.Body.String())
				b, _ := ioutil.ReadFile(tt.ExpectedOutputFile)
				require.Equal(t, string(b), rec.Body.String())
				require.NoError(t, err)
			}
		})
	}

}
