package course

import (
	"bytes"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/constant/elasticsearch"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	esRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch"
	elasticsearchUtils "github.com/samithiwat/elastic-with-go/src/internal/utils/elasticsearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repository struct {
	esRepo esRepo.Repository
}

func NewRepository(esRepo esRepo.Repository) Repository {
	return &repository{esRepo: esRepo}
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

	if err := r.esRepo.Search(elasticsearchConstant.CourseIndexName, &req, &queryResultMap); err != nil {
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
	return r.esRepo.Insert(elasticsearchConstant.CourseIndexName, docID, course)
}

func (r repository) BulkInsert(courseList *[]*course.Course) error {
	var (
		nDocList     = len(*courseList)
		buf          = bytes.Buffer{}
		currentBatch = 0
	)

	for pos, datum := range *courseList {
		pos := pos
		doc := datum.ToDoc()
		docID := datum.GetID()

		if err := elasticsearchUtils.AppendDocToBuffer(docID, doc, &buf); err != nil {
			log.Error().
				Err(err).
				Msg("Error while create request body")
			return err
		}

		currentBatch = pos / elasticsearchConstant.DocPerBatch
		if pos == nDocList-1 {
			currentBatch++
		}

		if pos > 0 && pos%elasticsearchConstant.DocPerBatch == 0 || pos == nDocList-1 {
			if err := r.esRepo.InsertBulk(elasticsearchConstant.CourseIndexName, &buf); err != nil {
				log.Error().
					Err(err).
					Str("service", "course elasticsearch repository").
					Str("module", "insert bulk").
					Msg("Error while insert bulk data")
			}
			buf.Reset()
		}
	}

	log.Info().
		Str("service", "course elasticsearch repository").
		Str("module", "insert bulk").
		Msg("Successfully insert bulk data")

	return nil
}
