package elasticsearch

import (
	"bytes"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
)

type Repository interface {
	Search(indexName string, req *search.Request, result *map[string]interface{}, meta *entity.PaginationMetadata) error
	Insert(indexName string, docID string, docData any) error
	InsertBulk(indexName string, buf *bytes.Buffer) error
}
