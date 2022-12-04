package class

import (
	"github.com/samithiwat/elastic-with-go/src/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Class struct {
	Type      entity.ClassType `json:"type"`
	DayOfWeek string           `json:"dayOfWeek"`
	Period    entity.Period    `json:"period"`
	Building  string           `json:"building"`
	Room      string           `json:"room"`
	Teachers  []string         `json:"teachers"`
}

func (e *Class) ToProto() *pb.Class {
	return &pb.Class{
		Type:      e.Type,
		DayOfWeek: e.DayOfWeek,
		Period:    e.Period,
		Building:  "",
		Room:      "",
		Teachers:  nil,
	}
}
