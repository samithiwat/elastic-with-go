package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
)

type Repository interface {
	Search(string, *search.Request, *map[string]interface{}) error
	Insert(indexName string, docID string, docData any) error
	InsertBulk(indexName string, docID string, docData any, pos int, nDocList int) error
}
