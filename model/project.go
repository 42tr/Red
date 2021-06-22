package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string  `gorm:"comment:'项目名称'"`
	Description string  `gorm:"comment:'描述'"`
	CreateBy    uint    `json:"createBy"`
	Type        string  `gorm:"comment:'项目类别'"`
	Price       float32 `sql:"type:decimal(10,4);" gorm:"comment:'中标价（万元）'"`
	PartyId     uint    `gorm:"comment:'甲方ID'"`
	Area        string  `gorm:"comment:'地区'"`
}
