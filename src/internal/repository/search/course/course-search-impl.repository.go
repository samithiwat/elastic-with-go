package course

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/mitchellh/mapstructure"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	searchRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/search"
)

var IndexName = "course_4"

type repository struct {
	searchRepo searchRepo.Repository
}

func NewRepository(searchRepo searchRepo.Repository) SearchRepository {
	return &repository{searchRepo: searchRepo}
}

func (r repository) Search(queryString string, result *[]*course.Course) error {
	queryResultMap := map[string]interface{}{}

	req := search.Request{
		Query: &types.Query{
			QueryString: &types.QueryStringQuery{
				Query: queryString,
			},
		},
	}

	if err := r.searchRepo.Search(IndexName, &req, &queryResultMap); err != nil {
		return status.Error(codes.Unavailable, err.Error())
	}

	queryResult := courseDto.QueryResult{}

	if err := mapstructure.Decode(queryResultMap, &queryResult); err != nil {
		return status.Error(codes.Internal, "Cannot decoded from map to struct: "+err.Error())
	}

	for _, hit := range queryResult.Hits.Hits {
		*result = append(*result, hit.Source.RawData)
	}

	return nil
}
