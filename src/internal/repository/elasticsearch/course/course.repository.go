package course

import "github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"

type SearchRepository interface {
	Search(string, *[]*course.Course) error
}
