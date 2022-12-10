package course

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/constant/elasticsearch"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	courseEsRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type repository struct {
	courseEsRepo courseEsRepo.Repository
}

func NewRepository(courseEsRepo courseEsRepo.Repository) Repository {
	return &repository{courseEsRepo: courseEsRepo}
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

	if err := r.courseEsRepo.Search(elasticsearchConstant.CourseIndexName, &req, &queryResultMap); err != nil {
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

func (r repository) Insert(docID string, course *course.Course) error {
	return r.courseEsRepo.Insert(elasticsearchConstant.CourseIndexName, docID, course)
}

func (r repository) BulkInsert(courseList *[]*course.Course) error {
	var (
		courseDocList []*course.CourseDoc
		wg            sync.WaitGroup
		nDocList      = len(*courseList)
	)

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

		pos := pos
		docID := c.ID.OID
		go func() {
			wg.Add(1)
			defer wg.Done()
			if err := r.courseEsRepo.InsertBulk(elasticsearchConstant.CourseIndexName, docID, &courseDoc, pos, nDocList); err != nil {
				log.Error().
					Err(err).
					Msg("Error while insert course data to elasticsearch database")
			}
		}()
	}

	wg.Wait()
	return nil
}
