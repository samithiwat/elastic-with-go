package section

import (
	"github.com/samithiwat/elastic-with-go/src/domain/entity"
	"github.com/samithiwat/elastic-with-go/src/domain/entity/class"
)

type Section struct {
	SectionNo string           `json:"sectionNo"`
	Closed    bool             `json:"closed"`
	Capacity  entity.Capacity  `json:"capacity"`
	Note      string           `json:"note"`
	Classes   []*class.Class   `json:"classes"`
	GenEdType entity.GenEdType `json:"genEdType"`
}
