package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/rs/zerolog/log"
	elasticsearchConstant "github.com/samithiwat/elastic-with-go/src/constant/elasticsearch"
	elasticsearchUtils "github.com/samithiwat/elastic-with-go/src/internal/utils/elasticsearch"
	"time"
)

type repository struct {
	esClient *elasticsearch.TypedClient
}

func NewRepository(esClient *elasticsearch.TypedClient) Repository {
	return &repository{esClient: esClient}
}

func (r repository) Search(indexName string, req *search.Request, result *map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.esClient.Search().Index(indexName).Request(req).Do(ctx)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return err
	}

	return nil
}

func (r repository) Insert() error {
	//TODO implement me
	panic("implement me")
}

func (r repository) InsertBulk(doc any, nDocList int) error {
	buf := bytes.Buffer{}
	currentBatch := 0

	if err := elasticsearchUtils.AppendDocToBuffer(c.ID.OID, courseDoc, &buf); err != nil {
		log.Error().
			Err(err).
			Msg("Error while create request body")
	}

	currentBatch = pos / elasticsearchConstant.DocPerBatch
	if pos == nDocList-1 {
		currentBatch++
	}

	if pos > 0 && pos%elasticsearchConstant.DocPerBatch == 0 || pos == nDocList-1 {
		go func() {
			wg.Add(1)
			defer wg.Done()

			res, err := h.client.Bulk(bytes.NewReader(buf.Bytes()), h.client.Bulk.WithIndex(elasticsearchConstant.CourseIndexName))
			if err != nil {
				log.Error().
					Err(err).
					Msg("Error while create data to elasticsearch database")
			}

			if res.IsError() {
				raw := map[string]interface{}{}

				if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
					log.Error().
						Err(err).
						Msgf("Failure to to parse response body")
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
		}()
	}
}
