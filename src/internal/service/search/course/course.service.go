package course

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	courseRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type service struct {
	courseRepo courseRepo.Repository
	cacheTTL   int
}

func NewService(courseRepo courseRepo.Repository, cacheTTL int) pb.SearchServiceServer {
	return &service{
		courseRepo: courseRepo,
		cacheTTL:   cacheTTL,
	}
}

func (s *service) Search(_ context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var (
		result      []*pb.Course
		queryResult []*course.Course
	)

	if err := s.courseRepo.Search(req.QueryString, &queryResult); err != nil {
		log.Error().
			Err(err).
			Str("service", "course_search").
			Str("module", "search").
			Msg("error while query from elasticsearch")

		return nil, err
	}

	for _, c := range queryResult {
		result = append(result, c.ToProto())
	}

	return &pb.SearchResponse{Course: result}, nil
}
