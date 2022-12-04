package course

import (
	"github.com/samithiwat/elastic-with-go/src/domain/entity"
	common "github.com/samithiwat/elastic-with-go/src/domain/entity/chula-course/common"
	"github.com/samithiwat/elastic-with-go/src/domain/entity/chula-course/section"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Course struct {
	entity.BaseMongo
	StudyProgram    common.StudyProgram `json:"studyProgram"`
	Semester        common.Semester     `json:"semester"`
	AcademicYear    string              `json:"academicYear"`
	CourseNo        string              `json:"courseNo"`
	CourseNameTh    string              `json:"courseNameTh"`
	CourseNameEn    string              `json:"courseNameEn"`
	CourseDescTh    string              `json:"courseDescTh"`
	CourseDescEn    string              `json:"courseDescEn"`
	AbbrName        string              `json:"abbrName"`
	Faculty         string              `json:"faculty"`
	Department      string              `json:"department"`
	Credit          uint                `json:"credit"`
	CreditHour      uint                `json:"creditHour"`
	CourseCondition string              `json:"courseCondition"`
	GenEdType       common.GenEdType    `json:"genEdType"`
	Rating          string              `json:"rating"`
	Midterm         common.ExamPeriod   `json:"midterm"`
	Final           common.ExamPeriod   `json:"final"`
	Sections        []*section.Section  `json:"sections"`
}

func (e *Course) ToProto() *pb.Course {
	var sectionList []*pb.Section

	if e.Sections != nil {
		for _, s := range e.Sections {
			sectionList = append(sectionList, s.ToProto())
		}
	}

	return &pb.Course{
		StudyProgram:    string(e.StudyProgram),
		Semester:        string(e.Semester),
		AcademicYear:    e.AcademicYear,
		CourseNo:        e.CourseNo,
		CourseNameTh:    e.CourseNameTh,
		CourseNameEn:    e.CourseNameEn,
		CourseDescTh:    e.CourseDescEn,
		CourseDescEn:    e.CourseDescTh,
		AbbrName:        e.AbbrName,
		Faculty:         e.Faculty,
		Department:      e.Department,
		Credit:          uint32(e.Credit),
		CreditHour:      uint32(e.CreditHour),
		CourseCondition: e.CourseCondition,
		GenEdType:       string(e.GenEdType),
		Rating:          e.Rating,
		Midterm:         e.Midterm.ToProto(),
		Final:           e.Final.ToProto(),
		Sections:        sectionList,
	}
}
