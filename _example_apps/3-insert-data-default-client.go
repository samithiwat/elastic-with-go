package main

import (
	"encoding/json"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/rs/zerolog/log"
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

		if _, err := client.Create("course", c.ID.OID, esutil.NewJSONReader(courseDoc)); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while create data to elasticsearch database")
		}
	}

}
