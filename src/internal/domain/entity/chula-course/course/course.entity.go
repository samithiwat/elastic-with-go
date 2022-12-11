package course

import (
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	common "github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/common"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/section"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Course struct {
	entity.BaseMongo
	StudyProgram    common.StudyProgram `json:"studyProgram" mapstructure:"studyProgram"`
	Semester        common.Semester     `json:"semester" mapstructure:"semester"`
	AcademicYear    string              `json:"academicYear" mapstructure:"academicYear"`
	CourseNo        string              `json:"courseNo" mapstructure:"courseNo"`
	CourseNameTh    string              `json:"courseNameTh" mapstructure:"courseNameTh"`
	CourseNameEn    string              `json:"courseNameEn" mapstructure:"courseNameEn"`
	CourseDescTh    string              `json:"courseDescTh" mapstructure:"courseDescTh"`
	CourseDescEn    string              `json:"courseDescEn" mapstructure:"courseDescEn"`
	AbbrName        string              `json:"abbrName" mapstructure:"abbrName"`
	Faculty         string              `json:"faculty" mapstructure:"faculty"`
	Department      string              `json:"department" mapstructure:"department"`
	Credit          float32             `json:"credit" mapstructure:"credit"`
	CreditHour      uint                `json:"creditHour" mapstructure:"creditHour"`
	CourseCondition string              `json:"courseCondition" mapstructure:"courseCondition"`
	GenEdType       common.GenEdType    `json:"genEdType" mapstructure:"genEdType"`
	Rating          string              `json:"rating" mapstructure:"rating"`
	Midterm         common.ExamPeriod   `json:"midterm" mapstructure:"midterm"`
	Final           common.ExamPeriod   `json:"final" mapstructure:"final"`
	Sections        []*section.Section  `json:"sections" mapstructure:"sections"`
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

func (e *Course) ToDoc() any {
	return &CourseDoc{
		AbbrName:     e.AbbrName,
		CourseNo:     e.CourseNo,
		CourseNameTh: e.CourseNameTh,
		CourseNameEn: e.CourseNameEn,
		CourseDescTh: e.CourseDescTh,
		CourseDescEn: e.CourseDescEn,
		GenEdType:    string(e.GenEdType),
		Semester:     string(e.Semester),
		AcademicYear: e.AcademicYear,
		RawData:      e,
	}
}

func (e *Course) GetID() string {
	return e.ID.OID
}
