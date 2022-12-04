package course

import (
	"github.com/samithiwat/elastic-with-go/src/domain/dto"
	"github.com/samithiwat/elastic-with-go/src/domain/entity/course"
)

type QueryResult struct {
	*dto.QueryResult
	Hits *QueryHits `json:"hits"`
}

type QueryHits struct {
	Total    TotalHits `json:"total"`
	MaxScore float32   `json:"max_score"`
	Hits     []*Hits   `json:"hits"`
}

type TotalHits struct {
	Value    uint   `json:"value"`
	Relation string `json:"relation"`
}

type Hits struct {
	Index  string            `json:"_index"`
	ID     string            `json:"_id"`
	Score  float32           `json:"_score"`
	Ignore []string          `json:"_ignored"`
	Source *course.CourseDoc `json:"_source"`
}
