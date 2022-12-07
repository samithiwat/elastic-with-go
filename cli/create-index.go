package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/config"
	"github.com/samithiwat/elastic-with-go/src/database"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	fileName  string
	indexName string
)

const BaseIndexFilePath = "./src/database/elasticsearch-index"

func init() {
	flag.StringVar(&indexName, "index-name", "", "Index name")
	flag.StringVar(&fileName, "filename", "", "Index filename")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	conf, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start service")
	}

	client, err := database.InitElasticDefaultClient(conf.App.Debug)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init elasticsearch client")
	}

	if err := createIndex(client); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while creating the index")
	}

	fmt.Println(strings.Repeat("_", 65))
	fmt.Printf("successfully create the index, %s \n", indexName)
}

func createIndex(client *elasticsearch.Client) error {
	indexJsonRaw, err := os.ReadFile(BaseIndexFilePath + "/" + fileName)
	if err != nil {
		return err
	}

	res, err := client.Indices.Create(
		indexName,
		client.Indices.Create.WithBody(bytes.NewReader(indexJsonRaw)),
	)
	if err != nil {
		return err
	}

	if res.IsError() {
		return err
	}

	return nil
}
