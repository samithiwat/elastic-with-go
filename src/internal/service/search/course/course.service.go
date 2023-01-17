package course

import (
	"context"
	"github.com/rs/zerolog/log"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	commonChulaCourse "github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/common"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	courseRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type service struct {
	courseRepo courseRepo.Repository
}

func NewService(courseRepo courseRepo.Repository) pb.SearchServiceServer {
	return &service{
		courseRepo: courseRepo,
	}
}

func (s *service) Search(_ context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var (
		result      []*pb.Course
		queryResult []*course.Course
	)

	meta, filter := createMetaDataAndFilter(req)

	if err := s.courseRepo.Search(filter, &queryResult, meta); err != nil {
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

	return &pb.SearchResponse{Pagination: &pb.CoursePagination{
		Items: result,
		Meta:  meta.ToProto(),
	}}, nil
}

func createMetaDataAndFilter(req *pb.SearchRequest) (*entity.PaginationMetadata, *courseDto.Filter) {
	meta := &entity.PaginationMetadata{
		ItemsPerPage: int(req.PaginationQuery.Limit),
		CurrentPage:  int(req.PaginationQuery.Page),
	}

	filter := &courseDto.Filter{
		Keyword: req.Keyword,
	}

	if req.Filter != nil {
		filter.DayOfWeeks = req.Filter.DayOfWeek
		filter.GenEdTypes = req.Filter.GenEdType
		filter.Semester = req.Filter.Semester
		filter.AcademicYear = req.Filter.AcademicYear
		filter.StudyProgram = req.Filter.StudyProgram

		if req.Filter.PeriodRange != nil {
			filter.PeriodRange = &commonChulaCourse.Period{
				Start: req.Filter.PeriodRange.Start,
				End:   req.Filter.PeriodRange.End,
			}
		}
	}

	return meta, filter
}
