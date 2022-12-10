package course

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/mitchellh/mapstructure"
	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/constant/elasticsearch"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	searchRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type repository struct {
	searchRepo searchRepo.Repository
}

func NewRepository(searchRepo searchRepo.Repository) Repository {
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

	if err := r.searchRepo.Search(elasticsearchConstant.CourseIndexName, &req, &queryResultMap); err != nil {
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

func (r repository) Insert(s string, c *course.Course) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) BulkInsert(docID string, courseList *[]*course.Course) error {
	var courseDocList []*course.CourseDoc

	var wg sync.WaitGroup

	for pos, c := range *courseList {
		courseDoc := &course.CourseDoc{
			AbbrName:     c.AbbrName,
			CourseNo:     c.CourseNo,
			CourseNameTh: c.CourseNameTh,
			CourseNameEn: c.CourseNameEn,
			CourseDescTh: c.CourseDescTh,
			CourseDescEn: c.CourseDescEn,
			RawData:      c,
		}
		courseDocList = append(courseDocList, courseDoc)

	}

	wg.Wait()
}
