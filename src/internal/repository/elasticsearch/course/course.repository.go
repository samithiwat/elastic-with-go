package course

import "github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"

type Repository interface {
	Search(string, *[]*course.Course) error
	Insert(string, *course.Course) error
	BulkInsert(*[]*course.Course) error
}
