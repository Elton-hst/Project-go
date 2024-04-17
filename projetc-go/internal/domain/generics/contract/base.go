package contract

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	PK        string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{
		PK:        uuid.NewString(),
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
