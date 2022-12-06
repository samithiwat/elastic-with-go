package section

import (
	"github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/class"
	common "github.com/samithiwat/elastic-with-go/src/internal/domain/entity/chula-course/common"
	"github.com/samithiwat/elastic-with-go/src/pb"
)

type Section struct {
	SectionNo string           `json:"sectionNo"`
	Closed    bool             `json:"closed"`
	Capacity  common.Capacity  `json:"capacity"`
	Note      string           `json:"note"`
	Classes   []*class.Class   `json:"classes"`
	GenEdType common.GenEdType `json:"genEdType"`
}

func (e *Section) ToProto() *pb.Section {
	var classes []*pb.Class

	if e.Classes != nil {
		for _, c := range e.Classes {
			classes = append(classes, c.ToProto())
		}
	}

	return &pb.Section{
		SectionNo: e.SectionNo,
		Closed:    e.Closed,
		Capacity:  e.Capacity.ToProto(),
		Note:      e.Note,
		Classes:   classes,
		GenEdType: string(e.GenEdType),
	}
}
