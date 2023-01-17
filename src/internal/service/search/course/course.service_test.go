package course

import (
	"context"
	"github.com/go-faker/faker/v4"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	courseMock "github.com/samithiwat/elastic-with-go/src/mock/course"
	courseSearchMock "github.com/samithiwat/elastic-with-go/src/mock/search/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"testing"
)

type CourseSearchServiceTest struct {
	suite.Suite
	Metadata           *entity.PaginationMetadata
	Course             *course.Course
	CourseList         []*course.Course
	CourseDtoList      []*pb.Course
	EmptyCourseList    []*course.Course
	EmptyCourseDtoList []*pb.Course
	CacheTTL           int
}

func TestCourseSearchService(t *testing.T) {
	suite.Run(t, new(CourseSearchServiceTest))
}

func (t *CourseSearchServiceTest) SetupTest() {
	t.Metadata = &entity.PaginationMetadata{
		ItemsPerPage: rand.Intn(100),
		ItemCount:    rand.Intn(100),
		TotalItem:    rand.Intn(10000),
		CurrentPage:  rand.Intn(99),
		TotalPage:    rand.Intn(100),
	}

	t.CourseList = courseMock.CreateCourseList(rand.Intn(5)+1, false)
	t.Course = t.CourseList[0]
	t.EmptyCourseList = []*course.Course{}
	t.CacheTTL = rand.Intn(100000)

	t.CourseDtoList = []*pb.Course{}
	for _, c := range t.CourseList {
		t.CourseDtoList = append(t.CourseDtoList, c.ToProto())
	}
}

func (t *CourseSearchServiceTest) TestSearchSuccess() {
	want := &pb.SearchResponse{Pagination: &pb.CoursePagination{
		Items: t.CourseDtoList,
		Meta:  t.Metadata.ToProto(),
	}}

	testSearchSuccess(t.T(), want, &t.CourseList, 1, 20)
	testSearchSuccess(t.T(), want, &t.CourseList, 10, 20)
	testSearchSuccess(t.T(), want, &t.CourseList, 50, 10)
	testSearchSuccess(t.T(), want, &t.CourseList, 100, 10)

}

func testSearchSuccess(t *testing.T, want *pb.SearchResponse, courseList *[]*course.Course, page int32, limit int32) {
	query := faker.Word()
	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", &courseDto.Filter{Keyword: query},
		&emptyCourseList, &entity.PaginationMetadata{
			ItemsPerPage: int(limit),
			CurrentPage:  int(page),
		}).
		Return(courseList,
			&entity.PaginationMetadata{
				ItemsPerPage: int(want.Pagination.Meta.ItemsPerPage),
				ItemCount:    int(want.Pagination.Meta.ItemCount),
				TotalItem:    int(want.Pagination.Meta.TotalItem),
				CurrentPage:  int(want.Pagination.Meta.CurrentPage),
				TotalPage:    int(want.Pagination.Meta.TotalPage),
			}, nil)

	courseSearchSrv := NewService(&courseSearchRepo)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		Keyword: query,
		PaginationQuery: &pb.PaginationQuery{
			Limit: limit,
			Page:  page,
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, want, actual)
}

func (t *CourseSearchServiceTest) TestSearchElasticsearchConnectionError() {
	query := faker.Word()

	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", &courseDto.Filter{Keyword: query},
		&emptyCourseList, &entity.PaginationMetadata{
			ItemsPerPage: t.Metadata.ItemsPerPage,
			CurrentPage:  t.Metadata.CurrentPage,
		}).
		Return(nil, nil, status.Error(codes.Unavailable, "cannot connect to elasticsearch"))

	courseSearchSrv := NewService(&courseSearchRepo)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		Keyword: query,
		PaginationQuery: &pb.PaginationQuery{
			Limit: int32(t.Metadata.ItemsPerPage),
			Page:  int32(t.Metadata.CurrentPage),
		},
	})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Unavailable, st.Code())
}

func (t *CourseSearchServiceTest) TestSearchElasticsearchDecodeError() {
	query := faker.Word()

	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", &courseDto.Filter{Keyword: query},
		&emptyCourseList, &entity.PaginationMetadata{
			ItemsPerPage: t.Metadata.ItemsPerPage,
			CurrentPage:  t.Metadata.CurrentPage,
		}).
		Return(nil, nil, status.Error(codes.Internal, "cannot decode to struct"))

	courseSearchSrv := NewService(&courseSearchRepo)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		Keyword: query,
		PaginationQuery: &pb.PaginationQuery{
			Limit: int32(t.Metadata.ItemsPerPage),
			Page:  int32(t.Metadata.CurrentPage),
		},
	})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}
