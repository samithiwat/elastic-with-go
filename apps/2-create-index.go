package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
	"os"
)

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

	indexJsonRaw, err := os.ReadFile("./inter-summer-course-index.json")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while read the course index file")
	}

	res, err := client.Index("course_3", bytes.NewReader(indexJsonRaw))

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while creating the course index")
	}

	defer res.Body.Close()
	fmt.Println("Creating Index:")
	fmt.Println(res)
}
