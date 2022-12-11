package course

type StudyProgramType string
type SemesterType string
type GenEdType string
type ClassType string
type DayOfWeek string

const (
	SO GenEdType = "SO"
	SC           = "SC"
	HU           = "HU"
	IN           = "IN"
	NO           = "NO"
)

const (
	Monday    DayOfWeek = "MO"
	Tuesday             = "TU"
	Wednesday           = "WE"
	Thursday            = "TH"
	Friday              = "FR"
	Saturday            = "SA"
	Sunday              = "SU"
	IA                  = "IA"
	AR                  = "AR"
)

const (
	Semester      StudyProgramType = "S"
	Trisemter                      = "T"
	International                  = "I"
)

const (
	First  SemesterType = "1"
	Second              = "2"
	Third               = "3"
)
