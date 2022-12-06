package main

import (
	"context"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/_example_apps/utils"
	"os"
)

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

	req := create.Request{
		Aliases: map[string]types.Alias{
			"course": {},
		},
		Mappings: &types.TypeMapping{
			Properties: map[string]types.Property{
				"rawData":  types.NewObjectProperty(),
				"abbrName": types.NewTextProperty(),
				"courseNo": types.NewTextProperty(),
				"courseNameTh": types.TextProperty{
					Type:     "text",
					Analyzer: utils.StringAdr("thai"),
				},
				"courseNameEn": types.NewTextProperty(),
				"courseDescTh": types.TextProperty{
					Type:     "text",
					Analyzer: utils.StringAdr("thai"),
				},
				"courseDescEn": types.NewTextProperty(),
			},
		},
		Settings: &types.IndexSettings{
			NumberOfReplicas: "1",
			NumberOfShards:   "1",
		},
	}

	res, err := client.Indices.Create("course_5").
		Request(&req).
		Do(context.Background())

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while creating the course index")
	}

	defer res.Body.Close()
}
