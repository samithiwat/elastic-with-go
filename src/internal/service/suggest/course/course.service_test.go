package suggest

import (
	"context"
	"github.com/go-faker/faker/v4"
	courseSearchMock "github.com/samithiwat/elastic-with-go/src/mock/search/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type CourseSuggestServiceTest struct {
	suite.Suite
}

func TestCourseSuggestService(t *testing.T) {
	suite.Run(t, new(CourseSuggestServiceTest))
}

func (t *CourseSuggestServiceTest) SetupTest() {

}

func (t *CourseSuggestServiceTest) TestSuggestSuccess() {
	text := faker.Word()
	result := []string{"success"}
	want := pb.SuggestResponse{Suggests: result}

	var resultIn []string

	repo := courseSearchMock.RepositoryMock{}
	repo.On("Suggest", text, &resultIn).Return(&result, nil).Once()

	srv := NewService(&repo)

	actual, err := srv.Suggest(context.Background(), &pb.SuggestRequest{
		Keyword: text,
	})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), &want, actual)
}

func (t *CourseSuggestServiceTest) TestSuggestInternalErr() {
	text := faker.Word()

	var resultIn []string

	repo := courseSearchMock.RepositoryMock{}
	repo.On("Suggest", text, &resultIn).Return(nil, status.Error(codes.Internal, "Internal error")).Once()

	srv := NewService(&repo)

	actual, err := srv.Suggest(context.Background(), &pb.SuggestRequest{
		Keyword: text,
	})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}
