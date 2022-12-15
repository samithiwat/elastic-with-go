package course

import (
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Repository interface {
	Search(*pb.SearchRequest, *[]*course.Course) error
	Insert(string, *course.Course) error
	BulkInsert(*[]*course.Course) error
}
