package models

import uuid "github.com/satori/go.uuid"

type Category struct {
	Base
	Name   string    `json:"name"`
	UserId uuid.UUID `gorm:"type:char(36);not null;"`
	Task   Task      `gorm:"foreignKey:CategoryId"`
}
