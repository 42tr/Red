package model

import "gorm.io/gorm"

type Apply struct {
	gorm.Model
	ProjectId  uint
	ApprovalId uint
	UserId     uint
	Money      float32
	Status     int
}