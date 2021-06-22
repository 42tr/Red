package service

import (
	"red/model"
	"red/serializer"

	"github.com/gin-gonic/gin"
)

type ProjectAddService struct {
	Name        string  `form:"name" json:"name" binding:"required,min=2,max=30"`
	Description string  `form:"description" json:"description" binding:"required,min=2,max=40"`
	Area        string  `form:"area" json:"area" binding:"required"`
	PartyId     uint    `form:"party" json:"party" binding:"required"`
	Type        string  `form:"type" json:"type" binding:"required"`
	Price       float32 `form:"price" json:"price" binding:"required"`
}

func (service *ProjectAddService) AddProject(c *gin.Context, uid uint) serializer.Response {
	project := model.Project{
		Name:        service.Name,
		Description: service.Description,
		CreateBy:    uid,
		Area:        service.Area,
		PartyId:     service.PartyId,
		Type:        service.Type,
		Price:       service.Price,
	}

	model.DB.Create(&project)
	return serializer.Success()
}

func (service *ProjectAddService) ChangeProject(c *gin.Context) serializer.Response {
	id := c.Param("id")
	var project model.Project
	model.DB.Where("id = ?", id).First(&project)

	project.Name = service.Name
	project.Description = service.Description
	project.PartyId = service.PartyId
	project.Type = service.Type
	project.Area = service.Area
	project.Price = service.Price

	model.DB.Save(&project)
	return serializer.Success()
}
