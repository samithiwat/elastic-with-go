package course

type CourseDoc struct {
	AbbrName     string  `json:"abbrName"`
	CourseNo     string  `json:"courseNo"`
	CourseNameTh string  `json:"courseNameTh"`
	CourseNameEn string  `json:"courseNameEn"`
	CourseDescTh string  `json:"courseDescTh"`
	CourseDescEn string  `json:"courseDescEn"`
	GenEdType    string  `json:"genEdType"`
	StudyProgram string  `json:"studyProgram"`
	Semester     string  `json:"semester"`
	AcademicYear string  `json:"academicYear"`
	RawData      *Course `json:"rawData"`
}
