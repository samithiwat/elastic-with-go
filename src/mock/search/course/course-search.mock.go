package course

import (
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
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

func (r *RepositoryMock) Search(in *pb.SearchRequest, result *[]*course.Course, meta *entity.PaginationMetadata) error {
	args := r.Called(in, result, meta)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*[]*course.Course)
	}

	if args.Get(1) != nil {
		*meta = *args.Get(1).(*entity.PaginationMetadata)
	}

	return args.Error(2)
}
