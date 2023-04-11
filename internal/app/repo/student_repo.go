package repo

import "github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"

//go:generate mockgen -source=$GOFILE -destination=$PROJ/internal/generated/mock_$GOPACKAGE/$GOFILE

type (
	Student struct {
		Name     string
		Gender   string
		Level    string
		State    string
		Major    string
		Activity string
	}
	StudentRepo interface {
		FindOne(gsheet string, row int) (*Student, error)
		Update(student *Student, gsheet string, row int) error
	}
	StudentRepoImpl struct{}
)

func init() {
	di.Provide(func() StudentRepo {
		return &StudentRepoImpl{}
	})
}

func (s *StudentRepoImpl) FindOne(gsheet string, row int) (*Student, error) {
	return nil, nil
}

func (s *StudentRepoImpl) Update(student *Student, gsheet string, row int) error {
	return nil
}
