package class

import (
	common "github.com/samithiwat/elastic-with-go/src/domain/entity/chula-course/common"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Class struct {
	Type      common.ClassType `json:"type"`
	DayOfWeek string           `json:"dayOfWeek"`
	Period    common.Period    `json:"period"`
	Building  string           `json:"building"`
	Room      string           `json:"room"`
	Teachers  []string         `json:"teachers"`
}

func (e *Class) ToProto() *pb.Class {
	return &pb.Class{
		Type:      string(e.Type),
		DayOfWeek: e.DayOfWeek,
		Period:    e.Period.ToProto(),
		Building:  e.Building,
		Room:      e.Room,
		Teachers:  e.Teachers,
	}
}
