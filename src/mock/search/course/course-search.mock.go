package course

import (
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Insert(s string, c *course.Course) error {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryMock) BulkInsert(i *[]*course.Course) error {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryMock) Search(filter *courseDto.Filter, result *[]*course.Course, meta *entity.PaginationMetadata) error {
	args := r.Called(filter, result, meta)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*[]*course.Course)
	}

	if args.Get(1) != nil {
		*meta = *args.Get(1).(*entity.PaginationMetadata)
	}

	return args.Error(2)
}

func (r *RepositoryMock) Suggest(text string, result *[]string) error {
	args := r.Called(text, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*[]string)
	}

	return args.Error(1)
}
