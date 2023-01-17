package course

import (
	courseDto "github.com/samithiwat/elastic-with-go/src/internal/domain/dto/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
)

type Repository interface {
	Search(filter *courseDto.Filter, result *[]*course.Course, meta *entity.PaginationMetadata) error
	Suggest(text string, result *[]string) error
	Insert(indexName string, in *course.Course) error
	BulkInsert(in *[]*course.Course) error
}
