package suggest

import (
	"context"
	"github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type service struct {
	repo course.Repository
}

func NewService(repo course.Repository) pb.SuggestServiceServer {
	return &service{
		repo: repo,
	}
}

func (s *service) Suggest(_ context.Context, req *pb.SuggestRequest) (*pb.SuggestResponse, error) {
	var result []string
	if err := s.repo.Suggest(req.Keyword, &result); err != nil {
		return nil, err
	}

	return &pb.SuggestResponse{
		Suggests: result,
	}, nil
}
