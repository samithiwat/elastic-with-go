package course

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	cacheRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/cache"
	courseRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	courseRepo courseRepo.Repository
	cacheRepo  cacheRepo.Repository
	cacheTTL   int
}

func NewService(courseRepo courseRepo.Repository, cacheRepo cacheRepo.Repository, cacheTTL int) pb.SearchServiceServer {
	return &service{
		courseRepo: courseRepo,
		cacheRepo:  cacheRepo,
		cacheTTL:   cacheTTL,
	}
}

func (s *service) Search(_ context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var result []*pb.Course

	if err := s.cacheRepo.GetCache(req.QueryString, &result); err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Error().
				Err(err).
				Str("service", "course_search").
				Str("module", "search").
				Msg("error while getting cached from database")

			return nil, status.Error(codes.Unavailable, "error while getting cached from database")
		}

		var queryResult []*course.Course

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

		if err := s.cacheRepo.SaveCache(req.QueryString, &result, s.cacheTTL); err != nil {
			log.Error().
				Err(err).
				Str("service", "course_search").
				Str("module", "search").
				Msg("error while saving cached to database")

			return nil, status.Error(codes.Unavailable, "error while saving cached to database")
		}
	}

	return &pb.SearchResponse{Course: result}, nil
}
