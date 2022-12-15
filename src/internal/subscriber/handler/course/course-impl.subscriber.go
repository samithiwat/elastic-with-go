package course

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	courseRepo "github.com/samithiwat/elastic-with-go/src/internal/repository/elasticsearch/course"

	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
)

type handler struct {
	courseRepo courseRepo.Repository
}

func NewCourseSubscriberHandler(courseRepo courseRepo.Repository) Handler {
	return &handler{
		courseRepo: courseRepo,
	}
}

func (h *handler) InsertData(ctx context.Context) {
	raw := ctx.Value("message")
	var courseList []*course.Course

	if err := json.Unmarshal(raw.([]byte), &courseList); err != nil {
		log.Error().
			Err(err).
			Msg("Error while parsing course to struct")
		return
	}

	if err := h.courseRepo.BulkInsert(&courseList); err != nil {
		log.Error().
			Err(err).
			Str("service", "search").
			Str("module", "insert data subscriber handler").
			Msg("Error while insert course data to elasticsearch")
	}
}
