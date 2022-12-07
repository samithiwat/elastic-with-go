package course

type CourseDoc struct {
	AbbrName     string  `json:"abbrName"`
	CourseNo     string  `json:"courseNo"`
	CourseNameTh string  `json:"courseNameTh"`
	CourseNameEn string  `json:"courseNameEn"`
	CourseDescTh string  `json:"courseDescTh"`
	CourseDescEn string  `json:"courseDescEn"`
	GenEdType    string  `json:"genEdType"`
	DayOfWeek    string  `json:"dayOfWeek"`
	PeriodRange  string  `json:"periodRange"`
	RawData      *Course `json:"rawData"`
}
