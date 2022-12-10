package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/rs/zerolog/log"
	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/constant/elasticsearch"
	elasticsearchUtils "github.com/samithiwat/elastic-with-go/src/internal/utils/elasticsearch"
	"time"
)

type repository struct {
	esTypedClient   *elasticsearch.TypedClient
	esDefaultClient *elasticsearch.Client
}

func NewRepository(esTypedClient *elasticsearch.TypedClient, esDefaultClient *elasticsearch.Client) Repository {
	return &repository{
		esTypedClient:   esTypedClient,
		esDefaultClient: esDefaultClient,
	}
}

func (r repository) Search(indexName string, req *search.Request, result *map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.esTypedClient.Search().Index(indexName).Request(req).Do(ctx)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return err
	}

	return nil
}

func (r repository) Insert(indexName string, docID string, doc any) error {
	if _, err := r.esDefaultClient.Create(indexName, docID, esutil.NewJSONReader(doc)); err != nil {
		log.Error().
			Err(err).
			Msg("Error while create data to elasticsearch database")
		return err
	}
	return nil
}

func (r repository) InsertBulk(indexName string, docID string, doc any, pos int, nDocList int) error {
	buf := bytes.Buffer{}
	currentBatch := 0

	if err := elasticsearchUtils.AppendDocToBuffer(docID, doc, &buf); err != nil {
		log.Error().
			Err(err).
			Msg("Error while create request body")
		return err
	}

	currentBatch = pos / elasticsearchConstant.DocPerBatch
	if pos == nDocList-1 {
		currentBatch++
	}

	if pos > 0 && pos%elasticsearchConstant.DocPerBatch == 0 || pos == nDocList-1 {

		res, err := r.esDefaultClient.Bulk(bytes.NewReader(buf.Bytes()), r.esDefaultClient.Bulk.WithIndex(indexName))
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error while create data to elasticsearch database")
			return err
		}

		if res.IsError() {
			raw := map[string]interface{}{}

			if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
				log.Error().
					Err(err).
					Msgf("Failure to to parse response body")
				return err
			}

			log.Error().Msgf("  Error: [%d] %s: %s",
				res.StatusCode,
				raw["error"].(map[string]interface{})["type"],
				raw["error"].(map[string]interface{})["reason"],
			)

		}

		resMap := map[string]interface{}{}

		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Error().
				Err(err).
				Msgf("Failure to to parse response body")
			return err
		}

		for _, item := range resMap["items"].([]interface{}) {
			status := item.(map[string]interface{})["index"].(map[string]interface{})["status"].(float64)

			if status > 201 {
				resErr := item.(map[string]interface{})["index"].(map[string]interface{})["error"].(map[string]interface{})

				log.Error().Msgf("  Error: [%.0f]: %s: %s",
					status,
					resErr["type"],
					resErr["reason"],
				)
			}
		}

		buf.Reset()
	}

	return nil
}
