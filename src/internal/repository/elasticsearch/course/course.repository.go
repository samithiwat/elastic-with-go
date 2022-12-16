package course

import (
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Repository interface {
	Search(req *pb.SearchRequest, result *[]*course.Course, meta *entity.PaginationMetadata) error
	Insert(indexName string, in *course.Course) error
	BulkInsert(in *[]*course.Course) error
}
