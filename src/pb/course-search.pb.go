// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: course-search.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudyProgram    string      `protobuf:"bytes,1,opt,name=StudyProgram,proto3" json:"StudyProgram,omitempty"`
	Semester        string      `protobuf:"bytes,2,opt,name=Semester,proto3" json:"Semester,omitempty"`
	AcademicYear    string      `protobuf:"bytes,3,opt,name=AcademicYear,proto3" json:"AcademicYear,omitempty"`
	CourseNo        string      `protobuf:"bytes,4,opt,name=CourseNo,proto3" json:"CourseNo,omitempty"`
	CourseNameTh    string      `protobuf:"bytes,5,opt,name=CourseNameTh,proto3" json:"CourseNameTh,omitempty"`
	CourseNameEn    string      `protobuf:"bytes,6,opt,name=CourseNameEn,proto3" json:"CourseNameEn,omitempty"`
	CourseDescTh    string      `protobuf:"bytes,7,opt,name=CourseDescTh,proto3" json:"CourseDescTh,omitempty"`
	CourseDescEn    string      `protobuf:"bytes,8,opt,name=CourseDescEn,proto3" json:"CourseDescEn,omitempty"`
	AbbrName        string      `protobuf:"bytes,9,opt,name=AbbrName,proto3" json:"AbbrName,omitempty"`
	Faculty         string      `protobuf:"bytes,10,opt,name=Faculty,proto3" json:"Faculty,omitempty"`
	Department      string      `protobuf:"bytes,11,opt,name=Department,proto3" json:"Department,omitempty"`
	Credit          uint32      `protobuf:"varint,12,opt,name=Credit,proto3" json:"Credit,omitempty"`
	CreditHour      uint32      `protobuf:"varint,13,opt,name=CreditHour,proto3" json:"CreditHour,omitempty"`
	CourseCondition string      `protobuf:"bytes,14,opt,name=CourseCondition,proto3" json:"CourseCondition,omitempty"`
	GenEdType       string      `protobuf:"bytes,15,opt,name=GenEdType,proto3" json:"GenEdType,omitempty"`
	Rating          string      `protobuf:"bytes,16,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Midterm         *ExamPeriod `protobuf:"bytes,17,opt,name=Midterm,proto3" json:"Midterm,omitempty"`
	Final           *ExamPeriod `protobuf:"bytes,18,opt,name=Final,proto3" json:"Final,omitempty"`
	Sections        []*Section  `protobuf:"bytes,19,rep,name=Sections,proto3" json:"Sections,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetStudyProgram() string {
	if x != nil {
		return x.StudyProgram
	}
	return ""
}

func (x *Course) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *Course) GetAcademicYear() string {
	if x != nil {
		return x.AcademicYear
	}
	return ""
}

func (x *Course) GetCourseNo() string {
	if x != nil {
		return x.CourseNo
	}
	return ""
}

func (x *Course) GetCourseNameTh() string {
	if x != nil {
		return x.CourseNameTh
	}
	return ""
}

func (x *Course) GetCourseNameEn() string {
	if x != nil {
		return x.CourseNameEn
	}
	return ""
}

func (x *Course) GetCourseDescTh() string {
	if x != nil {
		return x.CourseDescTh
	}
	return ""
}

func (x *Course) GetCourseDescEn() string {
	if x != nil {
		return x.CourseDescEn
	}
	return ""
}

func (x *Course) GetAbbrName() string {
	if x != nil {
		return x.AbbrName
	}
	return ""
}

func (x *Course) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *Course) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *Course) GetCredit() uint32 {
	if x != nil {
		return x.Credit
	}
	return 0
}

func (x *Course) GetCreditHour() uint32 {
	if x != nil {
		return x.CreditHour
	}
	return 0
}

func (x *Course) GetCourseCondition() string {
	if x != nil {
		return x.CourseCondition
	}
	return ""
}

func (x *Course) GetGenEdType() string {
	if x != nil {
		return x.GenEdType
	}
	return ""
}

func (x *Course) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *Course) GetMidterm() *ExamPeriod {
	if x != nil {
		return x.Midterm
	}
	return nil
}

func (x *Course) GetFinal() *ExamPeriod {
	if x != nil {
		return x.Final
	}
	return nil
}

func (x *Course) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionNo string    `protobuf:"bytes,1,opt,name=SectionNo,proto3" json:"SectionNo,omitempty"`
	Closed    bool      `protobuf:"varint,2,opt,name=Closed,proto3" json:"Closed,omitempty"`
	Capacity  *Capacity `protobuf:"bytes,3,opt,name=Capacity,proto3" json:"Capacity,omitempty"`
	Note      string    `protobuf:"bytes,4,opt,name=Note,proto3" json:"Note,omitempty"`
	Classes   []*Class  `protobuf:"bytes,5,rep,name=Classes,proto3" json:"Classes,omitempty"`
	GenEdType string    `protobuf:"bytes,6,opt,name=GenEdType,proto3" json:"GenEdType,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{1}
}

func (x *Section) GetSectionNo() string {
	if x != nil {
		return x.SectionNo
	}
	return ""
}

func (x *Section) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *Section) GetCapacity() *Capacity {
	if x != nil {
		return x.Capacity
	}
	return nil
}

func (x *Section) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Section) GetClasses() []*Class {
	if x != nil {
		return x.Classes
	}
	return nil
}

func (x *Section) GetGenEdType() string {
	if x != nil {
		return x.GenEdType
	}
	return ""
}

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	DayOfWeek string   `protobuf:"bytes,2,opt,name=DayOfWeek,proto3" json:"DayOfWeek,omitempty"`
	Period    *Period  `protobuf:"bytes,3,opt,name=Period,proto3" json:"Period,omitempty"`
	Building  string   `protobuf:"bytes,4,opt,name=Building,proto3" json:"Building,omitempty"`
	Room      string   `protobuf:"bytes,5,opt,name=Room,proto3" json:"Room,omitempty"`
	Teachers  []string `protobuf:"bytes,6,rep,name=Teachers,proto3" json:"Teachers,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{2}
}

func (x *Class) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Class) GetDayOfWeek() string {
	if x != nil {
		return x.DayOfWeek
	}
	return ""
}

func (x *Class) GetPeriod() *Period {
	if x != nil {
		return x.Period
	}
	return nil
}

func (x *Class) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *Class) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *Class) GetTeachers() []string {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type ExamPeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   string  `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	Period *Period `protobuf:"bytes,2,opt,name=Period,proto3" json:"Period,omitempty"`
}

func (x *ExamPeriod) Reset() {
	*x = ExamPeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExamPeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamPeriod) ProtoMessage() {}

func (x *ExamPeriod) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamPeriod.ProtoReflect.Descriptor instead.
func (*ExamPeriod) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{3}
}

func (x *ExamPeriod) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ExamPeriod) GetPeriod() *Period {
	if x != nil {
		return x.Period
	}
	return nil
}

type Capacity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current uint32 `protobuf:"varint,1,opt,name=Current,proto3" json:"Current,omitempty"`
	Max     uint32 `protobuf:"varint,2,opt,name=Max,proto3" json:"Max,omitempty"`
}

func (x *Capacity) Reset() {
	*x = Capacity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capacity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capacity) ProtoMessage() {}

func (x *Capacity) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capacity.ProtoReflect.Descriptor instead.
func (*Capacity) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{4}
}

func (x *Capacity) GetCurrent() uint32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Capacity) GetMax() uint32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type Period struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=Start,proto3" json:"Start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=End,proto3" json:"End,omitempty"`
}

func (x *Period) Reset() {
	*x = Period{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Period) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Period) ProtoMessage() {}

func (x *Period) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Period.ProtoReflect.Descriptor instead.
func (*Period) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{5}
}

func (x *Period) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Period) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword      string   `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
	GenEdType    []string `protobuf:"bytes,2,rep,name=GenEdType,proto3" json:"GenEdType,omitempty"`
	DayOfWeek    []string `protobuf:"bytes,3,rep,name=DayOfWeek,proto3" json:"DayOfWeek,omitempty"`
	PeriodRange  *Period  `protobuf:"bytes,4,opt,name=PeriodRange,proto3" json:"PeriodRange,omitempty"`
	StudyProgram string   `protobuf:"bytes,5,opt,name=StudyProgram,proto3" json:"StudyProgram,omitempty"`
	Semester     string   `protobuf:"bytes,6,opt,name=Semester,proto3" json:"Semester,omitempty"`
	AcademicYear string   `protobuf:"bytes,7,opt,name=AcademicYear,proto3" json:"AcademicYear,omitempty"`
	Limit        uint32   `protobuf:"varint,8,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset       uint32   `protobuf:"varint,9,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{6}
}

func (x *SearchRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchRequest) GetGenEdType() []string {
	if x != nil {
		return x.GenEdType
	}
	return nil
}

func (x *SearchRequest) GetDayOfWeek() []string {
	if x != nil {
		return x.DayOfWeek
	}
	return nil
}

func (x *SearchRequest) GetPeriodRange() *Period {
	if x != nil {
		return x.PeriodRange
	}
	return nil
}

func (x *SearchRequest) GetStudyProgram() string {
	if x != nil {
		return x.StudyProgram
	}
	return ""
}

func (x *SearchRequest) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *SearchRequest) GetAcademicYear() string {
	if x != nil {
		return x.AcademicYear
	}
	return ""
}

func (x *SearchRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course []*Course `protobuf:"bytes,1,rep,name=Course,proto3" json:"Course,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_search_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_search_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_course_search_proto_rawDescGZIP(), []int{7}
}

func (x *SearchResponse) GetCourse() []*Course {
	if x != nil {
		return x.Course
	}
	return nil
}

var File_course_search_proto protoreflect.FileDescriptor

var file_course_search_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0xa0, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x6f, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x54, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x45, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x54, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x73, 0x63, 0x54, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x73, 0x63, 0x45, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x73, 0x63, 0x45, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x62, 0x62, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x41, 0x62, 0x62, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46,
	0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x61,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x28, 0x0a,
	0x0f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x45, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x65, 0x6e, 0x45,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x33, 0x0a,
	0x07, 0x4d, 0x69, 0x64, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x07, 0x4d, 0x69, 0x64, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x2f, 0x0a, 0x05, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x05, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x52, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x45, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x65, 0x6e, 0x45, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x2d, 0x0a, 0x06,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x4f, 0x0a, 0x0a, 0x45, 0x78, 0x61, 0x6d, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x36, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x4d, 0x61, 0x78,
	0x22, 0x30, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45,
	0x6e, 0x64, 0x22, 0xb0, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x47, 0x65, 0x6e, 0x45, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x47, 0x65, 0x6e, 0x45, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x37, 0x0a, 0x0b, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x0b, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x79,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59,
	0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x63, 0x61, 0x64, 0x65,
	0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x32, 0x58, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_search_proto_rawDescOnce sync.Once
	file_course_search_proto_rawDescData = file_course_search_proto_rawDesc
)

func file_course_search_proto_rawDescGZIP() []byte {
	file_course_search_proto_rawDescOnce.Do(func() {
		file_course_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_search_proto_rawDescData)
	})
	return file_course_search_proto_rawDescData
}

var file_course_search_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_course_search_proto_goTypes = []interface{}{
	(*Course)(nil),         // 0: course_search.Course
	(*Section)(nil),        // 1: course_search.Section
	(*Class)(nil),          // 2: course_search.Class
	(*ExamPeriod)(nil),     // 3: course_search.ExamPeriod
	(*Capacity)(nil),       // 4: course_search.Capacity
	(*Period)(nil),         // 5: course_search.Period
	(*SearchRequest)(nil),  // 6: course_search.SearchRequest
	(*SearchResponse)(nil), // 7: course_search.SearchResponse
}
var file_course_search_proto_depIdxs = []int32{
	3,  // 0: course_search.Course.Midterm:type_name -> course_search.ExamPeriod
	3,  // 1: course_search.Course.Final:type_name -> course_search.ExamPeriod
	1,  // 2: course_search.Course.Sections:type_name -> course_search.Section
	4,  // 3: course_search.Section.Capacity:type_name -> course_search.Capacity
	2,  // 4: course_search.Section.Classes:type_name -> course_search.Class
	5,  // 5: course_search.Class.Period:type_name -> course_search.Period
	5,  // 6: course_search.ExamPeriod.Period:type_name -> course_search.Period
	5,  // 7: course_search.SearchRequest.PeriodRange:type_name -> course_search.Period
	0,  // 8: course_search.SearchResponse.Course:type_name -> course_search.Course
	6,  // 9: course_search.SearchService.Search:input_type -> course_search.SearchRequest
	7,  // 10: course_search.SearchService.Search:output_type -> course_search.SearchResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_course_search_proto_init() }
func file_course_search_proto_init() {
	if File_course_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Section); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExamPeriod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capacity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Period); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_search_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_course_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_search_proto_goTypes,
		DependencyIndexes: file_course_search_proto_depIdxs,
		MessageInfos:      file_course_search_proto_msgTypes,
	}.Build()
	File_course_search_proto = out.File
	file_course_search_proto_rawDesc = nil
	file_course_search_proto_goTypes = nil
	file_course_search_proto_depIdxs = nil
}
