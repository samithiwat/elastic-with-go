package course

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	courseRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch/course"

	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/constant/elasticsearch"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	elasticsearchUtils "github.com/samithiwat/elastic-with-go/src/internal/utils/elasticsearch"
	"sync"
)

type handler struct {
	courseRepo courseRepo.Repository
}

func NewCourseSubscriberHandler(courseRepo courseRepo.Repository) Handler {
	return handler{
		courseRepo: courseRepo,
	}
}

func (h handler) InsertData(ctx context.Context) {
	raw := ctx.Value("message")
	var courseList []*course.Course

	if err := json.Unmarshal(raw.([]byte), &courseList); err != nil {
		log.Error().
			Err(err).
			Msg("Error while parsing course to struct")
	}

	var courseDocList []*course.CourseDoc
	buf := bytes.Buffer{}
	currentBatch := 0
	nCourseList := len(courseList)

	var wg sync.WaitGroup

	for pos, c := range courseList {
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

		if err := elasticsearchUtils.AppendDocToBuffer(c.ID.OID, courseDoc, &buf); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while create request body")
		}

		currentBatch = pos / elasticsearchConstant.DocPerBatch
		if pos == nCourseList-1 {
			currentBatch++
		}

		if pos > 0 && pos%elasticsearchConstant.DocPerBatch == 0 || pos == nCourseList-1 {
			go func() {
				wg.Add(1)
				defer wg.Done()

				res, err := h.client.Bulk(bytes.NewReader(buf.Bytes()), h.client.Bulk.WithIndex(elasticsearchConstant.CourseIndexName))
				if err != nil {
					log.Error().
						Err(err).
						Msg("Error while create data to elasticsearch database")
				}

				if res.IsError() {
					raw := map[string]interface{}{}

					if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
						log.Error().
							Err(err).
							Msgf("Failure to to parse response body")
					}

					log.Error().Msgf("  Error: [%d] %s: %s",
						res.StatusCode,
						raw["error"].(map[string]interface{})["type"],
						raw["error"].(map[string]interface{})["reason"],
					)

				}

				resMap := map[string]interface{}{}

				if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
					log.Error().
						Err(err).
						Msgf("Failure to to parse response body")
				}

				for _, item := range resMap["items"].([]interface{}) {
					status := item.(map[string]interface{})["index"].(map[string]interface{})["status"].(float64)

					if status > 201 {
						resErr := item.(map[string]interface{})["index"].(map[string]interface{})["error"].(map[string]interface{})

						log.Error().Msgf("  Error: [%.0f]: %s: %s",
							status,
							resErr["type"],
							resErr["reason"],
						)
					}
				}

				buf.Reset()
			}()
		}
	}

	wg.Wait()
}
