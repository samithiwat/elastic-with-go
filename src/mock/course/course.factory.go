package course

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/samithiwat/elastic-with-go/_example_apps/utils"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/class"
	common "github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/common"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/course"
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/section"
	"math/rand"
)

func CreateCourseList(n int, isIncludedID bool) []*course.Course {
	var result []*course.Course

	for i := 0; i < n; i++ {
		result = append(result, CreateCourse(isIncludedID))
	}

	return result
}

func CreateCourse(isIncludedID bool) *course.Course {
	result := &course.Course{
		StudyProgram:    common.StudyProgram(faker.Word()),
		Semester:        common.Semester(faker.Word()),
		AcademicYear:    faker.Word(),
		CourseNo:        faker.Word(),
		CourseNameTh:    faker.Word(),
		CourseNameEn:    faker.Word(),
		CourseDescTh:    faker.Paragraph(),
		CourseDescEn:    faker.Paragraph(),
		AbbrName:        faker.Word(),
		Faculty:         faker.Word(),
		Department:      faker.Word(),
		Credit:          float32(rand.Intn(5)) + 1,
		CreditHour:      uint(rand.Intn(5)) + 1,
		CourseCondition: faker.Sentence(),
		GenEdType:       common.GenEdType(faker.Word()),
		Rating:          faker.Word(),
		Midterm: common.ExamPeriod{
			Date: utils.StringAdr(faker.Word()),
			Period: common.Period{
				Start: faker.Word(),
				End:   faker.Word(),
			},
		},
		Final: common.ExamPeriod{
			Date: utils.StringAdr(faker.Word()),
			Period: common.Period{
				Start: faker.Word(),
				End:   faker.Word(),
			},
		},
		Sections: CreateSectionList(rand.Intn(5) + 1),
	}

	if isIncludedID {
		result.BaseMongo = entity.BaseMongo{ID: entity.ObjectID{
			OID: uuid.NewString(),
		}}
	}

	return result
}

func CreateSectionList(n int) []*section.Section {
	var result []*section.Section

	for i := 0; i < n; i++ {
		result = append(result, CreateSection())
	}

	return result
}

func CreateSection() *section.Section {
	return &section.Section{
		SectionNo: faker.Word(),
		Closed:    false,
		Capacity: common.Capacity{
			Current: uint(rand.Intn(500)),
			Max:     uint(rand.Intn(500)),
		},
		Note:      faker.Sentence(),
		Classes:   CreateClassList(rand.Intn(5) + 1),
		GenEdType: common.GenEdType(faker.Word()),
	}
}

func CreateClassList(n int) []*class.Class {
	var result []*class.Class

	for i := 0; i < n; i++ {
		result = append(result, CreateClass())
	}

	return result
}

func CreateClass() *class.Class {
	return &class.Class{
		Type:      common.ClassType(faker.Word()),
		DayOfWeek: faker.DayOfWeek(),
		Period: common.Period{
			Start: faker.Word(),
			End:   faker.Word(),
		},
		Building: faker.Word(),
		Room:     faker.Word(),
		Teachers: []string{faker.Word()},
	}
}
