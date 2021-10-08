package models

import (
	"time"
)

type Task struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	AssingedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	Deadline   time.Time `json:"deadline"`
	Done       bool      `json:"done"`
	UserId     uint      `json:"user_id"`
	CategoryId uint      `json:"category_id"`
	Base
}
