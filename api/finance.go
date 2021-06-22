package api

import (
	"github.com/gin-gonic/gin"
	"red/service"
)

func GetFinanceList(c *gin.Context) {
	var service service.FinanceService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetFinanceList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetFinanceAggregation(c *gin.Context) {
	var service service.FinanceService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetFinanceAggregation(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}