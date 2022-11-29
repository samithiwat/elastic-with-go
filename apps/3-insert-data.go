package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/domain/entity/course"
	"os"
)

// Assume we have index name test

func main() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "admin",
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

	var courses []*course.Course

	if err := json.Unmarshal(courseIn, &courses); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while parsing course to struct")
	}

	for _, c := range courses {
		courseDoc := course.CourseDoc{
			AbbrName:     c.AbbrName,
			CourseNo:     c.CourseNo,
			CourseNameTh: c.CourseNameTh,
			CourseNameEn: c.CourseNameEn,
			CourseDescTh: c.CourseDescTh,
			CourseDescEn: c.CourseDescEn,
			RawData:      c,
		}

		courseDocByte, _ := json.Marshal(courseDoc)

		res, err := client.Create("course_2", c.ID.OID, bytes.NewReader(courseDocByte))
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while create data to elasticsearch database")
		}
		fmt.Println(res)
	}

}
