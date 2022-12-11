package course

import (
	"context"
	"github.com/go-faker/faker/v4"
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
	queryString := faker.Word()

	want := &pb.SearchResponse{Course: t.CourseDtoList}

	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", queryString, &emptyCourseList).Return(&t.CourseList, nil)

	courseSearchSrv := NewService(&courseSearchRepo, t.CacheTTL)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		QueryString: queryString,
	})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CourseSearchServiceTest) TestSearchWithFilterSuccess() {
	queryString := faker.Word()

	want := &pb.SearchResponse{Course: t.CourseDtoList}

	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", queryString, &emptyCourseList).Return(&t.CourseList, nil)

	courseSearchSrv := NewService(&courseSearchRepo, t.CacheTTL)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		QueryString: queryString,
	})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CourseSearchServiceTest) TestSearchElasticsearchConnectionError() {
	queryString := faker.Word()

	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", queryString, &emptyCourseList).Return(nil, status.Error(codes.Unavailable, "cannot connect to elasticsearch"))

	courseSearchSrv := NewService(&courseSearchRepo, t.CacheTTL)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		QueryString: queryString,
	})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Unavailable, st.Code())
}

func (t *CourseSearchServiceTest) TestSearchElasticsearchDecodeError() {
	queryString := faker.Word()

	var emptyCourseList []*course.Course

	courseSearchRepo := courseSearchMock.RepositoryMock{}
	courseSearchRepo.On("Search", queryString, &emptyCourseList).Return(nil, status.Error(codes.Internal, "cannot decode to struct"))

	courseSearchSrv := NewService(&courseSearchRepo, t.CacheTTL)

	actual, err := courseSearchSrv.Search(context.Background(), &pb.SearchRequest{
		QueryString: queryString,
	})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}
