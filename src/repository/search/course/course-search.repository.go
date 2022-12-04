package course

import "github.com/samithiwat/elastic-with-go/src/domain/entity/chula-course/course"

type SearchRepository interface {
	Search(string, *[]*course.Course) error
}
