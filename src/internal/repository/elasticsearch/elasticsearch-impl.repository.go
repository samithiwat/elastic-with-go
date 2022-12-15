package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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

	if res.StatusCode > 200 {
		// TODO add log error
		return errors.New("Invalid query")
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

func (r repository) InsertBulk(indexName string, buf *bytes.Buffer) error {
	res, err := r.esDefaultClient.Bulk(bytes.NewReader(buf.Bytes()), r.esDefaultClient.Bulk.WithIndex(indexName))
	if err != nil {
		return err
	}

	if res.IsError() {
		raw := map[string]interface{}{}

		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			return err
		}

		log.Error().Msgf("Error: [%d] %s: %s",
			res.StatusCode,
			raw["error"].(map[string]interface{})["type"],
			raw["error"].(map[string]interface{})["reason"],
		)

	}

	return nil
}
