package search

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
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
