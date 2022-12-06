package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/_example_apps/constant"
	"github.com/samithiwat/elastic-with-go/_example_apps/domain/entity/chula-course/course"
	"os"
)

// Assume we already had an index

func main() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "admin",
		Logger: &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while init elastic client")
	}

	{
		if _, err := client.Info(); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while connect to elastic client")
		}
	}

	courseIn, err := os.ReadFile("./inter-summer-courses.json")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while loading course from json")
	}

	var courseList []*course.Course

	if err := json.Unmarshal(courseIn, &courseList); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while parsing course to struct")
	}

	var courseDocList []*course.CourseDoc
	buf := bytes.Buffer{}
	currentBatch := 0
	nCourseList := len(courseList)

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

		if err := appendDocToBuffer(c.ID.OID, courseDoc, &buf); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while create request body")
		}

		currentBatch = pos / constant.DocPerBatch
		if pos == nCourseList-1 {
			currentBatch++
		}

		if pos > 0 && pos%constant.DocPerBatch == 0 || pos == nCourseList-1 {
			res, err := client.Bulk(bytes.NewReader(buf.Bytes()), client.Bulk.WithIndex(constant.IndexName))
			if err != nil {
				log.Fatal().
					Err(err).
					Msg("Error while create data to elasticsearch database")
			}

			if res.IsError() {
				raw := map[string]interface{}{}

				if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
					log.Fatal().Msgf("Failure to to parse response body: %s", err)
				}

				log.Printf("  Error: [%d] %s: %s",
					res.StatusCode,
					raw["error"].(map[string]interface{})["type"],
					raw["error"].(map[string]interface{})["reason"],
				)

			}

			resMap := map[string]interface{}{}

			if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
				log.Fatal().
					Err(err).
					Msgf("Failure to to parse response body: %s", err)
			}

			for _, item := range resMap["items"].([]interface{}) {
				status := item.(map[string]interface{})["index"].(map[string]interface{})["status"].(float64)

				if status > 201 {
					resErr := item.(map[string]interface{})["index"].(map[string]interface{})["error"].(map[string]interface{})

					log.Printf("  Error: [%.0f]: %s: %s",
						status,
						resErr["type"],
						resErr["reason"],
					)
				}
			}

			buf.Reset()
		}

	}
}

func appendDocToBuffer(docId string, docData interface{}, buf *bytes.Buffer) error {
	meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, docId, "\n"))
	data, err := json.Marshal(docData)
	if err != nil {
		return err
	}
	data = append(data, "\n"...)

	buf.Grow(len(meta) + len(data))
	buf.Write(meta)
	buf.Write(data)

	return nil
}
