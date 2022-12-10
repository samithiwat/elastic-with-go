package elasticsearch

import "github.com/elastic/go-elasticsearch/v8/typedapi/core/search"

type Repository interface {
	Search(string, *search.Request, *map[string]interface{}) error
	Insert() error
	InsertBulk() error
}
