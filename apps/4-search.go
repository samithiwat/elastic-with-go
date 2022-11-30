package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/rs/zerolog/log"
	courseDto "github.com/samithiwat/elastic-with-go/src/domain/dto/course"
	"os"
	"time"
)

// Assume that we have data in the elasticsearch database

func main() {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
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

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	res, err := client.Search().Index("course_3").Request(&search.Request{
		Query: &types.Query{
			Match: map[string]types.MatchQuery{
				"courseNameEn": {
					Query: "marketing",
				},
			},
		},
	}).Do(ctx)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while query the result")
	}
	result := courseDto.QueryResult{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while parsing json to struct")
	}

	for pos, hit := range result.Hits.Hits {
		fmt.Printf("Result %d: ", pos)
		fmt.Printf("AbbrName: %s\n", hit.Source.AbbrName)
		fmt.Printf("CourseNo: %s\n", hit.Source.CourseNo)
		fmt.Printf("CourseNameTh: %s\n", hit.Source.CourseNameTh)
		fmt.Printf("CourseNameEn: %s\n", hit.Source.CourseNameEn)
		fmt.Printf("CourseDescTh: %s\n", hit.Source.CourseDescTh)
		fmt.Printf("CourseDescEn: %s\n", hit.Source.CourseDescEn)
		//fmt.Printf("RawData: %s", hit.Source.AbbrName)
	}
}
