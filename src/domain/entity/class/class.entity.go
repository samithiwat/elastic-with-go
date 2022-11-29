package class

import "github.com/samithiwat/elastic-with-go/src/domain/entity"

type Class struct {
	Type      entity.ClassType `json:"type"`
	DayOfWeek string           `json:"dayOfWeek"`
	Period    entity.Period    `json:"period"`
	Building  string           `json:"building"`
	Room      string           `json:"room"`
	Teachers  []string         `json:"teachers"`
}
