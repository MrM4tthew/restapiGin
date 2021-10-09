package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Task struct {
	Base
	AssingedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	Deadline   time.Time `json:"deadline"`
	Done       bool      `json:"done"`
	UserId     uuid.UUID `gorm:"type:char(36);not null;"`
	CategoryId uuid.UUID `gorm:"type:char(36);"`
}
