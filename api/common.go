package api

import (
	"red/serializer"
	"red/service"

	"github.com/gin-gonic/gin"
)

func GetDic(c *gin.Context) {
	c.JSON(200, serializer.Success(service.GetDic()))
}
