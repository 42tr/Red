package model

import "gorm.io/gorm"

type Approval struct {
	gorm.Model
	Name        string
}
