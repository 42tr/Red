package model

type Income struct {
	ID        uint    `json:"id"`
	ProjectId uint    `json:"projectId"`
	Type      string  `json:"type"`
	Amount    float32 `json:"amount" sql:"type:decimal(10,2);"`
	Date      string  `json:"date"`
	CreateBy  uint    `json:"createBy"`
}
