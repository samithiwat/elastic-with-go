package course

import (
	"github.com/samithiwat/elastic-with-go/src/domain/entity/chula-course/course"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Search(queryString string, result *[]*course.Course) error {
	args := r.Called(queryString, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*[]*course.Course)
	}

	return args.Error(1)
}