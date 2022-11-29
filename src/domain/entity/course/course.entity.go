package course

import (
	"github.com/samithiwat/elastic-with-go/src/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/domain/entity/section"
)

type Course struct {
	entity.BaseMongo
	StudyProgram    entity.StudyProgram `json:"studyProgram"`
	Semester        entity.Semester     `json:"semester"`
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
	GenEdType       entity.GenEdType    `json:"genEdType"`
	Rating          string              `json:"rating"`
	Midterm         entity.ExamPeriod   `json:"midterm"`
	Final           entity.ExamPeriod   `json:"final"`
	Sections        []*section.Section  `json:"sections"`
}
