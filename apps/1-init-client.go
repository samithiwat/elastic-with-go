package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
)

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
		res, err := client.Info()
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while connect to elastic client")
		}

		fmt.Println("Info:")
		fmt.Println(res)
	}
}
