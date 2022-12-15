package entity

import (
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type StudyProgram string
type Semester string
type GenEdType string
type ClassType string
type DayOfWeek string

type ExamPeriod struct {
	Date   *string `json:"date"`
	Period Period  `json:"period"`
}

func (e *ExamPeriod) ToProto() *pb.ExamPeriod {
	result := &pb.ExamPeriod{
		Period: e.Period.ToProto(),
	}

	if e.Date != nil {
		result.Date = *e.Date
	}

	return result
}

type Capacity struct {
	Current uint `json:"current"`
	Max     uint `json:"max"`
}

func (e *Capacity) ToProto() *pb.Capacity {
	return &pb.Capacity{
		Current: uint32(e.Current),
		Max:     uint32(e.Max),
	}
}

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (e *Period) ToProto() *pb.Period {
	return &pb.Period{
		Start: e.Start,
		End:   e.End,
	}
}
