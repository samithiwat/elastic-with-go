package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/database"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	fileName  string
	indexName string
	isDebug   bool
	indexCmd  *flag.FlagSet
)

const BaseIndexFilePath = "./src/database/elasticsearch-index"

func init() {
	indexCmd = flag.NewFlagSet("index", flag.ExitOnError)

	indexCmd.StringVar(&indexName, "index-name", "", "Index name")
	indexCmd.StringVar(&fileName, "filename", "", "Index filename")
	indexCmd.BoolVar(&isDebug, "debug", false, "Enable the debug log")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	client, err := database.InitElasticDefaultClient(isDebug)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init elasticsearch client")
	}

	if len(os.Args) < 2 {
		log.Fatal().
			Msg("Invalid input (need command)")
	}

	switch os.Args[1] {
	case "index":
		switch os.Args[2] {
		case "create":
			if err := handleCreateIndex(client); err != nil {
				log.Fatal().
					Err(err).
					Msg("Error while creating the index")
			}

			fmt.Println(strings.Repeat("_", 65))
			fmt.Printf("successfully create the index, %s \n", indexName)
		default:
			log.Fatal().
				Msg("Invalid input (invalid command)")
		}
	default:
		log.Fatal().
			Msg("Invalid input (invalid command)")
	}

}

func handleCreateIndex(client *elasticsearch.Client) error {
	if err := indexCmd.Parse(os.Args[3:]); err != nil {
		return err
	}

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
