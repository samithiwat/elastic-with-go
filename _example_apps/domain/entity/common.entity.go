package entity

import (
	"github.com/google/uuid"
	"github.com/samithiwat/elastic-with-go/_example_apps/utils"
	"github.com/samithiwat/elastic-with-go/src/pb"
	"gorm.io/gorm"
	"time"
)

type BaseMongo struct {
	ID ObjectID `json:"_id"`
}

type ObjectID struct {
	OID string `json:"$oid"`
}

type Base struct {
	ID        *uuid.UUID     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp;autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;type:timestamp"`
}

func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if b.ID == nil {
		b.ID = utils.UUIDAdr(uuid.New())
	}

	return nil
}

type PaginationMetadata struct {
	ItemsPerPage int
	ItemCount    int
	TotalItem    int
	CurrentPage  int
	TotalPage    int
}

func (p *PaginationMetadata) GetOffset() int {
	return (p.GetCurrentPage() - 1) * p.GetItemPerPage()
}

func (p *PaginationMetadata) GetItemPerPage() int {
	if p.ItemsPerPage < 10 {
		p.ItemsPerPage = 10
	}
	if p.ItemsPerPage > 100 {
		p.ItemsPerPage = 100
	}

	return p.ItemsPerPage
}

func (p *PaginationMetadata) GetCurrentPage() int {
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}

func (p *PaginationMetadata) ToProto() *pb.PaginationMetadata {
	return &pb.PaginationMetadata{
		TotalItem:    int64(p.TotalItem),
		ItemCount:    int64(p.ItemCount),
		ItemsPerPage: int64(p.ItemsPerPage),
		TotalPage:    int64(p.TotalPage),
		CurrentPage:  int64(p.CurrentPage),
	}
}
