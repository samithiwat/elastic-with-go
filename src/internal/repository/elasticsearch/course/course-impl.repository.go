package course

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/common/constant/elasticsearch"
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	esRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch"
	elasticsearchUtils "github.com/samithiwat/elastic-with-go/src/internal/utils/elasticsearch"
	"github.com/samithiwat/elastic-with-go/src/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repository struct {
	esRepo esRepo.Repository
}

func NewRepository(esRepo esRepo.Repository) Repository {
	return &repository{esRepo: esRepo}
}

func (r *repository) Search(in *pb.SearchRequest, result *[]*course.Course) error {
	queryResultMap := map[string]interface{}{}

	req := search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must: []types.Query{
					{
						MultiMatch: &types.MultiMatchQuery{
							Query:  in.Keyword,
							Fields: []string{"abbrName^5", "courseNo^5", "courseNameEn^3", "courseDescEn", "courseNameTh^3", "courseDescTh"},
						},
					},
					{
						Term: map[string]types.TermQuery{
							"semester": {Value: in.Semester},
						},
					},
					{
						Term: map[string]types.TermQuery{
							"studyProgram": {Value: in.StudyProgram},
						},
					},
					{
						Term: map[string]types.TermQuery{
							"academicYear": {Value: in.AcademicYear},
						},
					},
					{
						Terms: &types.TermsQuery{
							TermsQuery: map[string]types.TermsQueryField{
								"genEdType": in.GenEdType,
							},
						},
					},
					{
						Nested: &types.NestedQuery{
							Path: "rawData",
							Query: &types.Query{
								Nested: &types.NestedQuery{
									Path: "rawData.sections",
									Query: &types.Query{
										Nested: &types.NestedQuery{
											Path: "rawData.sections.classes",
											Query: &types.Query{
												Bool: &types.BoolQuery{
													Must: []types.Query{
														{
															Terms: &types.TermsQuery{
																TermsQuery: map[string]types.TermsQueryField{
																	"rawData.sections.classes.dayOfWeek": in.DayOfWeek,
																},
															},
														},
														{
															Nested: &types.NestedQuery{
																Path: "rawData.sections.classes.period",
																Query: &types.Query{
																	QueryString: &types.QueryStringQuery{
																		Query: fmt.Sprintf("rawData.sections.classes.period.start:[%s TO %s] AND rawData.sections.classes.period.end:[* TO %s]", in.PeriodRange.Start, in.PeriodRange.End, in.PeriodRange.End),
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
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

func (r *repository) Insert(docID string, course *course.Course) error {
	return r.esRepo.Insert(elasticsearchConstant.CourseIndexName, docID, course)
}

func (r *repository) BulkInsert(courseList *[]*course.Course) error {
	var (
		nDocList     = len(*courseList)
		buf          = bytes.Buffer{}
		currentBatch = 0
	)

	for pos, datum := range *courseList {
		pos := pos
		doc := datum.ToDoc()
		docID := datum.GetID()

		if err := elasticsearchUtils.AppendDocToBuffer(elasticsearchConstant.CourseIndexName, docID, doc, &buf); err != nil {
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
