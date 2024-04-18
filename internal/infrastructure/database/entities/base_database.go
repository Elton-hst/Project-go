package entities

import "time"

type BaseEntityDatabase struct {
	PK        string `gorm:"type:string;primary_key"`
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
