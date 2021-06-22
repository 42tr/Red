package model

import "gorm.io/gorm"

type Remark struct {
	gorm.Model
	ApplyId uint
	Remark  string
	UserId  uint
}
