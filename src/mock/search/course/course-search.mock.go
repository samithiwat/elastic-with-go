package course

import (
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
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

func (r *RepositoryMock) Search(in *pb.SearchRequest, result *[]*course.Course) error {
	args := r.Called(in, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*[]*course.Course)
	}

	return args.Error(1)
}
