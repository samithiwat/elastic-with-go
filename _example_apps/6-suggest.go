package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/_example_apps/utils"
	"os"
	"time"
)

const Text = "01231"

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

	res, err := client.Search().
		Index("course").
		Request(&search.Request{
			Suggest: &types.Suggester{
				Suggesters: map[string]types.FieldSuggester{
					"name": {Phrase: &types.PhraseSuggester{
						Field: "courseNameEn",
						DirectGenerator: []types.DirectGenerator{
							{
								Field:       "courseNameEn",
								SuggestMode: &suggestmode.Always,
							},
						},
					}},
					"no": {Term: &types.TermSuggester{
						Field: "courseNo",
					}},
				},
				Text: utils.StringAdr(Text),
			},
		}).
		Do(ctx)

	if err != nil {
		log.Fatal().
			Err(err).
			Str("text", Text).
			Msg("Error while suggest the word")
	}

	result := map[string]interface{}{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while parsing json to struct")
	}
	defer res.Body.Close()

	for name, suggest := range result["suggest"].(map[string]interface{}) {
		log.Print(name)
		for _, word := range suggest.([]interface{}) {
			//log.Print(word.(map[string]interface{})["text"])
			for _, option := range word.(map[string]interface{})["options"].([]interface{}) {
				log.Print(option.(map[string]interface{})["text"].(string))
			}

		}
	}
}
